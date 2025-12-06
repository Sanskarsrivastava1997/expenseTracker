package cmds

import (
	filehandlers "expense-tracker/fileHandlers"
	"fmt"

	"github.com/hashicorp/cli"
)

type List struct{ UI cli.ColoredUi }

// Help implements cli.Command.
func (a *List) Help() string {
	message := "Usage: expense-tracker list"
	return message
}

// Run implements cli.Command.
func (a *List) Run(args []string) int {
	data, err := filehandlers.ListEntry(FileName, SheetName)
	if err != nil {
		a.UI.Error(fmt.Sprintln(err))
	}
	for _, data := range data {
		a.UI.Output(fmt.Sprintf("# %-3s %-20s %-10s %-3s", data[0], data[1], data[2], data[3]))
	}

	return 0
}

// Synopsis implements cli.Command.
func (a *List) Synopsis() string {
	message := "Use list to get all expenses."
	return message
}
