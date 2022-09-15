package service

import (
	"encoding/json"
	"net/http"
)

func Flights(w http.ResponseWriter, r *http.Request) {
	var raw [][]string

	err := json.NewDecoder(r.Body).Decode(&raw)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("500 - Incorrect data"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(analyzeFlights(raw))
}

func analyzeFlights(raw [][]string) [2]string {
	exist := map[string]int{}
	result := [2]string{}
	for _, v := range raw {
		for e := range v {
			if _, ok := exist[v[e]]; !ok {
				exist[v[e]] = e
			} else {
				delete(exist, v[e])
			}
		}
	}

	for key, element := range exist {
		result[element] = key
	}
	return result
}
