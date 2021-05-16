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
	TitleThemePlayer    *audio.Player
	InitialsThemePlayer *audio.Player
)

func init() {
	// Set up audio context
	audioContext = audio.NewContext(sampleRate)

	// Initialize audio players
	var err error

	// Create title theme player
	TitleThemePlayer, err = initPlayer(TitleTheme_wav)
	if err != nil {
		log.Fatal(err)
	}

	// Create initials theme player
	InitialsThemePlayer, err = initPlayer(InitialsTheme_wav)
	if err != nil {
		log.Fatal(err)
	}
}

func initPlayer(b []byte) (*audio.Player, error) {
	// Load song data
	stream, err := wav.DecodeWithSampleRate(sampleRate, bytes.NewReader(b))
	if err != nil {
		log.Fatal(err)
	}

	// Create looping stream
	loop := audio.NewInfiniteLoop(stream, stream.Length())

	// Create theme song player
	return audio.NewPlayer(audioContext, loop)
}
