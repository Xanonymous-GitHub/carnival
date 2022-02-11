// Code generated by entc, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/Xanonymous-GitHub/carnival/db/ent/application"
	"github.com/Xanonymous-GitHub/carnival/db/ent/applicationassignmenthistory"
	"github.com/Xanonymous-GitHub/carnival/db/ent/applicationstatushistory"
	"github.com/Xanonymous-GitHub/carnival/db/ent/attachment"
	"github.com/Xanonymous-GitHub/carnival/db/ent/reviewer"
	"github.com/Xanonymous-GitHub/carnival/db/ent/schema"
	"github.com/Xanonymous-GitHub/carnival/db/ent/ticket"
	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	applicationFields := schema.Application{}.Fields()
	_ = applicationFields
	// applicationDescBasicID is the schema descriptor for basic_id field.
	applicationDescBasicID := applicationFields[1].Descriptor()
	// application.BasicIDValidator is a validator for the "basic_id" field. It is called by the builders before save.
	application.BasicIDValidator = applicationDescBasicID.Validators[0].(func(string) error)
	// applicationDescPremiumID is the schema descriptor for premium_id field.
	applicationDescPremiumID := applicationFields[2].Descriptor()
	// application.PremiumIDValidator is a validator for the "premium_id" field. It is called by the builders before save.
	application.PremiumIDValidator = applicationDescPremiumID.Validators[0].(func(string) error)
	// applicationDescBotDisplayName is the schema descriptor for bot_display_name field.
	applicationDescBotDisplayName := applicationFields[3].Descriptor()
	// application.BotDisplayNameValidator is a validator for the "bot_display_name" field. It is called by the builders before save.
	application.BotDisplayNameValidator = applicationDescBotDisplayName.Validators[0].(func(string) error)
	// applicationDescBotMid is the schema descriptor for bot_mid field.
	applicationDescBotMid := applicationFields[4].Descriptor()
	// application.BotMidValidator is a validator for the "bot_mid" field. It is called by the builders before save.
	application.BotMidValidator = applicationDescBotMid.Validators[0].(func(string) error)
	// applicationDescApplicantName is the schema descriptor for applicant_name field.
	applicationDescApplicantName := applicationFields[6].Descriptor()
	// application.ApplicantNameValidator is a validator for the "applicant_name" field. It is called by the builders before save.
	application.ApplicantNameValidator = applicationDescApplicantName.Validators[0].(func(string) error)
	// applicationDescApplicantEmail is the schema descriptor for applicant_email field.
	applicationDescApplicantEmail := applicationFields[7].Descriptor()
	// application.ApplicantEmailValidator is a validator for the "applicant_email" field. It is called by the builders before save.
	application.ApplicantEmailValidator = applicationDescApplicantEmail.Validators[0].(func(string) error)
	// applicationDescApplicationMid is the schema descriptor for application_mid field.
	applicationDescApplicationMid := applicationFields[8].Descriptor()
	// application.ApplicationMidValidator is a validator for the "application_mid" field. It is called by the builders before save.
	application.ApplicationMidValidator = applicationDescApplicationMid.Validators[0].(func(string) error)
	// applicationDescRemark is the schema descriptor for remark field.
	applicationDescRemark := applicationFields[9].Descriptor()
	// application.RemarkValidator is a validator for the "remark" field. It is called by the builders before save.
	application.RemarkValidator = applicationDescRemark.Validators[0].(func(string) error)
	// applicationDescReviewComment is the schema descriptor for review_comment field.
	applicationDescReviewComment := applicationFields[13].Descriptor()
	// application.ReviewCommentValidator is a validator for the "review_comment" field. It is called by the builders before save.
	application.ReviewCommentValidator = applicationDescReviewComment.Validators[0].(func(string) error)
	// applicationDescAssigner is the schema descriptor for assigner field.
	applicationDescAssigner := applicationFields[14].Descriptor()
	// application.AssignerValidator is a validator for the "assigner" field. It is called by the builders before save.
	application.AssignerValidator = applicationDescAssigner.Validators[0].(func(string) error)
	// applicationDescAssignee is the schema descriptor for assignee field.
	applicationDescAssignee := applicationFields[15].Descriptor()
	// application.AssigneeValidator is a validator for the "assignee" field. It is called by the builders before save.
	application.AssigneeValidator = applicationDescAssignee.Validators[0].(func(string) error)
	// applicationDescCreatedDtime is the schema descriptor for created_dtime field.
	applicationDescCreatedDtime := applicationFields[16].Descriptor()
	// application.DefaultCreatedDtime holds the default value on creation for the created_dtime field.
	application.DefaultCreatedDtime = applicationDescCreatedDtime.Default.(func() time.Time)
	// applicationDescID is the schema descriptor for id field.
	applicationDescID := applicationFields[0].Descriptor()
	// application.DefaultID holds the default value on creation for the id field.
	application.DefaultID = applicationDescID.Default.(func() uuid.UUID)
	applicationassignmenthistoryFields := schema.ApplicationAssignmentHistory{}.Fields()
	_ = applicationassignmenthistoryFields
	// applicationassignmenthistoryDescAssigner is the schema descriptor for assigner field.
	applicationassignmenthistoryDescAssigner := applicationassignmenthistoryFields[1].Descriptor()
	// applicationassignmenthistory.AssignerValidator is a validator for the "assigner" field. It is called by the builders before save.
	applicationassignmenthistory.AssignerValidator = applicationassignmenthistoryDescAssigner.Validators[0].(func(string) error)
	// applicationassignmenthistoryDescAssignee is the schema descriptor for assignee field.
	applicationassignmenthistoryDescAssignee := applicationassignmenthistoryFields[2].Descriptor()
	// applicationassignmenthistory.AssigneeValidator is a validator for the "assignee" field. It is called by the builders before save.
	applicationassignmenthistory.AssigneeValidator = applicationassignmenthistoryDescAssignee.Validators[0].(func(string) error)
	// applicationassignmenthistoryDescCreatedTime is the schema descriptor for created_time field.
	applicationassignmenthistoryDescCreatedTime := applicationassignmenthistoryFields[3].Descriptor()
	// applicationassignmenthistory.DefaultCreatedTime holds the default value on creation for the created_time field.
	applicationassignmenthistory.DefaultCreatedTime = applicationassignmenthistoryDescCreatedTime.Default.(func() time.Time)
	applicationstatushistoryFields := schema.ApplicationStatusHistory{}.Fields()
	_ = applicationstatushistoryFields
	// applicationstatushistoryDescCreatedTime is the schema descriptor for created_time field.
	applicationstatushistoryDescCreatedTime := applicationstatushistoryFields[2].Descriptor()
	// applicationstatushistory.DefaultCreatedTime holds the default value on creation for the created_time field.
	applicationstatushistory.DefaultCreatedTime = applicationstatushistoryDescCreatedTime.Default.(func() time.Time)
	attachmentFields := schema.Attachment{}.Fields()
	_ = attachmentFields
	// attachmentDescObsOid is the schema descriptor for obs_oid field.
	attachmentDescObsOid := attachmentFields[3].Descriptor()
	// attachment.ObsOidValidator is a validator for the "obs_oid" field. It is called by the builders before save.
	attachment.ObsOidValidator = attachmentDescObsOid.Validators[0].(func(string) error)
	// attachmentDescObsHash is the schema descriptor for obs_hash field.
	attachmentDescObsHash := attachmentFields[4].Descriptor()
	// attachment.ObsHashValidator is a validator for the "obs_hash" field. It is called by the builders before save.
	attachment.ObsHashValidator = attachmentDescObsHash.Validators[0].(func(string) error)
	// attachmentDescCreatedDtime is the schema descriptor for created_dtime field.
	attachmentDescCreatedDtime := attachmentFields[5].Descriptor()
	// attachment.DefaultCreatedDtime holds the default value on creation for the created_dtime field.
	attachment.DefaultCreatedDtime = attachmentDescCreatedDtime.Default.(func() time.Time)
	reviewerFields := schema.Reviewer{}.Fields()
	_ = reviewerFields
	// reviewerDescReviewerID is the schema descriptor for reviewer_id field.
	reviewerDescReviewerID := reviewerFields[0].Descriptor()
	// reviewer.ReviewerIDValidator is a validator for the "reviewer_id" field. It is called by the builders before save.
	reviewer.ReviewerIDValidator = reviewerDescReviewerID.Validators[0].(func(string) error)
	// reviewerDescReviewerName is the schema descriptor for reviewer_name field.
	reviewerDescReviewerName := reviewerFields[1].Descriptor()
	// reviewer.ReviewerNameValidator is a validator for the "reviewer_name" field. It is called by the builders before save.
	reviewer.ReviewerNameValidator = reviewerDescReviewerName.Validators[0].(func(string) error)
	// reviewerDescCreatedDtime is the schema descriptor for created_dtime field.
	reviewerDescCreatedDtime := reviewerFields[3].Descriptor()
	// reviewer.DefaultCreatedDtime holds the default value on creation for the created_dtime field.
	reviewer.DefaultCreatedDtime = reviewerDescCreatedDtime.Default.(func() time.Time)
	ticketFields := schema.Ticket{}.Fields()
	_ = ticketFields
	// ticketDescCreator is the schema descriptor for creator field.
	ticketDescCreator := ticketFields[3].Descriptor()
	// ticket.CreatorValidator is a validator for the "creator" field. It is called by the builders before save.
	ticket.CreatorValidator = ticketDescCreator.Validators[0].(func(string) error)
	// ticketDescReplier is the schema descriptor for replier field.
	ticketDescReplier := ticketFields[6].Descriptor()
	// ticket.ReplierValidator is a validator for the "replier" field. It is called by the builders before save.
	ticket.ReplierValidator = ticketDescReplier.Validators[0].(func(string) error)
	// ticketDescReviewer is the schema descriptor for reviewer field.
	ticketDescReviewer := ticketFields[7].Descriptor()
	// ticket.ReviewerValidator is a validator for the "reviewer" field. It is called by the builders before save.
	ticket.ReviewerValidator = ticketDescReviewer.Validators[0].(func(string) error)
	// ticketDescCreatedDtime is the schema descriptor for created_dtime field.
	ticketDescCreatedDtime := ticketFields[11].Descriptor()
	// ticket.DefaultCreatedDtime holds the default value on creation for the created_dtime field.
	ticket.DefaultCreatedDtime = ticketDescCreatedDtime.Default.(func() time.Time)
}
