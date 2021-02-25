package main

import (
	"encoding/json"
	"errors"
	"fib-api/fib"
	"fib-api/response"
	"math/big"
	"net/http"
	"strconv"
)

func Fibo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	resp := response.New()
	// First Fibonacchi index
	firstNumber := r.URL.Query().Get("x")

	// Last Fibonacci index
	lastNumber := r.URL.Query().Get("y")

	// Convert numbers to indexes
	lastIndex, err := strconv.ParseInt(lastNumber, 10, 64)
	if check(err, w, resp, "incorrent data type", 100) {
		return
	}
	firstIndex, err := strconv.ParseInt(firstNumber,10,64)
	if check(err, w, resp, "incorrect data type", 101) {
		return
	}
	// Validate indices
	if firstIndex >= lastIndex {
		if check(errors.New("incorrect indices"), w, resp,"",102) {
			return
		}
	}

	fibList := make([]*big.Int, 0, 0)
	for i := 0; i <= int(lastIndex); i++ {
		fibs := fib.FibonacciBig(uint(i))
		fibList = append(fibList, fibs)
	}
	index := int(firstIndex)
	fibList = fibList[index:]

	content, err := json.Marshal(struct {
		StatusResponse *response.StatusResponse `json:"status_response"`
		Fibonacci      []*big.Int              `json:"fibonacci"`
	}{resp, fibList})
	if check(err, w, resp, "error marshalling response",103) {
		return
	}
	w.Write(content)
	return
}

func check(err error, w http.ResponseWriter, resp *response.StatusResponse, message string, code int) bool {
	if err != nil {
		resp.Status = false
		if len(message) == 0 {
			resp.ResponseMessage = err.Error()
		} else {
			resp.ResponseMessage = message + ": " + err.Error()
		}
		resp.ResponseInternalCode = code
		w.Write(resp.GetJSON())
		return true
	}
	return false
}
