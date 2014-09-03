package link

import "os/exec"

type LocalLink struct{}

func NewLocalLink() *LocalLink {
	return &LocalLink{}
}

func (s *LocalLink) Run(command string) (string, error) {
	cmd := exec.Command("sh", "-c", command)

	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return string(output), nil
}

func (s *LocalLink) Ready() bool {
	return true
}
