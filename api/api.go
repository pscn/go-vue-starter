package api

import "github.com/pscn/go-vue-starter/models"

// API -
type API struct {
	users  *models.UserFactory
	quotes *models.QuoteManager
}

// NewAPI -
func NewAPI(db *models.DB) *API {

	usermgr, _ := models.NewUserFactory(db)
	quotemgr, _ := models.NewQuoteManager(db)

	return &API{
		users:  usermgr,
		quotes: quotemgr,
	}
}
