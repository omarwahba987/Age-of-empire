# Age-of-empire
Age of empire command line program which takes a name of Empires name as a user input 
and search for information about this unit then render this unit’s data to the user.

-Installation

	Before first use :

	1- Age of empire  use golang so you should install golang in your machine and two external packages so you should import them 
	

	2- there alot of databeses but this is a small program so i choose a fast and light database  
    LevelDB key/value database in the Go 
		you can get by enter in terminal 
		go get github.com/syndtr/goleveldb/leveldb

Features :
There is 2 features 
1- get information about empire you want

	in windows : you can open terminal and get to project directory and type 
	 go run empirecl.go empire_name 
     Ex: go run empirecl.go Bizantines
	
	in linux : you can open terminal and get to project directory and type 
	go build empirecl.go 
	to build it just at first time then for usage
	./empirecl empire_name 
     Ex: ./empirecl Bizantines

	
2- make unit test for the program

    just type 
    go test 
    in cmd for both linux and windows
    
Dependencies in project :
  not built in :
  
   . levelDB
   
  built in packages :
  
 
. fmt
. os
. strings
. encoding/json
. log
. net/http
. testing
