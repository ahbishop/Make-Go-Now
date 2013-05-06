package main

import (
	"fmt"
	"os"
	"strings"
	"time"
	"math"
	"encoding/json"
	"log"
	"io"
)

func a_function(data []float32, comm_link chan float32) {
	result := float32(0)
	for _, val := range data {
		result += val
	}
	time.Sleep(1000 * time.Millisecond)
	comm_link <- result
}

func thingy(derp, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case derp <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("Die die die")
			return
	  default:
			fmt.Print("   .")
			time.Sleep(4e8)
		}
	}
}

type Weasel struct{
	Num, Power float64
}

var m map[string]Weasel
var n = map[string]Weasel{
	"Mr. White" : { -13180123.0124, 109.1813106743},
	"Mr. Pink" : { -12.0001238513487, 12.139877430},
}

func main() {
	da_var := "Sup sup"
	if len(os.Args) > 1 {
		da_var = strings.Join(os.Args[1:], " ")
	}
	fmt.Println("I can haz: ", da_var)

	data := []float32{3.14, 5.3, 13, 50.32, 10.13, 1513.123132143, 3234.1231}
	c := make(chan float32)
	go a_function(data, c)
	result := <-c

	fmt.Println(result)

	derp := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-derp)
		}
		quit <- 0
	}()
	thingy(derp, quit)


	m = make(map[string]Weasel)

	m["whut"] = Weasel{
		3.14589762, -1.011151304,
	}
	delete(m,"whut")
	delete(m,"nope")

	shrimp_taco := func(x,y float64) float64 {
		return math.Sqrt(x*x*1.43 + y*3.1+ 0.5*y*y)
	}

	fmt.Println(shrimp_taco(12.123123862485,3.9876542776))
	
	
	const da_json = ` {"Type":42, "Text":"Lulzorz"}
					  {"Type":"Gnarly", "Text":"Hang loose"}
					 `
	type Mail struct{
		Type, Text string
	}
	
	json_decoder := json.NewDecoder(strings.NewReader(da_json))
	log.Print("Doing some things")
	for{
		var m Mail
		if err:= json_decoder.Decode(&m); err == io.EOF{
			break
		}
		fmt.Printf("I formatted this -> %s: %s\n", m.Type, m.Text)
		
		input := Mail{ Type: "Sweet", Text:"Narf"}
		gimp, derr := json.Marshal(input)
		if derr != nil{
			fmt.Println("Massive failure:", derr)
		}
		os.Stdout.Write(gimp)
	}
	

}
