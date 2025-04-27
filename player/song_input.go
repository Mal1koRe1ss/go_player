package player

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func RunPlayer(filePath string) {
	player, err := NewMusicPlayer(filePath)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer player.Stop()

	player.Play()
	fmt.Println("Music playing...")
	fmt.Printf("Currently playing : %s\n", filePath)
	printCommands()

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		scanner.Scan()
		processCommand(scanner.Text(), player)
		
		select {
		case <-player.Done:
			fmt.Println("Music finished.")
			return
		default:
		}
	}
}

func processCommand(cmd string, p *MusicPlayer) {
	switch strings.ToLower(strings.TrimSpace(cmd)) {
	case "pause":
		p.Pause()
		fmt.Println("Paused")
	case "resume":
		p.Resume()
		fmt.Println("Resuming")
	case "stop":
		fmt.Println("Stoping...")
		os.Exit(0)
	case "help":
		printCommands()
	default:
		fmt.Println("Invalid command")
	}
}

func printCommands() {
	fmt.Println(`
Commands:
  pause    - Pauses the player.
  resume   - Resumes the player.
  stop     - Stop & exit.
  help     - List of commands...
`)
}