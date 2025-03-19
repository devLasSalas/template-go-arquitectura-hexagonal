package adaptador_servidor

import (
	"fmt"
	"genesis/pos/reportes_pos/dominio/constantes"
	config "genesis/pos/reportes_pos/presentacion/api/gin/config/routes"
	"log"
)

func Start() error {

	servidor, err := config.GinConfig()

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
