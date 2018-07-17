package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/fizzBuzz", fizzBuzzHandler).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func fizzBuzzHandler(w http.ResponseWriter, r *http.Request) {
	string1 := getParamOrDefault(r, "string1", "fizz")
	string2 := getParamOrDefault(r, "string2", "buzz")
	int1, err1 := strconv.Atoi(getParamOrDefault(r, "int1", "3"))
	int2, err2 := strconv.Atoi(getParamOrDefault(r, "int2", "5"))
	limit, err3 := strconv.Atoi(getParamOrDefault(r, "limit", "100"))

	if err1 != nil || err2 != nil || err3 != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("among expected integer parameters (int1, int2 and limit) some were not integers."))
	} else {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(FizzBuzz(string1, string2, int1, int2, limit))
	}
}

func getParamOrDefault(r *http.Request, name string, def string) string {
	param := r.URL.Query().Get(name)
	if param == "" {
		param = def
	}
	return param
}
