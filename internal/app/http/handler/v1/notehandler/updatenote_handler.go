package notehandler

import (
	"net/http"

	"altechomega.lukma.dev/notes/internal/domain"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type updateNoteForm struct {
	Title   string `form:"title"`
	Content string `form:"content"`
}

type UpdateNoteHandler struct {
	NoteUsecase domain.NoteUsecase
}

func (handler *UpdateNoteHandler) UpdateNote(ctx *gin.Context) {
	var request updateNoteForm

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

	err = handler.NoteUsecase.UpdateNote(ctx, domain.Note{
		ID:      uuid.MustParse(ctx.Param("id")),
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
