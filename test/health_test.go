// +build e2e
package test

import (
	"fmt"
	"testing"

	"github.com/go-resty/resty/v2"
)

func TestHealthendpoint(t *testing.T) {
	fmt.Println("Running Test for health")
	client := resty.New()
	resp, err := client.R().Get("http://localhost:8080/api/health"); err != nil {
		t.Fail()
	}
	fmt.Println(resp.StatusCode())
}