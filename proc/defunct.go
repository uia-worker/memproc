package main

import (
    "os"
    "fmt"
)

type Config struct {
    Debug bool
}

func main() {
    config := Config{
    	    Debug: true,
    }
    mypid := os.Getpid()
    if mypid != 1 {
	    if config.Debug {
		    fmt.Printf("PID er %d og ikke 1.\n", mypid)
		    return
	    }
    }
    fmt.Println("PID er 1")
}
