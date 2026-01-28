package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main(){
	//jeito bonito de pegar env do ambiente, so que se nao tiver, ele retornar 0
	// fmt.Println(os.Getenv("LOCAL_SERVER"))

	//jeito para analisar se tem coisa dentro
	// val, exist := os.LookupEnv("LOCAL_SERVER")
	// if !exist {
	// 	fmt.Println("A váriavel de ambiente não esta setado")
	// } else {
	// 	fmt.Println(val)
	// }

	//agora usando uma biblioteca do golangadssads
	//aqui podemos colocar as variaveis de ambiente em um arquivo .env
	err := godotenv.Load()
	if err != nil {panic(err)}

	//agora podemos usar o getenv normalmente

	if val, aux := os.LookupEnv("PORT"); !aux {
		panic("Valor não encontrado")
	} else {
		fmt.Println(val)
	}

}