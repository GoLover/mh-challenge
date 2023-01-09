package cmd

import (
	"mhlv/apps/center/delivery"
	"mhlv/apps/center/usecase"
	"mhlv/domain"

	"github.com/gin-gonic/gin"
)

var cu domain.CenterUsecase

func Boot(ginRouter *gin.Engine) {
	cu := usecase.NewCenterUsecase(5)
	delivery.New(ginRouter, cu)
}
