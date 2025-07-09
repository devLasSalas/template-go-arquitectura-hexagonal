package dominio_repositorios_categorias_entrada_salida

import entidades_categoria_entrada_salida "genesis/pos/reportes_pos/dominio/entidades/categoria_entrada_salida"

type GetCategoriasEntradaSalidaInterface interface {
	Get() (*[]entidades_categoria_entrada_salida.DataResponse, error)
}
