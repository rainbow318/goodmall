package main

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/suuyh/demo/demo_proto/biz/dal"
	"github.com/suuyh/demo/demo_proto/biz/dal/mysql"
	"github.com/suuyh/demo/demo_proto/biz/model"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	dal.Init()
	// CURD
	// Create
	mysql.DB.Create(&model.User{Email: "demo@example.com", Password: "123kjdf"})

	// Update
	mysql.DB.Model(&model.User{}).Where("email=?", "demo@example.com").Update("password", "2222")

	// Read
	var row model.User
	mysql.DB.Model(&model.User{}).Where("email=?", "demo@example.com").First(&row)
	fmt.Printf("row: %+v\n", row)

	// Delete
	// mysql.DB.Where("email=?", "demo@example.com").Delete(&model.User{}) // 软删
	mysql.DB.Unscoped().Where("email=?", "demo@example.com").Delete(&model.User{}) // 完全删
}
