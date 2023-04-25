package grpc

////go:generate protoc -I. -I$GOPATH/src --go_out=plugins=grpc:. --go_opt=paths=source_relative logic/logic.proto
//go:generate protoc -I ./api/ --go_out=plugins=grpc:./api/ --go_opt=paths=source_relative ./api/hello.proto
