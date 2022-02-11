package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

type ApplicationStatus string

const (
	AppWip       ApplicationStatus = "wip"
	AppReviewing ApplicationStatus = "reviewing"
	AppVerified  ApplicationStatus = "verified"
	AppRejected  ApplicationStatus = "rejected"
	AppWaiting   ApplicationStatus = "waiting"
	AppReplied   ApplicationStatus = "replied"
	AppRevoked   ApplicationStatus = "revoked"
	AppCanceling ApplicationStatus = "canceling"
)

// Values get all enum values in ApplicationStatus.
// TODO(TU): extract enum values function
func (ApplicationStatus) Values() (kinds []string) {
	for _, k := range []ApplicationStatus{
		AppWip, AppReviewing, AppVerified, AppRejected, AppWaiting, AppReplied, AppRevoked, AppCanceling,
	} {
		kinds = append(kinds, string(k))
	}
	return
}

type BotActiveStatus string

const (
	Active  BotActiveStatus = "active"
	Suspend BotActiveStatus = "suspend"
	Delete  BotActiveStatus = "delete"
)

// Values get all enum values in BotActiveStatus.
// TODO(TU): extract enum values function
func (BotActiveStatus) Values() (kinds []string) {
	for _, k := range []BotActiveStatus{Active, Suspend, Delete} {
		kinds = append(kinds, string(k))
	}
	return
}

type StoreType string

const (
	OnlineStore   StoreType = "online_store"
	PhysicalStore StoreType = "physical_store"
)

// Values get all enum values in StoreType.
// TODO(TU): extract enum values function
func (StoreType) Values() (kinds []string) {
	for _, k := range []StoreType{OnlineStore, PhysicalStore} {
		kinds = append(kinds, string(k))
	}
	return
}

// Application holds the schema definition for the Application entity.
type Application struct {
	ent.Schema
}

// Fields of the Application.
func (Application) Fields() []ent.Field {
	const (
		basicIdSizeLimit        = 10
		premiumIdSizeLimit      = 32
		botDisplayNameSizeLimit = 32
		botMidSizeLimit         = 64
		applicantNameSizeLimit  = 16
		applicantEmailSizeLimit = 64
		applicantMidSizeLimit   = 64
		remarkSizeLimit         = 256
		reviewCommentSizeLimit  = 256
		assignerSizeLimit       = 32
		assigneeSizeLimit       = 32
	)

	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			StorageKey("application_id"),
		field.String("basic_id").
			MaxLen(basicIdSizeLimit).
			Unique(),
		field.String("premium_id").
			MaxLen(premiumIdSizeLimit).
			Unique(),
		field.String("bot_display_name").
			MaxLen(botDisplayNameSizeLimit),
		field.String("bot_mid").
			MaxLen(botMidSizeLimit).
			Unique(),
		field.Enum("bot_active_status").
			GoType(BotActiveStatus("")),
		field.String("applicant_name").
			MaxLen(applicantNameSizeLimit),
		field.String("applicant_email").
			MaxLen(applicantEmailSizeLimit),
		field.String("application_mid").
			Sensitive().
			MaxLen(applicantMidSizeLimit),
		field.String("remark").
			MaxLen(remarkSizeLimit),
		field.Enum("store_type").
			GoType(StoreType("")),
		field.Text("website_url"),
		field.Enum("application_status").
			GoType(ApplicationStatus("")),
		field.String("review_comment").
			MaxLen(reviewCommentSizeLimit),
		field.String("assigner").
			MaxLen(assignerSizeLimit),
		field.String("assignee").
			MaxLen(assigneeSizeLimit),
		field.Time("created_dtime").
			Default(time.Now).
			Immutable(),
		field.Time("update_dtime"),
	}
}

// Edges of the Application.
func (Application) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("tickets", Ticket.Type).
			Unique(),
		edge.To("assignment_histories", ApplicationAssignmentHistory.Type),
		edge.To("status_histories", ApplicationStatusHistory.Type),
		edge.To("attachments", Attachment.Type),
	}
}
