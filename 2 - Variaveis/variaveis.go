package main

func main() {
	var variavel1 string = "Variável 1"
	variavel2 := "Variável 2"

	fmt.Println(variavel1)
	fmt.Println(variavel2)

	var (
		variavel3 string = "fgsadfsadf"
		variavel4 string = "dsfasfsadfsd"
	)

	fmt.Println(variavel3, variavel 4)

	variavel5, variavel6 := "Variável 5", "Variável 6"
	fmt.Println(variavel5, variavel6)

	const constante1 string = "Constante 1"
	fmt.Println(constante1)

	variavel5, variavel6 = variavel6, variavel5
	fmt.Println(variavel5, variavel6)
}