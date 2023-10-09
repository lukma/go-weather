package main

import (
	"altechomega.lukma.dev/notes/config"
	"altechomega.lukma.dev/notes/internal/app/http/router"
	"altechomega.lukma.dev/notes/internal/pkg/sqlutil"
)

func main() {
	c := config.NewConfig()

	db := sqlutil.ConnectDB(c)
	defer db.Close()

	r := router.NewRouter(c, db)
	r.RunHttpServer()
}
