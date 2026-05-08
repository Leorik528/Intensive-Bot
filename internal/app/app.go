package app

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"intensive-bot/internal/bot"
	"intensive-bot/internal/config"
	"intensive-bot/internal/domain"
	"intensive-bot/internal/service"
)

type App struct {
	cfg              *config.Config
	bot              *bot.Bot
	intensiveService *service.IntensiveService
}

func New() (*App, error) {
	cfg, err := config.Load()
	if err != nil {
		return nil, err
	}
	is := service.NewIntensiveService()
	b, err := bot.New(cfg, is)
	if err != nil {
		return nil, err
	}
	return &App{cfg: cfg, bot: b, intensiveService: is}, nil
}

func (a *App) Run() error {
	go func() { _ = a.bot.Run(context.Background()) }()
	return a.runAdminHTTP()
}

func (a *App) runAdminHTTP() error {
	r := gin.Default()
	r.GET("/admin/intensives", func(c *gin.Context) { c.JSON(http.StatusOK, a.intensiveService.ListAll()) })
	r.POST("/admin/intensives", func(c *gin.Context) {
		var req domain.Intensive
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, a.intensiveService.Create(req))
	})
	r.PUT("/admin/intensives/:id", func(c *gin.Context) {
		var req domain.Intensive
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var id int64
		_, _ = fmt.Sscan(c.Param("id"), &id)
		updated, err := a.intensiveService.Update(id, req)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, updated)
	})
	r.POST("/admin/intensives/:id/toggle", func(c *gin.Context) {
		var id int64
		_, _ = fmt.Sscan(c.Param("id"), &id)
		isOpen := c.Query("open") == "true"
		if err := a.intensiveService.SetOpen(id, isOpen); err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.Status(http.StatusNoContent)
	})
	r.POST("/admin/intensives/:id/broadcast", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"status": "queued", "intensive_id": c.Param("id")}) })
	return r.Run(a.cfg.AdminHTTPAddr)
}
