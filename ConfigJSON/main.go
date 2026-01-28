package main

import (
	"encoding/json"
	"fmt"
	"os"
)

//uma struct pra gente pegar o que tiver dentro do json
type Config struct {
	Server struct {
		Port int
		Host string
		StaticDir string
	}
}

func main(){
	//aqui estamos abrindo o arquivo
	file, err := os.Open("config.json")
	if err != nil {panic(err)}
	defer file.Close()

	//inicializando a struct config para que possamos pegar que esta no arquivo json
	config := Config{Server: struct{Port int; Host string; StaticDir string}{}}

	if err := json.NewDecoder(file).Decode(&config); err != nil {panic(err)}

	fmt.Printf("StaticDir : %s --- Port: %d --- Host: %s\n", config.Server.StaticDir, config.Server.Port, config.Server.Host)
}