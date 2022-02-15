// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/Xanonymous-GitHub/carnival/db/pkg/ent/application"
	"github.com/Xanonymous-GitHub/carnival/db/pkg/ent/applicationassignmenthistory"
	"github.com/google/uuid"
)

// ApplicationAssignmentHistory is the model entity for the ApplicationAssignmentHistory schema.
type ApplicationAssignmentHistory struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// ApplicationID holds the value of the "application_id" field.
	ApplicationID uuid.UUID `json:"application_id,omitempty"`
	// Assigner holds the value of the "assigner" field.
	Assigner string `json:"assigner,omitempty"`
	// Assignee holds the value of the "assignee" field.
	Assignee string `json:"assignee,omitempty"`
	// CreatedTime holds the value of the "created_time" field.
	CreatedTime time.Time `json:"created_time,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ApplicationAssignmentHistoryQuery when eager-loading is set.
	Edges ApplicationAssignmentHistoryEdges `json:"edges"`
}

// ApplicationAssignmentHistoryEdges holds the relations/edges for other nodes in the graph.
type ApplicationAssignmentHistoryEdges struct {
	// Applications holds the value of the applications edge.
	Applications *Application `json:"applications,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ApplicationsOrErr returns the Applications value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ApplicationAssignmentHistoryEdges) ApplicationsOrErr() (*Application, error) {
	if e.loadedTypes[0] {
		if e.Applications == nil {
			// The edge applications was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: application.Label}
		}
		return e.Applications, nil
	}
	return nil, &NotLoadedError{edge: "applications"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ApplicationAssignmentHistory) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case applicationassignmenthistory.FieldID:
			values[i] = new(sql.NullInt64)
		case applicationassignmenthistory.FieldAssigner, applicationassignmenthistory.FieldAssignee:
			values[i] = new(sql.NullString)
		case applicationassignmenthistory.FieldCreatedTime:
			values[i] = new(sql.NullTime)
		case applicationassignmenthistory.FieldApplicationID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type ApplicationAssignmentHistory", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ApplicationAssignmentHistory fields.
func (aah *ApplicationAssignmentHistory) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case applicationassignmenthistory.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			aah.ID = int(value.Int64)
		case applicationassignmenthistory.FieldApplicationID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field application_id", values[i])
			} else if value != nil {
				aah.ApplicationID = *value
			}
		case applicationassignmenthistory.FieldAssigner:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field assigner", values[i])
			} else if value.Valid {
				aah.Assigner = value.String
			}
		case applicationassignmenthistory.FieldAssignee:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field assignee", values[i])
			} else if value.Valid {
				aah.Assignee = value.String
			}
		case applicationassignmenthistory.FieldCreatedTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_time", values[i])
			} else if value.Valid {
				aah.CreatedTime = value.Time
			}
		}
	}
	return nil
}

// QueryApplications queries the "applications" edge of the ApplicationAssignmentHistory entity.
func (aah *ApplicationAssignmentHistory) QueryApplications() *ApplicationQuery {
	return (&ApplicationAssignmentHistoryClient{config: aah.config}).QueryApplications(aah)
}

// Update returns a builder for updating this ApplicationAssignmentHistory.
// Note that you need to call ApplicationAssignmentHistory.Unwrap() before calling this method if this ApplicationAssignmentHistory
// was returned from a transaction, and the transaction was committed or rolled back.
func (aah *ApplicationAssignmentHistory) Update() *ApplicationAssignmentHistoryUpdateOne {
	return (&ApplicationAssignmentHistoryClient{config: aah.config}).UpdateOne(aah)
}

// Unwrap unwraps the ApplicationAssignmentHistory entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (aah *ApplicationAssignmentHistory) Unwrap() *ApplicationAssignmentHistory {
	tx, ok := aah.config.driver.(*txDriver)
	if !ok {
		panic("ent: ApplicationAssignmentHistory is not a transactional entity")
	}
	aah.config.driver = tx.drv
	return aah
}

// String implements the fmt.Stringer.
func (aah *ApplicationAssignmentHistory) String() string {
	var builder strings.Builder
	builder.WriteString("ApplicationAssignmentHistory(")
	builder.WriteString(fmt.Sprintf("id=%v", aah.ID))
	builder.WriteString(", application_id=")
	builder.WriteString(fmt.Sprintf("%v", aah.ApplicationID))
	builder.WriteString(", assigner=")
	builder.WriteString(aah.Assigner)
	builder.WriteString(", assignee=")
	builder.WriteString(aah.Assignee)
	builder.WriteString(", created_time=")
	builder.WriteString(aah.CreatedTime.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// ApplicationAssignmentHistories is a parsable slice of ApplicationAssignmentHistory.
type ApplicationAssignmentHistories []*ApplicationAssignmentHistory

func (aah ApplicationAssignmentHistories) config(cfg config) {
	for _i := range aah {
		aah[_i].config = cfg
	}
}
