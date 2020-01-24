package table

import (
	"github.com/IacopoMelani/Go-Starter-Project/pkg/manager/db"
	record "github.com/IacopoMelani/Go-Starter-Project/pkg/models/table_record"
)

// Costanti relative alla tabella users
const (
	UsersColRecordID = "record_id"
	UsersColName     = "name"
	UsersColLastname = "lastname"
	UsersColGender   = "gender"

	UsersTableName = "users"
)

// User - Struct che definisce la tabella "users"
// implementa TableRecordInterface
type User struct {
	tr       *record.TableRecord
	RecordID int64   `json:"id" db:"record_id"`
	Name     *string `json:"name" db:"name"`
	Lastname *string `json:"lastname" db:"lastname"`
	Gender   *string `json:"gender" db:"gender"`
}

var du = &User{}

// LoadAllUsers - Si occupa di restituire tutti gli utenti presenti nel database
func LoadAllUsers() ([]*User, error) {

	query := "SELECT " + record.AllField(du) + " FROM " + du.GetTableName()

	rows := db.QueryOrPanic(query)

	defer rows.Close()

	var result []*User

	for rows.Next() {

		u := NewUser(db.GetConnection())

		if err := record.LoadFromRow(rows, u); err != nil {
			return nil, err
		}

		result = append(result, u)
	}

	return result, nil
}

// NewUser - Si occupa di istanziare un nuovo oggetto User istanziando il relativo TableRecord e impostandolo come "nuovo"
// È consigliato utilizzare sempre questo metodo per creare una nuova istanza di User
func NewUser(db db.SQLConnector) *User {

	u := new(User)
	u.tr = record.NewTableRecord(true, false)
	u.tr.SetSQLConnection(db)

	return u
}

// GetTableRecord - Restituisce l'istanza di TableRecord
func (u User) GetTableRecord() *record.TableRecord {
	return u.tr
}

// GetPrimaryKeyName - Restituisce il nome della chiave primaria
func (u User) GetPrimaryKeyName() string {
	return UsersColRecordID
}

// GetPrimaryKeyValue - Restituisce l'indirizzo di memoria del valore della chiave primaria
func (u User) GetPrimaryKeyValue() int64 {
	return u.RecordID
}

// GetTableName - Restituisce il nome della tabella
func (u User) GetTableName() string {
	return UsersTableName
}
