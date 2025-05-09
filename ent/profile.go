// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/dlukt/graphql-backend-starter/ent/profile"
	"github.com/rs/xid"
)

// Profile is the model entity for the Profile schema.
type Profile struct {
	config `json:"-"`
	// ID of the ent.
	ID xid.ID `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// Sub holds the value of the "sub" field.
	Sub string `json:"sub,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Gender holds the value of the "gender" field.
	Gender       string `json:"gender,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Profile) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case profile.FieldSub, profile.FieldName, profile.FieldGender:
			values[i] = new(sql.NullString)
		case profile.FieldCreateTime, profile.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		case profile.FieldID:
			values[i] = new(xid.ID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Profile fields.
func (pr *Profile) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case profile.FieldID:
			if value, ok := values[i].(*xid.ID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				pr.ID = *value
			}
		case profile.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				pr.CreateTime = value.Time
			}
		case profile.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				pr.UpdateTime = value.Time
			}
		case profile.FieldSub:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field sub", values[i])
			} else if value.Valid {
				pr.Sub = value.String
			}
		case profile.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				pr.Name = value.String
			}
		case profile.FieldGender:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field gender", values[i])
			} else if value.Valid {
				pr.Gender = value.String
			}
		default:
			pr.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Profile.
// This includes values selected through modifiers, order, etc.
func (pr *Profile) Value(name string) (ent.Value, error) {
	return pr.selectValues.Get(name)
}

// Update returns a builder for updating this Profile.
// Note that you need to call Profile.Unwrap() before calling this method if this Profile
// was returned from a transaction, and the transaction was committed or rolled back.
func (pr *Profile) Update() *ProfileUpdateOne {
	return NewProfileClient(pr.config).UpdateOne(pr)
}

// Unwrap unwraps the Profile entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pr *Profile) Unwrap() *Profile {
	_tx, ok := pr.config.driver.(*txDriver)
	if !ok {
		panic("ent: Profile is not a transactional entity")
	}
	pr.config.driver = _tx.drv
	return pr
}

// String implements the fmt.Stringer.
func (pr *Profile) String() string {
	var builder strings.Builder
	builder.WriteString("Profile(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pr.ID))
	builder.WriteString("create_time=")
	builder.WriteString(pr.CreateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("update_time=")
	builder.WriteString(pr.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("sub=")
	builder.WriteString(pr.Sub)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(pr.Name)
	builder.WriteString(", ")
	builder.WriteString("gender=")
	builder.WriteString(pr.Gender)
	builder.WriteByte(')')
	return builder.String()
}

// Profiles is a parsable slice of Profile.
type Profiles []*Profile
