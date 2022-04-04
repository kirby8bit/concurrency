package example2

import "time"

func timer(d time.Duration) <-chan int {
	//duration - продолжительность. в данном случае сна))
	c := make(chan int)
	go func() {
		time.Sleep(d)
		c <- 1
	}()
	return c
}

func SomeFunc() {
	for i := 0; i < 24; i++ {
		c := timer(1 * time.Second)
		<-c
	}
}
