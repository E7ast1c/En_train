package main

import (
	"os"
	"os/signal"
	"syscall"

	"en_train/internal/config"
	"en_train/internal/handler"
	"en_train/internal/repository"
	"en_train/internal/server"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	appConfig, err := config.GetAppConfig()
	if err != nil || appConfig == nil {
		logrus.Fatalf("error initializing configs: %s", err)
	}

	db, err := repository.NewPostgresDB(appConfig.DB)
	if err != nil || appConfig == nil{
		logrus.Fatalf("failed to initialize db: %s", err)
	}

	repos := repository.NewRepository(db)
	handlers := handler.NewHandler(repos)
	srv := new(server.HttpServer)

	go func() {
		if err := srv.Run(appConfig.Api, handlers.InitRoutes()); err != nil {
			logrus.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	logrus.Print(repos)
	logrus.Print("TodoApp Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("App Shutting Down")

	// TODO graceful shutdown
	// if err := srv.Shutdown(context.Background()); err != nil {
	// 	logrus.Errorf("error occured on server shutting down: %s", err.Error())
	// }

	if err := db.Close(); err != nil {
		logrus.Errorf("error occured on db connection close: %s", err)
	}
}




