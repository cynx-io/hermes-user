package entity

import (
	pb "github.com/cynxees/hermes-user/api/proto/gen/hermes"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type TblUser struct {
	EssentialEntity
	Username string `gorm:"size:255;not null;unique" json:"username"`
	Password string `gorm:"size:255;not null" json:"password"`
	Coin     int    `gorm:"default:0" json:"coin"`
}

func (u TblUser) ToUserResponse() *pb.User {
	return &pb.User{
		Id:          int32(u.Id),
		Username:    u.Username,
		Coin:        int32(u.Coin),
		CreatedDate: timestamppb.New(u.CreatedDate),
		UpdatedDate: timestamppb.New(u.UpdatedDate),
	}
}
