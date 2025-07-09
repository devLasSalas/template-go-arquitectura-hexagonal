package dominio_casosusos_categorias_entradas_salidas

import entidades_categoria_entrada_salida "genesis/pos/reportes_pos/dominio/entidades/categoria_entrada_salida"

type GetCategoriasEntradaSalidaUseCase interface {
	Execute() (*[]entidades_categoria_entrada_salida.DataResponse, error)
}
