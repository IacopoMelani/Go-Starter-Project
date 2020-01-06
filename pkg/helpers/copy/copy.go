package copy

// Float32 - Restituisce l'indirizzo di un valore di tipo float32
func Float32(val float32) *float32 {
	return &val
}

// Float64 - Restituiscce l'indirizzo di un valore di tipo float64
func Float64(val float64) *float64 {
	return &val
}

// Int - Restituisce l'indirizzo di valore di tipo int
func Int(val int) *int {
	return &val
}

// Int64 - Restituisce l'indirizzo di un valore di tipo int64
func Int64(val int64) *int64 {
	return &val
}

// String - Restituisce l'indirizzo di un valore di tipo string
func String(val string) *string {
	return &val
}
