package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/extra/redisotel"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"kratos-demo/app/demo/internal/conf"
	"kratos-demo/app/demo/internal/entity"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewDB, NewData, NewRedis, NewUserRepo)

// Data .
type Data struct {
	db  *gorm.DB
	rdb *redis.Client
	log *log.Helper
}

func NewDB(conf *conf.Data, logger log.Logger) *gorm.DB {
	log := log.NewHelper(log.With(logger, "module", "demo-service/data/gorm"))

	db, err := gorm.Open(mysql.Open(conf.Database.Source), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	if err := db.AutoMigrate(&entity.User{}); err != nil {
		log.Fatal(err)
	}

	return db
}
func NewRedis(conf *conf.Data, logger log.Logger) *redis.Client {

	rdb := redis.NewClient(&redis.Options{
		Addr:         conf.Redis.Addr,
		Password:     conf.Redis.Password,
		DB:           int(conf.Redis.Db),
		DialTimeout:  conf.Redis.DialTimeout.AsDuration(),
		WriteTimeout: conf.Redis.WriteTimeout.AsDuration(),
		ReadTimeout:  conf.Redis.ReadTimeout.AsDuration(),
	})
	rdb.AddHook(redisotel.TracingHook{})
	return rdb
}

// NewData .
func NewData(db *gorm.DB, rdb *redis.Client, logger log.Logger) (*Data, func(), error) {
	log := log.NewHelper(log.With(logger, "module", "demo-service/data"))

	d := &Data{
		db:  db,
		rdb: rdb,
		log: log,
	}
	return d, func() {
		log.Info("message", "closing the data resources")
		//if err := d.db.Close(); err != nil {
		//	log.Error(err)
		//}
		if err := d.rdb.Close(); err != nil {
			log.Error(err)
		}
	}, nil
}
