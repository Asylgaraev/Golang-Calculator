package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)    // read input from console
	text, _ := reader.ReadString('\n')     // I don't know what do this line
	raw := strings.Split(text, "\n")       // my way to del /n ^-^
	mathLine := strings.Split(raw[0], " ") // create list with three figure
	if len(mathLine) == 3 {                // except 4 figure cases
		figure1, error1 := strconv.Atoi(mathLine[0])
		figureOps := mathLine[1]
		figure2, error2 := strconv.Atoi(mathLine[2])

		// except use Rome and Arab num cases
		testService := 0
		if error1 != nil {
			testService++
			figure1 = 1
		}
		if error2 != nil {
			testService++
			figure2 = 1
		}
		if testService == 1 {
			panic("используются одновременно разные системы счисления")
		}

		// except zero case
		if figure1 == 0 || figure2 == 0 {
			panic("этот калькулятор работает только с числами от 1 до 10")
		}

		// convert Rome to Arab for Rome num case
		if testService == 2 {
			switch mathLine[0] {
			case "I":
				figure1 = 1
			case "II":
				figure1 = 2
			case "III":
				figure1 = 3
			case "IV":
				figure1 = 4
			case "V":
				figure1 = 5
			case "VI":
				figure1 = 6
			case "VII":
				figure1 = 7
			case "VIII":
				figure1 = 8
			case "IX":
				figure1 = 9
			case "X":
				figure1 = 10
			default:
				panic("неправильно набрана римская цифра")
			}
			switch mathLine[2] {
			case "I":
				figure2 = 1
			case "II":
				figure2 = 2
			case "III":
				figure2 = 3
			case "IV":
				figure2 = 4
			case "V":
				figure2 = 5
			case "VI":
				figure2 = 6
			case "VII":
				figure2 = 7
			case "VIII":
				figure2 = 8
			case "IX":
				figure2 = 9
			case "X":
				figure2 = 10
			default:
				panic("неправильно набрана римская цифра")

			}
		}

		// create temp value for output
		answer := 0
		switch figureOps {
		case "+":
			answer = figure1 + figure2
		case "-":
			answer = figure1 - figure2
		case "*":
			answer = figure1 * figure2
		case "/":
			answer = figure1 / figure2
		default:
			panic("оператором ожет бытьт только (+, -, /, *)") // except incorrect operand
		}

		// create output value for Rome case
		if testService == 2 {
			if answer <= 0 {
				panic("в римской системе нет и отрицательных чисел")
			}
			altAnswer := "" // empty str for concatenate
			if answer >= 10 {
				altAnswer += "X"
				answer = answer - 10
			}
			switch answer {
			case 10:
				altAnswer += "X"
			case 9:
				altAnswer += "IX"
			case 8:
				altAnswer += "VIII"
			case 7:
				altAnswer += "VII"
			case 6:
				altAnswer += "VI"
			case 5:
				altAnswer += "V"
			case 4:
				altAnswer += "IV"
			case 3:
				altAnswer += "III"
			case 2:
				altAnswer += "II"
			}
			fmt.Println(altAnswer)
		} else {
			fmt.Print(answer) // output value for Arab case
		}
	} else {
		panic("формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *)")
	}
}
