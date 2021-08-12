package environments

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

var (
	AppPort                = ""
	MongoConnectionString  = ""
	MongoPoolMin           = 0
	MongoPoolMax           = 0
	MongoMaxIdleTimeSecond = 0

	RedisConnectionString = ""
	RedisPassword         = ""

	RabbitMQConnection = ""
	RabbitMQQueue      = ""

	RegistrationDatabase   = ""
	RegistrationCollection = ""

	AccessSecret  = ""
	RefreshSecret = ""

	MidtransServerKey = ""
	MidtransClientKey = ""
)

func New() {
	var err error

	//viper.SetConfigName(fmt.Sprintf("appsettings.%s", os.Getenv("ENVIRONMENT")))
	viper.SetConfigName(fmt.Sprintf("appsettings.%s", "Dev"))
	viper.SetConfigType("json")
	viper.AddConfigPath(".")

	err = viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	AppPort = viper.GetString("app.port")

	MongoConnectionString = viper.GetString("mongo.connection")
	MongoPoolMin = viper.GetInt("mongo.poolMin")
	MongoPoolMax = viper.GetInt("mongo.poolMin")
	MongoMaxIdleTimeSecond = viper.GetInt("mongo.maxIdleTimeSecond")

	RedisConnectionString = viper.GetString("redis.connection")
	RedisPassword = viper.GetString("redis.password")

	RabbitMQConnection = viper.GetString("rabbitMQ.connection")
	RabbitMQQueue = viper.GetString("rabbitMQ.queue")

	RegistrationDatabase = viper.GetString("registration.database")
	RegistrationCollection = viper.GetString("registration.collection")

	AccessSecret = viper.GetString("security.accessSecret")
	RefreshSecret = viper.GetString("security.refreshSecret")

	MidtransServerKey = viper.GetString("midtrans.serverKey")
	MidtransClientKey = viper.GetString("midtrans.clientKey")
}
