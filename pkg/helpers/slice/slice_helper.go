package slice

// checkIfIndexExist - Return if the index for the passed slice exists
func checkIfIndexExist(slice []interface{}, index int) bool {
	return len(slice) > index
}

// Remove - Removes an element from the slice at certain index, if the index not exists returns the slice unchanged
func Remove(slice []interface{}, index int) []interface{} {
	if ok := checkIfIndexExist(slice, index); ok {
		return append(slice[:index], slice[index+1:]...)
	}
	return slice
}

// RemoveString -  Removes an element from the slice of strings at certain index, if the index not exists returns the slice unchanged
func RemoveString(slice []string, index int) []string {
	if ok := checkIfIndexExist(StringToInterface(slice), index); ok {
		return append(slice[:index], slice[index+1:]...)
	}
	return slice
}

// StringToInterface - Transform a slice of interface to slice of string
func StringToInterface(slice []string) []interface{} {

	newSlice := make([]interface{}, len(slice))

	for i, v := range slice {
		newSlice[i] = v
	}

	return newSlice
}
