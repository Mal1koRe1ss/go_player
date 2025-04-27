package main

import (
	"fmt"
	"os"
	"go_player/player"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go_player <mp3-path>")
		return
	}

	player.RunPlayer(os.Args[1])
}