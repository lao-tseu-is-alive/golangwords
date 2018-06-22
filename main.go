package main

/**
testing go with the golang tour https://tour.golang.org/
 */

import (
	"fmt"
	"time"
	"unicode/utf8"
	"os"
	"strings"
)

func add(a , b int) int {
	return a + b
}

func addf(a , b float64) float64{
	return a + b
}

func swap(s1, s2 string) (string, string) {
	return s2, s1
}

func reverseUtfWord(word string) string {
	var reversed  string
	i := 0
	for len(word) > 0 {
		runeValue, size := utf8.DecodeLastRuneInString(word)
		i += size
		fmt.Printf("negative index=[%v] char[%q]\n",i, runeValue)
		reversed += string(runeValue)
		word = word[:len(word) - size]
	}
	return reversed
}

func sayHello(name string) {
	const myName string = "Carlos ☯"
	fmt.Printf("Hello, my name is %v !\n", myName)
	fmt.Printf("My name reversed  %v !\n", reverseUtfWord(myName))
	fmt.Println("I know how to swap :-) let's try with the words 'Hello' 'World' :")
	fmt.Println(swap("Hello", "World"))
}

func nextWeekend()  {
	today := time.Now().Weekday()
	fmt.Printf("Today is %v, %v\n", today, time.Now().Format(time.RFC3339))
	switch time.Saturday {
	case today + 0:
		fmt.Println("It's week-end !")
	case today + 1:
		fmt.Println("Tomorrow it's week-end !")
	case today + 2:
		fmt.Println("In two days it's week-end !")
	default:
		fmt.Println("Week-end is far away ...")


	}
}

func readTextFile(filename string) string {
	// variants of reading : file https://kgrz.io/reading-files-in-go-an-overview.html
	textError := "## ERROR "
	file, err := os.Open(filename)
	if err != nil {
		textError += "OPENING FILE"
		fmt.Println(textError, err)
		return textError
	}
	defer file.Close()

	fileinfo, err := file.Stat()
	if err != nil {
		textError += "GETTING INFO ON FILE"
		fmt.Println(textError, err)
		return textError
	}
	fileSize := fileinfo.Size()
	fmt.Printf("The FILE is %v bytes\n", fileSize)
	buffer := make([]byte, fileSize)

	bytesread, err := file.Read(buffer)
	if err != nil {
		textError += "READING FILE "
		fmt.Println(textError, err)
		return textError
	}
	fmt.Printf(" SUCCESSFULLY READ %v \n", bytesread)
	return string(buffer)
}

func main()  {
	// picked a 
	const inputFilename string = "./data/LES MISÉRABLES -- Tome I -- FANTINE.utf8_txt"
	var text = readTextFile(inputFilename)
	lines := strings.Split(text, "\n")
	fmt.Printf("The text is %v bytes long\n", len(text))
	fmt.Printf("and have %v lines\n", len(lines))
	fmt.Printf(lines[45])

}

