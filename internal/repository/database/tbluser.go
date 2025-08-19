package database

import (
	"context"
	"github.com/cynx-io/cynx-core/src/helper"
	"github.com/cynx-io/hermes-user/internal/model/entity"
	"gorm.io/gorm"
	"time"
)

type UserRepo struct {
	DB *gorm.DB
}

func NewUserRepo(DB *gorm.DB) *UserRepo {
	return &UserRepo{
		DB: DB,
	}
}

func (r *UserRepo) UpsertUser(ctx context.Context, user *entity.TblUser) (*entity.TblUser, error) {
	var existingUser entity.TblUser

	// Try to find existing user by auth0_id
	err := r.DB.WithContext(ctx).Where("auth0_id = ?", user.Auth0Id).First(&existingUser).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	if err == gorm.ErrRecordNotFound {
		// Create new user
		user.CreatedDate = time.Now()
		user.UpdatedDate = time.Now()
		err = r.DB.WithContext(ctx).Create(user).Error
		if err != nil {
			return nil, err
		}
		return user, nil
	}

	// Update existing user
	existingUser.Auth0Id = user.Auth0Id
	existingUser.Email = user.Email
	existingUser.UpdatedDate = time.Now()

	// Use helper to update optional fields only if provided
	helper.SetIfNotNil(&existingUser.Name, &user.Name)
	if user.SubscriptionTier != "" {
		existingUser.SubscriptionTier = user.SubscriptionTier
	}
	helper.SetIfNotNil(&existingUser.LastLoginAt, &user.LastLoginAt)
	if user.IsActive != existingUser.IsActive {
		existingUser.IsActive = user.IsActive
	}

	err = r.DB.WithContext(ctx).Save(&existingUser).Error
	if err != nil {
		return nil, err
	}

	return &existingUser, nil
}
