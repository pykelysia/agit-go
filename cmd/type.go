package cmd

type (
	mode    int
	Command struct {
		argsWithoutKeys []string
		args            map[string]string
		path            string
		modeName        mode
	}
)
