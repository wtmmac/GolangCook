package grpc

// 生成方法1
// go generate
//go:generate /opt/homebrew/bin/protoc -I ./proto/ --go_out=plugins=grpc:./api/ --go_opt=paths=source_relative ./proto/hello.proto
