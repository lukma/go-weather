package router

import (
	"altechomega.lukma.dev/notes/internal/app/http/handler/v1/notehandler"
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
	getNotesHandler := &notehandler.GetNotesHandler{
		NoteUsecase: r.noteUsecase,
	}
	createNoteHandler := &notehandler.CreateNoteHandler{
		NoteUsecase: r.noteUsecase,
	}
	updateNoteHandler := &notehandler.UpdateNoteHandler{
		NoteUsecase: r.noteUsecase,
	}
	deleteNoteHandler := &notehandler.DeleteNoteHandler{
		NoteUsecase: r.noteUsecase,
	}

	router := group.Group("/notes")
	{
		router.GET("", getNotesHandler.GetNotes)
		router.POST("", createNoteHandler.CreateNote)
		router.PATCH("/:id", updateNoteHandler.UpdateNote)
		router.DELETE("/:id", deleteNoteHandler.DeleteNote)
	}
}
