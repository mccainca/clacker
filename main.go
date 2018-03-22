package main

import (
	"fmt"

	"bitbucket.org/StephenPatrick/go-winaudio/winaudio"
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/keyboard"
)

func main() {
	keys := keyboard.NewDriver()

	work := func() {
		keys.On(keyboard.Key, func(data interface{}) {
			key := data.(keyboard.KeyEvent)
			fmt.Println("keyboard event!", key, key.Char)
			err := playSound("_sounds\\balloon.wav")
			if err != nil {
				fmt.Println(err)
			}
		})
	}

	robot := gobot.NewRobot("clackerBot",
		[]gobot.Connection{},
		[]gobot.Device{keys},
		work,
	)

	robot.Start()
}

func playSound(path string) error {
	winaudio.InitWinAudio()
	wavfile, _ := winaudio.LoadWav(path)
	wavfile.SetLooping(false)
	wavfile.Play()
	return nil
}
