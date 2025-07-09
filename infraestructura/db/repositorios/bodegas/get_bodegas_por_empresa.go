package infraestructura_repos_bodegas

import (
	"encoding/json"
	dominio_adaptadores_clientes_db "genesis/pos/reportes_pos/dominio/adaptadores/clientes/db"
	entidades_bodegas "genesis/pos/reportes_pos/dominio/entidades/bodegas"
	"log"
)

type GetBodegasPorEmpresaRepository struct {
	Cliente dominio_adaptadores_clientes_db.IClienteDB
}

func (GBPER *GetBodegasPorEmpresaRepository) Get(id_empresa int) (*[]entidades_bodegas.DataResponse, error) {

	query := "select * from tienda.fnc_obtener_bodegas_por_empresa($1);"
	resultados, err := GBPER.Cliente.Exec(query, []any{id_empresa})
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var resultadoGeneral *[]entidades_bodegas.DataResponse
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
