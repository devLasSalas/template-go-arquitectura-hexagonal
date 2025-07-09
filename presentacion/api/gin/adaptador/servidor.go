package api_adaptador_servidor

import (
	api_routes "genesis/pos/reportes_pos/presentacion/api/gin/routes"
	"log"
	"os"
)

func Start() error {

	servidor, err := api_routes.GinConfig()

	if err != nil {
		log.Fatal(err)
		return err
	}

	port := ":" + os.Getenv("SERVER_PORT")
	dev := os.Getenv("DEV") == "true"

	if dev {
		log.Println("INICIANDO SERVIDOR SIN SSL en el puerto", port)
		if err := servidor.Run(port); err != nil {
			log.Fatal(err)
		}
	} else {
		certFile := os.Getenv("SSL_CERT_FILE")
		keyFile := os.Getenv("SSL_KEY_FILE")

		log.Println("INICIANDO SERVIDOR CON SSL en el puerto")
		if err := servidor.RunTLS(port, certFile, keyFile); err != nil {
			log.Fatal(err)
		}
	}

	return nil
}
