package model

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand/v2"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/patrickmn/go-cache"
	"github.com/redis/go-redis/v9"
	mycache "github.com/suutest/app/product/biz/cache"
	"gorm.io/gorm"
)

type Product struct {
	Base
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Picture     string  `json:"picture"`
	Price       float32 `json:"price"`

	Categories []Category `json:"categories" gorm:"many2many:product_category"`
}

func (p Product) TableName() string {
	return "product"
}

type ProductQuery struct {
	ctx context.Context
	db  *gorm.DB
}

func (p ProductQuery) GetById(productId int) (product Product, err error) {
	err = p.db.WithContext(p.ctx).Model(&Product{}).First(&product, productId).Error // 这里调用WithContext方法主要是便于后面做链路追踪
	return
}

func (p ProductQuery) SearchProduct(q string) (products []*Product, err error) {
	err = p.db.WithContext(p.ctx).Model(&Product{}).Find(&products, "name like ? or description like ?",
		"%"+q+"%", "%"+q+"%",
	).Error // 模糊查询
	return
}

func NewProductQuery(ctx context.Context, db *gorm.DB) *ProductQuery {
	return &ProductQuery{
		ctx: ctx,
		db:  db,
	}
}

func (p ProductQuery) BatchGetByIds(productIds []uint32) (products []Product, err error) {
	err = p.db.WithContext(p.ctx).Debug().Model(&Product{}).Where("id IN ?", productIds).Find(&products).Error // 这里调用WithContext方法主要是便于后面做链路追踪
	return
}

type CachedProductQuery struct {
	productQuery ProductQuery
	redisClient  *redis.Client
	localCache   *cache.Cache
	prefix       string
}

func (c CachedProductQuery) GetById(productId int) (product Product, err error) {
	cachedKey := fmt.Sprintf("%s_%s_%d", c.prefix, "product_by_id", productId) // 用这个key从redis中获取数据
	cachedResult := c.redisClient.Get(c.productQuery.ctx, cachedKey)

	// 用闭包构建一个错误链，如果中间有任何一个发生错误就往下走
	err = func() error {
		if err := cachedResult.Err(); err != nil {
			return err
		}
		cachedResultByte, err := cachedResult.Bytes()
		if err != nil {
			return err
		}
		err = json.Unmarshal(cachedResultByte, &product)
		if err != nil {
			return err
		}
		return nil
	}()
	if err != nil { // 如果上面这些步骤里有错误发生，我们就尝试从数据库中获取数据
		product, err = c.productQuery.GetById(productId)
		if err != nil {
			return Product{}, err
		}
		// 如果数据库中的数据获取成功，就把它做序列化然后存在缓存里
		encoded, err := json.Marshal(product)
		if err != nil {
			return product, nil
		}
		_ = c.redisClient.Set(c.productQuery.ctx, cachedKey, encoded, time.Hour)
	}
	return
}

func (c CachedProductQuery) TwoLevelCacheGetById(productId int) (product Product, err error) {
	cachedKey := fmt.Sprintf("%s_%s_%d", c.prefix, "product_by_id", productId)
	// 1.先查本地缓存
	if item, found := c.localCache.Get(cachedKey); found {
		return item.(Product), nil
	}

	klog.Info("Local memory cache miss")

	// 2.本地未命中则查Redis
	cachedResult := c.redisClient.Get(c.productQuery.ctx, cachedKey)

	err = func() error {
		if err := cachedResult.Err(); err != nil {
			return err
		}
		cachedResultByte, err := cachedResult.Bytes()
		if err != nil {
			return err
		}
		err = json.Unmarshal(cachedResultByte, &product)
		if err != nil {
			return err
		}
		return nil
	}()
	if err != nil { // 3. 如果上面这些步骤里有错误发生，我们就尝试从数据库中获取数据
		product, err = c.productQuery.GetById(productId)
		if err != nil {
			return Product{}, err
		}
		// 如果数据库中的数据获取成功，就把它做序列化然后存在缓存里
		encoded, err := json.Marshal(product)
		if err != nil {
			return product, nil
		}
		d := time.Duration(rand.IntN(30)) * time.Second // 带有随机值防止缓存雪崩
		_ = c.redisClient.Set(c.productQuery.ctx, cachedKey, encoded, time.Hour+d)
		// 存在本地缓存
		c.localCache.Set(cachedKey, product, 5*time.Minute+d)
	}
	return
}

func (c CachedProductQuery) SearchProduct(q string) (products []*Product, err error) {
	// 这里假设商品搜索的命中率比较低，所以还是直接从数据库中获取数据
	return c.productQuery.SearchProduct(q)
}

func NewCachedProductQuery(ctx context.Context, db *gorm.DB, cachedClient *redis.Client) *CachedProductQuery {
	return &CachedProductQuery{
		productQuery: *NewProductQuery(ctx, db),
		redisClient:  cachedClient,
		localCache:   mycache.LocalCache,
		prefix:       "shop",
	}
}

func (c CachedProductQuery) BatchGetByIds(productIds []uint32) (products []Product, err error) {
	var missed_ids []uint32
	for _, i := range productIds {
		cachedKey := fmt.Sprintf("%s_%s_%d", c.prefix, "product_by_id", i)
		cachedResult := c.redisClient.Get(c.productQuery.ctx, cachedKey)
		err = func() error {
			if err := cachedResult.Err(); err != nil {
				return err
			}
			cachedResultByte, err := cachedResult.Bytes()
			if err != nil {
				return err
			}
			var p Product
			err = json.Unmarshal(cachedResultByte, &p)
			products = append(products, p)
			if err != nil {
				return err
			}
			return nil
		}()
		if err != nil {
			missed_ids = append(missed_ids, i)
		}
	}
	resp, err := c.productQuery.BatchGetByIds(missed_ids)
	if err != nil {
		return products, err
	}
	products = append(products, resp...)

	for _, i := range missed_ids {
		cachedKey := fmt.Sprintf("%s_%s_%d", c.prefix, "product_by_id", i)
		for _, p := range resp {
			if uint32(p.ID) == i {
				encoded, err := json.Marshal(p)
				if err != nil {
					return products, err
				}
				_ = c.redisClient.Set(c.productQuery.ctx, cachedKey, encoded, time.Hour)
				break
			}
		}
	}
	return
}

// 读写分离：给ProductQuery传读库的db，给ProductMutation传写库的db，就可以实现简单的读写分离
type ProductMutation struct {
	ctx context.Context
	db  *gorm.DB
}
