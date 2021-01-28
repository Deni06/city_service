package main

import (
	"github.com/afex/hystrix-go/hystrix"
	"github.com/go-redis/redis"
	"github.com/micro/cli"
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	tracing "github.com/micro/go-plugins/wrapper/trace/opentracing"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	jaegerlog "github.com/uber/jaeger-client-go/log"
	"github.com/uber/jaeger-lib/metrics"
	"gitlab.visionet.co.id/pokota/xanadu/CityService/constant"
	"gitlab.visionet.co.id/pokota/xanadu/CityService/driver"
	"gitlab.visionet.co.id/pokota/xanadu/CityService/handler"
	"gitlab.visionet.co.id/pokota/xanadu/CityService/helper"
	cb "gitlab.visionet.co.id/pokota/xanadu/CityService/middleware"
	siteservice "gitlab.visionet.co.id/pokota/xanadu/CityService/proto/cityservice"
	"gitlab.visionet.co.id/pokota/xanadu/CityService/util"
	"os"
	"strconv"
)

func main() {
	cfg := jaegercfg.Configuration{
		ServiceName: constant.SERVICE_NAME,
		Sampler:     &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter:    &jaegercfg.ReporterConfig{
			LogSpans: true,
		},
	}
	jLogger := jaegerlog.StdLogger
	jMetricsFactory := metrics.NullFactory
	hystrix.ConfigureCommand(constant.SERVICE_NAME, hystrix.CommandConfig{
		Timeout:               1,
	})

	// Initialize tracer with a logger and a metrics factory
	tracer, closer, errJaeger := cfg.NewTracer(
		jaegercfg.Logger(jLogger),
		jaegercfg.Metrics(jMetricsFactory),
	)
	opentracing.SetGlobalTracer(tracer)
	if errJaeger!=nil{
		log.Fatal(errJaeger)
	}
	defer closer.Close()

	// New Service
	service := micro.NewService(
		micro.Name(constant.SERVICE_NAME),
		micro.Version(constant.SERVICE_VERSION),
		micro.WrapClient(tracing.NewClientWrapper(tracer),cb.NewClientWrapper() ),
		micro.WrapHandler(tracing.NewHandlerWrapper(tracer)),
		micro.WrapSubscriber(tracing.NewSubscriberWrapper(tracer)),
		micro.Flags(
			cli.BoolFlag{
				Name:  "db_usage",
				Usage: "Config for flag usage",
			},
			cli.StringFlag{
				Name:  "db_host",
				Usage: "Config for db host",
			},
			cli.StringFlag{
				Name:  "db_host_param",
				Usage: "Config for db port and db instance, db port ex: :8080 or db instance ex: /sql2014 for sql server" +
					" and db port ex : 8080 for other than sql server",
			},
			cli.StringFlag{
				Name:  "db_username",
				Usage: "Config for username",
			},
			cli.StringFlag{
				Name:  "db_password",
				Usage: "Config for db password",
			},
			cli.StringFlag{
				Name:  "db_name",
				Usage: "Config for db name",
			},
			cli.StringFlag{
				Name:  "db_ssl_mode",
				Usage: "Config for db ssl mode",
			},
			cli.StringFlag{
				Name:  "db_dialect",
				Usage: "Config for db dialect",
			},

		),
	)

	// Initialise service
	var dialect string
	service.Init(
		micro.Action(func(c *cli.Context) {
			dialect = c.String("db_dialect")
			driver.SetParamGorm(driver.Parameter{
				UseCli:c.Bool("db_usage"),
				SslMode:c.String("db_ssl_mode"),
				DbName:c.String("db_name"),
				Password:c.String("db_password"),
				User:c.String("db_username"),
				HostParam:c.String("db_host_param"),
				Host:c.String("db_host"),
				Dialect:dialect,
			})
		}),
	)

	db,errInitDB := driver.InitGorm()
	if errInitDB!=nil{
		log.Fatal(errInitDB.Error())
	}else{
		if dialect==""{
			dialect = constant.POSTGRESQL_DIALECT
		}
		err := helper.InitDriver(dialect)
		if err!=nil{
			log.Fatal(err)
		}
		defer db.Close()
	}

	// Register Handler
	siteservice.RegisterCityHandler(service.Server(), new(handler.CityService))

	redisHost := constant.REDIS_HOST
	if os.Getenv(constant.REDIS_HOST_ENV) != "" {
		redisHost = os.Getenv(constant.REDIS_HOST_ENV)
	}
	redisDB := constant.REDIS_DB
	if os.Getenv(constant.REDIS_DB_ENV) != "" {
		redisDB, _ = strconv.Atoi(os.Getenv(constant.REDIS_DB_ENV))
	}
	redisPassword := constant.REDIS_PASSWORD
	if os.Getenv(constant.REDIS_PASSWORD_ENV) != "" {
		redisHost = os.Getenv(constant.REDIS_PASSWORD_ENV)
	}

	redisClient := redis.NewClient(&redis.Options{
		Addr:     redisHost,
		Password: redisPassword, // no password set
		DB:       redisDB,       // use default DB
	})

	_, err := redisClient.Ping().Result()
	if err != nil {
		log.Fatal(err.Error())
	}
	util.SetRedisClient(redisClient)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
