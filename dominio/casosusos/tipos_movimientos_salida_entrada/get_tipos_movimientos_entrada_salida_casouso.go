package dominio_casosusos_tipos_mov_entradas_salidas

import (
	entidades_tipos_movimientos_entrada_salida "genesis/pos/reportes_pos/dominio/entidades/tipos_movimientos_entrada_salida"
)

type GetTiposMovEntradaSalidaUseCase interface {
	Execute(id_categoria int) (*[]entidades_tipos_movimientos_entrada_salida.DataResponse, error)
}
