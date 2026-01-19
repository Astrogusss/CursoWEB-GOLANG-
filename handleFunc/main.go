package main

import(
	"net/http"
)

func HandlerBravo(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Vai todo mundo tomar no cu"))
}

func main(){

	//podemos criar nossas proprias funcoes dentro do handlefunc sem precisar de serveHTTP
	//tambem podemos criar as funcoes fora ou dentro
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("Hello amigo"))
	})

	http.HandleFunc("/world", func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("World sjbksj"))
	})

	http.HandleFunc("/bravo", HandlerBravo)

	http.ListenAndServe(":5000", nil)

}