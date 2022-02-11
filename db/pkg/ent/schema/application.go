package schema

import (
	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/schema"
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

func (ApplicationStatus) ToMap() map[string]int32 {
	r := make(map[string]int32)
	for i, k := range []ApplicationStatus{
		AppWip, AppReviewing, AppVerified, AppRejected, AppWaiting, AppReplied, AppRevoked, AppCanceling,
	} {
		r[string(k)] = int32(i) + 1
	}
	return r
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

func (BotActiveStatus) ToMap() map[string]int32 {
	r := make(map[string]int32)
	for i, k := range []BotActiveStatus{Active, Suspend, Delete} {
		r[string(k)] = int32(i) + 1
	}
	return r
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

func (StoreType) ToMap() map[string]int32 {
	r := make(map[string]int32)
	for i, k := range []StoreType{OnlineStore, PhysicalStore} {
		r[string(k)] = int32(i) + 1
	}
	return r
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
			StorageKey("application_id").
			Annotations(entproto.Field(1)),

		field.String("basic_id").
			MaxLen(basicIdSizeLimit).
			Unique().
			Annotations(entproto.Field(2)),

		field.String("premium_id").
			MaxLen(premiumIdSizeLimit).
			Unique().
			Annotations(entproto.Field(3)),

		field.String("bot_display_name").
			MaxLen(botDisplayNameSizeLimit).
			Annotations(entproto.Field(4)),

		field.String("bot_mid").
			MaxLen(botMidSizeLimit).
			Unique().
			Annotations(entproto.Field(5)),

		field.Enum("bot_active_status").
			Values(BotActiveStatus("").Values()...).
			Annotations(
				entproto.Field(6),
				entproto.Enum(BotActiveStatus("").ToMap()),
			),

		field.String("applicant_name").
			MaxLen(applicantNameSizeLimit).
			Annotations(entproto.Field(7)),

		field.String("applicant_email").
			MaxLen(applicantEmailSizeLimit).
			Annotations(entproto.Field(8)),

		field.String("application_mid").
			Sensitive().
			MaxLen(applicantMidSizeLimit).
			Annotations(entproto.Field(9)),

		field.String("remark").
			MaxLen(remarkSizeLimit).
			Annotations(entproto.Field(10)),

		field.Enum("store_type").
			Values(StoreType("").Values()...).
			Annotations(
				entproto.Field(11),
				entproto.Enum(StoreType("").ToMap()),
			),

		field.Text("website_url").
			Annotations(entproto.Field(12)),

		field.Enum("application_status").
			Values(ApplicationStatus("").Values()...).
			Annotations(
				entproto.Field(13),
				entproto.Enum(ApplicationStatus("").ToMap()),
			),

		field.String("review_comment").
			MaxLen(reviewCommentSizeLimit).
			Annotations(entproto.Field(14)),

		field.String("assigner").
			MaxLen(assignerSizeLimit).
			Annotations(entproto.Field(15)),

		field.String("assignee").
			MaxLen(assigneeSizeLimit).
			Annotations(entproto.Field(16)),

		field.Time("created_dtime").
			Default(time.Now).
			Immutable().
			Annotations(entproto.Field(17)),

		field.Time("update_dtime").
			Annotations(entproto.Field(18)),
	}
}

// Edges of the Application.
func (Application) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("tickets", Ticket.Type).
			Unique().
			Annotations(entproto.Field(19)),

		edge.To("assignment_histories", ApplicationAssignmentHistory.Type).
			Annotations(entproto.Field(20)),

		edge.To("status_histories", ApplicationStatusHistory.Type).
			Annotations(entproto.Field(21)),

		edge.To("attachments", Attachment.Type).
			Annotations(entproto.Field(22)),
	}
}

func (Application) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(),
	}
}
