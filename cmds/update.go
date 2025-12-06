package cmds

import "github.com/hashicorp/cli"

type Update struct{ UI cli.ColoredUi }

// Help implements cli.Command.
func (a *Update) Help() string {
	message := "Usage: expense-tracker update --id <001> --description <hello world> --amount <20>"
	return message
}

// Run implements cli.Command.
func (a *Update) Run(args []string) int {
	return 0
}

// Synopsis implements cli.Command.
func (a *Update) Synopsis() string {
	message := "Use update to update any expenses."
	return message
}
