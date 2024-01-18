package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	dictionary "projet1/directory/directories"
)

func TestAddHandler(t *testing.T) {
	testDictionary := dictionary.NewDictionary()

	entry := dictionary.Entry{Word: "test", Definition: "testing"}
	jsonEntry, _ := json.Marshal(entry)
	req := httptest.NewRequest("POST", "/add", bytes.NewReader(jsonEntry))
	req.Header.Set("Content-Type", "application/json")

	res := httptest.NewRecorder()

	AddHandler(testDictionary)(res, req)

	if res.Code != http.StatusCreated {
		t.Errorf("Expected status code %d, got %d", http.StatusCreated, res.Code)
	}

	definition, err := testDictionary.Get("test")
	if err != nil || definition != "testing" {
		t.Errorf("Expected dictionary entry 'test' with definition 'testing', got '%s'", definition)
	}
}

func TestGetHandler(t *testing.T) {
	testDictionary := dictionary.NewDictionary()
	testDictionary.Add("test", "testing")

	req := httptest.NewRequest("GET", "/get/test", nil)

	res := httptest.NewRecorder()

	GetHandler(testDictionary)(res, req)

	if res.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, res.Code)
	}

	expectedJSON := `{"word":"test","definition":"testing"}`
	if res.Body.String() != expectedJSON {
		t.Errorf("Expected JSON response %s, got %s", expectedJSON, res.Body.String())
	}
}

func TestRemoveHandler(t *testing.T) {
	testDictionary := dictionary.NewDictionary()
	testDictionary.Add("test", "testing")

	req := httptest.NewRequest("DELETE", "/remove/test", nil)

	res := httptest.NewRecorder()

	RemoveHandler(testDictionary)(res, req)

	if res.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, res.Code)
	}

	_, err := testDictionary.Get("test")
	if err == nil {
		t.Errorf("Expected dictionary entry 'test' to be removed, but it still exists")
	}
}

func TestListHandler(t *testing.T) {

	testDictionary := dictionary.NewDictionary()
	testDictionary.Add("apple", "fruit")
	testDictionary.Add("banana", "fruit")
	testDictionary.Add("carrot", "vegetable")

	req := httptest.NewRequest("GET", "/list", nil)

	res := httptest.NewRecorder()

	ListHandler(testDictionary)(res, req)

	if res.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, res.Code)
	}

	expectedResponse := "apple: fruit\nbanana: fruit\ncarrot: vegetable\n"
	if res.Body.String() != expectedResponse {
		t.Errorf("Expected response:\n%s\nGot:\n%s", expectedResponse, res.Body.String())
	}
}
