package util

import (
	"gitlab.visionet.co.id/pokota/xanadu/CityService/constant"
	"log"
	"os"
)

func WriteLogMain(v ...interface{}){
	log.Println(v...)

	allServiceLogPath := constant.ALL_SERVICE_LOG_PATH_MAIN
	rc := GetRedisClient()
	if rc.Get(constant.ALL_SERVICE_LOG_PATH_MAIN_REDIS).Val()!=""{
		allServiceLogPath = rc.Get(constant.ALL_SERVICE_LOG_PATH_MAIN_REDIS).Val()
	}

	f, err := os.OpenFile(allServiceLogPath,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Printf("error opening file: %v", err)
	}
	defer f.Close()

	loggerAll := log.New(f, constant.PREFIX_LOG, log.LstdFlags)

	loggerAll.Println(v...)
}

func WriteLogApi(v ...interface{}){
	log.Println(v...)

	allServiceLogPath := constant.ALL_SERVICE_LOG_PATH_API
	rc := GetRedisClient()
	if rc.Get(constant.ALL_SERVICE_LOG_PATH_API_REDIS).Val()!=""{
		allServiceLogPath = rc.Get(constant.ALL_SERVICE_LOG_PATH_API_REDIS).Val()
	}

	f, err := os.OpenFile(allServiceLogPath,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Printf("error opening file: %v", err)
	}
	defer f.Close()

	loggerAll := log.New(f, constant.PREFIX_LOG, log.LstdFlags)

	loggerAll.Println(v...)
}
