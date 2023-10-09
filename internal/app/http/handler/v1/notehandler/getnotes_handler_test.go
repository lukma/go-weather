package notehandler_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"altechomega.lukma.dev/notes/internal/app/http/handler/v1/notehandler"
	"altechomega.lukma.dev/notes/internal/domain"
	"altechomega.lukma.dev/notes/internal/domain/mocks"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetNotes(t *testing.T) {
	noteUsecase := mocks.NewNoteUsecase()
	handler := &notehandler.GetNotesHandler{
		NoteUsecase: noteUsecase,
	}
	router := gin.Default()
	router.GET("/req", handler.GetNotes)

	t.Run("success", func(t *testing.T) {
		createdAt, _ := time.Parse("2006-01-02T15:04:05-0700", "2023-10-09T12:00:00.000124Z")
		updatedAt := createdAt.Add(time.Hour * time.Duration(2))
		expected := []domain.Note{
			{
				ID:        uuid.MustParse("6ba7b811-9dad-11d1-80b4-00c04fd430c8"),
				Title:     "Hello World",
				Content:   "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
				CreatedAt: createdAt,
				UpdatedAt: &updatedAt,
			},
		}

		mockGetNotes := noteUsecase.
			On("GetNotes", mock.Anything).
			Return(expected, nil)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/req", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		data, _ := json.Marshal(domain.SuccessResult[[]domain.Note]{Data: expected})
		assert.Equal(t, string(data), w.Body.String())

		mockGetNotes.Unset()
	})

	t.Run("failed", func(t *testing.T) {
		mockGetNotes := noteUsecase.
			On("GetNotes", mock.Anything).
			Return([]domain.Note(nil), errors.New("some error"))

		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/req", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)

		mockGetNotes.Unset()
	})
}
