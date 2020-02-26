// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"github.com/google/wire"
	"hz_project/internal/app/ratings"
	"hz_project/internal/app/ratings/controllers"
	"hz_project/internal/app/ratings/grpcservers"
	"hz_project/internal/app/ratings/repositories"
	"hz_project/internal/app/ratings/services"
	"hz_project/internal/pkg/app"
	"hz_project/internal/pkg/config"
	"hz_project/internal/pkg/consul"
	"hz_project/internal/pkg/database"
	"hz_project/internal/pkg/jaeger"
	"hz_project/internal/pkg/log"
	"hz_project/internal/pkg/transports/grpc"
	"hz_project/internal/pkg/transports/http"
)

// Injectors from wire.go:

func CreateApp(cf string) (*app.Application, error) {
	viper, err := config.New(cf)
	if err != nil {
		return nil, err
	}
	options, err := log.NewOptions(viper)
	if err != nil {
		return nil, err
	}
	logger, err := log.New(options)
	if err != nil {
		return nil, err
	}
	ratingsOptions, err := ratings.NewOptions(viper, logger)
	if err != nil {
		return nil, err
	}
	httpOptions, err := http.NewOptions(viper)
	if err != nil {
		return nil, err
	}
	databaseOptions, err := database.NewOptions(viper, logger)
	if err != nil {
		return nil, err
	}
	db, err := database.New(databaseOptions)
	if err != nil {
		return nil, err
	}
	ratingsRepository := repositories.NewMysqlRatingsRepository(logger, db)
	ratingsService := services.NewRatingService(logger, ratingsRepository)
	ratingsController := controllers.NewRatingsController(logger, ratingsService)
	initControllers := controllers.CreateInitControllersFn(ratingsController)
	configuration, err := jaeger.NewConfiguration(viper, logger)
	if err != nil {
		return nil, err
	}
	tracer, err := jaeger.New(configuration)
	if err != nil {
		return nil, err
	}
	engine := http.NewRouter(httpOptions, logger, initControllers, tracer)
	consulOptions, err := consul.NewOptions(viper)
	if err != nil {
		return nil, err
	}
	client, err := consul.New(consulOptions)
	if err != nil {
		return nil, err
	}
	server, err := http.New(httpOptions, logger, engine, client)
	if err != nil {
		return nil, err
	}
	serverOptions, err := grpc.NewServerOptions(viper)
	if err != nil {
		return nil, err
	}
	ratingsServer, err := grpcservers.NewRatingsServer(logger, ratingsService)
	if err != nil {
		return nil, err
	}
	initServers := grpcservers.CreateInitServersFn(ratingsServer)
	grpcServer, err := grpc.NewServer(serverOptions, logger, initServers, client, tracer)
	if err != nil {
		return nil, err
	}
	application, err := ratings.NewApp(ratingsOptions, logger, server, grpcServer)
	if err != nil {
		return nil, err
	}
	return application, nil
}

// wire.go:

var providerSet = wire.NewSet(log.ProviderSet, config.ProviderSet, database.ProviderSet, services.ProviderSet, consul.ProviderSet, jaeger.ProviderSet, http.ProviderSet, grpc.ProviderSet, ratings.ProviderSet, repositories.ProviderSet, controllers.ProviderSet, grpcservers.ProviderSet)