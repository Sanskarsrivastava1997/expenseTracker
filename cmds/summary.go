package cmds

import "github.com/hashicorp/cli"

type Summary struct{ UI cli.ColoredUi }

// Help implements cli.Command.
func (a *Summary) Help() string {
	message := "Usage: expense-tracker summary"
	return message
}

// Run implements cli.Command.
func (a *Summary) Run(args []string) int {
	return 0
}

// Synopsis implements cli.Command.
func (a *Summary) Synopsis() string {
	message := "Use summary to summarize the total amount"
	return message
}
