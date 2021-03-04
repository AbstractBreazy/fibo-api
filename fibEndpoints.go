package main

import (
	"encoding/json"
	"errors"
	"fibo-api/fib"
	"fibo-api/response"
	"math/big"
	"net/http"
	"strconv"
)

// Fibo endpoint calculates the Fibonacci Sequence
func Fibo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	resp := response.New()
	// First Fibonacchi index
	firstNumber := r.URL.Query().Get("x")

	// Last Fibonacci index
	lastNumber := r.URL.Query().Get("y")

	// Convert numbers to indexes
	lastIndex, err := strconv.ParseInt(lastNumber, 10, 64)
	if check(err, w, resp, "parameter `y` has an invalid data type", 100) {
		return
	}
	firstIndex, err := strconv.ParseInt(firstNumber, 10, 64)
	if check(err, w, resp, "parameter `x` has an invalid data type", 101) {
		return
	}

	// Validate indices
	// If x = y, then the answer will contain only one number from the sequence
	if firstIndex > lastIndex {
		if check(errors.New("parameter `x` cannot be more than parameter `y`"), w, resp, "failed to validtate indices", 102) {
			return
		}
	}

	// Calculating the fibonacci sequence based on the x,y
	fibList := make([]*big.Int, 0, 0)
	for i := 0; i <= int(lastIndex); i++ {
		fibs := fib.Fibonacci(uint(i))
		fibList = append(fibList, fibs)
	}
	fibList = fibList[int(firstIndex):]

	content, err := json.Marshal(struct {
		StatusResponse *response.StatusResponse `json:"status_response"`
		Fibonacci      []*big.Int               `json:"fibonacci"`
	}{resp, fibList})
	if check(err, w, resp, "error marshalling response", 103) {
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
