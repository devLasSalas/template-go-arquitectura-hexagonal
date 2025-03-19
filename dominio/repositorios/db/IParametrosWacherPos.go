package dominio_repositorios

import "genesis/pos/reportes_pos/dominio/entidades"

type IRecuperarWacher interface {
	Consultar(codigo string) (*entidades.ParametrosWatcher, error)
}
