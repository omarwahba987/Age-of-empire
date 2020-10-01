package main

import (
	"ageofempire/domain"
	"fmt"
	"os"
	"strings"
)

func main() {
	inputEmpireName := os.Args[1:]
	empireName := strings.Join(inputEmpireName," ")
	empireResponseObj, success, _ := domain.GetEmpireInformation(empireName)
	if success {
		fmt.Println(empireResponseObj)
	}

}
