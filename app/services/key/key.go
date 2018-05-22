package key

import (
	"database/sql"
	"time"

	"github.com/epointpayment/key_management_system/app/models"
	"github.com/epointpayment/key_management_system/app/services/keygen"

	dbx "github.com/go-ozzo/ozzo-dbx"
)

// DB is the database handler
var DB *dbx.DB

// KeyService is a service that manages keys
type KeyService struct {
	userID int
	g      *keygen.KeyGenService
}

// New creates an instance of the service
func New(userID int) *KeyService {
	kg := keygen.New()

	return &KeyService{
		userID: userID,
		g:      kg,
	}
}

// Generate creates a new key and stores it in the database
func (ks *KeyService) Generate() (key *models.Key, err error) {
	key = new(models.Key)

	// Generate a key
	str, err := ks.g.Generate()
	if err != nil {
		return
	}

	// Insert into database
	tx, err := DB.Begin()
	if err != nil {
		return
	}

	key.UserID = ks.userID
	key.Key = str
	key.DateCreated = time.Now().UTC()

	err = tx.Model(key).Insert()
	if err != nil {
		tx.Rollback()
		return
	}

	err = tx.Commit()
	return
}

// Get retrieves a key from the database
func (ks *KeyService) Get(keyID int) (key *models.Key, err error) {
	key = new(models.Key)

	err = DB.Select().
		Where(dbx.HashExp{"id": keyID, "user_id": ks.userID}).
		One(key)
	if err != nil {
		if err == sql.ErrNoRows {
			err = ErrKeyNotFound
		}
		return nil, err
	}

	return
}
