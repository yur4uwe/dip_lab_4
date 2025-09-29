package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	htgotts "github.com/hegedustibor/htgo-tts"
	"github.com/hegedustibor/htgo-tts/voices"
)

func listenCommand() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Tell your team: ")
	command, _ := reader.ReadString('\n')
	return strings.TrimSpace(command)
}

func doThisCommand(message string) bool {
	message = strings.ToLower(message)
	if strings.Contains(message, "hello") {
		sayMessage("hey buddy")
	} else if strings.Contains(message, "so long") {
		sayMessage("while a friend")
		return false
	} else {
		sayMessage(message)
	}
	return true
}

func sayMessage(message string) {
	fmt.Println("Voice assistant:", message)

	if err := os.MkdirAll("audio", os.ModePerm); err != nil {
		fmt.Println("Error creating audio folder:", err)
		return
	}

	msg_file_name := strings.ReplaceAll(message, " ", "_")

	speech := htgotts.Speech{Folder: "audio", Language: voices.Ukrainian}
	file_path, err := speech.CreateSpeechFile(message, msg_file_name)
	if err != nil {
		fmt.Println("Error generating speech file:", err)
		return
	}

	f, err := os.Open(file_path)
	if err != nil {
		fmt.Println("Error opening audio file:", err)
		return
	}
	defer f.Close()

	streamer, format, err := mp3.Decode(f)
	if err != nil {
		fmt.Println("Error decoding mp3:", err)
		return
	}
	defer streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))
	<-done
}

func main() {
	for {
		command := listenCommand()
		if !doThisCommand(command) {
			break
		}
	}
}
