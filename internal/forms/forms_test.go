package forms

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestForm_IsValid(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever-url", nil)
	form := New(r.PostForm)

	isValid := form.IsValid()
	if !isValid {
		t.Error("got invalid when should have been valid")
	}
}
func TestForm_Required(t *testing.T) {
	r := httptest.NewRequest("POST", "/anyUrl", nil)
	form := New(r.PostForm)

	form.Required("a", "b", "c")
	if form.IsValid() {
		t.Error("form shows valid when required fields missing")
	}

	postedData := url.Values{}
	postedData.Add("a", "a")
	postedData.Add("b", "b")
	postedData.Add("c", "c")

	r, _ = http.NewRequest("POST", "/whateverUrl", nil)
	r.PostForm = postedData
	form = New(r.PostForm)
	form.Required("a", "b", "c")
	if !form.IsValid() {
		t.Error("does not show required fields when it should")
	}
}

func TestForm_Has(t *testing.T) {
	r := httptest.NewRequest("POST", "/anyUrl", nil)
	form := New(r.PostForm)

	has := form.Has("whatever")
	if has {
		t.Error("form shows has field when it doesnt")
	}

	postedData := url.Values{}
	postedData.Add("a", "a")
	form = New(postedData)

	has = form.Has("a")
	if !has {
		t.Error("shows form does not have field when should")
	}
}

func TestForm_MinLength(t *testing.T) {
	r := httptest.NewRequest("POST", "/anyUrl", nil)
	form := New(r.PostForm)

	form.MinLength("X", 10)
	if form.IsValid() {
		t.Error("form shows min length for non existing field")
	}

	isError := form.Errors.Get("X")
	if isError == "" {
		t.Error("should have an error but did not get one")
	}

	// not valid data
	postedData := url.Values{}
	postedData.Add("some_field", "some value")
	form = New(postedData)

	form.MinLength("some_field", 100)
	if form.IsValid() {
		t.Error("shows minlength of 100 when data is shorter")
	}

	// valid data
	postedData = url.Values{}
	postedData.Add("another_field", "abc123")
	form = New(postedData)

	form.MinLength("another_field", 3)
	if !form.IsValid() {
		t.Error("form doesnt show field when should")
	}

	isError = form.Errors.Get("another_field")
	if isError != "" {
		t.Error("should not have an error but got one")
	}
}

func TestForm_IsEmail(t *testing.T) {
	postedData := url.Values{}
	form := New(postedData)

	form.IsEmail("x")
	if form.IsValid() {
		t.Error("shows email field when it is empty")
	}

	// valid data
	postedData = url.Values{}
	postedData.Add("email", "chamses@gmail.com")
	form = New(postedData)

	form.IsEmail("email")
	if !form.IsValid() {
		t.Error("doesnt accept email when should")
	}

	postedData = url.Values{}
	postedData.Add("email", "chamses")
	form = New(postedData)

	form.IsEmail("email")
	if form.IsValid() {
		t.Error("got valid email when should not")
	}
}
