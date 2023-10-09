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

func TestCreateNote(t *testing.T) {
	noteUsecase := mocks.NewNoteUsecase()
	handler := &notehandler.CreateNoteHandler{
		NoteUsecase: noteUsecase,
	}
	router := gin.Default()
	router.POST("/req", handler.CreateNote)

	t.Run("success", func(t *testing.T) {
		mockCreateNote := noteUsecase.
			On("CreateNote", mock.Anything, mock.Anything).
			Return(nil)

		v := url.Values{}
		v.Set("title", "Some Title")
		v.Set("content", "Lorem ipsum dolor sit amet, consectetur adipiscing elit.")
		body := bytes.NewBufferString(v.Encode())

		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/req", body)
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		mockCreateNote.Unset()
	})

	t.Run("failed any missing field", func(t *testing.T) {
		mockCreateNote := noteUsecase.
			On("CreateNote", mock.Anything, mock.Anything).
			Return(errors.New("some error"))

		v := url.Values{}
		v.Set("title", "Some Title")
		body := bytes.NewBufferString(v.Encode())

		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/req", body)
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)

		mockCreateNote.Unset()
	})

	t.Run("failed any datasource error", func(t *testing.T) {
		mockCreateNote := noteUsecase.
			On("CreateNote", mock.Anything, mock.Anything).
			Return(errors.New("some error"))

		v := url.Values{}
		v.Set("title", "Some Title")
		v.Set("content", "Lorem ipsum dolor sit amet, consectetur adipiscing elit.")
		body := bytes.NewBufferString(v.Encode())

		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/req", body)
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)

		mockCreateNote.Unset()
	})
}
