package fonts

import (
	"log"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

// Arcade font in various sizes.
var (
	ArcadeFont64 font.Face
	ArcadeFont32 font.Face
	ArcadeFont16 font.Face
)

// dpi is the font resolution.
const dpi = 72

func init() {
	// Loads the arcade font file
	arcadeTT, err := opentype.Parse(PressStart2PRegular)
	if err != nil {
		log.Fatal(err)
	}

	// Loads the arcade font face size 64 point
	ArcadeFont64, err = opentype.NewFace(arcadeTT, &opentype.FaceOptions{
		Size:    64,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}

	// Loads the arcade font face size 32 point
	ArcadeFont32, err = opentype.NewFace(arcadeTT, &opentype.FaceOptions{
		Size:    32,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}

	// Loads the arcade font face size 16 point
	ArcadeFont16, err = opentype.NewFace(arcadeTT, &opentype.FaceOptions{
		Size:    16,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
}
