// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/Xanonymous-GitHub/carnival/db/pkg/ent/application"
	"github.com/Xanonymous-GitHub/carnival/db/pkg/ent/ticket"
	"github.com/google/uuid"
)

// Application is the model entity for the Application schema.
type Application struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// BasicID holds the value of the "basic_id" field.
	BasicID string `json:"basic_id,omitempty"`
	// PremiumID holds the value of the "premium_id" field.
	PremiumID string `json:"premium_id,omitempty"`
	// BotDisplayName holds the value of the "bot_display_name" field.
	BotDisplayName string `json:"bot_display_name,omitempty"`
	// BotMid holds the value of the "bot_mid" field.
	BotMid string `json:"bot_mid,omitempty"`
	// BotActiveStatus holds the value of the "bot_active_status" field.
	BotActiveStatus application.BotActiveStatus `json:"bot_active_status,omitempty"`
	// ApplicantName holds the value of the "applicant_name" field.
	ApplicantName string `json:"applicant_name,omitempty"`
	// ApplicantEmail holds the value of the "applicant_email" field.
	ApplicantEmail string `json:"applicant_email,omitempty"`
	// ApplicationMid holds the value of the "application_mid" field.
	ApplicationMid string `json:"-"`
	// Remark holds the value of the "remark" field.
	Remark string `json:"remark,omitempty"`
	// StoreType holds the value of the "store_type" field.
	StoreType application.StoreType `json:"store_type,omitempty"`
	// WebsiteURL holds the value of the "website_url" field.
	WebsiteURL string `json:"website_url,omitempty"`
	// ApplicationStatus holds the value of the "application_status" field.
	ApplicationStatus application.ApplicationStatus `json:"application_status,omitempty"`
	// ReviewComment holds the value of the "review_comment" field.
	ReviewComment string `json:"review_comment,omitempty"`
	// Assigner holds the value of the "assigner" field.
	Assigner string `json:"assigner,omitempty"`
	// Assignee holds the value of the "assignee" field.
	Assignee string `json:"assignee,omitempty"`
	// CreatedDtime holds the value of the "created_dtime" field.
	CreatedDtime time.Time `json:"created_dtime,omitempty"`
	// UpdateDtime holds the value of the "update_dtime" field.
	UpdateDtime time.Time `json:"update_dtime,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ApplicationQuery when eager-loading is set.
	Edges ApplicationEdges `json:"edges"`
}

// ApplicationEdges holds the relations/edges for other nodes in the graph.
type ApplicationEdges struct {
	// Tickets holds the value of the tickets edge.
	Tickets *Ticket `json:"tickets,omitempty"`
	// AssignmentHistories holds the value of the assignment_histories edge.
	AssignmentHistories []*ApplicationAssignmentHistory `json:"assignment_histories,omitempty"`
	// StatusHistories holds the value of the status_histories edge.
	StatusHistories []*ApplicationStatusHistory `json:"status_histories,omitempty"`
	// Attachments holds the value of the attachments edge.
	Attachments []*Attachment `json:"attachments,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// TicketsOrErr returns the Tickets value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ApplicationEdges) TicketsOrErr() (*Ticket, error) {
	if e.loadedTypes[0] {
		if e.Tickets == nil {
			// The edge tickets was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: ticket.Label}
		}
		return e.Tickets, nil
	}
	return nil, &NotLoadedError{edge: "tickets"}
}

// AssignmentHistoriesOrErr returns the AssignmentHistories value or an error if the edge
// was not loaded in eager-loading.
func (e ApplicationEdges) AssignmentHistoriesOrErr() ([]*ApplicationAssignmentHistory, error) {
	if e.loadedTypes[1] {
		return e.AssignmentHistories, nil
	}
	return nil, &NotLoadedError{edge: "assignment_histories"}
}

// StatusHistoriesOrErr returns the StatusHistories value or an error if the edge
// was not loaded in eager-loading.
func (e ApplicationEdges) StatusHistoriesOrErr() ([]*ApplicationStatusHistory, error) {
	if e.loadedTypes[2] {
		return e.StatusHistories, nil
	}
	return nil, &NotLoadedError{edge: "status_histories"}
}

// AttachmentsOrErr returns the Attachments value or an error if the edge
// was not loaded in eager-loading.
func (e ApplicationEdges) AttachmentsOrErr() ([]*Attachment, error) {
	if e.loadedTypes[3] {
		return e.Attachments, nil
	}
	return nil, &NotLoadedError{edge: "attachments"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Application) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case application.FieldBasicID, application.FieldPremiumID, application.FieldBotDisplayName, application.FieldBotMid, application.FieldBotActiveStatus, application.FieldApplicantName, application.FieldApplicantEmail, application.FieldApplicationMid, application.FieldRemark, application.FieldStoreType, application.FieldWebsiteURL, application.FieldApplicationStatus, application.FieldReviewComment, application.FieldAssigner, application.FieldAssignee:
			values[i] = new(sql.NullString)
		case application.FieldCreatedDtime, application.FieldUpdateDtime:
			values[i] = new(sql.NullTime)
		case application.FieldID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Application", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Application fields.
func (a *Application) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case application.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				a.ID = *value
			}
		case application.FieldBasicID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field basic_id", values[i])
			} else if value.Valid {
				a.BasicID = value.String
			}
		case application.FieldPremiumID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field premium_id", values[i])
			} else if value.Valid {
				a.PremiumID = value.String
			}
		case application.FieldBotDisplayName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field bot_display_name", values[i])
			} else if value.Valid {
				a.BotDisplayName = value.String
			}
		case application.FieldBotMid:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field bot_mid", values[i])
			} else if value.Valid {
				a.BotMid = value.String
			}
		case application.FieldBotActiveStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field bot_active_status", values[i])
			} else if value.Valid {
				a.BotActiveStatus = application.BotActiveStatus(value.String)
			}
		case application.FieldApplicantName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field applicant_name", values[i])
			} else if value.Valid {
				a.ApplicantName = value.String
			}
		case application.FieldApplicantEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field applicant_email", values[i])
			} else if value.Valid {
				a.ApplicantEmail = value.String
			}
		case application.FieldApplicationMid:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field application_mid", values[i])
			} else if value.Valid {
				a.ApplicationMid = value.String
			}
		case application.FieldRemark:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field remark", values[i])
			} else if value.Valid {
				a.Remark = value.String
			}
		case application.FieldStoreType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field store_type", values[i])
			} else if value.Valid {
				a.StoreType = application.StoreType(value.String)
			}
		case application.FieldWebsiteURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field website_url", values[i])
			} else if value.Valid {
				a.WebsiteURL = value.String
			}
		case application.FieldApplicationStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field application_status", values[i])
			} else if value.Valid {
				a.ApplicationStatus = application.ApplicationStatus(value.String)
			}
		case application.FieldReviewComment:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field review_comment", values[i])
			} else if value.Valid {
				a.ReviewComment = value.String
			}
		case application.FieldAssigner:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field assigner", values[i])
			} else if value.Valid {
				a.Assigner = value.String
			}
		case application.FieldAssignee:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field assignee", values[i])
			} else if value.Valid {
				a.Assignee = value.String
			}
		case application.FieldCreatedDtime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_dtime", values[i])
			} else if value.Valid {
				a.CreatedDtime = value.Time
			}
		case application.FieldUpdateDtime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_dtime", values[i])
			} else if value.Valid {
				a.UpdateDtime = value.Time
			}
		}
	}
	return nil
}

// QueryTickets queries the "tickets" edge of the Application entity.
func (a *Application) QueryTickets() *TicketQuery {
	return (&ApplicationClient{config: a.config}).QueryTickets(a)
}

// QueryAssignmentHistories queries the "assignment_histories" edge of the Application entity.
func (a *Application) QueryAssignmentHistories() *ApplicationAssignmentHistoryQuery {
	return (&ApplicationClient{config: a.config}).QueryAssignmentHistories(a)
}

// QueryStatusHistories queries the "status_histories" edge of the Application entity.
func (a *Application) QueryStatusHistories() *ApplicationStatusHistoryQuery {
	return (&ApplicationClient{config: a.config}).QueryStatusHistories(a)
}

// QueryAttachments queries the "attachments" edge of the Application entity.
func (a *Application) QueryAttachments() *AttachmentQuery {
	return (&ApplicationClient{config: a.config}).QueryAttachments(a)
}

// Update returns a builder for updating this Application.
// Note that you need to call Application.Unwrap() before calling this method if this Application
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Application) Update() *ApplicationUpdateOne {
	return (&ApplicationClient{config: a.config}).UpdateOne(a)
}

// Unwrap unwraps the Application entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Application) Unwrap() *Application {
	tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Application is not a transactional entity")
	}
	a.config.driver = tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Application) String() string {
	var builder strings.Builder
	builder.WriteString("Application(")
	builder.WriteString(fmt.Sprintf("id=%v", a.ID))
	builder.WriteString(", basic_id=")
	builder.WriteString(a.BasicID)
	builder.WriteString(", premium_id=")
	builder.WriteString(a.PremiumID)
	builder.WriteString(", bot_display_name=")
	builder.WriteString(a.BotDisplayName)
	builder.WriteString(", bot_mid=")
	builder.WriteString(a.BotMid)
	builder.WriteString(", bot_active_status=")
	builder.WriteString(fmt.Sprintf("%v", a.BotActiveStatus))
	builder.WriteString(", applicant_name=")
	builder.WriteString(a.ApplicantName)
	builder.WriteString(", applicant_email=")
	builder.WriteString(a.ApplicantEmail)
	builder.WriteString(", application_mid=<sensitive>")
	builder.WriteString(", remark=")
	builder.WriteString(a.Remark)
	builder.WriteString(", store_type=")
	builder.WriteString(fmt.Sprintf("%v", a.StoreType))
	builder.WriteString(", website_url=")
	builder.WriteString(a.WebsiteURL)
	builder.WriteString(", application_status=")
	builder.WriteString(fmt.Sprintf("%v", a.ApplicationStatus))
	builder.WriteString(", review_comment=")
	builder.WriteString(a.ReviewComment)
	builder.WriteString(", assigner=")
	builder.WriteString(a.Assigner)
	builder.WriteString(", assignee=")
	builder.WriteString(a.Assignee)
	builder.WriteString(", created_dtime=")
	builder.WriteString(a.CreatedDtime.Format(time.ANSIC))
	builder.WriteString(", update_dtime=")
	builder.WriteString(a.UpdateDtime.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Applications is a parsable slice of Application.
type Applications []*Application

func (a Applications) config(cfg config) {
	for _i := range a {
		a[_i].config = cfg
	}
}