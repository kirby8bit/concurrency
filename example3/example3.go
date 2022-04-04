package example3

import "time"

func player(table chan int) {
	for {
		ball := <-table                    //отскок
		ball++                             //удар в ракетку
		time.Sleep(100 * time.Millisecond) //полет
		table <- ball                      //мяч ударился в стол
	}
}

func SomeFunc() {
	var Ball int
	table := make(chan int)
	go player(table)
	go player(table)

	table <- Ball               //удар от ракетки в стол
	time.Sleep(1 * time.Second) //мяч (0) летит
	<-table                     //отскок
}
