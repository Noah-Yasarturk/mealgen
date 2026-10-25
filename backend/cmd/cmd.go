package main

import (
	"context"
	"mealgen/internal/db"
	"mealgen/internal/handler"

	"go.uber.org/zap"
)

func main() {
	ctx := context.Background()

	// Start logger
	zapLogger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	defer zapLogger.Sync()
	logger := zapLogger.Sugar()

	// Connect to db
	dbConn, err := db.OpenDb(ctx)
	if err != nil {
		logger.Fatalf("connecting to db: %v", err)
	}
	defer dbConn.Close(ctx)
	dbBaseDao := db.NewDbBaseDao(logger, dbConn)

	// Create handler
	h := handler.NewHandler(logger, dbBaseDao)

	// Start server
	h.ServeRoutes(ctx)

}
