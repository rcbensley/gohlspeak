package main

import (
	"bytes"
	"embed"
	"time"

	_ "embed"

	"github.com/ebitengine/oto/v3"
	"github.com/hajimehoshi/go-mp3"
)

//go:embed vox/*.mp3
var files embed.FS

func main() {

	fp := "vox/attention.wav.mp3"
	fb, err := files.ReadFile(fp)
	if err != nil {
		panic(err)
	}

	fr := bytes.NewReader(fb)

	d, err := mp3.NewDecoder(fr)
	if err != nil {
		panic("mp3.NewDecoder failed: " + err.Error())
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

	player := otoCtx.NewPlayer(d)
	player.Play()
	for player.IsPlaying() {
		time.Sleep(time.Millisecond)
	}
	err = player.Close()
	if err != nil {
		panic(err)
	}
}
