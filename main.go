package main

import (
	"fmt"
	"encoding/json"
)
type Action struct {
	Action string `json:"action"`
	ObjectName string `json:"object"`
}

type GeneralObject interface {
	GetCreateAction() DefinedAction  
}



type DefinedAction interface {
	Process() 
	GetFromJSON( rawData []byte)
}


func main() {
	strjson := `
   {
      "action":"create",
      "object":"Teacher",
      "data":{
         "subject":"Chemistry",
         "salary":192.4,
         "classrooms":[
            "CLSR-001",
            "CLSR-002",
            "CLSR-003"
         ],
         "person":{
            "name":"Anatolij",
            "surname":"Kosolapenko",
            "personalCode":"101080-10300"
         }
      }
   }`
	rawData := []byte(strjson )
	var a Action
	err := json.Unmarshal(rawData, &a)
	if err != nil {
		fmt.Println(err )
	}
	var actionObj GeneralObject
	if a.ObjectName  == "Teacher" {
		actionObj = &Teacher{}
	}
	var toDo DefinedAction
	if a.Action == "create" {
		toDo = actionObj.GetCreateAction()
	}
	toDo.GetFromJSON(rawData)
	
	toDo.Process()
}
