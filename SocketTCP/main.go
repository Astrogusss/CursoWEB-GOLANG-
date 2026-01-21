package main

import (
	//biblioteca para criar buffers
	"bufio"
	"fmt"
	"net"
)

func main(){
	// aqui ele retonar um listener e um erro, entao temos que tratar ele
	ls, err := net.Listen("tcp", ":5000")
	if err != nil{
		panic(err)
	}
	//sempre recomendado fechar o listener sempre que abrir
	defer ls.Close()

	for {
		//aqui ele retonar a conexao / estabelecendo a conexão com o cliente

		conexao, err := ls.Accept()
		fmt.Println("Conexão estabelecida")
		if err != nil {
			panic(err)
		}
		
		//aqui estamos tratando uma requisicao do cliente
		//tem que ser uma goroutine porque se outra requisicao chegar no servidor, sem a goroutine, ele teria que esperar ate que a primeira requisição seja retornada
		//com a goroutine, ela ja seria executada em concorrencia com as requisicoes chegadas
		go func(con net.Conn){
			//newReader receber como paramentro um reader, que nesse caso tem o mesmo papel da conexão (uma interface que tem um reader)
			//o \n se colocar para que quando chegar no final da linha, seja quebrada
			frase, err := bufio.NewReader(con).ReadString('\n')
			if err != nil {
				panic(err)
			}

			fmt.Printf("Dado recebido %s", frase)
			con.Write([]byte("Sua mensagem foi recebida seu cuzão\n"))
			con.Close()
			fmt.Printf("Sua conexão foi fechada\n")
		}(conexao)
	}
}