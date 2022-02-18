package schema

import (
	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/Xanonymous-GitHub/carnival/db/pkg/ent/utils"
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

var applicationStatus = []ApplicationStatus{
	AppWip, AppReviewing, AppVerified, AppRejected, AppWaiting, AppReplied, AppRevoked, AppCanceling,
}

// Values get all enum values in ApplicationStatus.
// TODO(TU): extract enum values function
func (ApplicationStatus) Values() []string {
	return utils.ToValuesFromEnum(applicationStatus)
}

func (ApplicationStatus) ToMap() map[string]int32 {
	return utils.ToMapFromEnum(applicationStatus)
}

type BotActiveStatus string

const (
	Active  BotActiveStatus = "active"
	Suspend BotActiveStatus = "suspend"
	Delete  BotActiveStatus = "delete"
)

var botActiveStatus = []BotActiveStatus{Active, Suspend, Delete}

// Values get all enum values in BotActiveStatus.
// TODO(TU): extract enum values function
func (BotActiveStatus) Values() []string {
	return utils.ToValuesFromEnum(botActiveStatus)
}

func (BotActiveStatus) ToMap() map[string]int32 {
	return utils.ToMapFromEnum(botActiveStatus)
}

type StoreType string

const (
	OnlineStore   StoreType = "online_store"
	PhysicalStore StoreType = "physical_store"
)

var storeType = []StoreType{OnlineStore, PhysicalStore}

// Values get all enum values in StoreType.
// TODO(TU): extract enum values function
func (StoreType) Values() []string {
	return utils.ToValuesFromEnum(storeType)
}

func (StoreType) ToMap() map[string]int32 {
	return utils.ToMapFromEnum(storeType)
}

type BotSuspendReason string

const (
	UserResign  BotSuspendReason = "user_resign"
	ForceResign BotSuspendReason = "force_resign"
	Penalty     BotSuspendReason = "penalty"
	Unpaid      BotSuspendReason = "unpaid"
)

var botSuspendReason = []BotSuspendReason{UserResign, ForceResign, Penalty, Unpaid}

// Values get all enum values in BotSuspendReason.
// TODO(TU): extract enum values function
func (BotSuspendReason) Values() []string {
	return utils.ToValuesFromEnum(botSuspendReason)
}

func (BotSuspendReason) ToMap() map[string]int32 {
	return utils.ToMapFromEnum(botSuspendReason)
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
		applicantBizIdSizeLimit = 64
		applicantEmailSizeLimit = 64
		applicantMidSizeLimit   = 64
		remarkSizeLimit         = 64
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
			Values(new(BotActiveStatus).Values()...).
			Annotations(
				entproto.Field(6),
				entproto.Enum(new(BotActiveStatus).ToMap()),
			),

		field.Enum("bot_suspend_reason").
			Values(new(BotSuspendReason).Values()...).
			Annotations(
				entproto.Field(7),
				entproto.Enum(new(BotSuspendReason).ToMap()),
			),

		field.String("applicant_name").
			MaxLen(applicantNameSizeLimit).
			Annotations(entproto.Field(8)),

		field.String("applicant_biz_id").
			MaxLen(applicantBizIdSizeLimit).
			Annotations(entproto.Field(9)),

		field.String("applicant_mid").
			Sensitive().
			MaxLen(applicantMidSizeLimit).
			Annotations(entproto.Field(10)),

		field.String("applicant_email").
			MaxLen(applicantEmailSizeLimit).
			Annotations(entproto.Field(11)),

		field.String("remark").
			MaxLen(remarkSizeLimit).
			Annotations(entproto.Field(12)),

		field.Enum("store_type").
			Values(new(StoreType).Values()...).
			Annotations(
				entproto.Field(13),
				entproto.Enum(new(StoreType).ToMap()),
			),

		field.Text("website_url").
			Annotations(entproto.Field(14)),

		field.Enum("application_status").
			Values(new(ApplicationStatus).Values()...).
			Annotations(
				entproto.Field(15),
				entproto.Enum(new(ApplicationStatus).ToMap()),
			),

		field.Text("review_comment").
			Annotations(entproto.Field(16)),

		field.String("assigner").
			MaxLen(assignerSizeLimit).
			Annotations(entproto.Field(17)),

		field.String("assignee").
			MaxLen(assigneeSizeLimit).
			Annotations(entproto.Field(18)),

		field.Time("created_dtime").
			Default(time.Now).
			Immutable().
			Annotations(entproto.Field(19)),

		field.Time("updated_dtime").
			Annotations(entproto.Field(20)),
	}
}

// Edges of the Application.
func (Application) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("tickets", Ticket.Type).
			Annotations(entproto.Field(21)),

		edge.To("assignment_histories", ApplicationAssignmentHistory.Type).
			Annotations(entproto.Field(22)),

		edge.To("status_histories", ApplicationStatusHistory.Type).
			Annotations(entproto.Field(23)),

		edge.To("attachments", Attachment.Type).
			Annotations(entproto.Field(24)),
	}
}

func (Application) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(),
	}
}
