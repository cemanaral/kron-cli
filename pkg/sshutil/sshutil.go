package sshutil

type SshConnection interface {
	ExecuteCommand(string) string
	connect()
	disconnect()
}

// SshConnectionWithPrivateKey
type SshConnectionWithPrivateKey struct {
}

func (conn SshConnectionWithPrivateKey) ExecuteCommand(command string) string {
	conn.connect()
	defer conn.disconnect()
	return ""
}

func (conn SshConnectionWithPrivateKey) connect() {
	return
}

func (conn SshConnectionWithPrivateKey) disconnect() {
	return
}

// SshConnectionWithPassword
type SshConnectionWithPassword struct {
}

func (conn SshConnectionWithPassword) ExecuteCommand(command string) string {
	conn.connect()
	defer conn.disconnect()
	return ""
}

func (conn SshConnectionWithPassword) connect() {
	return
}

func (conn SshConnectionWithPassword) disconnect() {
	return
}
