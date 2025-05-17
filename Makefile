.PHONY: gen-frontend
gen-frontend:
	cd app/frontend && cwgo server --type HTTP --idl ../../idl/frontend/auth.proto --service frontend --module github.com/spark4862/smartmall/app/frontend -I ../../idl
	# cd 用于帮助--module找到正确module

.PHONY: gen-user
gen-user:
	cd rpc_gen && cwgo client --type RPC --service user --module github.com/spark4862/smartmall/rpc_gen -I ../idl --idl ../idl/user.proto
	cd app/user && cwgo server --type RPC --service user --module github.com/spark4862/smartmall/app/user -I ../../idl --idl ../../idl/user.proto --pass "--use github.com/spark4862/smartmall/rpc_gen/kitex_gen"
	# --pass让生成服务端代码时不再生成客户端代码