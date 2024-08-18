package main

import (
	_ "encoding/json"
	"fmt"

	_ "net/http"
	"os"
	"time"
)

func main() {

	fmt.Printf("The Current time is in Nano-Seconds %v\n", time.Now().Nanosecond())
	current_page_Size := os.Getpagesize()

	fmt.Printf("The Current Page Size is %v\n", current_page_Size)

	fmt.Printf("The Program Exited time is in Nano-Seconds %v\n", time.Now().Nanosecond())

}
