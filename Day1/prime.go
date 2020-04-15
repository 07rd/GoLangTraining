package main

import "fmt"

func isprime(number int){
    check := true
    if number == 0 || number == 1 {
        fmt.Printf("%d is not a  prime number\n", number)
    } else {
        for i := 2; i <= number/2; i++ {
            if number%i == 0 {
                fmt.Printf("%d is not a  prime number\n", number)
                check = false
                break
            }
        }
    }
     if check == true {
         fmt.Printf("%d is a prime number\n", number)
     }

}
