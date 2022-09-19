package main

import (
	"github.com/cherryReptile/GoSeeder/internal/base"
	"github.com/cherryReptile/GoSeeder/internal/models"
	"log"
	"time"
)

func main() {
	app := new(base.App)
	app.Init()

	fake := models.Fake{String: "test", Date: time.Now()}
	err := fake.Create(app.DB)

	if err != nil {
		log.Fatal("SUka")
	}
}
