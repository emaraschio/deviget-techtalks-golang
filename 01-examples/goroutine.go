package main

import (
    "runtime"
    "sync"
    "fmt"
)

func init() {
    //runtime.GOMAXPROCS(2)
    runtime.GOMAXPROCS(10)
}

func main() {
    var wg sync.WaitGroup
    wg.Add(4)

    fmt.Println("Init Goroutines")
    go func() {
        for count := 10; count >= 0; count-- {
            fmt.Printf("[A:%d]\n", count)
        }

        wg.Done()
    }()

    go func(){
        for count := 0; count < 10; count++ {
            fmt.Printf("[B:%d]\n", count)
        }

        wg.Done()
    }()

    go func(){
        for count := 0; count < 10; count++ {
            fmt.Printf("[C:%d]\n", count)
        }

        wg.Done()
    }()

    go func(){
        for count := 0; count < 10; count++ {
            fmt.Printf("[D:%d]\n", count)
        }

        wg.Done()
    }()

    fmt.Println("Waiting goroutines to finish")
    wg.Wait()

    fmt.Println("\nFinish the program")
}