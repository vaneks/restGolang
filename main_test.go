package main

import (
	"encoding/json"
	"net/http/httptest"
	"testing"
)

func TestAdd(t *testing.T){
	res := Response{Success: true, ErrCode: "",Value : 3}
	r := httptest.NewRequest("GET", "http://localhost:8181/add?a=1&b=2", nil)
	w := httptest.NewRecorder()
	add(w, r)
	response := Response{}
	json.Unmarshal(w.Body.Bytes(), &response)
	if res != response {
		t.Errorf("Test Failed")
	}
}

func TestSub(t *testing.T){
	res := Response{Success: true, ErrCode: "",Value : 4}
	r := httptest.NewRequest("GET", "http://localhost:8181/sub?a=6&b=2", nil)
	w := httptest.NewRecorder()
	sub(w, r)
	response := Response{}
	json.Unmarshal(w.Body.Bytes(), &response)
	if res != response {
		t.Errorf("Test Failed")
	}
}

func TestMul(t *testing.T){
	res := Response{Success: true, ErrCode: "",Value : 12}
	r := httptest.NewRequest("GET", "http://localhost:8181/mul?a=6&b=2", nil)
	w := httptest.NewRecorder()
	mul(w, r)
	response := Response{}
	json.Unmarshal(w.Body.Bytes(), &response)
	if res != response {
		t.Errorf("Test Failed")
	}
}

func TestDiv(t *testing.T){
	res := Response{Success: true, ErrCode: "",Value : 3}
	r := httptest.NewRequest("GET", "http://localhost:8181/mul?a=6&b=2", nil)
	w := httptest.NewRecorder()
	div(w, r)
	response := Response{}
	json.Unmarshal(w.Body.Bytes(), &response)
	if res != response {
		t.Errorf("Test Failed")
	}
}