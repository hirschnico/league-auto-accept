package main

import (
	"log"
	"os"
	"time"

	"github.com/hirschnico/league-auto-accept/pkg/riotapi"
)

func exit() {
	log.Println("Press Enter to exit...")
	_, _ = os.Stdin.Read([]byte{0})
	os.Exit(1)
}

func main() {
	log.Println("Starting League Auto-Accept...")

	creds, err := riotapi.ReadLockfile()
	if err != nil {
		log.Println("League Client is not running.")
		exit()
	}

	client := riotapi.NewClient(creds)

	log.Println("League Auto-Accept has started.")

	for {
		state, err := client.CheckReadyCheck()
		if err != nil {
			log.Println("Error doing ready check.")
			exit()
		}

		if state == "InProgress" {
			log.Println("Match found! Accepting...")

			if err := client.AcceptMatch(); err != nil {
				log.Println("Failed to accept match.")
				exit()
			} else {
				log.Println("Match accepted!")
				time.Sleep(13 * time.Second)
			}
		}

		inGame, err := client.IsInGame()
		if err != nil {
			log.Println("Error checking game state.")
			exit()
		}

		if inGame {
			log.Println("User is in a game.")
			exit()
		}

		time.Sleep(time.Second)
	}
}
