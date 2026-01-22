package main

import (
	"fmt"
	"html/template"
	"os"
)

type aux struct {
	Nome string
	Idade int
}

func main(){
	//aqui estamos criando um template com o nome teste, e o codigo html tem que ser colocado dentro do parse para analisar (tem que ser um codigo html válido)
	temp, _ := template.New("teste").Parse("h1>Isso é um título </h1>")

	//														  \/ jeito para fazer colocar uma variavel, depois no execute, tem que colocar o valor da variavel
	temp1, _ := template.New("temp1").Parse("<h1>Isso é um {{ .Nome }} e idade {{ .Idade }} </h1>")

	auxStruct := aux{Nome: "Gustavo", Idade: 10}
	//									\/ o que sera colocado na variavel
	err1 := temp1.Execute(os.Stdout, auxStruct)
	if err1 != nil {panic(err1)}

	fmt.Println()

	//criando outro template somente para ver a funcionalidade do executeTemplate
	temp.New("teste1").Parse("<h1>Isso é um título1</h1>")

	// temp.Execute(os.Stdout, nil) --> desse jeito serve somente para quando a variavel tem somente um template
	//o parse do template, faz com que a engine nele possa analisar o texto html e retornar erros se houver alguma inconsistencia
	err1 = temp.ExecuteTemplate(os.Stdout, "teste", nil)
	if err1 != nil  {panic(err1)}

	fmt.Println()

	err1 = temp.ExecuteTemplate(os.Stdout, "teste1", nil)
	if err1 != nil {panic(err1)}

}