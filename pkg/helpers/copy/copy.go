package copy

// Float32 - Restituisce l'indirizzo di un valore di tipo float32
func Float32(val float32) *float32 {
	return &val
}

// Int - Restituisce l'indirizzo di valore di tipo int
func Int(val int) *int {
	return &val
}

// String - Restituisce l'indirizzo di un valore di tipo string
func String(val string) *string {
	return &val
}
