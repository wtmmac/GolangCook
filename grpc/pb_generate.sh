# 生成方法2
# Tips：将--plugin= 后边的 protoc-gen-go 和 protoc-gen-go-rpc 的前缀路径换成自己实际的地址
# protoc --go-grpc_opt=require_unimplemented_servers=false 设置一个选项来生成没有向前兼容的代码
protoc --go-grpc_opt=require_unimplemented_servers=false -I=./ ./proto/*.proto --plugin=$(which protoc-gen-go) --go_out=./api --plugin=$(which protoc-gen-go-grpc) --go-grpc_out=./api
