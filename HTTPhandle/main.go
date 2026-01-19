package main

import "net/http"




type HandlerHello struct{}
type HandlerWorld struct{}

func (HandlerHello) ServeHTTP(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hello"))
}
func (HandlerWorld) ServeHTTP(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("World"))
}

func main(){
	hello := HandlerHello{}
	world := HandlerWorld{}	

	// aqui estamos tratando dois pradros de url diferente com dois handlers diferentes
	//sendo essa funcao do handler, tratar diferentes urls (sendo cada um para tratar de processar cada requisicao especifica)
	http.Handle("/hello", hello)
	http.Handle("/world", world)

	//tem que colocar nil par gente fazer com que os handlers anteriores sejam usados
	http.ListenAndServe(":5000", nil)
}