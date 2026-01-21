package main

import (
	"net/http"
)

type Myhandler int
type MyHandler2 int

func (Myhandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// o tratamendo dos verbos esta atrelado ao request
	if r.Method != "POST"{
		//podemos enviar um header para que explique do erro 405 (method not allowed)
		//para que o cliente qual/quais verbos são permitidos
		w.Header().Set("Allow", "POST")

		//pra setar qual foi o erro
		w.WriteHeader(405) // so pode ser chamada uma vez por requisicao

		//importante retornar a request
		return
	}
	w.Write([]byte("Bom dia pessoal dfjdfj"))
}

func (MyHandler2) ServeHTTP(w http.ResponseWriter, r *http.Request){
	if r.Method != "POST" {
		//pode-se usar o http error, ao inves de ficar setando header e mandando o erro manualmente
		w.Header().Set("Allow", "POST")
		http.Error(w, "Método não suportado", http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("Boa tarde pessoal"))
}



func main(){
	mux := http.NewServeMux()

	var testando Myhandler
	var testando2 MyHandler2

	mux.Handle("/", testando)
	mux.Handle("/bomdia", testando2)
	
	http.ListenAndServe(":5000", mux)
}