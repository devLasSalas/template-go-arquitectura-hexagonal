package infraestructura_repos_categorias_entrada_salida

import (
	"encoding/json"
	dominio_adaptadores_clientes_db "genesis/pos/reportes_pos/dominio/adaptadores/clientes/db"
	entidades_categoria_entrada_salida "genesis/pos/reportes_pos/dominio/entidades/categoria_entrada_salida"
	"log"
)

type GetCategoriesRepository struct {
	Cliente dominio_adaptadores_clientes_db.IClienteDB
}

func (CC *GetCategoriesRepository) Get() (*[]entidades_categoria_entrada_salida.DataResponse, error) {

	query := "select * from tienda.fnc_obtener_categorias_entrada_salida();"
	resultados, err := CC.Cliente.Exec(query, []any{})
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var resultadoGeneral *[]entidades_categoria_entrada_salida.DataResponse
	jsonResultados, err := json.Marshal(resultados[0][0])
	if err != nil {
		log.Println("❌ Error al serializar resultados:", err)
		return nil, err
	}

	if err := json.Unmarshal([]byte(jsonResultados), &resultadoGeneral); err != nil {
		log.Println("❌ Error al deserializar JSON general:", err)
		return nil, err
	}

	return resultadoGeneral, nil

}
