package main

import (
	"github.com/krmsaeed/barber-api/api"
	"github.com/krmsaeed/barber-api/config"
	"github.com/krmsaeed/barber-api/data/cache"
	"github.com/krmsaeed/barber-api/data/db"
	"github.com/krmsaeed/barber-api/data/db/migrations"
	"github.com/krmsaeed/barber-api/pkg/logging"
)

// @securityDefinitions.apikey AuthBearer
// @in header
// @name Authorization
func main() {

	cfg := config.GetConfig()
	logger := logging.NewLogger(cfg)

	err := cache.InitRedis(cfg)
	defer cache.CloseRedis()
	if err != nil {
		logger.Fatal(logging.Redis, logging.Startup, err.Error(), nil)
	}

	err = db.InitDb(cfg)
	defer db.CloseDb()
	if err != nil {
		logger.Fatal(logging.Postgres, logging.Startup, err.Error(), nil)
	}
	migrations.Up_1()

	api.InitServer(cfg)
}
