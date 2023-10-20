package main

import (
	"bytes"
	"embed"
	"fmt"
	"os"
	"time"

	_ "embed"

	"github.com/ebitengine/oto/v3"
	"github.com/hajimehoshi/go-mp3"
)

//go:embed vox
var files embed.FS

func main() {
	l := len(os.Args)

	if l <= 1 {
		fmt.Println("no words")
		os.Exit(1)
	}

	words := os.Args[1:]

	a := []*mp3.Decoder{}

	for _, w := range words {
		fName := "vox/" + w + ".mp3"
		var fb []byte
		var err error
		fb, err = files.ReadFile(fName)
		if err != nil {
			fb, _ = files.ReadFile("vox/error.mp3")
		}
		fr := bytes.NewReader(fb)
		d, err := mp3.NewDecoder(fr)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		a = append(a, d)

	}

	op := &oto.NewContextOptions{}
	op.SampleRate = 44100
	op.ChannelCount = 2
	op.Format = oto.FormatSignedInt16LE

	otoCtx, readyChan, err := oto.NewContext(op)
	if err != nil {
		panic("oto.NewContext failed: " + err.Error())
	}
	<-readyChan

	for _, w := range a {
		player := otoCtx.NewPlayer(w)
		player.Play()
		for player.IsPlaying() {
			time.Sleep(time.Millisecond)
		}
		err = player.Close()
		if err != nil {
			panic(err)
		}
		time.Sleep(time.Millisecond * 333)
	}
}
