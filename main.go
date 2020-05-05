package main

import (
	"fmt"
	"go-protobuf/pkg/service"
	"log"
	"time"
)

var (
	username = "admin"
	password = "admin"
)

func main() {
	fmt.Println("Protobuf example")
	fmt.Println()

	rows := 100

	ms := time.Now()
	for i := 0; i < rows; i++ {
		protobuf := service.NewUserService(username, password)
		if protobuf == nil {
			panic("protobuf is nil")
		}

		protobuf.ToBinary()
		// fmt.Println(bytesData)
	}
	bytesTime := time.Since(ms)

	ms = time.Now()
	for i := 0; i < rows; i++ {
		protobuf := service.NewUserService(username, password)
		if protobuf == nil {
			panic("protobuf is nil")
		}

		protobuf.ToJSON(false)
		// fmt.Println(jsonData)
	}

	fmt.Println("Bin", bytesTime)
	fmt.Println("Json", time.Since(ms))
	fmt.Println()

	protobuf := service.NewUserService(username, password)
	if protobuf == nil {
		panic("protobuf is nil")
	}

	bytesData, _ := protobuf.ToBinary()
	fmt.Println(bytesData)

	jsonData, _ := protobuf.ToJSON(false)
	fmt.Println(jsonData)

	//

	if err := protobuf.WriteFile("data.bin", bytesData); err != nil {
		panic(err)
	}

	if err := protobuf.WriteFile("data.json", []byte(jsonData)); err != nil {
		panic(err)
	}

	//

	user, err := protobuf.ReadFile("data.bin", service.BINARY)
	if err != nil {
		panic(err)
	}

	fmt.Println(user)

	user, err = protobuf.ReadFile("data.json", service.JSON)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(user)

	user, err = protobuf.ReadFile("data.json", service.UNKNOW)
	if err != nil {
		panic(err)
	}

	fmt.Println(user)
}
