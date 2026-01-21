package main

import (
	"fmt"
	"net/http"
)

func main(){
	mux := http.NewServeMux()

	mux.HandleFunc("/notas", func(w http.ResponseWriter, r *http.Request){
		if r.Method != http.MethodPost {
			//          \/ permite que setamos o header
			w.Header().Set("Allow", "GET")
			http.Error(w, "Metodo não suportado", 405)
			return
		}
		// fmt.Fprint(w, "<h1> Bom dia pessoal, aqui ta foda </h1>") ---> aqui o content type é html
		fmt.Fprint(w, "Bom dia pessoal, aqui ta foda") // aqui o content type é plain
	})

	// alterando para emitir um cabeçalho do tipo json
	mux.HandleFunc("/notas/unifei", func(w http.ResponseWriter, r *http.Request){
		if r.Method == http.MethodPost {
			w.Header().Set("Allow", "GET")
			http.Error(w, "Método não suportado", 405)
			return
		}
		// agora estamos dizendo que a resposta esta sendo enviada no formato json
		w.Header().Set("Content-Type", "aplication/json")

		// podemos tambem deletar alguma informação do header, mas primeiro temos que setar eles para nil
		w.Header()["Date"] = nil

		//podemos tambem adicionar informações ao header
		w.Header().Add("SouBonito", "Óbvio")
		fmt.Fprint(w, "{id : 10}")
	})

	http.ListenAndServe(":5000", mux)
}