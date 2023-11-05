package menu

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/google/uuid"
	"github.com/subbbbbaru/whitesoftdz/record"
)

type ButtonMenu uint

const (
	SEARCHRECORDBYID ButtonMenu = iota
	SEARCHRECORDBYNAME
	EXIT
)

func welcomeMSG() {
	fmt.Println("Menu:")
	fmt.Println("0. Print object by ID")
	fmt.Println("1. Search objects by name")
	fmt.Println("2. Exit")
}

func scanButtonMenu() (ButtonMenu, error) {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	number, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return ButtonMenu(number), fmt.Errorf("wrong enter")
	}
	return ButtonMenu(number), nil
}

func MainMenu(recordMap *record.MapRecord) {
	for {
		welcomeMSG()

		fmt.Print("Enter choice: ")
		choice, err := scanButtonMenu()
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		switch choice {
		case SEARCHRECORDBYID:
			searchRecordById(*recordMap)
		case SEARCHRECORDBYNAME:
			searchRecordByName(*recordMap)
		case EXIT:
			return
		default:
			fmt.Println("Invalid choice!")
		}
	}
}

func searchRecordById(recordMap record.MapRecord) {
	var strUUID string
	fmt.Println("Enter ID:")
	if _, err := fmt.Scanln(&strUUID); err == nil {
		id, err := uuid.Parse(strUUID)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		rec, findBool := recordMap.FindByID(id)
		printRecordById(rec, findBool, id.String())
	} else {
		fmt.Println("Error: ", err.Error())
	}
}

func printRecordById(rec record.Record, flag bool, searchID string) {
	if flag {
		fmt.Println(rec.String())
	} else {
		fmt.Println(`No ID "`, searchID, `" found!`)
	}
}

func searchRecordByName(recordMap record.MapRecord) {
	var Name string
	fmt.Println("Enter a Name or part of a Name:")
	if _, err := fmt.Scanln(&Name); err == nil {
		recs, findBool := recordMap.FindByName(Name)
		printRecordByName(recs, findBool, Name)
	} else {
		fmt.Println("Error: ", err.Error())
	}
}

func printRecordByName(recs []record.Record, flag bool, searchName string) {
	if flag {
		for _, rec := range recs {
			fmt.Println(rec.String())
		}
	} else {
		fmt.Println(`No Name "`, searchName, `" found!`)
	}
}
