package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
	"sync"
)

// See https://github.com/swaywm/sway/blob/master/sway/sway-ipc.7.scd
// or https://man.archlinux.org/man/sway-ipc.7.en for details.
const(
	swayGetOutputs = 3
)

var byteOrder binary.ByteOrder = binary.LittleEndian

var currentSocket struct {
	mutex sync.Mutex
	connection net.Conn
}

// Send IPC call to sway
func sendIpc(messageName uint32) (string, error) {
	currentSocket.mutex.Lock()
	defer currentSocket.mutex.Unlock()

	socketAddress, err := getSwaySock()
	if (err != nil) {
		return "", err
	}

	err = openSocket(socketAddress)
	if (err != nil) {
		return "", err
	}

	response, err := socketSendMessage(currentSocket.connection, messageName)

	if(currentSocket.connection != nil) {
		currentSocket.connection.Close()
	}

	return response, err
}

func getSwaySock() (string, error) {
	socket_from_env := strings.TrimSpace(os.Getenv("SWAYSOCK"))

	if (len(socket_from_env) == 0) {
		return "", fmt.Errorf("Failed to get socket path from SWAYSOCK environment variable")
	}

	return socket_from_env, nil
}

func openSocket(socketAdress string) (error) {
	if currentSocket.connection != nil {
		currentSocket.connection.Close()
	}

	socketConnection, err := net.Dial("unix", socketAdress)
	if (err == nil) {
		currentSocket.connection = socketConnection
		return nil
	} else {
		return err
	}
}

func socketSendMessage(socketConnection net.Conn, messageName uint32) (string, error) {
	// Send
	message := ipcHeader(messageName)
	err := binary.Write(socketConnection, byteOrder, &message)
	if (err != nil) {
		return "", err
	}

	// Receive header
	var header ipcMessageHeader
	err = binary.Read(socketConnection, byteOrder, &header)
	if (err != nil) {
		return "", err
	}

	// Receive payload
	responsePayload := make([]byte, header.Length)
	_, err = io.ReadFull(socketConnection, responsePayload)

	return	string(responsePayload), err
}

type ipcMessageHeader struct {
	MagicString [6]byte
	Length uint32
	Type uint32
}

func ipcHeader(messageName uint32) (ipcMessageHeader) {
	return ipcMessageHeader{
		MagicString: [6]byte{'i', '3', '-', 'i', 'p', 'c'},
		Length: 0,
		Type: messageName,
	}
}

