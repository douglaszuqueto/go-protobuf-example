package service

import (
	"go-protobuf/pb"
	"io/ioutil"
	"math/rand"
	"strconv"
	"time"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
)

const (
	// BINARY BINARY
	BINARY = iota
	// JSON JSON
	JSON
	// UNKNOW UNKNOW
	UNKNOW
)

// UserService UserService
type UserService struct {
	User *pb.User
}

// NewUserService NewUserService
func NewUserService(username, password string) *UserService {
	s := &UserService{}

	user := s.newUser(username, password)
	if user == nil {
		return nil
	}

	s.User = user

	return s
}

// NewUser NewUser
func (s *UserService) NewUser(username, password string) *pb.User {
	user := s.newUser(username, password)
	if user == nil {
		return nil
	}

	return user
}

func (s *UserService) newUser(username, password string) *pb.User {
	user := &pb.User{
		Id:       strconv.Itoa(rand.Intn(100)),
		Username: username,
		Password: password,
	}

	var err error

	user.CreatedAt, err = ptypes.TimestampProto(time.Now())
	if err != nil {
		return nil
	}

	user.UpdatedAt, err = ptypes.TimestampProto(time.Now())
	if err != nil {
		return nil
	}

	return user
}

// ToBinary ToBinary
func (s *UserService) ToBinary() ([]byte, error) {
	return proto.Marshal(s.User)
}

// ToJSON ToJSON
func (s *UserService) ToJSON(indent bool) (string, error) {
	marshaler := jsonpb.Marshaler{
		EnumsAsInts:  false,
		EmitDefaults: true,
		OrigName:     true,
	}

	if indent {
		marshaler.Indent = "  "
	}

	return marshaler.MarshalToString(s.User)
}

// WriteFile WriteFile
func (s *UserService) WriteFile(filename string, data []byte) error {
	path := "./tmp/" + filename

	err := ioutil.WriteFile(path, data, 0644)

	return err
}

// ReadFile ReadFile
func (s *UserService) ReadFile(filename string, format int) (*pb.User, error) {
	path := "./tmp/" + filename

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	user := &pb.User{}

	switch format {
	case BINARY:
		err = proto.Unmarshal(data, user)
		break
	case JSON:
		err = jsonpb.UnmarshalString(string(data), user)
		break
	default:
		panic("unknow format:")
	}

	return user, err
}
