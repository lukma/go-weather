package notehandler_test

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"altechomega.lukma.dev/notes/internal/app/http/handler/v1/notehandler"
	"altechomega.lukma.dev/notes/internal/domain/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUpdateNote(t *testing.T) {
	noteUsecase := mocks.NewNoteUsecase()
	handler := &notehandler.UpdateNoteHandler{
		NoteUsecase: noteUsecase,
	}
	router := gin.Default()
	router.PATCH("/req/:id", handler.UpdateNote)

	t.Run("success", func(t *testing.T) {
		mockUpdateNote := noteUsecase.
			On("UpdateNote", mock.Anything, mock.Anything).
			Return(nil)

		v := url.Values{}
		v.Set("title", "Some Title")
		body := bytes.NewBufferString(v.Encode())

		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPatch, "/req/6ba7b811-9dad-11d1-80b4-00c04fd430c8", body)
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		mockUpdateNote.Unset()
	})

	t.Run("failed", func(t *testing.T) {
		mockUpdateNote := noteUsecase.
			On("UpdateNote", mock.Anything, mock.Anything).
			Return(errors.New("some error"))

		v := url.Values{}
		v.Set("title", "Some Title")
		v.Set("content", "Lorem ipsum dolor sit amet, consectetur adipiscing elit.")
		body := bytes.NewBufferString(v.Encode())

		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPatch, "/req/6ba7b811-9dad-11d1-80b4-00c04fd430c8", body)
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)

		mockUpdateNote.Unset()
	})
}
