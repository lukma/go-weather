package notehandler

import (
	"net/http"

	"altechomega.lukma.dev/notes/internal/domain"
	"github.com/gin-gonic/gin"
)

type createNoteForm struct {
	Title   string `form:"title" binding:"required"`
	Content string `form:"content" binding:"required"`
}

type CreateNoteHandler struct {
	NoteUsecase domain.NoteUsecase
}

func (handler *CreateNoteHandler) CreateNote(ctx *gin.Context) {
	var request createNoteForm

	err := ctx.ShouldBind(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, domain.ErrorResponse{
			Error: domain.ErrorDesc{
				Code:    domain.ErrMissingField.Code,
				Message: err.Error(),
			},
		})
		return
	}

	err = handler.NoteUsecase.CreateNote(ctx, domain.Note{
		Title:   request.Title,
		Content: request.Content,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, domain.ErrorResponse{
			Error: domain.ErrorDesc{
				Code:    domain.ErrServerFailure.Code,
				Message: err.Error(),
			},
		})
		return
	}

	ctx.JSON(http.StatusOK, nil)
}
