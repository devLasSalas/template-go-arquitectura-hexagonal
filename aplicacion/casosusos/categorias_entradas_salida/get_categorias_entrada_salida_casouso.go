package aplicacion_casosusos_ctg_entrada_salida

import (
	entidades_categoria_entrada_salida "genesis/pos/reportes_pos/dominio/entidades/categoria_entrada_salida"
	dominio_repositorios_categorias_entrada_salida "genesis/pos/reportes_pos/dominio/repositorios/db/categorias_entrada_salida"
)

type GetCategoriasEntradaSalidaUseCase struct {
	Repositorio dominio_repositorios_categorias_entrada_salida.GetCategoriasEntradaSalidaInterface
}

func (GCESU *GetCategoriasEntradaSalidaUseCase) Execute() (*[]entidades_categoria_entrada_salida.DataResponse, error) {
	return GCESU.Repositorio.Get()
}
