package aplicacion_servicios_ctg_entrada_salida

import (
	dominio_casosusos_categorias_entradas_salidas "genesis/pos/reportes_pos/dominio/casosusos/categorias_entradas_salida"
	entidades_categoria_entrada_salida "genesis/pos/reportes_pos/dominio/entidades/categoria_entrada_salida"
)

type GetCategoriasEntradaSalidaService struct {
	GetCategoriasEntradaSalidaUseCase dominio_casosusos_categorias_entradas_salidas.GetCategoriasEntradaSalidaUseCase
}

func (GCES *GetCategoriasEntradaSalidaService) Execute() (*[]entidades_categoria_entrada_salida.DataResponse, error) {
	datos, err := GCES.GetCategoriasEntradaSalidaUseCase.Execute()
	if err != nil {
		return nil, err
	}

	return datos, nil
}
