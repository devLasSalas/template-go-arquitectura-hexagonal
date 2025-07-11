package main

import (
	"genesis/pos/reportes_pos/dominio/constantes"
	api_adaptador_servidor "genesis/pos/reportes_pos/presentacion/api/gin/adaptador"
	"genesis/pos/reportes_pos/presentacion/contenedor"
	"log"
	"os"
)

func main() {

	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.SetOutput(os.Stdout)

	// config_logger.Init()

	if err := contenedor.InicializarContenedor(); err != nil {
		log.Fatal(err)
		panic(err)
	}

	log.Println("INICIANDO SERVICIO REPORTES POS: ", constantes.HOST_IP, ":", constantes.HOST_PORT)
	if err := api_adaptador_servidor.Start(); err != nil {
		log.Fatal(err)
		panic(err)
	}
}
