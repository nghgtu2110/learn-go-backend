package stringPkg

func IsEvenLengthString(source string) bool {
	var length int = len(source)
	// if length % 2 == 0 {
	// 	return true
	// } else {
	// 	return false
	// }

	return length%2 == 0
}
