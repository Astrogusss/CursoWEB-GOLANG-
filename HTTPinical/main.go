package main

import(
	"net/http"
)

type MyHandler struct{}


//nesse caso, temos que o "w" é usando pra gente responder conforme o que foi passado no request "r"
func (MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Quero que todo mundo se foda"))
}

func main(){
	//escuta e serve - ela sobe o servidor com as configurações que passamos a ela
	handler := MyHandler{}
	http.ListenAndServe(":5000", handler)
}