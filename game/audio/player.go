package audio

import (
	"bytes"
	"log"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
)

var (
	audioContext *audio.Context
)

const (
	sampleRate = 44100
)

// Audio players
var (
	ThemePlayer *audio.Player
)

func init() {
	// Set up audio context
	audioContext = audio.NewContext(sampleRate)

	// Load theme song data
	themeStream, err := wav.Decode(audioContext, bytes.NewReader(SoupTheMoonTheme_wav))
	if err != nil {
		log.Fatal(err)
	}

	// Create looping theme stream
	themeLoop := audio.NewInfiniteLoop(themeStream, themeStream.Length())

	// Create theme song player
	ThemePlayer, err = audio.NewPlayer(audioContext, themeLoop)
	if err != nil {
		log.Fatal(err)
	}
}
