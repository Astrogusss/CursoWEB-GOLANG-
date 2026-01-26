package main

import (
	"html/template"
	"os"
)

func main(){
	// nao pode passar caminho de arquivo relativo, so pode ser com base na raiz do projeto
	// OBS: o nome do template ja é o proprio nome do arquivo
	temp, err := template.ParseFiles("ParseArqHTML/index.html")
	if err != nil {panic(err)}
	

	err = temp.Execute(os.Stdout, "Gustavo")
	if err != nil {(panic(err))}
}