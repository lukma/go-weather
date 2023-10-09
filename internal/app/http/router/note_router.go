package router

import (
	notehandlerV1 "altechomega.lukma.dev/notes/internal/app/http/handler/v1/notehandler"
	"altechomega.lukma.dev/notes/internal/domain"
	"github.com/gin-gonic/gin"
)

type noteRoute struct {
	noteUsecase domain.NoteUsecase
}

func NewNoteRouter(noteUsecase domain.NoteUsecase) *noteRoute {
	return &noteRoute{
		noteUsecase: noteUsecase,
	}
}

func (r *noteRoute) RegisterIn(group *gin.RouterGroup) {
	getNotesHandler := &notehandlerV1.GetNotesHandler{
		NoteUsecase: r.noteUsecase,
	}
	createNoteHandler := &notehandlerV1.CreateNoteHandler{
		NoteUsecase: r.noteUsecase,
	}
	updateNoteHandler := &notehandlerV1.UpdateNoteHandler{
		NoteUsecase: r.noteUsecase,
	}
	deleteNoteHandler := &notehandlerV1.DeleteNoteHandler{
		NoteUsecase: r.noteUsecase,
	}

	router := group.Group("/v1/notes")
	{
		router.GET("", getNotesHandler.GetNotes)
		router.POST("", createNoteHandler.CreateNote)
		router.PATCH("/:id", updateNoteHandler.UpdateNote)
		router.DELETE("/:id", deleteNoteHandler.DeleteNote)
	}
}
