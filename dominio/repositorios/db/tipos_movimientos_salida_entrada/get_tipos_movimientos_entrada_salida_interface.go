package dominio_repositorios_tipos_mov_entrada_salida

import (
	entidades_tipos_movimientos_entrada_salida "genesis/pos/reportes_pos/dominio/entidades/tipos_movimientos_entrada_salida"
)

type GetTiposMovimientosEntradaSalidaInterface interface {
	Get(id_categoria int) (*[]entidades_tipos_movimientos_entrada_salida.DataResponse, error)
}
