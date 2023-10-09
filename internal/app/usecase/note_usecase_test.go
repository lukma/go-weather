package usecase_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"altechomega.lukma.dev/notes/internal/app/usecase"
	"altechomega.lukma.dev/notes/internal/domain"
	"altechomega.lukma.dev/notes/internal/domain/mocks"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetNotes(t *testing.T) {
	noteRepository := mocks.NewNoteRepository()
	usecase := usecase.NewNoteUsecase(time.Second, noteRepository)

	t.Run("success", func(t *testing.T) {
		createdAt, _ := time.Parse("2006-01-02T15:04:05-0700", "2023-10-09T12:00:00.000124Z")
		updatedAt := createdAt.Add(time.Hour * time.Duration(2))
		expected := []domain.Note{
			{
				ID:        uuid.MustParse("6ba7b811-9dad-11d1-80b4-00c04fd430c8"),
				Title:     "Some Title",
				Content:   "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
				CreatedAt: createdAt,
				UpdatedAt: &updatedAt,
			},
		}

		mockGetNotes := noteRepository.
			On("GetNotes", mock.Anything).
			Return(expected, nil)

		actual, err := usecase.GetNotes(context.TODO())

		assert.Equal(t, expected, actual)
		assert.NoError(t, err)

		mockGetNotes.Unset()
	})

	t.Run("failed", func(t *testing.T) {
		mockGetNotes := noteRepository.
			On("GetNotes", mock.Anything).
			Return([]domain.Note(nil), errors.New("some error"))

		actual, err := usecase.GetNotes(context.TODO())

		assert.Empty(t, actual)
		assert.Error(t, err)

		mockGetNotes.Unset()
	})
}

func TestCreateNote(t *testing.T) {
	noteRepository := mocks.NewNoteRepository()
	usecase := usecase.NewNoteUsecase(time.Second, noteRepository)

	t.Run("success", func(t *testing.T) {
		mockCreateNote := noteRepository.
			On("CreateNote", mock.Anything, mock.Anything).
			Return(nil)

		err := usecase.CreateNote(context.TODO(), domain.Note{})

		assert.NoError(t, err)

		mockCreateNote.Unset()
	})

	t.Run("failed", func(t *testing.T) {
		mockCreateNote := noteRepository.
			On("CreateNote", mock.Anything, mock.Anything).
			Return(errors.New("some error"))

		err := usecase.CreateNote(context.TODO(), domain.Note{})

		assert.Error(t, err)

		mockCreateNote.Unset()
	})
}

func TestUpdateNote(t *testing.T) {
	noteRepository := mocks.NewNoteRepository()
	usecase := usecase.NewNoteUsecase(time.Second, noteRepository)

	t.Run("success", func(t *testing.T) {
		mockUpdateNote := noteRepository.
			On("UpdateNote", mock.Anything, mock.Anything).
			Return(nil)

		err := usecase.UpdateNote(context.TODO(), domain.Note{})

		assert.NoError(t, err)

		mockUpdateNote.Unset()
	})

	t.Run("failed", func(t *testing.T) {
		mockUpdateNote := noteRepository.
			On("UpdateNote", mock.Anything, mock.Anything).
			Return(errors.New("some error"))

		err := usecase.UpdateNote(context.TODO(), domain.Note{})

		assert.Error(t, err)

		mockUpdateNote.Unset()
	})
}

func TestDeleteNote(t *testing.T) {
	noteRepository := mocks.NewNoteRepository()
	usecase := usecase.NewNoteUsecase(time.Second, noteRepository)

	t.Run("success", func(t *testing.T) {
		mockDeleteNote := noteRepository.
			On("DeleteNote", mock.Anything, mock.Anything).
			Return(nil)

		err := usecase.DeleteNote(context.TODO(), uuid.MustParse("6ba7b811-9dad-11d1-80b4-00c04fd430c8"))

		assert.NoError(t, err)

		mockDeleteNote.Unset()
	})

	t.Run("failed", func(t *testing.T) {
		mockDeleteNote := noteRepository.
			On("DeleteNote", mock.Anything, mock.Anything).
			Return(errors.New("some error"))

		err := usecase.DeleteNote(context.TODO(), uuid.MustParse("6ba7b811-9dad-11d1-80b4-00c04fd430c8"))

		assert.Error(t, err)

		mockDeleteNote.Unset()
	})
}
