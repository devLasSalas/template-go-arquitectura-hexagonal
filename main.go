package main

import (
	api_adaptador_servidor "genesis/pos/reportes_pos/presentacion/api/gin/adaptador"
	"genesis/pos/reportes_pos/presentacion/contenedor"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.SetOutput(os.Stdout)
	log.Println("¡MS MOVIMIENTO INVENTARIO!")

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	if err := contenedor.InicializarContenedor(); err != nil {
		log.Fatal(err)
		panic(err)
	}

	log.Println("INICIANDO SERVICIO REPORTES POS")
	if err := api_adaptador_servidor.Start(); err != nil {
		log.Fatal(err)
		panic(err)
	}
}
