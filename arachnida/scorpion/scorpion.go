package scorpion

import (
	"fmt"
	"os"
)

// valid ext

func Usage() {
	fmt.Fprintln(os.Stderr, "usage: "+
		""+
		"./scorpion FILE [...FILE]")
}

func main() {
	if len(os.Args[1:]) < 1 {
		fmt.Fprintln(os.Stderr, "Not enough arguments")
		Usage()
		os.Exit(1)
	}
	for _, file := range os.Args[1:] {
		// si extension no en valid ext wont do
		_, err := os.Stat(file)
		if err != nil {

		}
	}
	// gorutinas y un mutex pa imprimir
}
