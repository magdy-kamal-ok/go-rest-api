//go:build e2e
// +build e2e

package test

import (
	"fmt"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
)

func TestGetComments(t *testing.T) {
	fmt.Println("Running Test for health")
	client := resty.New()
	resp, err := client.R().Get(BASE_URL + "/api/comment")
	if err != nil {
		t.Fail()
	}
	fmt.Println(resp.StatusCode())
	assert.Equal(t, 200, resp.StatusCode())
}

func TestPostComment(t *testing.T) {
	fmt.Println("Running Test for health")
	client := resty.New()
	resp, err := client.R().
		SetBody(`{"slug": "/", "author": "Abc"}`).
		Post(BASE_URL + "/api/comment")
	if err != nil {
		t.Fail()
	}
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode())
}
