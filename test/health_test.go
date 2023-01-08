//go:build e2e
// +build e2e

package test

import (
	"fmt"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
)

func TestHealthendpoint(t *testing.T) {
	fmt.Println("Running Test for health")
	client := resty.New()
	resp, err := client.R().Get("http://localhost:8080/api/health")
	if err != nil {
		t.Fail()
	}
	fmt.Println(resp.StatusCode())
	assert.Equal(t, 200, resp.StatusCode())
}
