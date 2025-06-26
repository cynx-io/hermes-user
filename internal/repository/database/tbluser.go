package database

import (
	"context"
	"fmt"
	pb "github.com/cynxees/hermes-user/api/proto/gen/hermes"
	"github.com/cynxees/hermes-user/internal/model/entity"
	"strings"

	"gorm.io/gorm"
)

type TblUser struct {
	DB *gorm.DB
}

func NewTblUser(DB *gorm.DB) *TblUser {
	return &TblUser{
		DB: DB,
	}
}

func (db *TblUser) InsertUser(ctx context.Context, user *entity.TblUser) (int, error) {
	err := db.DB.WithContext(ctx).Create(&user).Error
	if err != nil {
		return 0, err
	}

	return user.Id, nil
}

func (db *TblUser) UpdateUserByUserId(ctx context.Context, userId int, user *entity.TblUser) error {
	err := db.DB.WithContext(ctx).Where("id = ?", userId).Updates(user).Error
	if err != nil {
		return err
	}

	return nil
}

func (db *TblUser) CheckUserExists(ctx context.Context, key string, value string) (bool, error) {
	var count int64
	err := db.DB.WithContext(ctx).Model(&entity.TblUser{}).Where(key+" = ?", value).Count(&count).Error
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (db *TblUser) GetUser(ctx context.Context, key string, value string) (*entity.TblUser, error) {
	var user entity.TblUser
	err := db.DB.WithContext(ctx).Where(key+" = ?", value).First(&user).Error
	if err != nil {
		return &user, err
	}

	return &user, nil
}

func (db *TblUser) PaginateUser(ctx context.Context, req *pb.PaginateRequest) ([]entity.TblUser, int64, error) {
	var users []entity.TblUser
	var total int64

	// Get total count
	if err := db.DB.Model(&entity.TblUser{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Build query
	query := db.DB.Model(&entity.TblUser{})

	// Apply sorting
	if req.SortBy != nil {
		sortOrder := "ASC"
		sortOrder = "ASC"
		if req.SortOrder != nil {
			sortOrder = strings.ToUpper(*req.SortOrder)
			if sortOrder != "ASC" && sortOrder != "DESC" {
				sortOrder = "ASC"
			}
		}
		query = query.Order(fmt.Sprintf("%s %s", *req.SortBy, sortOrder))
	} else {
		// Default sorting by id
		query = query.Order("id ASC")
	}

	// Apply pagination
	offset := 0
	if req.Offset != nil {
		offset = int(*req.Offset)
	}

	limit := int(10)
	if req.Limit != nil {
		limit = int(*req.Limit)
	}
	query = query.Limit(limit).Offset(offset)

	// Execute query
	if err := query.Find(&users).Error; err != nil {
		return nil, 0, err
	}

	return users, total, nil
}

func (db *TblUser) CountIp(ctx context.Context, ipAddress string) (count int64, err error) {
	err = db.DB.WithContext(ctx).Model(&entity.TblUser{}).Where("ip_address = ?", ipAddress).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (db *TblUser) GetUsersByIp(ctx context.Context, ipAddress string) ([]*entity.TblUser, error) {
	var users []*entity.TblUser
	err := db.DB.WithContext(ctx).Where("ip_address = ?", ipAddress).Find(&users).Error
	if err != nil {
		return nil, err
	}

	if len(users) == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return users, nil
}
