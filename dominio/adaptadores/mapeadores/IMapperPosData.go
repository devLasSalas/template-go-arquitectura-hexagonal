package mapeadores

import "genesis/pos/reportes_pos/dominio/entidades"

type MapearDatosPos interface {
	MapearA(pos []interface{}) (*entidades.PosData, error)
}
