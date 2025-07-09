package dominio_adaptadores_clientes_db

type IClienteDB interface {
	Select(query string, argumentos []any) ([][]interface{}, error)
	Exec(query string, argumentos []any) ([][]interface{}, error)
}
