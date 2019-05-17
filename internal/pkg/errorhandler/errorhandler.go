package errorhandler

import "fmt"

//Check just checks and prints the error. does not panic
func Check(e error) bool {
	if e != nil {
		fmt.Printf("\nError in module %s: ", e.Error())
		return true
	}
	return false
}

//Checkandpanic panics on error
func Checkandpanic(e error) {
	if e != nil {
		panic(e)
	}
}

//Checkandlog checks and prints the error along with module details. does not panic
func Checkandlog(e error, module string) bool {
	if e != nil {
		fmt.Printf("\nError in module %s: %s", module, e.Error())
		return true
	}
	return false
}
