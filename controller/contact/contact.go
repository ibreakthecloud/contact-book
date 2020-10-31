package contact

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func New() *Contact {
	return &Contact{}
}

// Add is used to create new contact with valid and unique email
func (c *Contact) Add(gc *gin.Context) {
	// verify if request payload is valid
	err := gc.BindJSON(c)
	if err != nil {
		gc.JSON(http.StatusInternalServerError, gin.H{"error": "unable to parse request payload"})
		return
	}

	// validate the email

	// Add the contact

}

// Update is used to edit the contact present in the book
func (c *Contact) Update(gc *gin.Context) {}

// Delete is used to delete the contact present in the book
func (c *Contact) Delete(gc *gin.Context) {}