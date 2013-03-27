package main
import (
    "fmt"
    "os" 
    "strings"
)

func main() 
{ 
    da_var := "Sup sup" 
    if len(os.Args) > 1 {  
        da_var = strings.Join(os.Args[1:], " ")
    }
    fmt.Println("I can haz: ", da_var)
}
