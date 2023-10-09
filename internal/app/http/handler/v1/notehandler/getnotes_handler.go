package notehandler

import (
	"net/http"

	"altechomega.lukma.dev/notes/internal/domain"
	"github.com/gin-gonic/gin"
)

type GetNotesHandler struct {
	NoteUsecase domain.NoteUsecase
}

func (handler *GetNotesHandler) GetNotes(ctx *gin.Context) {
	notes, err := handler.NoteUsecase.GetNotes(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, domain.ErrorResponse{
			Error: domain.ErrorDesc{
				Code:    domain.ErrServerFailure.Code,
				Message: err.Error(),
			},
		})
		return
	}

	ctx.JSON(http.StatusOK, domain.SuccessResult[[]domain.Note]{Data: notes})
}
