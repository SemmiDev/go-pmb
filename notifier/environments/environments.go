package environments

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

var (
	RabbitMQConnection      = ""
	RabbitMQQueue           = ""
	EmailSmtpHost           = ""
	EmailSmtpPort           = 0
	EmailSenderName         = ""
	EmailAuth               = ""
	EmailPassword           = ""
	EmailAddressHeaderEmail = ""
	EmailAddressHeaderName  = ""
)

func New() {
	var err error

	//viper.SetConfigName(fmt.Sprintf("appsettings.%s", os.Getenv("ENVIRONMENT")))
	viper.SetConfigName(fmt.Sprintf("appsettings.Dev"))
	viper.SetConfigType("json")
	viper.AddConfigPath(".")

	err = viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	RabbitMQConnection = viper.GetString("rabbitMQ.connection")
	RabbitMQQueue = viper.GetString("rabbitMQ.queue")

	EmailSmtpHost = viper.GetString("email.smtp.host")
	EmailSmtpPort = viper.GetInt("email.smtp.port")
	EmailSenderName = viper.GetString("email.senderName")
	EmailAuth = viper.GetString("email.authEmail")
	EmailPassword = viper.GetString("email.authPassword")
	EmailAddressHeaderEmail = viper.GetString("email.addressHeaderEmail")
	EmailAddressHeaderName = viper.GetString("email.addressHeaderName")
}
