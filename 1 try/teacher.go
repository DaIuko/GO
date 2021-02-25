package main
import (
	"fmt"
	"encoding/json"
)
type Teacher struct {
    ID         string   `json:"id"`
    Subject    string   `json:"subject"`
    Salary     float64  `json:"salary"`
    Classrooms []string `json:"classrooms"`
    Person     struct {
        Name         string `json:"name"`
        Surname      string `json:"surname"`
        PersonalCode string `json:"personalCode"`
    } `json:"person"`
}
func (t Teacher) GetCreateAction() DefinedAction  {
	return &CreateTeacher{}
}
type CreateTeacher struct {
	T Teacher `json:"data"`
}
func (action *CreateTeacher) GetFromJSON( rawData []byte) {
	err := json.Unmarshal(rawData, action )	
	if err != nil {
		fmt.Println(err )
	}
}
func (action CreateTeacher) Process() {
	fmt.Println("Teacher created")
}

type UpdateTeacher struct {
	T Teacher `json:"data"`
}
func (action UpdateTeacher) Process() {
	fmt.Println("Teacher created")
}

type ReadTeacher struct {
	T Teacher `json:"data"`
}
func (action ReadTeacher ) Process() {
	fmt.Println("Teacher created")
}

type DeleteTeacher struct {
	T Teacher `json:"data"`
}
func (action DeleteTeacher) Process() {
	fmt.Println("Teacher created")
}
