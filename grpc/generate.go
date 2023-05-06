package grpc

// 生成方法1
// go generate
//go:generate protoc -I ./api/ --go_out=plugins=grpc:./api/ --go_opt=paths=source_relative ./api/hello.proto
