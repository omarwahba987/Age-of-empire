package main

import (
	"ageofempire/domain"
	"fmt"
	"log"
	"testing"
)

func TestEmpire(t *testing.T) {
	fmt.Println("\n\n ** you can run test twice then delete database to test again**")
	if domain.InternetConnected(){
		log.Println("internet connection is good")
	}else{
		log.Println("there is no internet conection")
	}



}

func TestRequestAndAddDb(t *testing.T) {
	log.Print("Test requests and adding in database")
	var testRequestCase = []string{"Japanese", "Koreans", "Mayans"}
	for _, empireName := range testRequestCase {
		empireResponseObj, success, dbExist := domain.GetEmpireInformation(empireName)
		if success {
			log.Println("Empire response \n", empireResponseObj)
			log.Print("success value : ", success)
			log.Print("exist in database : ", dbExist, "\n\n\n\n")
		}
	}
}

func TestGetFromDb(t *testing.T) {
	log.Print("Testgetting data from database")
	var testDbtCase = []string{"Japanese", "Koreans", "Mayans"}
	for _, empireName := range testDbtCase {
		empireResponseObj, success, dbExist := domain.GetEmpireInformation(empireName)
		if success {
			log.Println("Empire response \n", empireResponseObj)
			log.Print("success value : ", success)
			log.Print("exist in database : ", dbExist,"\n\n\n\n")
		}
	}
}
func TestRongValue(t *testing.T) {
	log.Print("Test requests rong values")
	var testFalsetCase = []string{"Japan", "Kor", "Maya"}
	for _, empireName := range testFalsetCase {
		empireResponseObj, success, dbExist := domain.GetEmpireInformation(empireName)
		if success {
			log.Println("Empire response \n", empireResponseObj)
			log.Print("success value : ", success)
			log.Print("exist in database : ", dbExist,"\n\n\n")
		} else {
			log.Print("success value : ", success, "\n there is no data available")

		}
	}
}



