package main

import (
	operations "crud2/handler"

	"github.com/gin-gonic/gin"
)

// restAPI amd API:
func main() {
	r := gin.Default()
	//db.DBConnection()
	r.POST("/insertone", operations.InsertOneorMore)
	r.GET("/fetch", operations.GetAll)
	r.POST("/deleteone/:id", operations.DeleteOne) //path variable
	r.POST("/getone", operations.GetOne)           //query param
	//comparasiton: path & query -->to make url clean we using path
	r.POST("/update", operations.Update)

	r.Run()

	// operations.Delete()

}
