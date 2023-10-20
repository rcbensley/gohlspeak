# HL Say command in Go

Play Half-Life 1 vox sound effects all from a single binary.

Inspired by [hlspeak](https://github.com/nickjanssen/hlspeak)

The pre-compiled binaries already contain the audio files.

The source audio files are here: `https://github.com/sourcesounds/hl1/tree/master/sound/vox`

## Usage

Missing or unknown words will always be 'error'
    
    ./gohlsay attention attention top shelf alien whiskey

List all words available at build time with the `-w` flag:

    ./gohlsay -w

