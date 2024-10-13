package main

import (
	"github.com/satjan/pggorm"
	"log"
	"os"
)

func main() {
	DB, err := pggorm.Init(&pggorm.Pgsql{
		Host:              os.Getenv("PG_HOST"),
		Port:              os.Getenv("PG_PORT"),
		Dbname:            os.Getenv("PG_DBNAME"),
		Username:          os.Getenv("PG_USERNAME"),
		Password:          os.Getenv("PG_PASSWORD"),
		ReplicaHost:       os.Getenv("PG_REPLICA_HOST"),
		MaxIdleCons:       pggorm.Atoi(os.Getenv("PG_MAX_IDLE_CONS")),
		MaxOpenCons:       pggorm.Atoi(os.Getenv("PG_MAX_OPEN_CONS")),
		MaxLifeTimeMinute: pggorm.Atoi(os.Getenv("PG_MAX_LIFE_TIME_MINUTE")),
		LogLevel:          pggorm.Atoi(os.Getenv("PG_LOG_LEVEL")),
	})
	if err != nil {
		panic(err)
	}

	var cnt int64
	DB.Table("user").Count(&cnt)
	log.Println(cnt)
}
