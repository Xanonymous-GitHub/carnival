// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// ApplicationsColumns holds the columns for the "applications" table.
	ApplicationsColumns = []*schema.Column{
		{Name: "application_id", Type: field.TypeUUID},
		{Name: "basic_id", Type: field.TypeString, Unique: true, Size: 10},
		{Name: "premium_id", Type: field.TypeString, Unique: true, Size: 32},
		{Name: "bot_display_name", Type: field.TypeString, Size: 32},
		{Name: "bot_mid", Type: field.TypeString, Unique: true, Size: 64},
		{Name: "bot_active_status", Type: field.TypeEnum, Enums: []string{"active", "suspend", "delete"}},
		{Name: "applicant_name", Type: field.TypeString, Size: 16},
		{Name: "applicant_email", Type: field.TypeString, Size: 64},
		{Name: "application_mid", Type: field.TypeString, Size: 64},
		{Name: "remark", Type: field.TypeString, Size: 256},
		{Name: "store_type", Type: field.TypeEnum, Enums: []string{"online_store", "physical_store"}},
		{Name: "website_url", Type: field.TypeString, Size: 2147483647},
		{Name: "application_status", Type: field.TypeEnum, Enums: []string{"wip", "reviewing", "verified", "rejected", "waiting", "replied", "revoked", "canceling"}},
		{Name: "review_comment", Type: field.TypeString, Size: 256},
		{Name: "assigner", Type: field.TypeString, Size: 32},
		{Name: "assignee", Type: field.TypeString, Size: 32},
		{Name: "created_dtime", Type: field.TypeTime},
		{Name: "update_dtime", Type: field.TypeTime},
	}
	// ApplicationsTable holds the schema information for the "applications" table.
	ApplicationsTable = &schema.Table{
		Name:       "applications",
		Columns:    ApplicationsColumns,
		PrimaryKey: []*schema.Column{ApplicationsColumns[0]},
	}
	// ApplicationAssignmentHistoriesColumns holds the columns for the "application_assignment_histories" table.
	ApplicationAssignmentHistoriesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "application_id", Type: field.TypeUUID},
		{Name: "assigner", Type: field.TypeString, Size: 32},
		{Name: "assignee", Type: field.TypeString, Size: 32},
		{Name: "created_time", Type: field.TypeTime},
		{Name: "application_assignment_histories", Type: field.TypeUUID, Nullable: true},
	}
	// ApplicationAssignmentHistoriesTable holds the schema information for the "application_assignment_histories" table.
	ApplicationAssignmentHistoriesTable = &schema.Table{
		Name:       "application_assignment_histories",
		Columns:    ApplicationAssignmentHistoriesColumns,
		PrimaryKey: []*schema.Column{ApplicationAssignmentHistoriesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "application_assignment_histories_applications_assignment_histories",
				Columns:    []*schema.Column{ApplicationAssignmentHistoriesColumns[5]},
				RefColumns: []*schema.Column{ApplicationsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ApplicationStatusHistoriesColumns holds the columns for the "application_status_histories" table.
	ApplicationStatusHistoriesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "application_id", Type: field.TypeUUID},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"wip", "reviewing", "verified", "rejected", "waiting", "replied", "revoked", "canceling"}},
		{Name: "created_time", Type: field.TypeTime},
		{Name: "application_status_histories", Type: field.TypeUUID, Nullable: true},
	}
	// ApplicationStatusHistoriesTable holds the schema information for the "application_status_histories" table.
	ApplicationStatusHistoriesTable = &schema.Table{
		Name:       "application_status_histories",
		Columns:    ApplicationStatusHistoriesColumns,
		PrimaryKey: []*schema.Column{ApplicationStatusHistoriesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "application_status_histories_applications_status_histories",
				Columns:    []*schema.Column{ApplicationStatusHistoriesColumns[4]},
				RefColumns: []*schema.Column{ApplicationsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// AttachmentsColumns holds the columns for the "attachments" table.
	AttachmentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "application_id", Type: field.TypeUUID},
		{Name: "ticket_id", Type: field.TypeInt32, Nullable: true},
		{Name: "a_type", Type: field.TypeEnum, Enums: []string{"biz_cert", "store_appearance", "id_document", "other", "supplement", "reference"}},
		{Name: "obs_oid", Type: field.TypeString, Size: 32},
		{Name: "obs_hash", Type: field.TypeString, Size: 128},
		{Name: "created_dtime", Type: field.TypeTime},
		{Name: "application_attachments", Type: field.TypeUUID, Nullable: true},
		{Name: "ticket_attachments", Type: field.TypeInt, Nullable: true},
	}
	// AttachmentsTable holds the schema information for the "attachments" table.
	AttachmentsTable = &schema.Table{
		Name:       "attachments",
		Columns:    AttachmentsColumns,
		PrimaryKey: []*schema.Column{AttachmentsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "attachments_applications_attachments",
				Columns:    []*schema.Column{AttachmentsColumns[7]},
				RefColumns: []*schema.Column{ApplicationsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "attachments_tickets_attachments",
				Columns:    []*schema.Column{AttachmentsColumns[8]},
				RefColumns: []*schema.Column{TicketsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ReviewersColumns holds the columns for the "reviewers" table.
	ReviewersColumns = []*schema.Column{
		{Name: "reviewer_id", Type: field.TypeString, Size: 10},
		{Name: "reviewer_name", Type: field.TypeString, Size: 32},
		{Name: "iims_role", Type: field.TypeEnum, Enums: []string{"oav_admin", "oav_reviewer"}},
		{Name: "created_dtime", Type: field.TypeTime},
	}
	// ReviewersTable holds the schema information for the "reviewers" table.
	ReviewersTable = &schema.Table{
		Name:       "reviewers",
		Columns:    ReviewersColumns,
		PrimaryKey: []*schema.Column{ReviewersColumns[0]},
	}
	// TicketsColumns holds the columns for the "tickets" table.
	TicketsColumns = []*schema.Column{
		{Name: "ticket_id", Type: field.TypeInt, Increment: true},
		{Name: "application_id", Type: field.TypeUUID},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"created", "replied", "completed"}},
		{Name: "creator", Type: field.TypeString, Size: 32},
		{Name: "content", Type: field.TypeString, Size: 2147483647},
		{Name: "reply", Type: field.TypeString, Size: 2147483647},
		{Name: "replier", Type: field.TypeString, Size: 32},
		{Name: "reviewer", Type: field.TypeString, Size: 32},
		{Name: "review_remark", Type: field.TypeString, Size: 2147483647},
		{Name: "replied_dtime", Type: field.TypeTime},
		{Name: "reviewed_dtime", Type: field.TypeTime},
		{Name: "created_dtime", Type: field.TypeTime},
		{Name: "updated_dtime", Type: field.TypeTime},
		{Name: "application_tickets", Type: field.TypeUUID, Unique: true, Nullable: true},
	}
	// TicketsTable holds the schema information for the "tickets" table.
	TicketsTable = &schema.Table{
		Name:       "tickets",
		Columns:    TicketsColumns,
		PrimaryKey: []*schema.Column{TicketsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "tickets_applications_tickets",
				Columns:    []*schema.Column{TicketsColumns[13]},
				RefColumns: []*schema.Column{ApplicationsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ApplicationsTable,
		ApplicationAssignmentHistoriesTable,
		ApplicationStatusHistoriesTable,
		AttachmentsTable,
		ReviewersTable,
		TicketsTable,
	}
)

func init() {
	ApplicationAssignmentHistoriesTable.ForeignKeys[0].RefTable = ApplicationsTable
	ApplicationStatusHistoriesTable.ForeignKeys[0].RefTable = ApplicationsTable
	AttachmentsTable.ForeignKeys[0].RefTable = ApplicationsTable
	AttachmentsTable.ForeignKeys[1].RefTable = TicketsTable
	TicketsTable.ForeignKeys[0].RefTable = ApplicationsTable
}