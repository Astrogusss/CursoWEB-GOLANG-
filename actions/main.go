package main

import (
	"html/template"
	"os"
)

func main(){
	//vamos definir alguns templates usando actions
	//pode colocar mais caminhos de html no parseFiles e definir os templates
	//cuidado com template corrente, o primeiro caminho que define os templates 
	t, err := template.ParseFiles("index.html")


	if err != nil {panic(err)}

	t.Execute(os.Stdout, nil)
	// outra maneira, poderia se executar com essa funcao
	// t.ExecuteTemplate(os.Stdout, "Pagrafo A", nil)
	
}