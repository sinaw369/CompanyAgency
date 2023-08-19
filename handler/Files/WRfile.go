package Files

import (
	"company/structs"
	"company/utility"
	"company/utility/constant"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Loadlogo(path string) string {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("can't open the file", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)
	var data = make([]byte, utility.CountNBOfile(path))

	_, oErr := file.Read(data)
	if oErr != nil {
		fmt.Println("can't read from the file", oErr)
		return ""
	}
	str := string(data)
	return str

}
func Load(path, serializationMode string) []structs.Agency {
	path = path + "companyStorage." + serializationMode
	var agstore []structs.Agency
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("can't open the file")

		return nil
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	var data = make([]byte, utility.CountNBOfile(path))
	_, oErr := file.Read(data)
	if oErr != nil {
		fmt.Println("can't read from the file", oErr)

		return nil
	}
	var dataStr = string(data)
	switch serializationMode {
	case constant.Mode_txt:
		agencySlice := strings.Split(dataStr, "\n")
		if len(agencySlice) >= 2 {
			agencySlice = agencySlice[:len(agencySlice)-1]
		}
		for _, a := range agencySlice {
			var agencyStruct = structs.Agency{}
			var dErr error
			agencyStruct, dErr = deserializeFormTxt(a)
			if dErr != nil {
				fmt.Println("can't deserialize agency record to agency struct", dErr)

				return nil
			}
			agstore = append(agstore, agencyStruct)
		}
	case constant.Mode_json:
		agencySlice := strings.Split(dataStr, "}")
		if len(agencySlice) >= 2 {
			agencySlice = agencySlice[:len(agencySlice)-1]
		}
		for _, a := range agencySlice {
			var agencyStruct = structs.Agency{}
			a += "}"
			err := json.Unmarshal([]byte(a), &agencyStruct)
			if err != nil {
				fmt.Println("can't deserialize agency record to agency struct with json mode", err)

				return nil
			}
			agstore = append(agstore, agencyStruct)
		}

	default:
		fmt.Println("invalid serialization mode ")

		return nil
	}
	return agstore
}
func Write(agency structs.Agency, path string, serializationMode string) {
	path = path + "companyStorage." + serializationMode
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("can't create or open file", err)

		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	var data []byte
	if serializationMode == constant.Mode_txt {
		data = []byte(fmt.Sprintf(
			"ID: %d, Name: %s, Address: %s, PhoneNumber: %s, DateOfMembership: %s, AgencyEmployees: %d, Region: %s\n",
			agency.ID,
			agency.Name,
			agency.Address,
			agency.PhoneNumber,
			agency.DateOfMembership,
			agency.AgencyEmployees,
			agency.Region,
		))
	} else if serializationMode == constant.Mode_json {
		data, err = json.MarshalIndent(agency, "", "\t")
		if err != nil {
			fmt.Println("can't marshal agency struct to json", err)

			return
		}
		data = append(data, 10) //append \n => []byte("\n")...
		//data = append(data, []byte("\n")...)

	} else {
		fmt.Println("invalid serialization mode ")

		return
	}

	_, wErr := file.Write(data)
	if wErr != nil {
		fmt.Println("can't write to the file", wErr)

		return
	}
}
func deserializeFormTxt(agencyStr string) (structs.Agency, error) {
	if agencyStr == "" {
		return structs.Agency{}, errors.New("agency string is empty")
	}
	var agency = structs.Agency{}
	userFields := strings.Split(agencyStr, ",")
	for _, field := range userFields {
		value := strings.Split(field, ": ")
		if len(value) != 2 {
			fmt.Println("field is not valid,skipping...", len(value))

			continue
		}
		fieldName := strings.ReplaceAll(value[0], " ", "")
		fieldValue := value[1]
		switch fieldName {
		case "ID":
			id, err := strconv.Atoi(fieldValue)
			if err != nil {
				return structs.Agency{}, errors.New("strconv error: ")
			}
			agency.ID = uint(id)
		case "Name":
			agency.Name = fieldValue
		case "Address":
			agency.Address = fieldValue
		case "PhoneNumber":
			agency.PhoneNumber = fieldValue
		case "DateOfMembership":
			agency.DateOfMembership = fieldValue
		case "AgencyEmployees":
			emp, err := strconv.Atoi(fieldValue)
			if err != nil {
				return structs.Agency{}, errors.New("strconv error: ")
			}
			agency.AgencyEmployees = uint(emp)

		case "Region":
			agency.Region = fieldValue
		}
	}
	return agency, nil
}
func EasyUpdate(path string, as []structs.Agency, serializationMode string) (bool, error) {
	updatePath := path + "companyStorage." + serializationMode
	f, _ := os.Open(updatePath)
	err := f.Close()
	if err != nil {
		return false, err
	}
	err = os.Truncate(updatePath, 0)
	if err != nil {
		return false, err
	}
	for _, a := range as {
		Write(a, path, serializationMode)
	}
	return true, nil
}
