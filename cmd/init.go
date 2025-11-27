package cmd

import (
	"os"

	"github.com/pykelysia/pyketools"
)

func NewCmd() *Command {
	cmd := &Command{}
	cmd.init()
	return cmd
}

func (cmd *Command) init() {

	// 初始化参数列表
	// 分为有键参数和无键参数
	cmd.args = make(map[string]string)
	for i := 1; i < len(os.Args); i++ {
		if os.Args[i][0] == '-' {
			if os.Args[i+1][0] == '-' {
				cmd.args[os.Args[i]] = ""
			} else {
				cmd.args[os.Args[i]] = os.Args[i+1]
				i++
			}
		} else {
			cmd.argsWithoutKeys = append(cmd.argsWithoutKeys, os.Args[i])
		}
	}

	// 根据有键参数设置相关模式
	shouldBreak := false
	for key, value := range cmd.args {
		if shouldBreak {
			break
		}

		switch key {

		// 设置模式
		case "-m", "--mode":
			switch value {
			case "normal":
				cmd.modeName = normalMode
			case "fast":
				cmd.modeName = fastMode
			default:
				pyketools.Errorf("unknown mode: %s", value)
			}

		// 设置路径
		case "-p", "--path":
			if value == "" {
				pyketools.Errorf("path cannot be empty.")
			} else {
				cmd.path = value
			}

		// 设置为默认模式
		case "-d", "--default":
			cmd.modeName = normalMode
			cmd.path = "./"
			shouldBreak = true
		}
	}
	if cmd.modeName == 0 {
		cmd.modeName = normalMode
	}
	if cmd.path == "" {
		cmd.path = "./"
	}
}
