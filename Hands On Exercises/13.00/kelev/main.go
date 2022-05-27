// Package kelev provides ageing conversion for between dogs and humans.
package kelev

// Years converts human years to dog years.
func Years(human int) int {
	return human * 7
}

//YearsTwo converts human years to dog years.
func YearsTwo(human int) int {
	count := 0
	for i := 0; i < human; i++ {
		count += 7
	}
	return count
}
