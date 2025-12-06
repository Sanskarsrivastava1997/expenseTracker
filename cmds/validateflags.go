package cmds

import "errors"

func validateAddFlags(desc string, amount float64) error {
	if desc == "" {
		return errors.New("description is required but no argument found")
	} else if amount == 0 {
		return errors.New("amount is required but no argument found")
	}
	return nil
}

func validateDeleteFlags(id int) error {
	if id == 0 {
		return errors.New("id is required but no argument found")
	}
	return nil
}
