package notehandler

import (
	"net/http"

	"altechomega.lukma.dev/notes/internal/domain"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type DeleteNoteHandler struct {
	NoteUsecase domain.NoteUsecase
}

func (handler *DeleteNoteHandler) DeleteNote(ctx *gin.Context) {
	err := handler.NoteUsecase.DeleteNote(ctx, uuid.MustParse(ctx.Param("id")))
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
