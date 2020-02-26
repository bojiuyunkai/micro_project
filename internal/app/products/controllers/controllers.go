package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"hz_project/internal/pkg/transports/http"
)

func CreateInitControllersFn(
	pc *ProductsController,
) http.InitControllers {
	return func(r *gin.Engine) {
		r.GET("/product/:id", pc.Get)
	}
}

var ProviderSet = wire.NewSet(NewProductsController, CreateInitControllersFn)
