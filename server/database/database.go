package database

import (
	"github.com/iamtonmoy0/go-react-blog-app/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Database() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("blog.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(new(models.Blog))
	return db

}
