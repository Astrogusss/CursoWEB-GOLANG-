package main

import (
	"fmt"
	"net/http"
)

func main(){
	mux := http.NewServeMux()

	mux.HandleFunc("/universidade", func(w http.ResponseWriter, r *http.Request){
		fmt.Println(r.URL.Path)
		fmt.Println(r.URL.RawQuery)

		//posso pegar um parametro que esta na query
		fmt.Printf(string(r.URL.Query().Get("id")))
		fmt.Fprint(w, "Conexão feita")
	})

	http.ListenAndServe(":5000", mux)
}