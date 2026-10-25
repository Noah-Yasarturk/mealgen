package main

import (
	"mealgen/internal/db"
	"mealgen/internal/handler"

	"go.uber.org/zap"
)

func main() {
	// Start logger
	zapLogger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	defer zapLogger.Sync()
	logger := zapLogger.Sugar()

	// Connect to db
	dbConn, err := db.OpenDb()
	if err != nil {
		logger.Fatalf("connecting to db: %v", err)
	}
	defer dbConn.Close()
	dbBaseDao := db.NewDbBaseDao(logger, dbConn)

	// Create handler
	h := handler.NewHandler(logger, dbBaseDao)

	// Start server
	h.ServeRoutes()

}
