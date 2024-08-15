package repository

import (
	"{appname}/pkg/log"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
	//rds    *redis.Client
	logger *log.Logger
}

func NewRepository(logger *log.Logger, db *gorm.DB) *Repository {
	return &Repository{
		db: db,
		//rds:    rds,
		logger: logger,
	}
}
func NewDb() *gorm.DB {
	// TODO: init db
	//db, err := gorm.Open(mysql.Open(conf.GetString("data.mysql.user")), &gorm.Config{})
	//if err != nil {
	//	panic(err)
	//}
	//return db
	return &gorm.DB{}
}
