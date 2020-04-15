package main
import "fmt"
 
func star_pt(){
    for i := 1; i <= 5; i++ { 
	for k := 1; k <= i; k++ {
            fmt.Printf("*")              
        }
        fmt.Println("")     
    }
 }

func number_pt()  {
    var temp int = 1

    for i := 1; i <= 5; i++ {

        for k := 1; k <= i; k++ {

            fmt.Printf(" %d",temp)
            temp++
        }
        fmt.Println("")
    }
}