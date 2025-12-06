package main

import (
	"expense-tracker/cmds"
	filehandlers "expense-tracker/fileHandlers"
	"fmt"
	"os"

	"github.com/hashicorp/cli"
)

func addCommandFactory() (cli.Command, error) {
	settings := cli.ColoredUi{
		OutputColor: cli.UiColorGreen,
		InfoColor:   cli.UiColorCyan,
		ErrorColor:  cli.UiColorRed,
		WarnColor:   cli.UiColorYellow,
		Ui: &cli.BasicUi{
			Writer: os.Stdout,
		},
	}
	return &cmds.Add{UI: settings}, nil
}

func updateCommandFactory() (cli.Command, error) {
	settings := cli.ColoredUi{
		OutputColor: cli.UiColorGreen,
		InfoColor:   cli.UiColorCyan,
		ErrorColor:  cli.UiColorRed,
		WarnColor:   cli.UiColorYellow,
		Ui: &cli.BasicUi{
			Writer: os.Stdout,
		},
	}
	return &cmds.Update{UI: settings}, nil
}

func deleteCommandFactory() (cli.Command, error) {
	settings := cli.ColoredUi{
		OutputColor: cli.UiColorGreen,
		InfoColor:   cli.UiColorCyan,
		ErrorColor:  cli.UiColorRed,
		WarnColor:   cli.UiColorYellow,
		Ui: &cli.BasicUi{
			Writer: os.Stdout,
		},
	}
	return &cmds.Delete{UI: settings}, nil
}

func listCommandFactory() (cli.Command, error) {
	settings := cli.ColoredUi{
		OutputColor: cli.UiColorGreen,
		InfoColor:   cli.UiColorCyan,
		ErrorColor:  cli.UiColorRed,
		WarnColor:   cli.UiColorYellow,
		Ui: &cli.BasicUi{
			Writer: os.Stdout,
		},
	}
	return &cmds.List{UI: settings}, nil
}

func summaryCommandFactory() (cli.Command, error) {
	settings := cli.ColoredUi{
		OutputColor: cli.UiColorGreen,
		InfoColor:   cli.UiColorCyan,
		ErrorColor:  cli.UiColorRed,
		WarnColor:   cli.UiColorYellow,
		Ui: &cli.BasicUi{
			Writer: os.Stdout,
		},
	}
	return &cmds.Summary{UI: settings}, nil
}

func main() {

	headers := &[]interface{}{"Id", "Description", "Amount", "Date", 0}
	fileName := cmds.FileName
	sheetName := cmds.SheetName

	// Create a new file if not exist
	err := filehandlers.New(fileName, sheetName, headers)
	if err != nil {
		fmt.Println(err)
	}

	c := cli.NewCLI("expense-tracker", "1.0.0")
	c.Args = os.Args[1:]
	c.Commands = map[string]cli.CommandFactory{
		"add ":    addCommandFactory,
		"update":  updateCommandFactory,
		"delete":  deleteCommandFactory,
		"list":    listCommandFactory,
		"summary": summaryCommandFactory,
	}

	exitStatus, err := c.Run()
	if err != nil {
		fmt.Println(err)
	}
	os.Exit(exitStatus)
}
