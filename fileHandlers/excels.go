package filehandlers

import (
	"fmt"
	"os"
	"strconv"

	"github.com/xuri/excelize/v2"
)

// check if the file exist or not
func isExist(filename string) bool {
	_, err := os.Stat(filename)
	if err != nil {
		return !os.IsNotExist(err)
	}
	return true
}

// Initilize file

// create a new file if not exist
func New(filename string, sheet string, headers *[]interface{}) error {

	// if file exist return
	if isExist(filename) {
		return nil
	}

	// Initilize and create a new file
	file := excelize.NewFile()

	// close the file after all operation done
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	// update the default sheet name to provided sheet name
	if err := file.SetSheetName("Sheet1", sheet); err != nil {
		return err
	}

	// add the header row
	err := file.SetSheetRow(sheet, "A1", headers)
	if err != nil {
		return err
	}

	// save the file
	if err := file.SaveAs(filename); err != nil {
		return err
	}
	return nil
}

// Get id
func GetId(filename string, sheet string) (int, error) {
	file, err := excelize.OpenFile(filename)
	if err != nil {
		return 0, err
	}
	// close the file after all operation done
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	data, err := file.GetCellValue(sheet, "E1")
	if err != nil {
		return 0, err
	}
	num, err := strconv.Atoi(data)
	if err != nil {
		return 0, err
	}
	return num, err

}

// Add a new Entry
func AddNewEntry(filename string, sheet string, data *[]interface{}) error {
	file, err := excelize.OpenFile(filename)
	if err != nil {
		return err
	}
	// close the file after all operation done
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	rows, err := file.GetRows(sheet)
	if err != nil {
		return err
	}
	nextRowNum := len(rows) + 1
	nextRowRef := fmt.Sprintf("A%d", nextRowNum)

	// Append new row
	if err := file.SetSheetRow(sheet, nextRowRef, data); err != nil {
		return err
	}
	// update the counter for id
	if err := file.SetCellValue(sheet, "E1", (*data)[0]); err != nil {
		return err
	}
	// save data
	if err := file.Save(); err != nil {
		return err
	}
	return nil
}

func RemoveEntry(filename string, sheet string, id int) error {
	file, err := excelize.OpenFile(filename)
	if err != nil {
		return err
	}

	// close the file after all operation done
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	data, _ := file.GetCols(sheet)
	for idx, val := range data[0][1:] {
		val, _ := strconv.Atoi(val)
		if val == id {

			if err := file.RemoveRow(sheet, idx+2); err != nil {
				return err
			}
			// save data
			if err := file.Save(); err != nil {
				return err
			}
			return nil
		}
	}
	return nil
}

func UpdateEntry(filename string, sheet string, id int) error {
	file, err := excelize.OpenFile(filename)
	if err != nil {
		return err
	}

	// close the file after all operation done
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	data, _ := file.GetCols(sheet)
	for idx, val := range data[0][1:] {
		val, _ := strconv.Atoi(val)
		if val == id {

			if err := file.RemoveRow(sheet, idx+2); err != nil {
				return err
			}
			// save data
			if err := file.Save(); err != nil {
				return err
			}
			return nil
		}
	}
	return nil
}

func ListEntry(filename string, sheet string) ([][]string, error) {
	file, err := excelize.OpenFile(filename)
	if err != nil {
		return [][]string{}, err
	}
	data, err := file.GetRows(sheet)
	if err != nil {
		return [][]string{}, err
	}
	return data, nil
}
