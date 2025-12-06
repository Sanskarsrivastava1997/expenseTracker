package cmds

import (
	filehandlers "expense-tracker/fileHandlers"
	"flag"
	"fmt"

	"github.com/hashicorp/cli"
)

type Delete struct{ UI cli.ColoredUi }

// Help implements cli.Command.
func (a *Delete) Help() string {
	message := "Usage: expense-tracker delete --id <001>"
	return message
}

// Run implements cli.Command.
func (a *Delete) Run(args []string) int {
	// declare a new flag set
	fs := flag.NewFlagSet("add", flag.ContinueOnError)
	id := fs.Int("id", 0, "Enter the Id for which expense needs to be deleted - (Required)")

	// parse the flag
	err := fs.Parse(args)
	if err != nil {
		return 1
	}

	// validate flag values
	if err := validateDeleteFlags(*id); err != nil {
		a.UI.Error(fmt.Sprintf("error: %v", err))
		return 1
	}

	// Remove the entry
	if err := filehandlers.RemoveEntry(FileName, SheetName, *id); err != nil {
		fmt.Println(err)
		return 1
	}
	output := "# Expense deleted successfully"
	a.UI.Output(output)

	return 0
}

// Synopsis implements cli.Command.
func (a *Delete) Synopsis() string {
	message := "Use delete to delete any expenses."
	return message
}
