package main

import (
	"fmt"
	"github.com/hajimehoshi/oto"
	"github.com/tosone/minimp3"
	"io/ioutil"
)

func main() {
	file, err := ioutil.ReadFile("./23.mp3")
	if err != nil {
		fmt.Println(err)
		return
	}
	dec, data, err := minimp3.DecodeFull(file)
	if err != nil {
		fmt.Println(err)
	}

	player, _ := oto.NewPlayer(dec.SampleRate, dec.Channels, 2, 1024)
	player.Write(data)

}
