package main

import (
	"log/slog"

	"Mars/database"
	"Mars/database/v2/controller"
	"Mars/server"
	"Mars/shared/configure"
	"Mars/shared/utils/json"
	"Mars/shared/utils/json/json_se/sonic"
)

func main() {
	slog.Info("Starting application....")
	configure.NewConfig()
	json.New()
	sonic.PreTouch()
	database.Database()
	controller.RemoveAllBundlerBuilds()
	server.Server()
}
