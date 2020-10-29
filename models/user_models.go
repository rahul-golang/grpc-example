package models

import "interview/grpc/proto"

type UserKey struct {
	Key int64
}

func NewUserKey(uKey *proto.UserKey) UserKey {
	return UserKey{
		Key: uKey.Key,
	}
}

type UserKeys struct {
	Keys []UserKey
}

type User struct {
	Key  int64
	Name string
}

func NewUserKeys(resp *proto.UserKeys) UserKeys {
	keys := []UserKey{}
	for _, v := range resp.Users {
		keys = append(keys, NewUserKey(v))
	}
	return UserKeys{
		Keys: keys,
	}

}
func (u User) MapToProto() *proto.User {
	return &proto.User{
		Key:  u.Key,
		Name: u.Name,
	}
}
