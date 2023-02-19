package sshutil

type SshConnection interface {
	executeCommand(string) string
	connect()
}

// SshConnectionWithPrivateKey
type SshConnectionWithPrivateKey struct {
}

func (conn SshConnectionWithPrivateKey) executeCommand(command string) string {
	return ""
}

func (conn SshConnectionWithPrivateKey) connect() {
	return
}

// SshConnectionWithPassword
type SshConnectionWithPassword struct {
}

func (conn SshConnectionWithPassword) executeCommand(command string) string {
	return ""
}

func (conn SshConnectionWithPassword) connect() {
	return
}
