package repositorios_infraestruture

import (
	dominio_adaptadores_clientes_db "genesis/pos/reportes_pos/dominio/adaptadores/clientes/db"
	"genesis/pos/reportes_pos/dominio/entidades"
)

type RecuperarWatcherParametors struct {
	Cliente dominio_adaptadores_clientes_db.IClienteDB
}

func (RWP *RecuperarWatcherParametors) Consultar(codigo string) (*entidades.ParametrosWatcher, error) {
	args := []any{codigo}
	respuesta, err := RWP.Cliente.Select("SELECT x.* FROM public.wacher_parametros x WHERE codigo = $1", args)

	if err != nil {
		return nil, err
	}

	parametro := &entidades.ParametrosWatcher{}
	for _, valor := range respuesta {
		parametro.Id = valor[0].(int64)
		parametro.Codigo = valor[1].(string)
		parametro.Tipo = valor[2].(int32)
		parametro.Valor = valor[3].(string)

	}

	return parametro, nil
}
