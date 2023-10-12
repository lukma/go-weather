package router

import (
	"database/sql"
	"time"

	"altechomega.lukma.dev/notes/config"
	repositoryDB "altechomega.lukma.dev/notes/internal/app/data/db/repository"
	"altechomega.lukma.dev/notes/internal/app/usecase"
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
	timeout := time.Duration(r.config.ContextTimeout) * time.Second

	noteRepository := repositoryDB.NewNoteRepository(r.db)
	noteUsecase := usecase.NewNoteUsecase(timeout, noteRepository)
	noteRouter := NewNoteRouter(noteUsecase)

	gin := gin.Default()
	gin.Use(CORSMiddleware())

	publicRouter := gin.Group("")
	{
		noteRouter.RegisterIn(publicRouter)
	}

	gin.Run(r.config.ServerAddress)
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
