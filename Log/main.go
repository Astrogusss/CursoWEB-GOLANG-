package main

import (
	"log/slog"
	"os"
)

func main(){
	// // \/ tipos de log (degub, info, error, warn)
	// slog.Debug("debug message")
	// slog.Info("Info message")
	// slog.Error("error message")
	// slog.Warn("warn message")

	//mas a gente tem que criar um handler para tratar os logs

	//aqui desse jeito ele vai sair estruturado da forma de json, podemos usar o NewTextHandler, tambem configurarmos nossos handlerOptions
	aux := slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
		Level: slog.LevelError,
	})

	//são separados por leveis, que nos dizer quais o leveis que podem ser impressos no log, aqui só estamos dizendo que somente pode imprimir do tipo error
	// aux := slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
	// 	Level: slog.LevelError,     <-----
	// })


	log := slog.New(aux)

	//agora podemos criar nosso logs
	log.Error("que caganeira")
}