package aplicacion_casosusos_tiposmov_entrada_salida

import (
	entidades_tipos_movimientos_entrada_salida "genesis/pos/reportes_pos/dominio/entidades/tipos_movimientos_entrada_salida"
	dominio_repositorios_tipos_mov_entrada_salida "genesis/pos/reportes_pos/dominio/repositorios/db/tipos_movimientos_salida_entrada"
	"log"
)

type GetTiposMovimientosEntradaSalidaUseCase struct {
	Repositorio dominio_repositorios_tipos_mov_entrada_salida.GetTiposMovimientosEntradaSalidaInterface
}

func (GTMESU *GetTiposMovimientosEntradaSalidaUseCase) Execute(id_categoria int) (*[]entidades_tipos_movimientos_entrada_salida.DataResponse, error) {
	log.Println("Entro aqui : 1")
	return GTMESU.Repositorio.Get(id_categoria)
}
