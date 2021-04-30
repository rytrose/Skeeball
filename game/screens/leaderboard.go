package screens

import (
	"fmt"
	"image/color"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/rytrose/soup-the-moon/game/fonts"
	"github.com/rytrose/soup-the-moon/game/input"
	"github.com/rytrose/soup-the-moon/game/state"
	"github.com/rytrose/soup-the-moon/game/util"
)

// leaderboardState maintains all state needed for the leaderboard screen.
type leaderboardState struct {
	index int
}

// theLeaderboardState is the state of the leaderboard screen.
var theLeaderboardState = &leaderboardState{}

// UpdateLeaderboard updates leaderboard screen state before every frame.
func UpdateLeaderboard() ScreenID {
	// Go back to the menu screen
	if input.Back() {
		// Reset scroll state
		theLeaderboardState.index = 0

		return ScreenMenu
	}

	// Scroll down menu
	if input.Down() {
		// Only allow scrolling down to reach the end of the list (i.e. index == len(entries) - 3)
		theLeaderboardState.index = util.Min(theLeaderboardState.index+1, util.Max(len(state.Global.Leaderboard.Entries)-3, 0))
	}

	// Scroll up menu
	if input.Up() {
		// Only allow scrolling to the top
		theLeaderboardState.index = util.Max(theLeaderboardState.index-1, 0)
	}

	return ScreenLeaderboard
}

// DrawLeaderboard draws one frame of the menu screen.
func DrawLeaderboard(count uint64, w, h int, screen *ebiten.Image) {
	drawLeaderboardTitle(w, screen)
	drawLeaders(w, h, screen)
}

// drawLeaderboardTitle draws the title to the top of the screen.
func drawLeaderboardTitle(w int, screen *ebiten.Image) {
	title := "Top Astronauts"

	// Draw the title centered
	titleX := (w - len(title)*32) / 2
	text.Draw(screen, title, fonts.ArcadeFont32, titleX, 4*16, color.White)
}

// drawLeaders draws the top scorers to the screen.
func drawLeaders(w, h int, screen *ebiten.Image) {
	if theLeaderboardState.index > len(state.Global.Leaderboard.Entries)-1 {
		return
	}

	i := theLeaderboardState.index
	for i < theLeaderboardState.index+3 {
		if i < len(state.Global.Leaderboard.Entries) {
			entry := state.Global.Leaderboard.Entries[i]

			// Draw rank
			text.Draw(screen, fmt.Sprintf("%d.", i+1), fonts.ArcadeFont16, w/16, ((10+((i-theLeaderboardState.index)*6))*16)-8, color.White)

			// Draw initials
			text.Draw(screen, strings.Join(getInitials(entry.Initials), " "), fonts.ArcadeFont32, (8*w)/32, (10+((i-theLeaderboardState.index)*6))*16, color.White)

			// Draw timestamp
			timestamp := entry.Timestamp.Format("1/2/06")
			text.Draw(screen, timestamp, fonts.ArcadeFont16, (18*w)/32, ((10+((i-theLeaderboardState.index)*6))*16)-8, color.White)

			// Draw score
			score := fmt.Sprintf("%07d", entry.Score)
			text.Draw(screen, score, fonts.ArcadeFont16, (25*w)/32, ((10+((i-theLeaderboardState.index)*6))*16)-8, color.White)
		} else {
			return
		}
		i++
	}

	if theLeaderboardState.index > 0 {
		// Draw up cursor
		text.Draw(screen, "/\\", fonts.ArcadeFont16, (7*w)/8, (7*h)/8, color.White)
	}

	if theLeaderboardState.index+3 < len(state.Global.Leaderboard.Entries) {
		// Draw down cursor
		text.Draw(screen, "\\/", fonts.ArcadeFont16, (7*w)/8, (7*h)/8+32, color.White)
	}
}
