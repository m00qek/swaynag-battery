package main

import (
	"os"
	"os/exec"
	"strings"

	ps "github.com/mitchellh/go-ps"
)

type MessageProcess struct {
	PID     int
	Display string
}

func show(message string, display string) (*MessageProcess, error) {
	cmd := exec.Command("swaynag", "--message", message, "--output", display)

	err := cmd.Start()
	if err != nil {
		return nil, err
	}

	go func() {
		cmd.Wait()
	}()

	return &MessageProcess{PID: cmd.Process.Pid, Display: display}, nil
}

func isSwaynag(process ps.Process, err error) bool {
	return err == nil &&
		process != nil &&
		strings.HasSuffix(process.Executable(), "swaynag")
}

func ShowMessage(message string, process MessageProcess) (*MessageProcess, error) {
	if isSwaynag(ps.FindProcess(process.PID)) {
		return &process, nil
	}
	return show(message, process.Display)
}

func CloseMessage(process MessageProcess) {
	systemProcess, err := os.FindProcess(process.PID)
	if err != nil {
		return
	}

	defer systemProcess.Release()
	systemProcess.Signal(os.Interrupt)
}

func ShowAll(message string, processes []MessageProcess) []MessageProcess {
	var openProcesses []MessageProcess
	for _, process := range processes {
		newProcess, _ := ShowMessage(message, process)
		if newProcess == nil {
			openProcesses = append(openProcesses, process)
		} else {
			openProcesses = append(openProcesses, *newProcess)
		}
	}
	return openProcesses
}

func CloseAll(processes []MessageProcess) {
	for _, process := range processes {
		CloseMessage(process)
	}
}
