package main

import (
	"internproject/controllers"
	"internproject/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/loans", controllers.CreateLoan)
	r.GET("/getloans", controllers.GetLoans)
	r.GET("/getloan/:id", controllers.GetLoan)
	r.GET("/getloans/query/", controllers.GetLoanparams)
	r.PATCH("/Delete/loans/:id", controllers.DeleteLoan)
	r.PATCH("/approve/:id", controllers.ApproveLoan)
	models.ConnectDataBase()
	r.Run(":8080")
}
