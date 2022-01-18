// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"api/app"
	"api/app/index"
	"api/app/pages"
	"api/app/roles"
	"api/app/users"
	"api/bootstrap"
	"api/common"
	"github.com/gin-gonic/gin"
	"github.com/weplanx/go/engine"
)

// Injectors from wire.go:

func App(value *common.Values) (*gin.Engine, error) {
	passport := bootstrap.UsePassport(value)
	client, err := bootstrap.UseMongoDB(value)
	if err != nil {
		return nil, err
	}
	database := bootstrap.UseDatabase(client, value)
	redisClient, err := bootstrap.UseRedis(value)
	if err != nil {
		return nil, err
	}
	conn, err := bootstrap.UseNats(value)
	if err != nil {
		return nil, err
	}
	cipher, err := bootstrap.UseCipher(value)
	if err != nil {
		return nil, err
	}
	iDx, err := bootstrap.UseIDx(value)
	if err != nil {
		return nil, err
	}
	inject := &common.Inject{
		Values:      value,
		MongoClient: client,
		Db:          database,
		Redis:       redisClient,
		Nats:        conn,
		Passport:    passport,
		Cipher:      cipher,
		Idx:         iDx,
	}
	service := &index.Service{
		Inject: inject,
	}
	usersService := &users.Service{
		Inject: inject,
	}
	pagesService := &pages.Service{
		Inject: inject,
	}
	controller := &index.Controller{
		Service: service,
		Users:   usersService,
		Pages:   pagesService,
	}
	jetStreamContext, err := bootstrap.UseJetStream(conn)
	if err != nil {
		return nil, err
	}
	engineEngine := bootstrap.UseEngine(value, jetStreamContext)
	engineService := &engine.Service{
		Engine: engineEngine,
		Db:     database,
	}
	engineController := &engine.Controller{
		Engine:  engineEngine,
		Service: engineService,
	}
	pagesController := &pages.Controller{
		Service: pagesService,
	}
	rolesService := &roles.Service{
		Inject: inject,
	}
	rolesController := &roles.Controller{
		Service: rolesService,
	}
	usersController := &users.Controller{
		Service: usersService,
	}
	ginEngine := app.New(value, passport, controller, engineController, pagesController, rolesController, usersController)
	return ginEngine, nil
}
