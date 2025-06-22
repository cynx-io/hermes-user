package entity

import (
	"github.com/cynxees/cynx-core/src/entity"
	"github.com/cynxees/cynx-core/src/types/usertype"
	pb "github.com/cynxees/hermes-user/api/proto/gen/hermes"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type TblUser struct {
	entity.EssentialEntity
	Id        int               `gorm:"primaryKey;index:idx_user_type_id" json:"id"`
	Username  string            `gorm:"size:255;not null;unique;index" json:"username"`
	UserType  usertype.UserType `gorm:"not null;index:idx_user_type_id" json:"user_type"`
	IpAddress string            `gorm:"not null;index" json:"ip_address"`
	Password  string            `gorm:"size:255;not null" json:"password"`
	Coin      int               `gorm:"default:0" json:"coin"`
}

func (u TblUser) Response() *pb.User {
	return &pb.User{
		Id:             int32(u.Id),
		Username:       u.Username,
		UserType:       int32(u.UserType),
		UserTypeString: u.UserType.String(),
		Coin:           int32(u.Coin),
		CreatedDate:    timestamppb.New(u.CreatedDate),
		UpdatedDate:    timestamppb.New(u.UpdatedDate),
	}
}
