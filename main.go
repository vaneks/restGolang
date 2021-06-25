package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type Response struct {
	Success bool
	ErrCode string
	Value int
}
func add (w http.ResponseWriter, r *http.Request){
	a := r.URL.Query().Get("a")
	b := r.URL.Query().Get("b")
	ax, _ := strconv.Atoi(a)
	bx, _ := strconv.Atoi(b)
	result := ax + bx
	res := Response{Success: true, ErrCode: "",Value :result}
	jsonData, _ := json.Marshal(res)
	w.Write(jsonData)
}
func sub(w http.ResponseWriter, r *http.Request){
	a := r.URL.Query().Get("a")
	b := r.URL.Query().Get("b")
	ax, _ := strconv.Atoi(a)
	bx, _ := strconv.Atoi(b)
	result := ax - bx
	res := Response{Success: true, ErrCode: "",Value :result}
	jsonData, _ := json.Marshal(res)
	w.Write(jsonData)
}
func mul(w http.ResponseWriter, r *http.Request){
	a := r.URL.Query().Get("a")
	b := r.URL.Query().Get("b")
	ax, _ := strconv.Atoi(a)
	bx, _ := strconv.Atoi(b)
	result := ax * bx
	res := Response{Success: true, ErrCode: "",Value :result}
	jsonData, _ := json.Marshal(res)
	w.Write(jsonData)
}
func div(w http.ResponseWriter, r *http.Request) {
	a := r.URL.Query().Get("a")
	b := r.URL.Query().Get("b")
	ax, _ := strconv.Atoi(a)
	bx, _ := strconv.Atoi(b)
	if bx == 0 {
		res := Response{Success: false, ErrCode: "Деление на 0",Value :0}
		jsonData, _ := json.Marshal(res)
		w.Write(jsonData)
	} else {
		result := ax / bx
		res := Response{Success: true, ErrCode: "", Value: result}
		jsonData, _ := json.Marshal(res)
		w.Write(jsonData)
	}
}
func handleRequests() {
	http.HandleFunc("/add", add)
	http.HandleFunc("/sub", sub)
	http.HandleFunc("/mul", mul)
	http.HandleFunc("/div", div)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	fmt.Println("Server listening!")
	handleRequests()
}