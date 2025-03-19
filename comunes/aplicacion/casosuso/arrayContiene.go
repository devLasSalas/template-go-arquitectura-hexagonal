package casosuso_comunes

func Contiene(val any, lista []any) bool {
	for _, valor := range lista {
		if valor == val {
			return true
		}
	}
	return false
}
