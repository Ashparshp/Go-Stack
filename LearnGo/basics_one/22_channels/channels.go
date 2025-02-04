package main

/*
func processNum(c chan int) {
	fmt.Println("Processing num: ", <-c)
}
*/

/*..reciving the data..
func processNum(c chan int) {

	for num := range c {
		fmt.Println("Processing num: ", num)
		time.Sleep(time.Second)
	}
}
*/

/*..sending result in channel after recieving..
func sum(c chan int, num1, num2 int) {
	result := num1 + num2
	c <- result

}
*/

/*..revisit.. (Waitgroup)
func task(isdone chan bool) {
	defer func() { isdone <- true }()
	fmt.Println("Processing...")

}
*/

/*
func emailSender(e <-chan string, d chan<- bool) {
	defer func() { d <- true }()
	for email := range e {
		fmt.Println("Sending email to: ", email)
		time.Sleep(time.Second)
	}
}
*/




func main() {

	/*..deadlock..
		messageChannel := make(chan string)
		messageChannel <- "ping"
		msg := <-messageChannel
		fmt.Println(msg)
	*/

	/*..processNum is ready..
		numChannel := make(chan int)
		go processNum(numChannel)
		numChannel <- 5
		time.Sleep(time.Second)
	*/

	/*..sending multiple data..
		numChannel := make(chan int)
		go processNum(numChannel)

		for {
			numChannel <- rand.Intn(100)
		}
	*/

	/*..receiving the result in res..
		result := make(chan int)
		go sum(result, 4, 5)
		res := <-result
		fmt.Println(res)
	*/

	/*..revisit..
	done := make(chan bool)
	go task(done)
	<-done
	*/

	/*..Buffered channel (non-blocking)..
	emailChannel := make(chan string, 100)
	emailChannel <- "a@gmail.com"
	emailChannel <- "b@gmail.com"

	fmt.Println(<-emailChannel )
	fmt.Println(<-emailChannel )
	*/

	/*
	emailChannel := make(chan string, 100)
	done := make(chan bool)
	go emailSender(emailChannel, done)

	for i:=1; i<100; i++ {
		emailChannel <- fmt.Sprintf("%d@gmail.com", i)
	}
	
	fmt.Println("done")
	close(emailChannel)
	<-done
	*/


	/*..goroutine synchronization using channels and the select statement..
	channelOne := make(chan int)
	channelTwo := make(chan string)

	go func() {
			channelOne <- 10
	} ()

	go func() {
		channelTwo <- "Hello"
	} ()

	for i:=0; i<2; i++ {
		select {
		case msg := <-channelOne:
			fmt.Println("Recieved: ", msg)
		case msg := <-channelTwo:
			fmt.Println("Recieved: ", msg)
		}
	}
	*/

}
