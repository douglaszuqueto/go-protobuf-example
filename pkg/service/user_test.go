package service_test

import (
	"fmt"
	"go-protobuf/pkg/service"
	"testing"
)

func TestUserService(t *testing.T) {
	t.Parallel()
}

var (
	username = "admin"
	password = "admin"
)

func BenchmarkBinary(b *testing.B) {
	loop(b, func() {
		toBinary()
	})
}

func BenchmarkJSON(b *testing.B) {
	loop(b, func() {
		toJSON()
	})
}

func BenchmarkWriteBinary(b *testing.B) {
	protobuf := service.NewUserService(username, password)
	if protobuf == nil {
		panic("protobuf is nil")
	}

	data, _ := protobuf.ToBinary()

	loop(b, func() {
		protobuf.WriteFile("user.bin", data)
	})
}

func BenchmarkWriteJSON(b *testing.B) {
	protobuf := service.NewUserService(username, password)
	if protobuf == nil {
		panic("protobuf is nil")
	}

	dataJson, _ := protobuf.ToJSON(true)

	loop(b, func() {
		protobuf.WriteFile("user.json", []byte(dataJson))
	})
}

func loop(b *testing.B, cb func()) {
	// for k := 0.; k <= 4; k++ {
	// n := int(math.Pow(10, k))
	b.Run(fmt.Sprintf("%d", 1), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			cb()
		}
	})
	// }
}

func toBinary() {
	protobuf := service.NewUserService(username, password)
	if protobuf == nil {
		panic("protobuf is nil")
	}

	protobuf.ToBinary()
}

func toJSON() {
	protobuf := service.NewUserService(username, password)
	if protobuf == nil {
		panic("protobuf is nil")
	}

	protobuf.ToJSON(false)
}
