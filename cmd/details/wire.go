// +build wireinject

package main

import (
	"github.com/google/wire"
	"hz_project/internal/app/details/controllers"
	"hz_project/internal/app/details/grpcservers"
	"hz_project/internal/app/details/repositories"
	"hz_project/internal/app/details/services"
	"hz_project/internal/app/details"
	"hz_project/internal/pkg/app"
	"hz_project/internal/pkg/config"
	"hz_project/internal/pkg/consul"
	"hz_project/internal/pkg/database"
	"hz_project/internal/pkg/jaeger"
	"hz_project/internal/pkg/log"
	"hz_project/internal/pkg/transports/grpc"
	"hz_project/internal/pkg/transports/http"
)

var providerSet = wire.NewSet(
	log.ProviderSet,
	config.ProviderSet,
	database.ProviderSet,
	services.ProviderSet,
	repositories.ProviderSet,
	consul.ProviderSet,
	jaeger.ProviderSet,
	http.ProviderSet,
	grpc.ProviderSet,
	details.ProviderSet,
	controllers.ProviderSet,
	grpcservers.ProviderSet,
)

func CreateApp(cf string) (*app.Application, error) {
	panic(wire.Build(providerSet))
}
