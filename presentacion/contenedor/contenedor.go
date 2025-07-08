package contenedor

import (
	dominio_adaptadores_clientes_db "genesis/pos/reportes_pos/dominio/adaptadores/clientes/db"
	dominio_adaptadores_clientes_http "genesis/pos/reportes_pos/dominio/adaptadores/clientes/http"
	"genesis/pos/reportes_pos/dominio/adaptadores/mapeadores"
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

var ClienteDB dominio_adaptadores_clientes_db.IClienteDB
var ClienteHttp dominio_adaptadores_clientes_http.IClienteHttp

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
