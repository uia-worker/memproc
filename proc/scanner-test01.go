package main

import (
	"bufio"
	"fmt"
	"os"
	"log"
	"runtime"
)

func main() {
	var input string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input = scanner.Text()
		if input == "q" || input == "exit" {
			fmt.Println("exit")
			os.Exit(0) 
		} else if input == "pid" {
			fmt.Println("Prosess id er ", checkProcessID())
			memoryInfo()
		} else {
			fmt.Println(input) // Println will add back the final '\n'
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}

func checkProcessID() int {
	return os.Getpid()
}

func memoryInfo() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	log.Println(float64(m.Sys))
	log.Println(float64(m.HeapAlloc))
}
