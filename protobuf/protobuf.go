package main

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	pb "protobuf/proto/user/pb"
)

func main(){
	person := &pb.Person{
		Name:"yz",
		Age:11,
		Emails:[]string{"782365461@qq.com","123456@163.com"},
		Phones:[]*pb.PhoneNumber{
			&pb.PhoneNumber{
				Number:"123456",
				Type:pb.PhoneType_HOME,
			},
			&pb.PhoneNumber{
				Number:"123456",
				Type:pb.PhoneType_MOBILE,
			},
			&pb.PhoneNumber{
				Number:"123456",
				Type:pb.PhoneType_WORK,
			},
		},
	}

	//marshal:  obj---[]byte
	data,err := proto.Marshal(person)
	if err != nil {
		fmt.Println(err)
	}

	//unmarshal : []byte---obj
	newPersonObj := &pb.Person{}
	err = proto.Unmarshal(data,newPersonObj)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(newPersonObj)
}