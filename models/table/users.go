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

// User - Define the users table
// Implements TableRecordInterface
type User struct {
	tr       *record.TableRecord
	RecordID int64   `json:"id" db:"record_id"`
	Name     *string `json:"name" db:"name"`
	Lastname *string `json:"lastname" db:"lastname"`
	Gender   *string `json:"gender" db:"gender"`
}

var du = &User{}

// LoadAllUsers - Returns all users
func LoadAllUsers() ([]*User, error) {

	query := "SELECT " + record.AllField(du) + " FROM " + du.GetTableName()

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}

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

// NewUser - Returns new instance of the struct with the init of TableRecord and set the flag "isNew"
// It's suggest to use always this to get new instance of the struct
func NewUser(db db.SQLConnector) *User {

	u := new(User)
	u.tr = record.NewTableRecord(true, false)
	u.tr.SetSQLConnection(db)

	return u
}

// GetTableRecord - Returns TableRecord instance of the User struct
func (u User) GetTableRecord() *record.TableRecord {
	return u.tr
}

// GetPrimaryKeyName - Returns primary key name
func (u User) GetPrimaryKeyName() string {
	return UsersColRecordID
}

// GetPrimaryKeyValue - Returns the value of the primary key
func (u User) GetPrimaryKeyValue() int64 {
	return u.RecordID
}

// GetTableName - Returns table name
func (u User) GetTableName() string {
	return UsersTableName
}
