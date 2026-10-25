package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"mealgen/internal/db"
)

type Handler struct {
	dbBaseDao *db.DbBaseDao
	logger    *zap.SugaredLogger
}

func NewHandler(l *zap.SugaredLogger, dbBaseDao *db.DbBaseDao) *Handler {
	return &Handler{
		dbBaseDao: dbBaseDao,
		logger:    l,
	}
}

func (h *Handler) createHealthCheckHandler(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		h.logger.Infof("server pinged")
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
}

func (h *Handler) createDbCheckHandler(r *gin.Engine) {
	r.GET("/ping-db", func(c *gin.Context) {
		h.logger.Info("db pinged")
		err := h.dbBaseDao.DbHealth()
		if err != nil {
			msg := fmt.Sprintf("db health check failure: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": msg,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message": "db up",
			})
		}

	})
}

func (h *Handler) ServeRoutes() {
	// Create a Gin router with default middleware (logger and recovery)
	r := gin.Default()
	h.createHealthCheckHandler(r)
	h.createDbCheckHandler(r)
	// Start server on port 8080 (default)
	r.Run()
}
