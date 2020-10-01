package domain

import (
	"encoding/json"
	"github.com/syndtr/goleveldb/leveldb"
	"log"
)

var DB *leveldb.DB
var Open = false

func opendatabase() bool {
	if !Open {
		Open = true
		DBpath := "Database/Empire"
		var err error
		DB, err = leveldb.OpenFile(DBpath, nil)
		if err != nil {
			log.Println("error open DB : ", err)
			return false
		}

		return true
	}
	return true

}
func closedatabase() bool {
	var err error
	Open = false
	err = DB.Close()
	if err != nil {
		log.Println("error close DB : ", err)
		return false
	}

	return true
}

func DbAddEmpire(data empireResponseStruct) bool {

	opendatabase()
	d, _ := json.Marshal(data)

	err := DB.Put([]byte(data.Name), d, nil)
	defer closedatabase()

	if err != nil {
		log.Println("error add DB : ", err)

		return false
	}

	return true
}
func DbGetEmpireByName(name string) (empire empireResponseStruct, exist bool) {
	opendatabase()
	data, err := DB.Get([]byte(name), nil)
	defer closedatabase()
	if err != nil {

		return empire, false
	}

	json.Unmarshal(data, &empire)

	return empire, true
}

