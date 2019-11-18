package main

type Watcher struct {
	messages map[string]MessageProcess
	status   Status
}

func NewWatcher() Watcher {
	messages := make(map[string]MessageProcess)
	return Watcher{
		messages: messages,
		status:   UNKNOWN,
	}
}

func (watcher *Watcher) Status() Status {
	return watcher.status
}

func (watcher *Watcher) Messages() []MessageProcess {
	var messages []MessageProcess
	for _, message := range watcher.messages {
		messages = append(messages, message)
	}
	return messages
}

func (watcher *Watcher) MessagesFor(displays StringSet) []MessageProcess {
	var messages []MessageProcess
	for display := range displays {
		message, contains := watcher.messages[display]
		if contains {
			messages = append(messages, message)
		} else {
			messages = append(messages, MessageProcess{PID: 0, Display: display})
		}
	}
	return messages
}

func (watcher *Watcher) Update(messages []MessageProcess, status Status) {
	for _, message := range messages {
		watcher.messages[message.Display] = message
	}
	watcher.status = status
}

func (watcher *Watcher) CleanUp(activeDisplays StringSet) {
	var currentDisplays []string
	for display := range watcher.messages {
		currentDisplays = append(currentDisplays, display)
	}

	for _, display := range currentDisplays {
		if !activeDisplays.Contains(display) {
			delete(watcher.messages, display)
		}
	}
}
