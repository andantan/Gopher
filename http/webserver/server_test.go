package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJsonHandler(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/students", nil)

	mux := MakeWebHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)

	var list []Student

	err := json.NewDecoder(res.Body).Decode(&list)

	assert.Nil(err)

	assert.Equal(2, len(list))
	assert.Equal("AAA", list[0].Name)
	assert.Equal("BBB", list[1].Name)
}

func TestJsonHander2(t *testing.T) {
	assert := assert.New(t)
	var student Student

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/students/1", nil)

	mux := MakeWebHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)

	err := json.NewDecoder(res.Body).Decode(&student)

	assert.Nil(err)

	assert.Equal("AAA", student.Name)
}
