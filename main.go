package main

import (
	"github.com/codingXiang/configer"
	cx "github.com/codingXiang/cxgateway/delivery/http"
	"github.com/codingXiang/go-logger"
	"github.com/codingXiang/go-orm"
	"sample_api/model"
	"sample_api/module/user/delivery/http"
	"sample_api/module/user/repository"
	repository2 "sample_api/module/user/service"
)

func init () {
	var err error
	//初始化 configer，設定預設讀取環境變數
	config := configer.NewConfigerCore("yaml", "config", "./config", ".")
	config.SetAutomaticEnv("")
	//初始化 Gateway
	cx.Gateway = cx.NewApiGateway("config", config)

	//初始化 db 參數
	db := configer.NewConfigerCore("yaml", "database", "./config", ".")
	db.SetAutomaticEnv("")
	configer.Config.AddCore("db", db)
	//設定資料庫
	if orm.DatabaseORM, err = orm.NewOrm("database", configer.Config.GetCore("db")); err == nil {
		// 建立 Table Schema (Module)
		logger.Log.Debug("setup table schema")
		{
			//設定 使用者資料
			orm.DatabaseORM.CheckTable(true, &model.User{})
		}
	} else {
		logger.Log.Error(err.Error())
		panic(err.Error())
	}
}

func main() {
	// 建立 Repository
	logger.Log.Debug("Create Repository Instance")
	var (
		db = orm.DatabaseORM.GetInstance()
		userRepo = repository.NewUserRepository(db)
	)
	// 建立 Service
	logger.Log.Debug("Create Service Instance")
	var (
		userSvc = repository2.NewUserService(userRepo)
	)
	// 建立 Handler (Module)
	logger.Log.Debug("Create Http Handler")
	{
		http.NewUserHttpHandler(cx.Gateway, userSvc)
	}
	cx.Gateway.Run()
}