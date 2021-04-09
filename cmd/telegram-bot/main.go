package main

import (
	"en_train/internal/config"
	"en_train/internal/repository"
	"en_train/internal/telegram"
	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
}

func main() {
	appConfig, err := config.GetAppConfig()
	if err != nil || appConfig == nil {
		logrus.Fatalf("error initializing configs: %s", err)
	}

	apiBot, err := telegram.InitTGBot(appConfig.Telegram.TelegramBotToken)
	if err != nil {
		logrus.Fatalf("init bot failure, error = %s\n",err)
	}

	db, err := repository.NewPostgresDB(appConfig.DB)
	if err != nil || appConfig == nil{
		logrus.Fatalf("failed to initialize db: %s", err)
	}

	repos := repository.NewRepository(db)

	bot := telegram.NewBot(apiBot, repos)

	err = bot.StartTGBot()
	if err != nil {
		logrus.Fatalf("start bot failure, error = %s\n",err)
	}
}