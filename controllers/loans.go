package controllers

import (
	"net/http"

	"internproject/models"

	"github.com/gin-gonic/gin"
)

//this strut is used for validation
type CreateLoanInput struct {
	CustomerName string `json:"customername" binding:"required"`
	PhoneNO      string `json:"phoneno" binding:"required"`
	Email        string `json:"email" binding:"required"`
	LoanAmount   int    `json:"loanamount" binding:"required"`
	CreditScore  int    `json:"creditscore" binding:"required"`
	Status       string `json:"status"`
}

//create the loan takes input from the body
func CreateLoan(c *gin.Context) {
	var check_loan CreateLoanInput
	if err := c.ShouldBindJSON(&check_loan); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	loan := models.Loan{CustomerName: check_loan.CustomerName, PhoneNO: check_loan.PhoneNO, Email: check_loan.Email,
		LoanAmount: check_loan.LoanAmount, CreditScore: check_loan.CreditScore, Status: "New"}
	models.DB.Create(&loan)
	c.JSON(200, gin.H{"Loan Request": loan})
}

//gets all the loans from the database
func GetLoans(c *gin.Context) {
	var loans []models.Loan
	models.DB.Find(&loans)

	c.JSON(200, gin.H{"All Loans": loans})
}

//gets loan with given id only (single loan)
func GetLoan(c *gin.Context) {
	var loan models.Loan
	if err := models.DB.Where("id = ?", c.Param("id")).First(&loan).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Loan Request": loan})
}

//gets loan with query parameter
func GetLoanparams(c *gin.Context) {
	var loan []models.Loan

	if err := models.DB.Where("status=? AND loan_amount >?", c.Query("status"), c.Query("loan_amount_gtr")).Find(&loan).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(200, gin.H{"results": loan})
}

//canclled the loans (apply status as canclled)
func DeleteLoan(c *gin.Context) {
	// Get model if exist
	var loan models.Loan
	if err := models.DB.Where("id = ?", c.Param("id")).First(&loan).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	var del = models.Loan{
		Status: "canclled",
	}
	models.DB.Model(&loan).Updates(del)

	c.JSON(http.StatusOK, gin.H{"Loan Request": "loan is Canclled"})
}

//approves the loans based on the credit score
func ApproveLoan(c *gin.Context) {
	var loan models.Loan
	if err := models.DB.Where("id = ?", c.Param("id")).First(&loan).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	if loan.CreditScore >= 750 {
		var del = models.Loan{
			Status: "Approved",
		}
		models.DB.Model(&loan).Updates(del)
		c.JSON(http.StatusOK, gin.H{"Loan Request": "loan is Approved"})
	} else {
		var del = models.Loan{
			Status: "Rejected",
		}
		models.DB.Model(&loan).Updates(del)
		c.JSON(http.StatusOK, gin.H{"Loan Request": "loan is rejected"})
	}
}
