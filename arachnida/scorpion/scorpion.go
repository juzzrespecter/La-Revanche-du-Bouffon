package scorpion

import (
	"fmt"
	"os"
)

func Usage() {

}

func main() {
	if len(os.Args[1:]) < 1 {
		fmt.Fprintln(os.Stderr, "")
		Usage()
		os.Exit(1)
	}
	// gorutinas y un mutex pa imprimir
}
