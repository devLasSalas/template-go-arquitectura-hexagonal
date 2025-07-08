package api_adaptador_servidor

import (
	"fmt"
	"genesis/pos/reportes_pos/dominio/constantes"
	api_routes "genesis/pos/reportes_pos/presentacion/api/gin/routes"
	"log"
)

func Start() error {

	servidor, err := api_routes.GinConfig()

	if err != nil {
		log.Fatal(err)
		return err
	}

	ip := constantes.HOST_IP
	port := constantes.HOST_PORT
	host := ip + ":" + port
	err = servidor.Run(host)
	if err != nil {
		log.Fatal(err)
		return err

	}
	fmt.Println("SERVIDOR CORRIENDO: ", host)
	return nil

}
