package slice

// Remove - Rimuovo un elemento ad un certo indice dello slice
func Remove(slice []interface{}, index int) []interface{} {
	return append(slice[:index], slice[index+1:]...)
}

// RemoveString - Rimuove un elemento ad un certo indice da un slice di stringhe
func RemoveString(slice []string, index int) []string {
	return append(slice[:index], slice[index+1:]...)
}
