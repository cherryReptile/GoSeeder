package main

import (
	"fmt"
	"github.com/cherryReptile/GoSeeder/internal/base"
	"github.com/cherryReptile/GoSeeder/internal/seeders"
	_ "go.uber.org/automaxprocs"
	"sync"
)

func main() {
	app := new(base.App)
	app.Init()

	fmt.Println("start")

	wg := &sync.WaitGroup{}
	for goroutineNum := 0; goroutineNum < 80; goroutineNum++ {
		wg.Add(1)
		go seeders.Run(app.DB, wg, goroutineNum)
		fmt.Println("start goroutine", goroutineNum+1)
	}
	wg.Wait()
	fmt.Println("db was filled")
}
