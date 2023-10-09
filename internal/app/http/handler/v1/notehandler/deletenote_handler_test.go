package notehandler_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"altechomega.lukma.dev/notes/internal/app/http/handler/v1/notehandler"
	"altechomega.lukma.dev/notes/internal/domain/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestDeleteNote(t *testing.T) {
	noteUsecase := mocks.NewNoteUsecase()
	handler := &notehandler.DeleteNoteHandler{
		NoteUsecase: noteUsecase,
	}
	router := gin.Default()
	router.DELETE("/req/:id", handler.DeleteNote)

	t.Run("success", func(t *testing.T) {
		mockDeleteNote := noteUsecase.
			On("DeleteNote", mock.Anything, mock.Anything).
			Return(nil)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodDelete, "/req/6ba7b811-9dad-11d1-80b4-00c04fd430c8", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		mockDeleteNote.Unset()
	})

	t.Run("failed", func(t *testing.T) {
		mockDeleteNote := noteUsecase.
			On("DeleteNote", mock.Anything, mock.Anything).
			Return(errors.New("some error"))

		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodDelete, "/req/6ba7b811-9dad-11d1-80b4-00c04fd430c8", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)

		mockDeleteNote.Unset()
	})
}
