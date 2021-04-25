package state

import (
	"encoding/gob"
	"log"
	"os"
	"time"
)

// State is the state of the game that is persisted to a local file.
type State struct {
	CurrentInitials []int
	Leaderboard     *Leaderboard
}

// Leaderboard maintains the state of the leaderboard.
type Leaderboard struct {
	Entries []*LeaderboardEntry
}

// LeaderboardEntry is a game entry on the leaderboard.
type LeaderboardEntry struct {
	Initials  []int
	Score     int
	Timestamp time.Time
}

// Global contains global game state that is persisted to a local file.
var Global *State

// statePath is the path to the state file.
const statePath = "state.gob"

func init() {
	load()
}

// load loads state from file, or creates an empty state.
func load() {
	stateFile, err := os.Open(statePath)
	if err != nil {
		log.Printf("unable to open stateFile: %s", err)
		log.Println("creating new state")

		// File doesn't exist or is corrupted, create new one
		Global = &State{
			CurrentInitials: make([]int, 3),
			Leaderboard: &Leaderboard{
				Entries: []*LeaderboardEntry{},
			},
		}

		// Save off new state
		Save()
		return
	}

	// Decode state from file
	s := &State{}
	Global = s
	stateDecoder := gob.NewDecoder(stateFile)
	err = stateDecoder.Decode(s)

	// If unable to decode, create a new state
	if err != nil {
		log.Printf("unable to decode stateFile: %s", err)
		log.Println("creating new state")
		Global = &State{
			CurrentInitials: make([]int, 3),
			Leaderboard: &Leaderboard{
				Entries: []*LeaderboardEntry{},
			},
		}
	}

	// Save off loaded state
	Save()
}

// Save saves the current state to a local file.
func Save() {
	stateFile, err := os.Create(statePath)
	if err != nil {
		log.Fatalf("unable to create stateFile: %s", err)
	}

	stateEncoder := gob.NewEncoder(stateFile)
	stateEncoder.Encode(Global)
}
