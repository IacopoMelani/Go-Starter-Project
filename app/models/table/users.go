package table

import (
	"github.com/IacopoMelani/Go-Starter-Project/app/models/dto"
	"github.com/IacopoMelani/Go-Starter-Project/pkg/manager/db"
	record "github.com/IacopoMelani/Go-Starter-Project/pkg/models/table_record"
	"gopkg.in/guregu/null.v4"
)

// Costanti relative alla tabella users
const (
	UsersColRecordID = "record_id"
	UsersColName     = "name"
	UsersColLastname = "lastname"
	UsersColGender   = "gender"

	UsersTableName = "users"
)

// MARK: User table model & constructor

// User - Define the users table
// Implements TableRecordInterface
type User struct {
	*record.TableRecord
	RecordID int64       `json:"id"       db:"record_id"`
	Name     null.String `json:"name"     db:"name"`
	Lastname null.String `json:"lastname" db:"lastname"`
	Gender   null.String `json:"gender"   db:"gender"`
}

// NewUser - Returns new instance of the struct with the init of TableRecord and set the flag "isNew"
// It's suggest to use always this to get new instance of the struct
func NewUser(db db.SQLConnector) *User {

	u := new(User)
	u.TableRecord = record.NewTableRecord(db, true, false)
	return u
}

// MARK: Mappable implementation

// Map - Implements Mappable interface
func (u User) Map() dto.DTO {
	return dto.UserDTO{
		ID:       u.RecordID,
		Name:     u.Name,
		Lastname: u.Lastname,
		Gender:   u.Gender,
	}
}

// MARK: TableRecordInterface implementation

// GetTableRecord - Returns TableRecord instance of the User struct
func (u User) GetTableRecord() *record.TableRecord {
	return u.TableRecord
}

// GetPrimaryKeyName - Returns primary key name
func (u User) GetPrimaryKeyName() string {
	return UsersColRecordID
}

// GetTableName - Returns table name
func (u User) GetTableName() string {
	return UsersTableName
}

// MARK: User exported

// LoadAllUsers - Returns all users
func LoadAllUsers(conn db.SQLConnector) ([]*User, error) {

	query := "SELECT " + record.AllField(&User{}) + " FROM " + UsersTableName

	var users []*User

	err := conn.Select(&users, query)
	if err != nil {
		return nil, err
	}

	return users, nil
}
