package main

import (
	"fmt"
	"net/http"
)

func main(){

	//ele ja faz o proprio roteamento
	mux := http.NewServeMux()


	//diferenca entre rota coringa e estática
	// quando se tem o barra no final, quer dizer que ela é coringa ex: /, /hello/, /sdfsdfd/sfsdfds/
	// o coringa ja é direcionado, como abaixo, a /universidade é estatica, so da pra acessar ela digitando na url do mesmo jeito
	// ja o / pode ser acessada como default se escrever a url que nao esta no roteamento, por exemplo: se vier a rota /peidoeletrico, vai ter um redirect para o barra
	//mas se tiver uma handle com /universidade/, se a url vier /universidade/sdfsdfgs, vai ter um redirect para o /universidade/

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprint(w, "top 5 universidades mais fodas")
	})

	mux.HandleFunc("/universidade", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprint(w, "quero que se foda todo mundo 1212",)
	})

	http.ListenAndServe(":5000", mux)


	// aqui estamos setando o servermux como padrão, sem ao menos nos setar,
	// fazendo com que se use variáveis globais, que pode causar impactos ao longo da execucao

	//qunado se instacia um serveMux, nao se pode usar o nil, porque nao estamos usando o defaultServeMux
	// http.ListenAndServe(":5000", nil)
}