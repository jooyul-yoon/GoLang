package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/jooyul-yoon/learngo/dict_proj/mydict"
)

func main() {
	// Initialize dictionary
	dictionary := mydict.Dictionary{"first": "hello", "second": "world"}
	choice := 0

	// Menu
	for {
		ShowMenu()
		fmt.Print("Enter a number: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			Add(dictionary)
		case 2:
			Update(dictionary)
		case 3:
			Delete(dictionary)
		case 4:
			Search(dictionary)
		case 5:
			dictionary.String()
		case 6:
			fmt.Println("Good Bye!")
			return
		}
	}

}

func ShowMenu() {
	fmt.Println(" ---------------")
	fmt.Println("|  1. Add	|")
	fmt.Println("|  2. Update	|")
	fmt.Println("|  3. Delete	|")
	fmt.Println("|  4. Search	|")
	fmt.Println("|  5. display	|")
	fmt.Println("|  6. Quit	|")
	fmt.Println(" ---------------")
}

func Add(dictionary mydict.Dictionary) {

	fmt.Print("Enter an exixting word: ")
	var word string
	fmt.Scanln(&word)

	fmt.Print("Enter the new definition: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	definition := strings.TrimSuffix(input, "\n")

	err := dictionary.Add(word, definition)
	if err != nil {
		fmt.Println(err)
	}
}

func Update(dictionary mydict.Dictionary) {
	fmt.Print("Enter a word to search: ")
	var word string
	fmt.Scanln(&word)

	fmt.Print("Enter the definition: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	definition := strings.TrimSuffix(input, "\n")

	err := dictionary.Update(word, definition)
	if err != nil {
		fmt.Println(err)
	}
}

func Delete(dictionary mydict.Dictionary) {
	fmt.Print("Enter a word to delete: ")
	var word string
	fmt.Scanln(&word)

	err := dictionary.Delete(word)
	if err != nil {
		fmt.Println(err)
	}
}

func Search(dictionary mydict.Dictionary) {
	fmt.Print("Enter a word to search: ")
	var word string
	fmt.Scanln(&word)

	def, err := dictionary.Search(word)
	if err == nil {
		fmt.Println("definition:", def)
	} else {
		fmt.Println(err)
	}
}
