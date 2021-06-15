package main_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddItemSuccess(t *testing.T) {
	client := &http.Client{}
	request, _ := http.NewRequest("GET", "http://localhost:3000/addItem/10", nil)
	resp, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	if err != nil {
		panic(err)
	}
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestRemoveItemSuccess(t *testing.T) {
	client := &http.Client{}
	request, _ := http.NewRequest("GET", "http://localhost:3000/removeItem/10", nil)
	resp, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	if err != nil {
		panic(err)
	}
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestHasItemError(t *testing.T) {
	client := &http.Client{}
	request, _ := http.NewRequest("GET", "http://localhost:3000/hasItem/879", nil)
	resp, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	if err != nil {
		panic(err)
	}
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func TestAddItemError(t *testing.T) {
	client := &http.Client{}
	request, _ := http.NewRequest("GET", "http://localhost:3000/addItem/10ghh", nil)
	resp, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	if err != nil {
		panic(err)
	}
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}
