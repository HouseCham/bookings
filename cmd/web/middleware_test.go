package main

import (
	"net/http"
	"testing"
)

func TestNoSurve(t *testing.T) {
	var myH myHandler
	hHandler := NoSurf(&myH)

	switch v := hHandler.(type) {
	case http.Handler:

	default:
		t.Errorf("type is not http.Handler -> %T", v)
	}
}

func TestSessionLoad(t *testing.T) {
	var myH myHandler
	h := SessionLoad(&myH)

	switch v := h.(type) {
	case http.Handler:

	default:
		t.Errorf("type is not http.Handler -> %T", v)
	}
}