package sshutil

type SshConnection interface {
	ExecuteCommand(string) string
	connect()
	disconnect()
}

// SshConnectionWithPrivateKey
type SshConnectionWithPrivateKey struct {
	Address string
	PrivateKeyPath string
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
	Address string
	Password string
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
