package contenedor

import (
	aplicacion_casosusos_bodegas "genesis/pos/reportes_pos/aplicacion/casosusos/bodegas"
	aplicacion_casosusos_ctg_entrada_salida "genesis/pos/reportes_pos/aplicacion/casosusos/categorias_entradas_salida"
	aplicacion_casosusos_tiposmov_entrada_salida "genesis/pos/reportes_pos/aplicacion/casosusos/tipos_movimientos_salida_entrada"
	aplicacion_servicios_bodegas "genesis/pos/reportes_pos/aplicacion/servicios/bodegas"
	aplicacion_servicios_ctg_entrada_salida "genesis/pos/reportes_pos/aplicacion/servicios/categorias_entradas_salida"
	aplicacion_servicios_tiposmov_entradas_salida "genesis/pos/reportes_pos/aplicacion/servicios/tipos_movimientos_salida_entrada"
	dominio_adaptadores_clientes_http "genesis/pos/reportes_pos/dominio/adaptadores/clientes/http"
	"genesis/pos/reportes_pos/dominio/adaptadores/mapeadores"
	dominio_casosusos_bodegas "genesis/pos/reportes_pos/dominio/casosusos/bodegas"
	dominio_casosusos_categorias_entradas_salidas "genesis/pos/reportes_pos/dominio/casosusos/categorias_entradas_salida"
	dominio_casosusos_tipos_mov_entradas_salidas "genesis/pos/reportes_pos/dominio/casosusos/tipos_movimientos_salida_entrada"
	dominio_repositorios "genesis/pos/reportes_pos/dominio/repositorios/db"
	dominio_repositorios_bodegas "genesis/pos/reportes_pos/dominio/repositorios/db/bodegas"
	dominio_repositorios_categorias_entrada_salida "genesis/pos/reportes_pos/dominio/repositorios/db/categorias_entrada_salida"
	dominio_repositorios_tipos_mov_entrada_salida "genesis/pos/reportes_pos/dominio/repositorios/db/tipos_movimientos_salida_entrada"
	infraestructura_db_cliente "genesis/pos/reportes_pos/infraestructura/db/cliente"
	infraestructura_repos_bodegas "genesis/pos/reportes_pos/infraestructura/db/repositorios/bodegas"
	infraestructura_repos_categorias_entrada_salida "genesis/pos/reportes_pos/infraestructura/db/repositorios/categorias_entrada_salida"
	infraestructura_repos_tipos_movimientos_salida_entrada "genesis/pos/reportes_pos/infraestructura/db/repositorios/tipos_movimientos_salida_entrada"
	cliente_infrastruc "genesis/pos/reportes_pos/infraestructura/http/cliente"
	mapper_infraestructura "genesis/pos/reportes_pos/infraestructura/mapper"
	"log"
	"os"
)

// SERVICIOS
var GetCategoriaEntradaSalidaService *aplicacion_servicios_ctg_entrada_salida.GetCategoriasEntradaSalidaService
var GetTiposMovEntradaSalidaService *aplicacion_servicios_tiposmov_entradas_salida.GetTiposMovEntradaSalidaService
var GetBodegasPorEmpresaService *aplicacion_servicios_bodegas.GetBodegasPorEmpresaService

// CASOS DE USO
var GetCategoriasEntradaSalidaUseCase dominio_casosusos_categorias_entradas_salidas.GetCategoriasEntradaSalidaUseCase
var GetTiposMovEntradaSalidaUseCase dominio_casosusos_tipos_mov_entradas_salidas.GetTiposMovEntradaSalidaUseCase
var GetBodegasPorEmpresaUseCase dominio_casosusos_bodegas.GetBodegasPorEmpresaUseCase

// REPOSITORIOS
var GetCategoriasEntradaSalidaRepo dominio_repositorios_categorias_entrada_salida.GetCategoriasEntradaSalidaInterface
var GetTiposMovEntradaSalidaRepo dominio_repositorios_tipos_mov_entrada_salida.GetTiposMovimientosEntradaSalidaInterface
var GetBodegasPorEmpresaRepo dominio_repositorios_bodegas.GetBodegasPorEmpresaInterface

// GENERALES
var RecuperarWacher dominio_repositorios.IRecuperarWacher

// Clientes
var ClienteHttp dominio_adaptadores_clientes_http.IClienteHttp

//MAPPER

var MapPosDatos mapeadores.MapearDatosPos

func InicializarContenedor() error {

	//Mapper
	MapPosDatos = &mapper_infraestructura.MapaerPosData{}

	//CLIENTES
	ClienteDB, err := infraestructura_db_cliente.InicializarCliente(os.Getenv("DATABASE_CONNECTION"))
	if err != nil {
		log.Fatal("Error inicializar bd cliente", err)
		return err

	}

	ClienteHttp, err = cliente_infrastruc.InicializarCliente()
	if err != nil {
		log.Fatal(err)
		return err

	}

	// REPOSITORIOS
	GetCategoriasEntradaSalidaRepo = &infraestructura_repos_categorias_entrada_salida.GetCategoriesRepository{Cliente: ClienteDB}
	GetTiposMovEntradaSalidaRepo = &infraestructura_repos_tipos_movimientos_salida_entrada.GetTiposMovimientosEntradaSalida{Cliente: ClienteDB}
	GetBodegasPorEmpresaRepo = &infraestructura_repos_bodegas.GetBodegasPorEmpresaRepository{Cliente: ClienteDB}

	// CASOS DE USOS
	GetCategoriasEntradaSalidaUseCase = &aplicacion_casosusos_ctg_entrada_salida.GetCategoriasEntradaSalidaUseCase{Repositorio: GetCategoriasEntradaSalidaRepo}
	GetTiposMovEntradaSalidaUseCase = &aplicacion_casosusos_tiposmov_entrada_salida.GetTiposMovimientosEntradaSalidaUseCase{Repositorio: GetTiposMovEntradaSalidaRepo}
	GetBodegasPorEmpresaUseCase = &aplicacion_casosusos_bodegas.GetBodegasPorEmpresaUseCase{Repositorio: GetBodegasPorEmpresaRepo}

	// SERVICIOS
	GetCategoriaEntradaSalidaService = &aplicacion_servicios_ctg_entrada_salida.GetCategoriasEntradaSalidaService{GetCategoriasEntradaSalidaUseCase: GetCategoriasEntradaSalidaUseCase}
	GetTiposMovEntradaSalidaService = &aplicacion_servicios_tiposmov_entradas_salida.GetTiposMovEntradaSalidaService{GetTiposMovEntradaSalidaUseCase: GetTiposMovEntradaSalidaUseCase}
	GetBodegasPorEmpresaService = &aplicacion_servicios_bodegas.GetBodegasPorEmpresaService{GetBodegasPorEmpresaUseCase: GetBodegasPorEmpresaUseCase}

	return nil
}
