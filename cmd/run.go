package cmd

import (
	"agit/decoder"
	"os/exec"
	"strings"

	"github.com/pykelysia/pyketools"
)

func (cmd *Command) Run() (bool, error) {
	cmdline := ""
	switch cmd.modeName {
	case normalMode, fastMode:
		cmdline = normalModeCommandLine
	}

	ps := exec.Command("PowerShell", "-Command", cmdline)
	ps.Dir = cmd.path

	output, err := ps.CombinedOutput()
	if err != nil {
		return false, err
	}

	if ok, s := cmd.checkOutput(output); !ok {
		pyketools.Errorf("git push failed. Output: %v", s)
		return false, nil
	} else {
		return true, nil
	}

}

func (cmd *Command) checkOutput(output []byte) (bool, string) {
	d, _ := decoder.NewDecoder(output)
	if strings.Contains(d, successOutput) {
		return true, d
	} else if strings.Contains(d, connectionReset) {
		return false, d
	} else if strings.Contains(d, notConnectServer) {
		return false, d
	} else if strings.Contains(d, notResolveHost) {
		return false, d
	}
	return false, d
}
