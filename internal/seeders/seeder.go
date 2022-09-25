package seeders

import (
	"fmt"
	"github.com/cherryReptile/GoSeeder/internal/models"
	"github.com/jaswdr/faker"
	"github.com/jmoiron/sqlx"
	"log"
	"os"
	"sync"
	"time"
)

func Run(db *sqlx.DB, wg *sync.WaitGroup, goroutineNum int) {
	start := time.Now()
	defer wg.Done()
	for counter := 0; counter < 6250; counter++ {
		create(db)
	}
	finish := time.Since(start)
	fmt.Println(fmt.Sprintf("****finished goroutine with num %v after %v****", goroutineNum+1, finish))
}

func create(db *sqlx.DB) {
	model := new(models.Fake)
	faker := faker.New()
	faker.Struct().Fill(model)
	err := model.Create(db)

	if err != nil {
		log.Fatalf("[ERROR] %v", err)
		os.Exit(1)
	}

}
