package cmds

import (
	filehandlers "expense-tracker/fileHandlers"
	"flag"
	"fmt"
	"time"

	"github.com/hashicorp/cli"
)

type Add struct{ UI cli.ColoredUi }

// Help implements cli.Command.
func (a Add) Help() string {
	message := `Usage: 
	expense-tracker add --description <hello world> --amount <20>`
	return message
}

// Run implements cli.Command.
func (a *Add) Run(args []string) int {

	// declare a new flag set
	fs := flag.NewFlagSet("add", flag.ContinueOnError)
	desc := fs.String("description", "", "Enter the description - (Required)")
	amount := fs.Float64("amount", 0, "Enter the amount - (Required)")

	// parse the flag
	err := fs.Parse(args)
	if err != nil {
		return 1
	}

	// validate flag values

	if err := validateAddFlags(*desc, *amount); err != nil {
		a.UI.Error(fmt.Sprintf("error: %v", err))
		return 1
	}

	// add data entry
	id, err := filehandlers.GetId("database.xlsx", "Expense")
	if err != nil {
		a.UI.Error(fmt.Sprintf("error: %v", err))
		return 1
	}
	data := &[]interface{}{id + 1, *desc, *amount, time.Now()}

	if err := filehandlers.AddNewEntry(FileName, SheetName, data); err != nil {
		a.UI.Error(fmt.Sprintf("error: %v", err))
		return 1
	}
	output := fmt.Sprintf("# Expense added successfully (ID: %v)", (*data)[0])
	a.UI.Output(output)

	return 0
}

// Synopsis implements cli.Command.
func (a *Add) Synopsis() string {
	message := "Use add to add new expenses."
	return message
}
