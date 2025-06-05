package entity

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	pb "hermes/api/proto/user"
)

type TblUser struct {
	EssentialEntity
	Username string `gorm:"size:255;not null;unique" json:"username"`
	Password string `gorm:"size:255;not null" json:"password"`
	Coin     int    `gorm:"default:0" json:"coin"`
}

func (u TblUser) ToUserResponse() *pb.UserData {
	return &pb.UserData{
		Id:          int32(u.Id),
		Username:    u.Username,
		Coin:        int32(u.Coin),
		CreatedDate: timestamppb.New(u.CreatedDate),
		UpdatedDate: timestamppb.New(u.UpdatedDate),
	}
}
