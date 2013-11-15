package main

import (
	"fmt"
	"os"
)

func fileWrite() {
        userFile := "./test.txt"
        fout,err := os.Create(userFile)
        
        defer fout.Close()
        if err != nil {
                fmt.Println(userFile,err)
                return
        }
        for i:= 0;i<10;i++ {
                fout.WriteString("Just a test测试!\r\n")
                fout.Write([]byte("Just a testa!\r\n"))
        }
}