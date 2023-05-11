package repository

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"lab4/user"
)

type DataRepository struct {
	path string
}

func NewDataRepository(path string) *DataRepository {
	return &DataRepository{path: path}
}

func (dr *DataRepository) Save(d *user.Data, path string) {
	data, err := json.MarshalIndent(&d, "", " ")
	if err != nil {
		fmt.Println("can not save")
		return
	}
	err = ioutil.WriteFile(path, data, 0644)
	if err != nil {
		fmt.Println("can not save 2")
		return
	}
}

func (dr *DataRepository) Open(d *user.Data, path string) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &d)
	if err != nil {
		return
	}
}
