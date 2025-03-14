export ROOT_MOD=github.com/suutest
.PHONY: gen-demo-proto
gen-demo-proto:
	@cd demo/demo_proto && cwgo server -I ../../idl --type RPC --module github.com/suuyh/demo/demo_proto --service demo_proto --idl ../../idl/echo.proto

.PHONY: gen-demo-thrift
gen-demo-thrift:
	@cd demo/demo_thrift && cwgo server -I ../../idl --type RPC --module github.com/suuyh/demo/demo_thrift --service demo_thrift --idl ../../idl/echo.thrift
	
.PHONY: gen-frontend
gen-frontend:
	@cd app/frontend && cwgo server --type HTTP --idl ../../idl/frontend/order_page.proto --service frontend -module ${ROOT_MOD}/app/frontend -I ../../idl

.PHONY: gen-user
gen-user:
	@cd rpc_gen && cwgo client --type RPC --service user --module ${ROOT_MOD}/rpc_gen --I ../idl --idl ../idl/user.proto
	@cd app/user && cwgo server --type RPC --service user --module ${ROOT_MOD}/app/user --pass "-use ${ROOT_MOD}/rpc_gen/kitex_gen" -I ../../idl --idl ../../idl/user.proto

.PHONY: gen-product
gen-product:
	@cd rpc_gen && cwgo client --type RPC --service product --module ${ROOT_MOD}/rpc_gen -I ../idl --idl ../idl/product.proto
	@cd app/product && cwgo server --type RPC --service product --module ${ROOT_MOD}/app/product --pass "-use ${ROOT_MOD}/rpc_gen/kitex_gen" -I ../../idl --idl ../../idl/product.proto

.PHONY: gen-cart
gen-cart:
	@cd rpc_gen && cwgo client --type RPC --service cart --module ${ROOT_MOD}/rpc_gen -I ../idl --idl ../idl/cart.proto
	@cd app/cart && cwgo server --type RPC --service cart --module ${ROOT_MOD}/app/cart --pass "-use ${ROOT_MOD}/rpc_gen/kitex_gen" -I ../../idl --idl ../../idl/cart.proto

.PHONY: gen-payment
gen-payment:
	@cd rpc_gen && cwgo client --type RPC --service payment --module ${ROOT_MOD}/rpc_gen -I ../idl --idl ../idl/payment.proto
	@cd app/payment && cwgo server --type RPC --service payment --module ${ROOT_MOD}/app/payment --pass "-use ${ROOT_MOD}/rpc_gen/kitex_gen" -I ../../idl --idl ../../idl/payment.proto

.PHONY: gen-checkout
gen-checkout:
	@cd rpc_gen && cwgo client --type RPC --service checkout --module ${ROOT_MOD}/rpc_gen -I ../idl --idl ../idl/checkout.proto
	@cd app/checkout && cwgo server --type RPC --service checkout --module ${ROOT_MOD}/app/checkout --pass "-use ${ROOT_MOD}/rpc_gen/kitex_gen" -I ../../idl --idl ../../idl/checkout.proto

.PHONY: gen-order
gen-order:
	@cd rpc_gen && cwgo client --type RPC --service order --module ${ROOT_MOD}/rpc_gen -I ../idl --idl ../idl/order.proto
	@cd app/order && cwgo server --type RPC --service order --module ${ROOT_MOD}/app/order --pass "-use ${ROOT_MOD}/rpc_gen/kitex_gen" -I ../../idl --idl ../../idl/order.proto

.PHONY: gen-email
gen-email:
	@cd rpc_gen && cwgo client --type RPC --service email --module ${ROOT_MOD}/rpc_gen -I ../idl --idl ../idl/email.proto
	@cd app/email && cwgo server --type RPC --service email --module ${ROOT_MOD}/app/email --pass "-use ${ROOT_MOD}/rpc_gen/kitex_gen" -I ../../idl --idl ../../idl/email.proto

.PHONY: gen-stock
gen-stock:
	@cd rpc_gen && cwgo client --type RPC --service stock --module ${ROOT_MOD}/rpc_gen -I ../idl --idl ../idl/stock.proto
	@cd app/stock && cwgo server --type RPC --service stock --module ${ROOT_MOD}/app/stock --pass "-use ${ROOT_MOD}/rpc_gen/kitex_gen" -I ../../idl --idl ../../idl/stock.proto


.PHONY: go-user
go-user:
	@cd app/user && go mod tidy && go run .

.PHONY: go-cart
go-cart:
	@cd app/cart && go mod tidy && go run .

.PHONY: go-product
go-product:
	@cd app/product && go mod tidy && go run .

.PHONY: go-checkout
go-checkout:
	@cd app/checkout && go mod tidy && go run .

.PHONY: go-order
go-order:
	@cd app/order && go mod tidy && go run .

.PHONY: go-email
go-email:
	@cd app/email && go mod tidy && go run .

.PHONY: go-frontend
go-frontend:
	@cd app/frontend && go mod tidy && go run .

.PHONY: go-payment
go-payment:
	@cd app/payment && go mod tidy && go run .

.PHONY: go-stock
go-stock:
	@cd app/stock && go mod tidy && go run .