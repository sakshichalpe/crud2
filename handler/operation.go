package operations

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

type employee struct {
	Name string `json:"name"`
	Id   string `json:"id"`
}

var e []employee

//var e emp //{ "Name": "Jane",     "Id": "1" }

// using ioutils
func InsertOneorMore(c *gin.Context) {
	jsonbyte, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Println("err at reading body")
	}
	err = json.Unmarshal(jsonbyte, &e)
	if err != nil {
		fmt.Println("error while unmarshallling", err)
	}
	c.JSON(200, e)
}
func GetAll(c *gin.Context) {
	fmt.Println(e)
	c.JSON(200, e)
}
func GetOne(c *gin.Context) {
	//get := c.Param("id")
	get := c.Query("id")
	var foundPerson *employee
	for _, v := range e {
		if v.Id == get {
			foundPerson = &v
			break
		}
	}
	if foundPerson == nil {
		fmt.Println("not found")
		return
	}
	c.JSON(200, foundPerson)
}
func DeleteOne(c *gin.Context) {
	del := c.Param("id")
	for k, v := range e {
		if v.Id == del {
			e = append(e[:k], e[k+1:]...)
		}
	}
	c.JSON(200, e)
}

func Update(c *gin.Context) {
	jsonbyte, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Println("error while reading the body")
		return
	}
	var ee employee
	if err = json.Unmarshal(jsonbyte, &ee); err != nil { //unmarshal: json to struct
		fmt.Println("error in unmarshall", err)
	}
	var updated employee
	for i := 0; i < len(e); i++ {
		if e[i].Id == ee.Id {
			e[i].Name = ee.Name
			updated = e[i]
		}
	}
	c.JSON(200, updated)
}
