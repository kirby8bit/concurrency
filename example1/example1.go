package example1

func SomeFunc() {
	// создаем новый канал типа int
	ch := make(chan int)
	// запускаем новую анонимную горутину
	go func() {
		// отправляем 42 в канал
		ch <- 42
	}()
	// ждем, читаем из канала
	<-ch
}
