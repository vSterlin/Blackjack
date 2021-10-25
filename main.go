package main

import (
	"fmt"

	"github.com/vSterlin/bj/game"
)

func main() {

	g := game.NewGame()
	for {
		fmt.Println("Press Q) to quit or any key to continue")
		var choice string
		fmt.Scanf("%s", &choice)

		if choice == "q" {
			fmt.Println("Game Over!!!")
			return
		}

		g.Loop()
	}

}
