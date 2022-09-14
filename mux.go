package main

import (
	"context"
	"net/http"

	"github.com/funasedaisuke/go-web-application/clock"
	"github.com/funasedaisuke/go-web-application/config"
	"github.com/funasedaisuke/go-web-application/handler"
	"github.com/funasedaisuke/go-web-application/service"
	"github.com/funasedaisuke/go-web-application/store"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

func NewMux(ctx context.Context,cfg *config.Config)( http.Handler,func(),error){
	mux := chi.NewRouter()
	mux.HandleFunc("/health",func(w http.ResponseWriter,r *http.Request){
		w.Header().Set("Content-Type","application/json; charset=utf-8")
		_,_=w.Write([]byte(`{"STATUS":"OK"}`))
	})
	v := validator.New()
	db,cleanup,err := store.New(ctx,cfg)
	if err != nil{
		return nil,cleanup,err
	}
	
	r := store.Repository{Clocker: clock.RealClocker{}}
	at := &handler.AddTask{
		Service: &service.AddTask{DB:db, Repo:&r},
		Validator: v,
	}
	mux.Post("/tasks",at.ServeHTTP)
	lt := &handler.ListTask{
		Service: &service.ListTask{DB: db,Repo:&r},}
	mux.Get("/tasks",lt.ServeHTTP)
	ru :=&handler.RegisterUser{
		Service: &service.RegisterUser{DB:db, Repo: &r},
	Validator: v, }
	mux.Get("/register",ru.ServeHTTP)
	return mux,cleanup,nil
}