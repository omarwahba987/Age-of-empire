package domain

import "strings"

func (p empireResponseStruct) String() string {

	return " Empire name :  " + p.Name + "\n Empire expansion is : " + p.Expansion + "\n Empire army type : " + p.Army_type +
		"\n More information : " +
		strings.Join(p.Civilization_bonus, ",\n   ") + "\n"

}
func GetEmpireInformation(empireName string) (empireResponseObj empireResponseStruct, success bool, dbExist bool) {
	empireResponseObj, dbExist = DbGetEmpireByName(empireName)
	if !dbExist {
		empireResponseObj, success := GetEmpireInformationReqest(empireName)
		return empireResponseObj, success, dbExist
	}
	return empireResponseObj, true, dbExist
}