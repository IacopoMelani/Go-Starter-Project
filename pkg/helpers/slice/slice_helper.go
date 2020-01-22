package slice

// checkIfIndexExist - Resitusice se l'indice per lo slice fornito esiste
func checkIfIndexExist(slice []interface{}, index int) bool {
	return len(slice) > index
}

// Remove - Rimuovo un elemento ad un certo indice dello slice
func Remove(slice []interface{}, index int) []interface{} {
	if ok := checkIfIndexExist(slice, index); ok {
		return append(slice[:index], slice[index+1:]...)
	}
	return slice
}

// RemoveString - Rimuove un elemento ad un certo indice da un slice di stringhe
func RemoveString(slice []string, index int) []string {
	if ok := checkIfIndexExist(StringToInterface(slice), index); ok {
		return append(slice[:index], slice[index+1:]...)
	}
	return slice
}

// StringToInterface - Trasformara uno slice di stringhe in uno slice di interface{}
func StringToInterface(slice []string) []interface{} {

	newSlice := make([]interface{}, len(slice))

	for i, v := range slice {
		newSlice[i] = v
	}

	return newSlice
}
