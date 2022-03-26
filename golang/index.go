package main

import (
	"os"
	"fmt"
	"log"
	"strconv"
)

func main() {

	fmt.Println("\n======================")
	fmt.Println(" KALKULATOR SEDERHANA ")
	fmt.Println("======================")

	// Call and change this operation mathemathic in here guys !
	operations("150", "30", "*")
	// operations("300", "30", "/")
	// operations("400", "30", "+")
}

func operations( number1 string, number2 string, signInput string)  {

	var operator string

	//Initial Value
	os.Setenv("FIRST_NUMBER", number1)
	// Get Value
	getNum1 := os.Getenv("FIRST_NUMBER")
	// Convert Value
	var convertNumb1, err1 = strconv.Atoi(getNum1)
	//Display Value
	fmt.Printf("\nFirst number: %d", convertNumb1)

	// Dan sama juga untuk nilai yang kedua
	os.Setenv("SECOND_NUMBER", number2)
	getNum2 := os.Getenv("SECOND_NUMBER")
	var convertNumb2, err2 = strconv.Atoi(getNum2)
	fmt.Printf("\n\nSecond number: %d", convertNumb2)

	// Check Operator Simbol
	os.Setenv("inputUser", signInput)
	operator = os.Getenv("inputUser")
	fmt.Printf("\n\nSelected Operator (+,-,/,%%,*) : %s\n", operator)

	if err1 == nil && err2 == nil {
		// Operator
		output := 0

		switch operator {
		case "+":
			output = convertNumb1 + convertNumb2
		case "-":
			output = convertNumb1 - convertNumb2
		case "/":
			output = convertNumb1 / convertNumb2
		case "*":
			output = convertNumb1 * convertNumb2
		case "%":
			output = convertNumb1 % convertNumb2
		default:
			fmt.Println("")
			error("Simbol Operator Tidak Ditemukan")
		}

		fmt.Print("\n/* Hasil Operasi */\n")
		fmt.Printf("%d %s %d = %d \n\n", convertNumb1, operator, convertNumb2, output)		
	} else if err1 != nil || err2 != nil {
		fmt.Println("")
		error("Perhitungan Gagal, First Number atau Second Number yang dimasukkan kurang tepat !")
	}
}

func error(value string) {
	log.Fatal("Error: " + value)	
}