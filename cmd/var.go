package cmd

const (
	normalMode mode = iota
	fastMode
)

const (
	normalModeCommandLine string = "git push"

	successOutput string = "Everything up-to-date"

	connectionReset string = "Connection was reset"

	notConnectServer string = "Could not connect to server"

	notResolveHost string = "Could not resolve host"
)
