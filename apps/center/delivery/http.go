package delivery

import (
	"mhlv/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	u domain.CenterUsecase
}

func New(rg *gin.Engine, u domain.CenterUsecase) (*Handler, error) {
	handler := Handler{
		u: u,
	}

	rg.POST("/coordinate-agent", handler.coordinate)
	return &handler, nil
}

// post coordinate agent
// @Summary coordinate agent
// @Schemes
// @Description coordinate agent
// @Tags coordinate agent
// @Accept json
// @Produce json
// @Success 200
// @Failure      500
// @Failure      400
// @Router /coordinate [post]
func (h *Handler) coordinate(ctx *gin.Context) {
	requestData := Coordinatation{}
	err := ctx.BindJSON(&requestData)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}
	go h.u.SubmitCoordinate(requestData.ToDomain())
	ctx.JSON(http.StatusOK, nil)
}
