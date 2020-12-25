package table

import (
	"github.com/IacopoMelani/Go-Starter-Project/app/models/dto"
	"github.com/IacopoMelani/Go-Starter-Project/pkg/manager/db"
	record "github.com/IacopoMelani/Go-Starter-Project/pkg/models/table_record"
	"github.com/jmoiron/sqlx"
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
	u.TableRecord = record.NewTableRecord(true, false)
	u.SetSQLConnection(db)

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

// MARK: User unexported

// loadAllUsersFromRows - Loads all users from sqlx rows result
func loadAllUsersFromRows(db db.SQLConnector, rows *sqlx.Rows) ([]*User, error) {

	var result []*User

	for rows.Next() {

		u := NewUser(db)

		if err := record.LoadFromRow(rows, u); err != nil {
			return nil, err
		}

		result = append(result, u)
	}

	return result, nil
}

// MARK: User exported

// LoadAllUsers - Returns all users
func LoadAllUsers(db db.SQLConnector) ([]*User, error) {

	query := "SELECT " + record.AllField(&User{}) + " FROM " + UsersTableName

	rows, err := db.Queryx(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	return loadAllUsersFromRows(db, rows)
}
