package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestCreateBookHandlerNoData(t *testing.T) {
	initializeBooks()
	jsonStr := []byte(``)
	request, err := http.NewRequest("POST", "/book", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	rec := httptest.NewRecorder()
	handler := http.HandlerFunc(createBookHandler)
	handler.ServeHTTP(rec, request)

	if status := rec.Code; status != http.StatusBadRequest {
		t.Errorf("Wrong status code. Returned: %v Expected: %v",
			status, http.StatusOK)
	}
}

func TestCreateBookHandlerBadData(t *testing.T) {
	initializeBooks()
	jsonStr := []byte(`{"Id":"4","Title":"","Author":""}`)
	request, err := http.NewRequest("POST", "/book", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	rec := httptest.NewRecorder()
	handler := http.HandlerFunc(createBookHandler)
	handler.ServeHTTP(rec, request)

	if status := rec.Code; status != http.StatusBadRequest {
		t.Errorf("Wrong status code. Returned: %v Expected: %v",
			status, http.StatusOK)
	}
}

func TestCreateBookHandlerCreatesBook(t *testing.T) {
	initializeBooks()
	jsonStr := []byte(`{"Id":"4","Title":"Hello Hello","Author":"Jane Doe","Publisher":"Publish House","PublishDate":"10/23/2017","Rating":3,"Status":"Published"}`)

	request, err := http.NewRequest("POST", "/book", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	rec := httptest.NewRecorder()
	handler := http.HandlerFunc(createBookHandler)
	handler.ServeHTTP(rec, request)

	if status := rec.Code; status != http.StatusCreated {
		t.Errorf("Wrong status code. Returned: %v Expected: %v",
			status, http.StatusOK)
	}
}

func TestGetBookHandlerBadParameter(t *testing.T) {
	initializeBooks()
	request, err := http.NewRequest("GET", "/book/abc", nil)
	if err != nil {
		t.Fatal(err)
	}

	request = mux.SetURLVars(request, map[string]string{"id": "abc"})

	rec := httptest.NewRecorder()
	handler := http.HandlerFunc(getBookHandler)
	handler.ServeHTTP(rec, request)

	if status := rec.Code; status != http.StatusBadRequest {
		t.Errorf("Wrong status code. Returned: %v Expected: %v",
			status, http.StatusOK)
	}
}

func TestGetBookHandlerNotFound(t *testing.T) {
	initializeBooks()
	request, err := http.NewRequest("GET", "/book/4", nil)
	if err != nil {
		t.Fatal(err)
	}

	request = mux.SetURLVars(request, map[string]string{"id": "4"})

	rec := httptest.NewRecorder()
	handler := http.HandlerFunc(getBookHandler)
	handler.ServeHTTP(rec, request)

	if status := rec.Code; status != http.StatusNotFound {
		t.Errorf("Wrong status code. Returned: %v Expected: %v",
			status, http.StatusOK)
	}
}

func TestGetBookHandlerReturnsBook(t *testing.T) {
	initializeBooks()
	request, err := http.NewRequest("GET", "/book", nil)
	if err != nil {
		t.Fatal(err)
	}

	request = mux.SetURLVars(request, map[string]string{"id": "1"})

	rec := httptest.NewRecorder()
	handler := http.HandlerFunc(getBookHandler)
	handler.ServeHTTP(rec, request)

	if status := rec.Code; status != http.StatusOK {
		t.Errorf("Wrong status code. Returned: %v Expected: %v",
			status, http.StatusOK)
	}

	var book Book
	requestBody, _ := ioutil.ReadAll(rec.Body)
	json.Unmarshal(requestBody, &book)
	if book.ID != "1" {
		t.Errorf("Unexpected Id. Returned: %v Expected: %v", book.ID, "1")
	}
}

func TestUpdateBookHandlerBadParameter(t *testing.T) {
	initializeBooks()
	jsonStr := []byte(`{"Id":"4","Title":"Hello Hello","Author":"Jane Doe","Publisher":"Publish House","PublishDate":"10/23/2017","Rating":3,"Status":"Published"}`)
	request, err := http.NewRequest("PUT", "/book/abc", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	request = mux.SetURLVars(request, map[string]string{"id": "abc"})

	rec := httptest.NewRecorder()
	handler := http.HandlerFunc(updateBookHandler)
	handler.ServeHTTP(rec, request)

	if status := rec.Code; status != http.StatusBadRequest {
		t.Errorf("Wrong status code. Returned: %v Expected: %v",
			status, http.StatusOK)
	}
}

func TestUpdateBookHandlerNoData(t *testing.T) {
	initializeBooks()
	jsonStr := []byte(``)
	request, err := http.NewRequest("PUT", "/book/1", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	request = mux.SetURLVars(request, map[string]string{"id": "1"})

	rec := httptest.NewRecorder()
	handler := http.HandlerFunc(updateBookHandler)
	handler.ServeHTTP(rec, request)

	if status := rec.Code; status != http.StatusBadRequest {
		t.Errorf("Wrong status code. Returned: %v Expected: %v",
			status, http.StatusOK)
	}
}

func TestUpdateBookHandlerBadData(t *testing.T) {
	initializeBooks()
	jsonStr := []byte(`{"Id":"4","Title":"","Author":""}`)
	request, err := http.NewRequest("PUT", "/book/1", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	request = mux.SetURLVars(request, map[string]string{"id": "1"})

	rec := httptest.NewRecorder()
	handler := http.HandlerFunc(updateBookHandler)
	handler.ServeHTTP(rec, request)

	if status := rec.Code; status != http.StatusBadRequest {
		t.Errorf("Wrong status code. Returned: %v Expected: %v",
			status, http.StatusOK)
	}
}

func TestUpdateBookHandlerNotFound(t *testing.T) {
	initializeBooks()
	jsonStr := []byte(`{"Id":"4","Title":"Hello Hello","Author":"Jane Doe","Publisher":"Publish House","PublishDate":"10/23/2017","Rating":3,"Status":"Published"}`)
	request, err := http.NewRequest("PUT", "/book/5", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	request = mux.SetURLVars(request, map[string]string{"id": "5"})

	rec := httptest.NewRecorder()
	handler := http.HandlerFunc(updateBookHandler)
	handler.ServeHTTP(rec, request)

	if status := rec.Code; status != http.StatusNotFound {
		t.Errorf("Wrong status code. Returned: %v Expected: %v",
			status, http.StatusOK)
	}
}

func TestUpdateBookHandlerUpdatesBook(t *testing.T) {
	initializeBooks()
	jsonStr := []byte(`{"Id":"4","Title":"Hello Hello","Author":"Jane Doe","Publisher":"Publish House","PublishDate":"10/23/2017","Rating":3,"Status":"Published"}`)
	request, err := http.NewRequest("PUT", "/book/1", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	request = mux.SetURLVars(request, map[string]string{"id": "1"})

	rec := httptest.NewRecorder()
	handler := http.HandlerFunc(updateBookHandler)
	handler.ServeHTTP(rec, request)

	if status := rec.Code; status != http.StatusCreated {
		t.Errorf("Wrong status code. Returned: %v Expected: %v",
			status, http.StatusOK)
	}
}

func TestDeleteBookHandlerBadParameter(t *testing.T) {
	initializeBooks()
	request, err := http.NewRequest("DELETE", "/book/abc", nil)
	if err != nil {
		t.Fatal(err)
	}

	request = mux.SetURLVars(request, map[string]string{"id": "abc"})

	rec := httptest.NewRecorder()
	handler := http.HandlerFunc(deleteBookHandler)
	handler.ServeHTTP(rec, request)

	if status := rec.Code; status != http.StatusBadRequest {
		t.Errorf("Wrong status code. Returned: %v Expected: %v",
			status, http.StatusOK)
	}
}

func TestDeleteBookHandlerNotFound(t *testing.T) {
	initializeBooks()
	request, err := http.NewRequest("DELETE", "/book/4", nil)
	if err != nil {
		t.Fatal(err)
	}

	request = mux.SetURLVars(request, map[string]string{"id": "4"})

	rec := httptest.NewRecorder()
	handler := http.HandlerFunc(deleteBookHandler)
	handler.ServeHTTP(rec, request)

	if status := rec.Code; status != http.StatusNotFound {
		t.Errorf("Wrong status code. Returned: %v Expected: %v",
			status, http.StatusOK)
	}
}

func TestDeleteBookHandlerReturnsBook(t *testing.T) {
	initializeBooks()
	request, err := http.NewRequest("DELETE", "/book", nil)
	if err != nil {
		t.Fatal(err)
	}

	request = mux.SetURLVars(request, map[string]string{"id": "1"})

	rec := httptest.NewRecorder()
	handler := http.HandlerFunc(deleteBookHandler)
	handler.ServeHTTP(rec, request)

	if status := rec.Code; status != http.StatusOK {
		t.Errorf("Wrong status code. Returned: %v Expected: %v",
			status, http.StatusOK)
	}
}
