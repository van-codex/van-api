// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"github.com/kainonly/go-bit"
	"lab-api/bootstrap"
	"lab-api/controller"
	"lab-api/service"
)

// Injectors from wire.go:

func Boot(config bit.Config) (*controller.Controllers, error) {
	db, err := bootstrap.InitializeDatabase(config)
	if err != nil {
		return nil, err
	}
	client, err := bootstrap.InitializeRedis(config)
	if err != nil {
		return nil, err
	}
	dependency := &service.Dependency{
		Db:    db,
		Redis: client,
	}
	crud := bit.InitializeCrud(db)
	cookie, err := bit.InitializeCookie(config)
	if err != nil {
		return nil, err
	}
	serviceDependency := service.Dependency{
		Db:    db,
		Redis: client,
	}
	admin := service.NewAdmin(serviceDependency)
	services := &service.Services{
		Dependency: dependency,
		Crud:       crud,
		Cooike:     cookie,
		Admin:      admin,
	}
	index := controller.NewIndex(services)
	resource := controller.NewResource(services)
	controllerAdmin := controller.NewAdmin(services)
	controllers := &controller.Controllers{
		Index:    index,
		Resource: resource,
		Admin:    controllerAdmin,
	}
	return controllers, nil
}
