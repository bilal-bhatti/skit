package skit

import (
	"io"
	"log"
	"net/http/httptest"
	"testing"
)

func TestSuccessRender(t *testing.T) {
	tests := []struct {
		in   interface{}
		code int
	}{
		{
			in:   []string{"one", "two"},
			code: 200,
		},
		{
			in:   make(chan int),
			code: 500,
		},
	}

	for _, test := range tests {
		resp := httptest.NewRecorder()

		Success(resp, test.in)

		log.Print(string(resp.Body.Bytes()))
		log.Println(resp.Header())
		log.Println(resp.Result().Status)
		log.Println(resp.Result().StatusCode)

		if test.code != resp.Result().StatusCode {
			t.Errorf("code mismatch - expected: %d, actual: %d", test.code, resp.Result().StatusCode)
		}
	}
}

func TestFailureRender(t *testing.T) {
	tests := []struct {
		in   error
		code int
	}{
		{
			in:   io.EOF,
			code: 500,
		},

		{
			in:   WithStatus(io.EOF, 500, "unexpected error"),
			code: 500,
		},
		{
			in:   WithStatus(io.EOF, 500, []string{"one", "two"}),
			code: 500,
		},
		{
			in:   WithStatus(io.EOF, 500, make(chan int)),
			code: 500,
		},
	}

	for _, test := range tests {
		resp := httptest.NewRecorder()

		Failure(resp, test.in)

		log.Print("body:", string(resp.Body.Bytes()))
		log.Println("headers:", resp.Header())
		log.Println("status:", resp.Result().Status)
		log.Println("code:", resp.Result().StatusCode)

		if test.code != resp.Result().StatusCode {
			t.Errorf("code mismatch - expected: %d, actual: %d", test.code, resp.Result().StatusCode)
		}
	}
}
