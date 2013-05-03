package main

import (
    "fmt"
    "os" 
    "strings"
	"time"
)

func a_function(data []float32, comm_link chan float32){
	result := float32(0)
	for _, val := range data {
		result += val
	}
	time.Sleep(1000 * time.Millisecond)
	comm_link <- result
}

func thingy(derp,quit chan int){
	x, y:=0,1
	for{
		select{
		case derp <- x:
			x, y = y, x+y
		case <- quit:
			fmt.Println("Die bitch")
			return
		}
	}

}

func main() {
    da_var := "Sup sup" 
    if len(os.Args) > 1 {  
        da_var = strings.Join(os.Args[1:], " ")
    }
    fmt.Println("I can haz: ", da_var)

	data := []float32{3.14, 5.3, 13, 50.32, 10.13, 1513.123132143, 531234.1231}
	c := make(chan float32)
	go a_function(data, c)
	result := <- c

	fmt.Println(result)

	derp := make(chan int)
	quit := make(chan int)
	go func() {
		for i:=0; i<10; i++ {
			fmt.Println(<-derp)
		}
		quit <-0
	}()
	thingy(derp, quit)
}
