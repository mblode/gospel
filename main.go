package main

import (
	"github.com/mblode/gospel/db"
	"github.com/mblode/gospel/router"
)

func main() {
	r := router.New()
	v1 := r.Group("/api")

	d := db.New()
	db.AutoMigrate(d)

	userStore := store.NewUserStore(d)
	locationStore := store.NewLocationStore(d)
	h := handler.NewHandler(userStore, locationStore)
	h.Register(v1)

	r.Logger.Fatal(r.Start("127.0.0.1:3000"))
}
