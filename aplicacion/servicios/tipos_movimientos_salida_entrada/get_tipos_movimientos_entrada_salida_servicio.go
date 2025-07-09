package aplicacion_servicios_tiposmov_entradas_salida

import (
	dominio_casosusos_tipos_mov_entradas_salidas "genesis/pos/reportes_pos/dominio/casosusos/tipos_movimientos_salida_entrada"
	entidades_tipos_movimientos_entrada_salida "genesis/pos/reportes_pos/dominio/entidades/tipos_movimientos_entrada_salida"
)

type GetTiposMovEntradaSalidaService struct {
	GetTiposMovEntradaSalidaUseCase dominio_casosusos_tipos_mov_entradas_salidas.GetTiposMovEntradaSalidaUseCase
}

func (GTMESS *GetTiposMovEntradaSalidaService) Execute(id_categoria int) (*[]entidades_tipos_movimientos_entrada_salida.DataResponse, error) {
	datos, err := GTMESS.GetTiposMovEntradaSalidaUseCase.Execute(id_categoria)
	if err != nil {
		return nil, err
	}

	return datos, nil
}
