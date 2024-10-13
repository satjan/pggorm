package pggorm

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/dbresolver"
	"log"
	"os"
	"time"
)

func Init(p *Pgsql) (*gorm.DB, error) {
	if p.Dbname == "" {
		return nil, fmt.Errorf("database name is empty")
	}

	gormConfig := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold:             time.Millisecond * 100,
				LogLevel:                  logger.LogLevel(p.LogLevel),
				IgnoreRecordNotFoundError: true,
				Colorful:                  true,
			},
		),
		SkipDefaultTransaction: true,
	}

	pgsqlConfig := postgres.Config{
		DSN:                  p.Dsn(),
		PreferSimpleProtocol: false,
	}

	pgsqlConfigReplica := postgres.Config{
		DSN:                  p.DsnReplica(),
		PreferSimpleProtocol: false,
	}

	db, err := gorm.Open(postgres.New(pgsqlConfig), gormConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	if err := db.Use(dbresolver.Register(dbresolver.Config{
		Replicas: []gorm.Dialector{postgres.New(pgsqlConfigReplica)},
		Policy:   dbresolver.RandomPolicy{},
	})); err != nil {
		return nil, fmt.Errorf("failed to register replica: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get DB instance: %w", err)
	}

	sqlDB.SetMaxIdleConns(p.MaxIdleCons)
	sqlDB.SetMaxOpenConns(p.MaxOpenCons)
	sqlDB.SetConnMaxLifetime(time.Duration(p.MaxLifeTimeMinute) * time.Minute)

	return db, nil
}
