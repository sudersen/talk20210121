package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	// Root Context
	rootCtx := context.Background()

	timeout := 1 * time.Microsecond
	timeoutDemo(rootCtx, timeout)
}

// START OMIT
// timeoutDemo demonstrates http timeout of 1 microsecond via context
func timeoutDemo(rootCtx context.Context, timeout time.Duration) {
	// The client
	client := &http.Client{}

	// Request
	req, err := http.NewRequest(http.MethodGet, "http://httpbin.org/get", nil)
	if err != nil {
		panic(err)
	}

	// Set timeout to the request via context
	timeoutCtx, cancel := context.WithTimeout(rootCtx, timeout)
	defer cancel()
	req = req.WithContext(timeoutCtx)

	// Response
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	out, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(out))
}
// END OMIT