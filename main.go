package main

//TODO makefileか、shellにgo run コマンドをまとめる
func main() {

	// ExcecuteTest1()

	// if err := run(); err != nil {
	// 	fmt.Println(err)
	// 	var c *CustomError
	// 	if errors.As(err, &c) {
	// 		fmt.Println(111111)
	// 	}
	// }

	// err1 := msgError("Error")
	// fmt.Println("Normal Error", err1)

	// err2 := WrapError("Test Error Message")
	// fmt.Println("[Wrapping Error]", err2)

	// var strErr *StrError
	// if errors.As(err2, &strErr) {
	// 	fmt.Printf("[Failed] %s\n", strErr)
	// }

	CallTestError3()
}
