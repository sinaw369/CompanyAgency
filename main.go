package main

import (
	"bufio"
	"company/handler/Files"
	"company/structs"
	"company/utility"
	"company/utility/color"
	"company/utility/constant"
	"company/utility/screen"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	agencyStorage     []structs.Agency
	firstTime         = true
	region            *string
	serializationMode string
	path              = "repository/"
	logoPath          = "repository/logo.txt"
	formPath          = "repository/form.txt"
)

func main() {
	agencyStorage = Files.Load(path, serializationMode)
	var logo = Files.Loadlogo(logoPath)
	region = flag.String("region", "tehran", "Place of registration of the agency")
	serializeMode := flag.String("serialize-mode", constant.Mode_json, "serialization mode to write data to file")
	command := flag.String("command", "no-command", "command to run")
	flag.Parse()
	scanner := bufio.NewScanner(os.Stdin)
	switch *serializeMode {
	case constant.Mode_txt:
		serializationMode = constant.Mode_txt
	default:
		serializationMode = constant.Mode_json
	}

	for {

		if firstTime == false {
			runCommand(*command)
			screen.ClearScreen()
			fmt.Println(color.Red + logo + color.Green)
			runCommand("default")
			fmt.Println("please enter a command")
			scanner.Scan()
			*command = scanner.Text()

		} else {
			screen.ClearScreen()
			fmt.Println(color.Red + logo + color.Green)
			runCommand(*command)
			//scanner := bufio.NewScanner(os.Stdin)
			fmt.Println("please enter a command")
			scanner.Scan()
			*command = scanner.Text()
			firstTime = false
		}

	}

}
func runCommand(command string) {
	switch command {
	case "create", "c":
		createAgency()
		utility.PressToContinue()
	case "get", "g":
		screen.ClearScreen()
		getAgency()
		utility.PressToContinue()
	case "edit", "e":
		screen.ClearScreen()
		editAgency()
		utility.PressToContinue()
	case "list", "l":
		screen.ClearScreen()
		listAgency()
		utility.PressToContinue()
	case "status", "s":
		screen.ClearScreen()
		statusAgency()
		utility.PressToContinue()
	case "change-region", "c-r":
		screen.ClearScreen()
		changeRegion()
		utility.PressToContinue()
	case "exit":
		screen.ClearScreen()
		fmt.Println(color.Reset)
		os.Exit(0)

	default:
		fmt.Println("the commands are:")
		fmt.Println("create | c            Registration of agency")
		fmt.Println("get | g               Receive agency information with ID")
		fmt.Println("edit | e              Edit agency information with ID")
		fmt.Println("list | l              List of all agencies")
		fmt.Println("status | s            Number of agencies and employees")
		fmt.Println("change-region | c-r   change region")
		fmt.Println("exit                  Exit")

	}
}
func createAgency() {
Again:
	screen.ClearScreen()
	screen.Moveto(1, 1)
	var msg = Files.Loadlogo(formPath)
	fmt.Println(color.Red + msg + color.Reset)
	scanner := bufio.NewScanner(os.Stdin)
	var agencyID string
	var Employees string
	var Name string
	var Address string
	var PhoneNumber string
	var DateOfMembership string
	fmt.Println(color.Green)
	screen.Moveto(7, 36)
	scanner.Scan()
	Name = scanner.Text()
	screen.Moveto(10, 36)
	scanner.Scan()
	Address = scanner.Text()
	screen.Moveto(13, 36)
	scanner.Scan()
	PhoneNumber = scanner.Text()
	screen.Moveto(16, 36)
	scanner.Scan()
	DateOfMembership = scanner.Text()
	screen.Moveto(19, 36)
	scanner.Scan()
	Employees = scanner.Text()
	integerAgencyEmployees, err := strconv.Atoi(Employees)
	agencyEmployees := uint(integerAgencyEmployees)
	if err != nil {
		fmt.Println(color.Yellow)
		screen.Moveto(26, 1)
		fmt.Println("agencyEmployees is not a valid integer", err)
		fmt.Println("Please fill out the form again")
		utility.PressToContinue()
		goto Again

	}

	screen.Moveto(22, 36)
	scanner.Scan()
	agencyID = scanner.Text()
	integerID, err := strconv.Atoi(agencyID)
	ID := uint(integerID)
	agency := structs.Agency{
		ID:               ID,
		Name:             Name,
		Address:          Address,
		PhoneNumber:      PhoneNumber,
		DateOfMembership: DateOfMembership,
		AgencyEmployees:  agencyEmployees,
		Region:           *region,
	}
	screen.Moveto(26, 1)
	fmt.Println("Agency successfully created")
	Files.Write(agency, path, serializationMode)
	agencyStorage = Files.Load(path, serializationMode)
}
func getAgency() {
	var agency structs.Agency
	//var i int
	fmt.Println("Please enter agency ID")
	var agencyID string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	agencyID = scanner.Text()
	aID, _ := strconv.Atoi(agencyID)
	uID := uint(aID)
	fmt.Println(color.Red)
	for _, agency = range agencyStorage {
		if agency.ID == uID {
			fmt.Printf("+%s+\n", strings.Repeat("-", 160))
			fmt.Printf("| %-30s | %-30s | %-30s | %-30s | %-30s | %-30s | %-30s |\n",
				fmt.Sprintf("%sRegion%s", color.Green, color.Red),
				fmt.Sprintf("%sID%s", color.Green, color.Red),
				fmt.Sprintf("%sName%s", color.Green, color.Red),
				fmt.Sprintf("%sAddress%s", color.Green, color.Red),
				fmt.Sprintf("%sPhoneNumber%s", color.Green, color.Red),
				fmt.Sprintf("%sDateOfMemberShip%s", color.Green, color.Red),
				fmt.Sprintf("%sAgencyEmployees%s", color.Green, color.Red))
			fmt.Printf("|%s|\n", strings.Repeat("-", 160))
			fmt.Printf("| %-30s | %-30s | %-30s | %-30s | %-30s | %-30s | %-30s |\n",
				fmt.Sprintf("%s%s%s", color.Green, *region, color.Red),
				fmt.Sprintf("%s%d%s", color.Green, agency.ID, color.Red),
				fmt.Sprintf("%s%s%s", color.Green, agency.Name, color.Red),
				fmt.Sprintf("%s%s%s", color.Green, agency.Address, color.Red),
				fmt.Sprintf("%s%s%s", color.Green, agency.PhoneNumber, color.Red),
				fmt.Sprintf("%s%s%s", color.Green, agency.DateOfMembership, color.Red),
				fmt.Sprintf("%s%d%s", color.Green, agency.AgencyEmployees, color.Red))
			fmt.Printf("+%s+\n", strings.Repeat("-", 160))
			fmt.Println(color.Green)

		}
	}
}
func editAgency() {
	fmt.Println("Please enter agency ID")
	var agencyID string
	var tempStr string
	var val string
	var flg = true
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	agencyID = scanner.Text()
	aID, _ := strconv.Atoi(agencyID)
	uID := uint(aID)
	fmt.Println(color.Red)
	for i, agency := range agencyStorage {
		if agency.ID == uID {
			flg = false
			fmt.Printf("+%s+\n", strings.Repeat("-", 160))
			fmt.Printf("| %-30s | %-30s | %-30s | %-30s | %-30s | %-30s | %-30s |\n",
				fmt.Sprintf("%sRegion%s", color.Green, color.Red),
				fmt.Sprintf("%sID%s", color.Green, color.Red),
				fmt.Sprintf("%sName%s", color.Green, color.Red),
				fmt.Sprintf("%sAddress%s", color.Green, color.Red),
				fmt.Sprintf("%sPhoneNumber%s", color.Green, color.Red),
				fmt.Sprintf("%sDateOfMemberShip%s", color.Green, color.Red),
				fmt.Sprintf("%sAgencyEmployees%s", color.Green, color.Red))
			fmt.Printf("|%s|\n", strings.Repeat("-", 160))
			fmt.Printf("| %-30s | %-30s | %-30s | %-30s | %-30s | %-30s | %-30s |\n",
				fmt.Sprintf("%s%s%s", color.Green, *region, color.Red),
				fmt.Sprintf("%s%d%s", color.Green, agency.ID, color.Red),
				fmt.Sprintf("%s%s%s", color.Green, agency.Name, color.Red),
				fmt.Sprintf("%s%s%s", color.Green, agency.Address, color.Red),
				fmt.Sprintf("%s%s%s", color.Green, agency.PhoneNumber, color.Red),
				fmt.Sprintf("%s%s%s", color.Green, agency.DateOfMembership, color.Red),
				fmt.Sprintf("%s%d%s", color.Green, agency.AgencyEmployees, color.Red))
			fmt.Printf("+%s+\n", strings.Repeat("-", 160))
			fmt.Println(color.Green)
			fmt.Println("Enter the field name")
			scanner.Scan()
			tempStr = scanner.Text()
			tempStr = strings.ToLower(tempStr)
			fmt.Println("Enter the new " + tempStr + " value")
			scanner.Scan()
			val = scanner.Text()
			switch tempStr {
			case "id":
				ID, _ := strconv.Atoi(val)
				uuID := uint(ID)
				agencyStorage[i].ID = uuID
			case "name":
				agencyStorage[i].Name = val
			case "address":
				agencyStorage[i].Address = val
			case "phonenumber":
				agencyStorage[i].PhoneNumber = val
			case "dateofmembership":
				agencyStorage[i].DateOfMembership = val
			case "agencyemployees":
				aEmployees, _ := strconv.Atoi(val)
				uEmployees := uint(aEmployees)
				agencyStorage[i].AgencyEmployees = uEmployees
			case "region":
				agencyStorage[i].Region = val

			default:
				fmt.Println("The entered value is invalid")
			}

		}
	}
	if flg {
		fmt.Println("There is no agency with this ID")
	}
	_, err := Files.EasyUpdate(path, agencyStorage, serializationMode)
	if err != nil {
		fmt.Println("can't update the file :", err)
	} else {
		agencyStorage = Files.Load(path, serializationMode)
		fmt.Println("The file has been updated")
	}
}
func listAgency() {
	fmt.Println(color.Red)
	fmt.Printf("+%s+\n", strings.Repeat("-", 160))
	fmt.Printf("| %-30s | %-30s | %-30s | %-30s | %-30s | %-30s | %-30s |\n",
		fmt.Sprintf("%sRegion%s", color.Green, color.Red),
		fmt.Sprintf("%sID%s", color.Green, color.Red),
		fmt.Sprintf("%sName%s", color.Green, color.Red),
		fmt.Sprintf("%sAddress%s", color.Green, color.Red),
		fmt.Sprintf("%sPhoneNumber%s", color.Green, color.Red),
		fmt.Sprintf("%sDateOfMemberShip%s", color.Green, color.Red),
		fmt.Sprintf("%sAgencyEmployees%s", color.Green, color.Red))
	fmt.Printf("|%s|\n", strings.Repeat("-", 160))

	for i, agency := range agencyStorage {
		if agency.Region == *region {
			fmt.Printf("| %-30s | %-30s | %-30s | %-30s | %-30s | %-30s | %-30s |\n",
				fmt.Sprintf("%s%s%s", color.Green, *region, color.Red),
				fmt.Sprintf("%s%d%s", color.Green, agency.ID, color.Red),
				fmt.Sprintf("%s%s%s", color.Green, agency.Name, color.Red),
				fmt.Sprintf("%s%s%s", color.Green, agency.Address, color.Red),
				fmt.Sprintf("%s%s%s", color.Green, agency.PhoneNumber, color.Red),
				fmt.Sprintf("%s%s%s", color.Green, agency.DateOfMembership, color.Red),
				fmt.Sprintf("%s%d%s", color.Green, agency.AgencyEmployees, color.Red))
			if len(agencyStorage) > i+1 && agencyStorage[i+1].Region == *region {
				fmt.Printf("|%s|\n", strings.Repeat("-", 160))
			}

		}
	}
	fmt.Printf("+%s+\n", strings.Repeat("-", 160))
	fmt.Println(color.Green)

}
func statusAgency() {
	var countEmp uint
	var countAgency uint
	for _, agency := range agencyStorage {
		if agency.Region == *region {
			countAgency++
			countEmp += agency.AgencyEmployees
		}
	}
	fmt.Println(color.Red)
	fmt.Printf("+%s+\n", strings.Repeat("-", 68))
	//To use the color in the string padding in the fmt.printf function, add the original value to 10
	fmt.Printf("| %-30s | %-30s | %-30s |\n",
		fmt.Sprintf("%sRegion%s", color.Green, color.Red),
		fmt.Sprintf("%sNumber of Agency%s", color.Green, color.Red),
		fmt.Sprintf("%sNumber of Employees%s", color.Green, color.Red))
	fmt.Printf("|%s|\n", strings.Repeat("-", 68))
	fmt.Printf("| %-30s | %-30s | %-30s |\n",
		fmt.Sprintf("%s%s%s", color.Green, *region, color.Red),
		fmt.Sprintf("%s%d%s", color.Green, countAgency, color.Red),
		fmt.Sprintf("%s%d%s", color.Green, countEmp, color.Red))
	fmt.Printf("+%s+\n", strings.Repeat("-", 68))
	fmt.Println(color.Green)
}
func changeRegion() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("please enter new region")
	scanner.Scan()
	testStr := scanner.Text()
	testStr = strings.ToLower(testStr)
	_, err := utility.Alpha(testStr)
	if err != nil {
		fmt.Println("The region format is incorrect")
		return
	}
	fmt.Println("The region was successfully changed")
	*region = testStr
}
