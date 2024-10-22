package link

import (
	"bytes"
	"io/ioutil"
	"os"

	"code.google.com/p/go.crypto/ssh"
)

type SSHLink struct {
	client *ssh.Client
	config *ssh.ClientConfig
}

func makeSigner(keyname string) (signer ssh.Signer, err error) {
	fp, err := os.Open(keyname)
	if err != nil {
		return
	}
	defer fp.Close()

	buf, _ := ioutil.ReadAll(fp)
	signer, _ = ssh.ParsePrivateKey(buf)
	return
}

func makeKeyring() ssh.AuthMethod {
	signers := []ssh.Signer{}
	keys := []string{os.Getenv("HOME") + "/.ssh/id_rsa", os.Getenv("HOME") + "/.ssh/id_dsa"}

	for _, keyname := range keys {
		signer, err := makeSigner(keyname)
		if err == nil {
			signers = append(signers, signer)
		}
	}

	return ssh.PublicKeys(signers...)
}

func generateConfig() *ssh.ClientConfig {
	return &ssh.ClientConfig{
		User: "jdiez",
		Auth: []ssh.AuthMethod{makeKeyring()},
	}
}

func NewSSHLink() *SSHLink {
	return &SSHLink{
		config: generateConfig(),
	}
}

func (s *SSHLink) Connect(host string) (err error) {
	// TODO: Handle errors
	s.client, err = ssh.Dial("tcp", host, s.config)
	return
}

func (s *SSHLink) Run(cmd string) (string, error) {
	session, err := s.client.NewSession()
	if err != nil {
		return "", err
	}

	var stdoutBuf bytes.Buffer
	session.Stdout = &stdoutBuf
	session.Run(cmd)

	return stdoutBuf.String(), nil
}

func (s *SSHLink) Disconnect() {
	s.client.Close()
}

func (s *SSHLink) Ready() bool {
	return s.client != nil
}
