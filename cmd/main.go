package main

import (
	"fmt"
	"github.com/cherryReptile/GoSeeder/internal/base"
	"github.com/cherryReptile/GoSeeder/internal/seeders"
	"runtime"
	"sync"
)

func main() {
	app := new(base.App)
	app.Init()

	fmt.Println("start")
	runtime.GOMAXPROCS(4)

	wg := &sync.WaitGroup{}
	for goroutineNum := 0; goroutineNum < 80; goroutineNum++ {
		wg.Add(1)
		go seeders.Run(app.DB, wg, goroutineNum)
		fmt.Println("start goroutine", goroutineNum+1)
	}
	wg.Wait()
	fmt.Println("db was filled")
}
