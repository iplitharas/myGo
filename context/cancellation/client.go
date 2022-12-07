package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
)

var client = http.Client{}

func callBoth(ctx context.Context, errVal, slowURL, fastURL string) {
	// ctx it's not the same, it's a child context that wraps the passed-in parent
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		err := callServer(ctx, "slow", slowURL)
		if err != nil {
			// cancel
			cancel()
		}
	}()
	go func() {
		defer wg.Done()
		err := callServer(ctx, "fast", fastURL+"?error="+errVal)
		if err != nil {
			// cancel
			cancel()
		}
	}()
	// wait for both
	wg.Wait()
	fmt.Println("done with both")
}

func callServer(ctx context.Context, label, url string) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		fmt.Println(label, "request err:", err)
		return err
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(label, "response err:", err)
		return err
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(label, "read err:", err)
		return err
	}
	result := string(data)
	if result != "" {
		fmt.Println("from:", label, ":response is:", result)
	}
	if result == "error" {
		fmt.Println("canceling from", label)
		return errors.New("error")
	}
	return nil
}

func main() {
	ss := SlowServer()
	defer ss.Close()
	fs := FastServer()
	defer fs.Close()

	ctx := context.Background()
	callBoth(ctx, os.Args[1], ss.URL, fs.URL)
}
