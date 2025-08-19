package dependencies

import (
	"context"
	"fmt"
	"github.com/cynx-io/cynx-core/src/logger"
	"github.com/cynx-io/hermes-user/internal/dependencies/config"
	"github.com/cynx-io/hermes-user/internal/model/entity"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type DatabaseClient struct {
	Db *gorm.DB
}

func NewDatabaseClient() (*DatabaseClient, error) {
	// Construct the DSN (Data Source Name)
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
		config.Config.Database.Username,
		config.Config.Database.Password,
		config.Config.Database.Host,
		config.Config.Database.Port,
		config.Config.Database.Database,
	)

	// Open a connection with GORM using the MySQL driver
	db, err := gorm.Open(mysql.Open(dataSourceName), &gorm.Config{
		Logger: gormlogger.Default.LogMode(gormlogger.Info),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// Check the connection
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get generic database object: %w", err)
	}
	if err = sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return &DatabaseClient{Db: db}, nil
}

func (client *DatabaseClient) Close() error {
	sqlDB, err := client.Db.DB()
	if err != nil {
		return fmt.Errorf("failed to get generic database object: %w", err)
	}
	return sqlDB.Close()
}

func (client *DatabaseClient) RunMigrations(ctx context.Context) error {

	logger.Info(ctx, "Running database migrations")
	err := client.Db.AutoMigrate(entity.TblUser{})
	if err != nil {
		return fmt.Errorf("failed to run migrations: %w", err)
	}

	logger.Info(ctx, "Migrations applied successfully")
	return nil
}
