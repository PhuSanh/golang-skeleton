package repository

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"golang-skeleton/config"
	"log"
	"sync"
	"time"
)

type repositoryImpl struct {
	db *gorm.DB
}

type IRepository interface {
	IUserRepo
}

var (
	once  sync.Once
	mutex sync.RWMutex
	db    *gorm.DB
)

func InitGORM(cfg *config.Config) {
	once.Do(func() {
		mutex.Lock()
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", cfg.Database.User, cfg.Database.Pass, cfg.Database.Host, cfg.Database.Port, cfg.Database.Name)
		var err error
		db, err = gorm.Open("mysql", dsn)
		if err != nil {
			log.Printf("[Error] Mysql [%s][%s:%s]: %s", cfg.Database.User, cfg.Database.Host, cfg.Database.Port, err)
			time.Sleep(1 * time.Second)
			mutex.Unlock()
			InitGORM(cfg)
		} else {
			db.DB().SetMaxIdleConns(10)
			db.DB().SetMaxOpenConns(100)

			err = db.DB().Ping()
			if err != nil {
				log.Printf("[Error] Mysql [%s][%s:%s]: %s\n", cfg.Database.User, cfg.Database.Host, cfg.Database.Port, err)
				time.Sleep(1 * time.Second)
				mutex.Unlock()
				InitGORM(cfg)
				return
			}

			log.Printf("[Success] Mysql [%s][%s:%s] connected\n", cfg.Database.User, cfg.Database.Host, cfg.Database.Port)
			mutex.Unlock()
		}
	})
}

func NewRepositoryImpl() IRepository {
	return &repositoryImpl{
		db: db,
	}
}
