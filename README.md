# go-webhookrelay

[Webhookrelay](https://webhookrelay.com/api-reference/) API client for Go

## Example usage

``` go
package main

import (
	"fmt"
	"os"

	"github.com/koalificationio/go-webhookrelay/pkg/client"
	api "github.com/koalificationio/go-webhookrelay/pkg/openapi/client/buckets"
	"github.com/koalificationio/go-webhookrelay/pkg/openapi/models"
)

func main() {
	cfg := client.Config{
		ApiKey:    os.Getenv("RELAY_KEY"),
		ApiSecret: os.Getenv("RELAY_SECRET"),
	}

	client := client.New(&cfg)

	// get some bukets
	buckets, err := client.Buckets.GetV1Buckets(api.NewGetV1BucketsParams())
	if err != nil {
		fmt.Printf("Error geting bucktes: %v", err)
		os.Exit(1)
	}

	fmt.Printf("First bucket name: %v", buckets.GetPayload()[0].Name)

	// create bucket and get input url
	params := api.NewPostV1BucketsParams().WithBody(&models.BucketRequest{Name: "test_bucket"})
	resp, err := client.Buckets.PostV1Buckets(params)
	if err != nil {
		fmt.Printf("Error creating bucktes: %v", err)
		os.Exit(1)
	}
	fmt.Printf("New bucket input url: https://my.webhookrelay.com/v1/webhooks/%s", resp.GetPayload().Inputs[0].ID)
}
```
