package main

//export testPanic
func testPanic() bool {
	panic("test panic message 123 321!")
	return true
}

func main() {
}
