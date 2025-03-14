## 库存服务
---
![截屏 2025-03-14 11 54 59](https://github.com/user-attachments/assets/397e9dea-5f16-4aaa-ae13-df8d429095ff)
TODO：redis的`stock:pre:[productId]`key，预扣时增加支付倒计时，到期后仍未支付的订单需取消（定时任务）？DB扣库存成功后，将`stock:pre:[productId]` key删除
