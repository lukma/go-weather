package router

import (
	"database/sql"

	"altechomega.lukma.dev/notes/config"
	"github.com/gin-gonic/gin"
)

type router struct {
	config *config.Config
	db     *sql.DB
}

func NewRouter(config *config.Config, db *sql.DB) *router {
	return &router{
		config: config,
		db:     db,
	}
}

func (r *router) RunHttpServer() {
	gin := gin.Default()

	gin.Run(r.config.ServerAddress)
}
