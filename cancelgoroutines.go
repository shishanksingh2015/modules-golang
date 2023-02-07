package main

func main() {
	quit := make(chan int)
	go func() {
		for {
			select {
			case <-quit:
				return
			default:
				//anycode
			}
		}
	}()

	quit <- 2
}
