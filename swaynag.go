package main

import (
	"os"
	"os/exec"
	"strings"

	ps "github.com/mitchellh/go-ps"
)

type Message struct {
	PID     int
	Display string
}

func show(text string, display string) (*Message, error) {
	cmd := exec.Command("swaynag", "--message", text, "--output", display, "--layer", "overlay")

	err := cmd.Start()
	if err != nil {
		logError("Unable to show swaynag in display '%s'.\n", display)
		return nil, err
	}

	go func() {
		cmd.Wait()
	}()

	return &Message{PID: cmd.Process.Pid, Display: display}, nil
}

func isSwaynag(process ps.Process, err error) bool {
	return err == nil &&
		process != nil &&
		strings.HasSuffix(process.Executable(), "swaynag")
}

func ShowMessage(text string, message Message) (*Message, error) {
	if isSwaynag(ps.FindProcess(message.PID)) {
		return &message, nil
	}

	return show(text, message.Display)
}

func CloseMessage(message Message) {
	process, err := os.FindProcess(message.PID)
	if err != nil {
		return
	}
	defer process.Release()
	process.Signal(os.Interrupt)
}

func ShowAll(text string, messages []Message) []Message {
	var openMessages []Message
	for _, message := range messages {
		newMessage, _ := ShowMessage(text, message)
		if newMessage == nil {
			openMessages = append(openMessages, message)
		} else {
			openMessages = append(openMessages, *newMessage)
		}
	}

	return openMessages
}

func CloseAll(messages []Message) {
	for _, message := range messages {
		CloseMessage(message)
	}
}
