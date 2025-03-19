package contenedor

import (
	comunes_db_clientes "genesis/pos/reportes_pos/comunes/dominio/adaptadores/clientes/db"
	comunes_http_clientes "genesis/pos/reportes_pos/comunes/dominio/adaptadores/clientes/http"
	"genesis/pos/reportes_pos/comunes/dominio/adaptadores/mapeadores"
	"genesis/pos/reportes_pos/dominio/constantes"
	dominio_repositorios "genesis/pos/reportes_pos/dominio/repositorios/db"
	infraestructura_db_cliente "genesis/pos/reportes_pos/infraestructura/db/cliente"
	cliente_infrastruc "genesis/pos/reportes_pos/infraestructura/http/cliente"
	mapper_infraestructura "genesis/pos/reportes_pos/infraestructura/mapper"
	"log"
)

// SERVICIOS

// CASOS DE USO

// REPOSITORIOS

// GENERALES
var RecuperarWacher dominio_repositorios.IRecuperarWacher

//Clientes

var ClienteDB comunes_db_clientes.IClienteDB
var ClienteHttp comunes_http_clientes.IClienteHttp

//MAPPER

var MapPosDatos mapeadores.MapearDatosPos

func InicializarContenedor() error {

	//Mapper
	MapPosDatos = &mapper_infraestructura.MapaerPosData{}

	//CLIENTES
	_, err := infraestructura_db_cliente.InicializarCliente(constantes.DB_CON)
	if err != nil {
		log.Fatal(err)
		return err

	}

	_, err = cliente_infrastruc.InicializarCliente()
	if err != nil {
		log.Fatal(err)
		return err

	}

	// REPOSITORIOS

	// CASOS DE USOS

	// SERVICIOS

	return nil
}
