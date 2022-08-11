package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func do(ctx context.Context, url string, timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	s, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(string(s))
	return nil
}

func main() {
	ctx := context.Background()
	err := do(ctx, "https://httpstat.us/200?sleep=3000", 1*time.Second)

	fmt.Println(err)
}
