package domain

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func InternetConnected() (ok bool) {
	_, err := http.Get("http://clients3.google.com/generate_204")
	if err != nil {
		return false
	}
	return true
}

func GetEmpireInformationReqest(empireName string) (empireResponseObj empireResponseStruct, success bool) {
	if !InternetConnected(){
		log.Println("there is no internet conection")
		return empireResponseObj,false
	}
	url := "https://age-of-empires-2-api.herokuapp.com/api/v1/civilization/"
	url += empireName
	resp, err := http.Get(url)
	if err != nil {
		fmt.Print(err)

	} else {

		decoder := json.NewDecoder(resp.Body)
		decoder.DisallowUnknownFields()
		err := decoder.Decode(&empireResponseObj)
		if err != nil {
			log.Println("Error request empire information : ", empireName, err)
			return empireResponseObj, false

		}

	}

	DbAddEmpire(empireResponseObj)
	return empireResponseObj, true
}
