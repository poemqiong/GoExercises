package main

import "fmt"

func main() {
    s := "A"
    for i := 1; i < 100; i++ {
        fmt.Printf("%s\n",s)
        s = s + "A"
    }
} 
