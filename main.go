package main

import (
	"internproject/controllers"
	"internproject/models"

	"github.com/gin-gonic/gin"
)

func main() {
	// gin router setup
	r := gin.Default()
	r.POST("/loans", controllers.CreateLoan)
	r.GET("/loans", controllers.GetLoans)
	r.GET("/loan/:id", controllers.GetLoan)
	r.GET("/getloans/query/", controllers.GetLoanparams)
	r.GET("/getloans/loans/", controllers.GetLoanparams2)
	r.PATCH("Delete/:id", controllers.DeleteLoan)
	r.PATCH("/loans/:id", controllers.ApproveLoan)
	models.ConnectDataBase() // call of database method
	r.Run(":8080")
}
