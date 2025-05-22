export ROOT_MOD=github.com/spark4862/smartmall

export FRONTEND_MOD="order_page"
.PHONY: gen-frontend
gen-frontend:
	cd app/frontend && cwgo server --type HTTP --idl ../../idl/frontend/${FRONTEND_MOD}.proto --service frontend --module ${ROOT_MOD}/app/frontend -I ../../idl
	# cd 用于帮助--module找到正确module

.PHONY: gen-user
gen-user:
	cd rpc_gen && cwgo client --type RPC --service user --module ${ROOT_MOD}/rpc_gen -I ../idl --idl ../idl/user.proto
	cd app/user && cwgo server --type RPC --service user --module ${ROOT_MOD}/app/user -I ../../idl --idl ../../idl/user.proto --pass "--use ${ROOT_MOD}/rpc_gen/kitex_gen"
	# --pass让生成服务端代码时不再生成客户端代码



export MOD=product
.PHONY: gen-product
gen-product:
	cd rpc_gen && cwgo client --type RPC --service ${MOD} --module ${ROOT_MOD}/rpc_gen -I ../idl --idl ../idl/${MOD}.proto
	cd app/${MOD} && cwgo server --type RPC --service ${MOD} --module ${ROOT_MOD}/app/${MOD} -I ../../idl --idl ../../idl/${MOD}.proto --pass "--use ${ROOT_MOD}/rpc_gen/kitex_gen"

export MOD=cart
.PHONY: gen-cart
gen-cart:
	cd rpc_gen && cwgo client --type RPC --service ${MOD} --module ${ROOT_MOD}/rpc_gen -I ../idl --idl ../idl/${MOD}.proto
	cd app/${MOD} && cwgo server --type RPC --service ${MOD} --module ${ROOT_MOD}/app/${MOD} -I ../../idl --idl ../../idl/${MOD}.proto --pass "--use ${ROOT_MOD}/rpc_gen/kitex_gen"

export MOD=payment
.PHONY: gen-payment
gen-payment:
	cd rpc_gen && cwgo client --type RPC --service ${MOD} --module ${ROOT_MOD}/rpc_gen -I ../idl --idl ../idl/${MOD}.proto
	cd app/${MOD} && cwgo server --type RPC --service ${MOD} --module ${ROOT_MOD}/app/${MOD} -I ../../idl --idl ../../idl/${MOD}.proto --pass "--use ${ROOT_MOD}/rpc_gen/kitex_gen"

export MOD=checkout
.PHONY: gen-checkout
gen-checkout:
	cd rpc_gen && cwgo client --type RPC --service ${MOD} --module ${ROOT_MOD}/rpc_gen -I ../idl --idl ../idl/${MOD}.proto
	cd app/${MOD} && cwgo server --type RPC --service ${MOD} --module ${ROOT_MOD}/app/${MOD} -I ../../idl --idl ../../idl/${MOD}.proto --pass "--use ${ROOT_MOD}/rpc_gen/kitex_gen"

export MOD=order
.PHONY: gen-order
gen-order:
	cd rpc_gen && cwgo client --type RPC --service ${MOD} --module ${ROOT_MOD}/rpc_gen -I ../idl --idl ../idl/${MOD}.proto
	cd app/${MOD} && cwgo server --type RPC --service ${MOD} --module ${ROOT_MOD}/app/${MOD} -I ../../idl --idl ../../idl/${MOD}.proto --pass "--use ${ROOT_MOD}/rpc_gen/kitex_gen"

export MOD=email
.PHONY: gen-email
gen-email:
	cd rpc_gen && cwgo client --type RPC --service ${MOD} --module ${ROOT_MOD}/rpc_gen -I ../idl --idl ../idl/${MOD}.proto
	cd app/${MOD} && cwgo server --type RPC --service ${MOD} --module ${ROOT_MOD}/app/${MOD} -I ../../idl --idl ../../idl/${MOD}.proto --pass "--use ${ROOT_MOD}/rpc_gen/kitex_gen"