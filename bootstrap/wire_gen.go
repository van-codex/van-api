// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package bootstrap

import (
	"github.com/weplanx/go/rest"
	"github.com/weplanx/go/sessions"
	"github.com/weplanx/go/values"
	"github.com/weplanx/server/api"
	"github.com/weplanx/server/api/index"
	"github.com/weplanx/server/common"
)

// Injectors from wire.go:

func NewAPI(values2 *common.Values) (*api.API, error) {
	client, err := UseMongoDB(values2)
	if err != nil {
		return nil, err
	}
	database := UseDatabase(values2, client)
	redisClient, err := UseRedis(values2)
	if err != nil {
		return nil, err
	}
	conn, err := UseNats(values2)
	if err != nil {
		return nil, err
	}
	jetStreamContext, err := UseJetStream(conn)
	if err != nil {
		return nil, err
	}
	keyValue, err := UseKeyValue(values2, jetStreamContext)
	if err != nil {
		return nil, err
	}
	cipher, err := UseCipher(values2)
	if err != nil {
		return nil, err
	}
	passport := UsePassport(values2)
	captcha := UseCaptcha(values2, redisClient)
	locker := UseLocker(values2, redisClient)
	inject := &common.Inject{
		V:         values2,
		Mgo:       client,
		Db:        database,
		RDb:       redisClient,
		JetStream: jetStreamContext,
		KeyValue:  keyValue,
		Cipher:    cipher,
		Passport:  passport,
		Captcha:   captcha,
		Locker:    locker,
	}
	hertz, err := UseHertz(values2)
	if err != nil {
		return nil, err
	}
	csrf := UseCsrf(values2)
	service := UseValues(values2, keyValue, cipher)
	controller := &values.Controller{
		Service: service,
	}
	sessionsService := UseSessions(values2, redisClient)
	sessionsController := &sessions.Controller{
		Service: sessionsService,
	}
	restService := UseRest(values2, client, database, redisClient, jetStreamContext, keyValue)
	restController := &rest.Controller{
		Service: restService,
	}
	indexService := &index.Service{
		Inject:   inject,
		Sessions: sessionsService,
	}
	indexController := &index.Controller{
		IndexService: indexService,
		Csrf:         csrf,
	}
	apiAPI := &api.API{
		Inject:       inject,
		Hertz:        hertz,
		Csrf:         csrf,
		Values:       controller,
		Sessions:     sessionsController,
		Rest:         restController,
		Index:        indexController,
		IndexService: indexService,
	}
	return apiAPI, nil
}
