package entity

import (
	"github.com/cynx-io/cynx-core/src/entity"
	pb "github.com/cynx-io/hermes-user/api/proto/gen/hermes"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type TblUser struct {
	entity.EssentialEntity
	Auth0Id          string     `gorm:"size:255;unique;not null" json:"auth0_id"`
	Email            string     `gorm:"size:255;not null" json:"email"`
	Name             *string    `gorm:"size:255" json:"name"`
	SubscriptionTier string     `gorm:"size:50;default:'free'" json:"subscription_tier"`
	LastLoginAt      *time.Time `json:"last_login_at"`
	IsActive         bool       `gorm:"default:true" json:"is_active"`
}

func (u TblUser) Response() *pb.User {
	resp := &pb.User{
		Id:               u.Id,
		Auth0Id:          u.Auth0Id,
		Email:            u.Email,
		SubscriptionTier: u.SubscriptionTier,
		CreatedAt:        timestamppb.New(u.EssentialEntity.CreatedDate),
		UpdatedAt:        timestamppb.New(u.EssentialEntity.UpdatedDate),
		IsActive:         u.IsActive,
	}

	if u.Name != nil {
		resp.Name = *u.Name
	}

	if u.LastLoginAt != nil {
		resp.LastLoginAt = timestamppb.New(*u.LastLoginAt)
	}

	return resp
}
