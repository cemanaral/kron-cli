package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"

	"golang.org/x/crypto/ssh"
)

func main() {
    pemBytes, err := ioutil.ReadFile("PATH_TO_PEM_FILE")
    if err != nil {
        log.Fatal(err)
    }
    signer, err := ssh.ParsePrivateKey(pemBytes)
    if err != nil {
        log.Fatalf("parse key failed:%v", err)
    }
    config := &ssh.ClientConfig{
        User: "ubuntu",
        Auth: []ssh.AuthMethod{ssh.PublicKeys(signer)},
        HostKeyCallback: ssh.InsecureIgnoreHostKey(),
    }
    
    conn, err := ssh.Dial("tcp", "IP_ADDR:22", config)
    if err != nil {
        log.Fatalf("dial failed:%v", err)
    }
    defer conn.Close()
    session, err := conn.NewSession()
    if err != nil {
        log.Fatalf("session failed:%v", err)
    }
    defer session.Close()
    var stdoutBuf bytes.Buffer
    session.Stdout = &stdoutBuf
    err = session.Run("hostname")
    if err != nil {
        log.Fatalf("Run failed:%v", err)
    }
    
    fmt.Printf("%s", stdoutBuf.String() )

}