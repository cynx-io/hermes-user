package database

import (
	"context"
	"fmt"
	pb "hermes/api/proto/user"
	"hermes/internal/model/entity"

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

func (db *TblUser) InsertUser(ctx context.Context, user entity.TblUser) (int, error) {
	err := db.DB.WithContext(ctx).Create(&user).Error
	if err != nil {
		return 0, err
	}

	return user.Id, nil
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
	if req.SortBy != "" {
		sortOrder := "ASC"
		if req.SortOrder == "desc" {
			sortOrder = "DESC"
		}
		query = query.Order(fmt.Sprintf("%s %s", req.SortBy, sortOrder))
	} else {
		// Default sorting by id
		query = query.Order("id ASC")
	}

	// Apply pagination
	offset := (req.Page - 1) * req.Limit
	query = query.Limit(int(req.Limit)).Offset(int(offset))

	// Execute query
	if err := query.Find(&users).Error; err != nil {
		return nil, 0, err
	}

	return users, total, nil
}
