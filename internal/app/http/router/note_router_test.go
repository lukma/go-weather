package router_test

import (
	"net/http"
	"testing"

	"altechomega.lukma.dev/notes/internal/app/http/router"
	"altechomega.lukma.dev/notes/internal/domain/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestNoteRouter(t *testing.T) {
	noteUsecase := mocks.NewNoteUsecase()

	r := gin.Default()
	router.NewNoteRouter(noteUsecase).RegisterIn(r.Group(""))

	paths := map[string]bool{}
	for _, v := range r.Routes() {
		paths[v.Path+"|"+v.Method] = true
	}

	assert.True(t, paths["/v1/notes|"+http.MethodGet])
	assert.True(t, paths["/v1/notes|"+http.MethodPost])
	assert.True(t, paths["/v1/notes/:id|"+http.MethodPatch])
	assert.True(t, paths["/v1/notes/:id|"+http.MethodDelete])
}
