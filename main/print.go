package main

import (
	"github.com/luthierycosta/N-puzzle/utils"
	"fmt"
	"time"
	"github.com/luthierycosta/N-puzzle/state"
	"github.com/luthierycosta/N-puzzle/path"

)

func print(initial state.State, path path.Path) {
	utils.ClearScreen()
	fmt.Println("Estado inicial:")
	initial.Print()
	time.Sleep(time.Second)
	for i, direction := range path {
		initial.Move(direction)
		initial.Print()
		fmt.Printf("\nPróximo movimento: %v \nVetor de movimentos:%v\n", path[i], path[i+1:])
		time.Sleep(100 * time.Millisecond)
		if (i != len(path)-1) {
			utils.ClearScreen()
		}
	}
}