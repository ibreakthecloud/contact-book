package contact

import (
	"github.com/gin-gonic/gin"
	"github.com/ibreakthecloud/contact-book/store"
	"github.com/ibreakthecloud/contact-book/util"
	"log"
	"net/http"
	"strconv"
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
	// 1. Check if email is valid -- done
	// 2. Check if email already exists?
	if !util.IsValidEmail(c.Email) {
		gc.JSON(http.StatusUnprocessableEntity, gin.H{"error":"invalid email"})
		return
	}

	if contactExists(c.Email) {
		gc.JSON(http.StatusUnprocessableEntity, gin.H{"error":"contact already exists"})
		return
	}

	// Add the contact
	err = store.NewStore.AddContact(c.Name, c.Email)
	if err != nil {
		log.Print(err)
		gc.JSON(http.StatusInternalServerError, gin.H{"error": "unable to add contact"})
		return
	}

	gc.JSON(http.StatusOK, gin.H{"message": "contact added successfully"})
	return
}

// Update is used to edit the contact present in the book
func (c *Contact) Update(gc *gin.Context) {}

// Delete is used to delete the contact present in the book
// Since the email is unique, delete is suppose to support only
// via email
func (c *Contact) Delete(gc *gin.Context) {
	// verify if request payload is valid
	err := gc.BindJSON(c)
	if err != nil {
		log.Print(err)
		gc.JSON(http.StatusInternalServerError, gin.H{"error": "unable to parse request payload"})
		return
	}

	if c.Email == "" {
		gc.JSON(http.StatusFailedDependency, gin.H{"error": "missing email address"})
		return
	}

	err = store.NewStore.DeleteContact(c.Email)
	if err != nil {
		log.Print(err)
		gc.JSON(http.StatusInternalServerError, gin.H{"error": "something went wrong"})
		return
	}

	gc.JSON(http.StatusOK, gin.H{"message": "contact deleted successfully"})
	return
}

// Get is used to get the contact
// if no parameters are provided, it returns the full book
// or searches on the basis of name or email
func (c Contact) Get(gc *gin.Context) {
	_ = gc.BindJSON(c)

	var err error
	var page = 0
	var name, email string

	queryParams := gc.Request.URL.Query()

	// check for page param
	if queryParams["page"] != nil {
		page, err = strconv.Atoi(queryParams["page"][0])
		if err != nil {
			gc.JSON(http.StatusUnprocessableEntity, gin.H{"error":"invalid page"})
			return
		}
	}

	// check for name param for searching
	if queryParams["name"] != nil {
		name = queryParams["name"][0]
	}

	// check for email param for searching
	if queryParams["email"] != nil {
		email = queryParams["email"][0]
	}

	contacts, err := store.NewStore.Get(name, email, page)
	if err != nil {
		log.Print(err)
		gc.JSON(http.StatusInternalServerError, gin.H{"error": "unable to get contact"})
		return
	}
	gc.JSON(http.StatusOK, contacts)
	return
}

// utils functions

func contactExists(email string) bool {
	contact, err := store.NewStore.Get("", email, 1)
	if err != nil {
		log.Print(err)
		return false
	}
	if contact != nil {
		return true
	}
	return false
}