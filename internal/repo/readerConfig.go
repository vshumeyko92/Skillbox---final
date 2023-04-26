package repo

import (
	"Skillbox-diploma/internal/struct"
	"github.com/spf13/viper"
	"log"
)

func ConfigReader() (config _struct.Config) {
	if err := initConfig(); err != nil {
		log.Printf("ошибка иницилиализации: %s", err.Error())
	}
	config.Debug = viper.GetBool("debug.enabled")
	config.SmsSource = viper.GetString("sms.sourceDirectory") + "/" + viper.GetString("sms.sourceFileName")
	config.SmsTarget = viper.GetString("sms.targetDirectory") + "/" + viper.GetString("sms.targetFileName")
	config.MmsSource = viper.GetString("mms.scheme") + "://" +
		viper.GetString("mms.address") + ":" +
		viper.GetString("mms.port") + "/" +
		viper.GetString("mms.endpoint")
	config.VoiceSource = viper.GetString("voice.sourceDirectory") + "/" + viper.GetString("voice.sourceFileName")
	config.VoiceTarget = viper.GetString("voice.targetDirectory") + "/" + viper.GetString("voice.targetFileName")
	config.MailSource = viper.GetString("mail.sourceDirectory") + "/" + viper.GetString("mail.sourceFileName")
	config.MailTarget = viper.GetString("mail.targetDirectory") + "/" + viper.GetString("mail.targetFileName")
	config.BillingSource = viper.GetString("billing.sourceDirectory") + "/" + viper.GetString("billing.sourceFileName")
	config.SupportSource = viper.GetString("support.scheme") + "://" +
		viper.GetString("support.address") + ":" +
		viper.GetString("support.port") + "/" +
		viper.GetString("support.endpoint")
	config.AccendendSource = viper.GetString("accendent.scheme") + "://" +
		viper.GetString("accendent.address") + ":" +
		viper.GetString("accendent.port") + "/" +
		viper.GetString("accendent.endpoint")
	config.WebListernerAddress = viper.GetString("web.address") + ":" + viper.GetString("web.port")
	return
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
