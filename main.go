package main

import (
	"fmt"
	"github.com/faiface/beep"
	"github.com/faiface/beep/effects"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/kindlyfire/go-keylogger"
)

const (
	delayKeyfetchMS = 5
)

func AudioPlayer(audio string) {
	f, err := os.Open(audio)
	streamer, format, err := mp3.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	defer func(streamer beep.StreamSeekCloser) {
		err := streamer.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(streamer)

	err = speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	if err != nil {
		panic(err)
	}
	ctrl := &beep.Ctrl{Streamer: beep.Loop(-1, streamer), Paused: false}
	volume := &effects.Volume{
		Streamer: ctrl,
		Base:     2,
		Volume:   0,
		Silent:   false,
	}
	speaker.Play(volume)

	kl := keylogger.NewKeylogger()
	var previousKey rune
	for {
		key := kl.GetKey()
		if !key.Empty {
			if previousKey == 0 {
				previousKey = key.Rune
			} else {
				if key.Rune == previousKey { // if the key is the same as the previous key
					if key.Keycode == 80 { // and if the key is the p key that will pause the music
						speaker.Lock()
						ctrl.Paused = !ctrl.Paused
						speaker.Unlock()
						return
					} else if key.Keycode == 85 { // 85 is u key that will increase volume
						speaker.Lock()
						volume.Volume += 0.5
						fmt.Println(volume.Volume, "Volume")
						speaker.Unlock()
					} else if key.Keycode == 89 { // 89 is y key that will lower volume
						speaker.Lock()
						volume.Volume -= 0.5
						fmt.Println(volume.Volume, "Volume")
						speaker.Unlock()
					}
				}
			}
			previousKey = key.Rune
		}
	}
}
func main() {
	var availableTracks []string
	// tracks
	availableTracks = append(availableTracks, "1665260018704_track.mp3")
	availableTracks = append(availableTracks, "1665260697904_track.mp3")
	availableTracks = append(availableTracks, "1665260698061_track.mp3")
	availableTracks = append(availableTracks, "1665260698061_track.mp3")
	availableTracks = append(availableTracks, "1665260698080_track.mp3")
	availableTracks = append(availableTracks, "1665260698090_track.mp3")
	availableTracks = append(availableTracks, "1665260873227_track.mp3")
	availableTracks = append(availableTracks, "1665260873228_track (1).mp3")
	availableTracks = append(availableTracks, "1665260873228_track (2).mp3")
	availableTracks = append(availableTracks, "1665260873228_track.mp3")
	availableTracks = append(availableTracks, "1665260698526_track.mp3")
	availableTracks = append(availableTracks, "1665260938798_track.mp3")
	availableTracks = append(availableTracks, "1665261191119_track.mp3")
	availableTracks = append(availableTracks, "1665261190998_track.mp3")
	kl := keylogger.NewKeylogger()
	var previousKey rune
	emptyCount := 0
	for {
		key := kl.GetKey()
		if !key.Empty {
			// if previousKey is null
			if previousKey == 0 {
				previousKey = key.Rune
			} else {
				if key.Rune == previousKey { // if the key is the same as the previous key
					if key.Keycode == 80 { // and if the key is the P key
						// do something
						fmt.Println("You pressed P, twice.")
						fmt.Println("You can press P **TWICE** to pause/resume the audio.")
						fmt.Println("You can press U **TWICE** to increase the volume.")
						fmt.Println("You can press Y **TWICE** to lower the volume.")
						// Pick a random track from the available tracks
						AudioPlayer(availableTracks[rand.Intn(len(availableTracks))])

					}
				}
			}
			previousKey = key.Rune
		}
		emptyCount++
		time.Sleep(delayKeyfetchMS * time.Millisecond)
	}
}
