package infraestructura_repos_tipos_movimientos_salida_entrada

import (
	"encoding/json"
	dominio_adaptadores_clientes_db "genesis/pos/reportes_pos/dominio/adaptadores/clientes/db"
	entidades_tipos_movimientos_entrada_salida "genesis/pos/reportes_pos/dominio/entidades/tipos_movimientos_entrada_salida"
	"log"
)

type GetTiposMovimientosEntradaSalida struct {
	Cliente dominio_adaptadores_clientes_db.IClienteDB
}

func (GTMES *GetTiposMovimientosEntradaSalida) Get(id_categoria int) (*[]entidades_tipos_movimientos_entrada_salida.DataResponse, error) {
	query := "select * from tienda.fnc_obtener_tipo_movimiento_entrada_salida_por_categoria($1);"
	resultados, err := GTMES.Cliente.Exec(query, []any{id_categoria})
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var resultadoGeneral *[]entidades_tipos_movimientos_entrada_salida.DataResponse
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
