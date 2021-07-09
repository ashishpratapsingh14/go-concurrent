package main

import (
	"fmt"
	_ "net/http/pprof"
	"sync"
	"time"
)

//func main() {
//	d := gomail.NewDialer("smtp.gmail.com", 587, "apratapfanatics@gmail.com", "ikwyejldrmpjmngs")
//	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
//
//    // Send emails using d.
//    m := gomail.NewMessage()
//    m.SetHeader("To","apratap@fanatics.com")
//    m.SetHeader("From", "apratapfanatics@gmail.com")
//    m.SetHeader("Subject", "Testing gomail")
//    m.SetHeader("text/plain","Hello Message")
//    m.
//    err := d.DialAndSend(m)
//    if err != nil {
//    	fmt.Println(err.Error())
//	}else{
//		fmt.Println("Success")
//	}
//}

//func main() {
//
//	tc := flag.Int("threads", 10, "Thread Count")
//	rut := flag.Int("rampup", 30, "Ramp up time in seconds")
//	et := flag.Int("etime", 1, "Execution time in minutes")
//	flag.Parse()
//
//	//Check if execution time is more than ramp up time
//	if *et*60 < *rut {
//		log.Fatalln("Total execution time needs to be more than ramp up time")
//	}
//
//	waitTime := *rut / *tc
//
//	log.Printf("Execution will happen with %d users with a ramp up time of %d seconds for %d minutes\n", *tc, *rut, *et)
//
//	tchan := make(chan int)
//	go func(c chan<- int) {
//		for ti := 1 ; ti <= *tc ; ti++ {
//			log.Printf("Thread Count %d", ti)
//			c <- ti
//			time.Sleep(time.Duration(waitTime) * time.Second)
//		}
//	}(tchan)
//
//	timeout := time.After(time.Duration(*et*60) * time.Second)
//
//	for {
//		select { //Select blocks the flow until one of the channels receives a message
//		case <-timeout: //receives a msg when execution duration is over
//			log.Printf("Execution completed")
//			return
//		case ts := <-tchan: //receives a message when a new user thread has to be initiated
//			log.Printf("Thread No %d started", ts)
//			go func(t int) {
//				//This is the place where you add all your tests
//				//In my case they were making rpc calls over rabbitmq with random inputs
//				//They keep running till the end of execution
//				for  {
//					fmt.Println("Hello")
//				}
//			}(ts)
//		}
//	}
//}

//func main() {
//	// we need a webserver to get the pprof webserver
//	go func() {
//		log.Println(http.ListenAndServe("localhost:6060", nil))
//	}()
//	fmt.Println("hello world")
//	var wg sync.WaitGroup
//	wg.Add(1)
//	go leakyFunction(wg)
//	wg.Wait()
//}

//func main() {
//	defer profile.Start().Stop()
//	fmt.Println("hello world")
//	var wg sync.WaitGroup
//	wg.Add(1)
//	go leakyFunction(wg)
//	wg.Wait()
//}

//func leakyFunction(wg sync.WaitGroup) {
//	defer wg.Done()
//	s := make([]string, 3)
//	for i := 0; i < 10000000; i++ {
//		s = append(s, "magical pandas")
//		if (i % 100000) == 0 {
//			time.Sleep(500 * time.Millisecond)
//		}
//	}
//}
//
//func main() {
//	//http.Handle("/metrics", promhttp.Handler())
//	runtime.SetCPUProfileRate(1000)
//	go func() {
//		http.ListenAndServe(":8080", nil)
//	}()
//	fmt.Println("hello world")
//	var wg sync.WaitGroup
//	wg.Add(1)
//	go leakyFunction(wg)
//	wg.Wait()
//}


func main() {
	JobQueue = make(chan Job, 10000)
	dispatcher := NewDispatcher(500)
	dispatcher.Run()
	doFunc := func(data interface{}) error {
		fmt.Println(data)
		return nil
	}
	var wg sync.WaitGroup
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go func(c int) {
			defer wg.Done()
			start:=c*100
			end :=(c+1)*100
			for j := start; j < end; j++ {
				// let's create a job with the payload
				work := Job{Payload: j, Executor:doFunc}
				// Push the work onto the queue.
				JobQueue <- work
			}
		}(i)
	}
	wg.Wait()
	time.Sleep(time.Minute)
}