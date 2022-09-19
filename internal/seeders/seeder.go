package seeders

//func Run(db *sqlx.DB) {
//	wg := &sync.WaitGroup{}
//	for counter := 0; counter < 500000; counter++ {
//		wg.Add(1)
//		go create(db)
//		fmt.Sprintf("created %v post", counter)
//		runtime.Gosched()
//	}
//	wg.Wait()
//	fmt.Println("success")
//}
//
//func create(db *sqlx.DB) {
//	model := new(models.Fake)
//	faker := faker.New()
//	faker.Struct().Fill(model)
//	err := model.Create(db)
//
//	if err != nil {
//		os.Exit(1)
//	}
//
//	fmt.Println("success")
//}
