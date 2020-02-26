// +build wireinject

package main

import (
	"github.com/google/wire"
	"hz_project/internal/app/products"
	"hz_project/internal/app/products/controllers"
	"hz_project/internal/app/products/services"
	"hz_project/internal/app/products/grpcclients"
	"hz_project/internal/pkg/config"
	"hz_project/internal/pkg/consul"
	"hz_project/internal/pkg/log"
	"hz_project/internal/pkg/jaeger"
	"hz_project/internal/pkg/app"
	"hz_project/internal/pkg/transports/grpc"
	"hz_project/internal/pkg/transports/http"
)

var providerSet = wire.NewSet(
	log.ProviderSet,
	config.ProviderSet,
	consul.ProviderSet,
	jaeger.ProviderSet,
	http.ProviderSet,
	grpc.ProviderSet,
	grpcclients.ProviderSet,
	controllers.ProviderSet,
	services.ProviderSet,
	products.ProviderSet,
)


func CreateApp(cf string) (*app.Application, error) {
	panic(wire.Build(providerSet))
}
