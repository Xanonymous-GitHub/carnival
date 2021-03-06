// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Xanonymous-GitHub/carnival/db/pkg/ent/application"
	"github.com/Xanonymous-GitHub/carnival/db/pkg/ent/applicationassignmenthistory"
	"github.com/Xanonymous-GitHub/carnival/db/pkg/ent/applicationstatushistory"
	"github.com/Xanonymous-GitHub/carnival/db/pkg/ent/attachment"
	"github.com/Xanonymous-GitHub/carnival/db/pkg/ent/ticket"
	"github.com/google/uuid"
)

// ApplicationCreate is the builder for creating a Application entity.
type ApplicationCreate struct {
	config
	mutation *ApplicationMutation
	hooks    []Hook
}

// SetBasicID sets the "basic_id" field.
func (ac *ApplicationCreate) SetBasicID(s string) *ApplicationCreate {
	ac.mutation.SetBasicID(s)
	return ac
}

// SetPremiumID sets the "premium_id" field.
func (ac *ApplicationCreate) SetPremiumID(s string) *ApplicationCreate {
	ac.mutation.SetPremiumID(s)
	return ac
}

// SetBotDisplayName sets the "bot_display_name" field.
func (ac *ApplicationCreate) SetBotDisplayName(s string) *ApplicationCreate {
	ac.mutation.SetBotDisplayName(s)
	return ac
}

// SetBotMid sets the "bot_mid" field.
func (ac *ApplicationCreate) SetBotMid(s string) *ApplicationCreate {
	ac.mutation.SetBotMid(s)
	return ac
}

// SetBotActiveStatus sets the "bot_active_status" field.
func (ac *ApplicationCreate) SetBotActiveStatus(aas application.BotActiveStatus) *ApplicationCreate {
	ac.mutation.SetBotActiveStatus(aas)
	return ac
}

// SetBotSuspendReason sets the "bot_suspend_reason" field.
func (ac *ApplicationCreate) SetBotSuspendReason(asr application.BotSuspendReason) *ApplicationCreate {
	ac.mutation.SetBotSuspendReason(asr)
	return ac
}

// SetApplicantName sets the "applicant_name" field.
func (ac *ApplicationCreate) SetApplicantName(s string) *ApplicationCreate {
	ac.mutation.SetApplicantName(s)
	return ac
}

// SetApplicantBizID sets the "applicant_biz_id" field.
func (ac *ApplicationCreate) SetApplicantBizID(s string) *ApplicationCreate {
	ac.mutation.SetApplicantBizID(s)
	return ac
}

// SetApplicantMid sets the "applicant_mid" field.
func (ac *ApplicationCreate) SetApplicantMid(s string) *ApplicationCreate {
	ac.mutation.SetApplicantMid(s)
	return ac
}

// SetApplicantEmail sets the "applicant_email" field.
func (ac *ApplicationCreate) SetApplicantEmail(s string) *ApplicationCreate {
	ac.mutation.SetApplicantEmail(s)
	return ac
}

// SetRemark sets the "remark" field.
func (ac *ApplicationCreate) SetRemark(s string) *ApplicationCreate {
	ac.mutation.SetRemark(s)
	return ac
}

// SetStoreType sets the "store_type" field.
func (ac *ApplicationCreate) SetStoreType(at application.StoreType) *ApplicationCreate {
	ac.mutation.SetStoreType(at)
	return ac
}

// SetWebsiteURL sets the "website_url" field.
func (ac *ApplicationCreate) SetWebsiteURL(s string) *ApplicationCreate {
	ac.mutation.SetWebsiteURL(s)
	return ac
}

// SetApplicationStatus sets the "application_status" field.
func (ac *ApplicationCreate) SetApplicationStatus(as application.ApplicationStatus) *ApplicationCreate {
	ac.mutation.SetApplicationStatus(as)
	return ac
}

// SetReviewComment sets the "review_comment" field.
func (ac *ApplicationCreate) SetReviewComment(s string) *ApplicationCreate {
	ac.mutation.SetReviewComment(s)
	return ac
}

// SetAssigner sets the "assigner" field.
func (ac *ApplicationCreate) SetAssigner(s string) *ApplicationCreate {
	ac.mutation.SetAssigner(s)
	return ac
}

// SetAssignee sets the "assignee" field.
func (ac *ApplicationCreate) SetAssignee(s string) *ApplicationCreate {
	ac.mutation.SetAssignee(s)
	return ac
}

// SetCreatedDtime sets the "created_dtime" field.
func (ac *ApplicationCreate) SetCreatedDtime(t time.Time) *ApplicationCreate {
	ac.mutation.SetCreatedDtime(t)
	return ac
}

// SetNillableCreatedDtime sets the "created_dtime" field if the given value is not nil.
func (ac *ApplicationCreate) SetNillableCreatedDtime(t *time.Time) *ApplicationCreate {
	if t != nil {
		ac.SetCreatedDtime(*t)
	}
	return ac
}

// SetUpdatedDtime sets the "updated_dtime" field.
func (ac *ApplicationCreate) SetUpdatedDtime(t time.Time) *ApplicationCreate {
	ac.mutation.SetUpdatedDtime(t)
	return ac
}

// SetID sets the "id" field.
func (ac *ApplicationCreate) SetID(u uuid.UUID) *ApplicationCreate {
	ac.mutation.SetID(u)
	return ac
}

// SetNillableID sets the "id" field if the given value is not nil.
func (ac *ApplicationCreate) SetNillableID(u *uuid.UUID) *ApplicationCreate {
	if u != nil {
		ac.SetID(*u)
	}
	return ac
}

// AddTicketIDs adds the "tickets" edge to the Ticket entity by IDs.
func (ac *ApplicationCreate) AddTicketIDs(ids ...int) *ApplicationCreate {
	ac.mutation.AddTicketIDs(ids...)
	return ac
}

// AddTickets adds the "tickets" edges to the Ticket entity.
func (ac *ApplicationCreate) AddTickets(t ...*Ticket) *ApplicationCreate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return ac.AddTicketIDs(ids...)
}

// AddAssignmentHistoryIDs adds the "assignment_histories" edge to the ApplicationAssignmentHistory entity by IDs.
func (ac *ApplicationCreate) AddAssignmentHistoryIDs(ids ...int) *ApplicationCreate {
	ac.mutation.AddAssignmentHistoryIDs(ids...)
	return ac
}

// AddAssignmentHistories adds the "assignment_histories" edges to the ApplicationAssignmentHistory entity.
func (ac *ApplicationCreate) AddAssignmentHistories(a ...*ApplicationAssignmentHistory) *ApplicationCreate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return ac.AddAssignmentHistoryIDs(ids...)
}

// AddStatusHistoryIDs adds the "status_histories" edge to the ApplicationStatusHistory entity by IDs.
func (ac *ApplicationCreate) AddStatusHistoryIDs(ids ...int) *ApplicationCreate {
	ac.mutation.AddStatusHistoryIDs(ids...)
	return ac
}

// AddStatusHistories adds the "status_histories" edges to the ApplicationStatusHistory entity.
func (ac *ApplicationCreate) AddStatusHistories(a ...*ApplicationStatusHistory) *ApplicationCreate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return ac.AddStatusHistoryIDs(ids...)
}

// AddAttachmentIDs adds the "attachments" edge to the Attachment entity by IDs.
func (ac *ApplicationCreate) AddAttachmentIDs(ids ...int) *ApplicationCreate {
	ac.mutation.AddAttachmentIDs(ids...)
	return ac
}

// AddAttachments adds the "attachments" edges to the Attachment entity.
func (ac *ApplicationCreate) AddAttachments(a ...*Attachment) *ApplicationCreate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return ac.AddAttachmentIDs(ids...)
}

// Mutation returns the ApplicationMutation object of the builder.
func (ac *ApplicationCreate) Mutation() *ApplicationMutation {
	return ac.mutation
}

// Save creates the Application in the database.
func (ac *ApplicationCreate) Save(ctx context.Context) (*Application, error) {
	var (
		err  error
		node *Application
	)
	ac.defaults()
	if len(ac.hooks) == 0 {
		if err = ac.check(); err != nil {
			return nil, err
		}
		node, err = ac.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ApplicationMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ac.check(); err != nil {
				return nil, err
			}
			ac.mutation = mutation
			if node, err = ac.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ac.hooks) - 1; i >= 0; i-- {
			if ac.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ac.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ac.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ac *ApplicationCreate) SaveX(ctx context.Context) *Application {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ac *ApplicationCreate) Exec(ctx context.Context) error {
	_, err := ac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ac *ApplicationCreate) ExecX(ctx context.Context) {
	if err := ac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ac *ApplicationCreate) defaults() {
	if _, ok := ac.mutation.CreatedDtime(); !ok {
		v := application.DefaultCreatedDtime()
		ac.mutation.SetCreatedDtime(v)
	}
	if _, ok := ac.mutation.ID(); !ok {
		v := application.DefaultID()
		ac.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ac *ApplicationCreate) check() error {
	if _, ok := ac.mutation.BasicID(); !ok {
		return &ValidationError{Name: "basic_id", err: errors.New(`ent: missing required field "Application.basic_id"`)}
	}
	if v, ok := ac.mutation.BasicID(); ok {
		if err := application.BasicIDValidator(v); err != nil {
			return &ValidationError{Name: "basic_id", err: fmt.Errorf(`ent: validator failed for field "Application.basic_id": %w`, err)}
		}
	}
	if _, ok := ac.mutation.PremiumID(); !ok {
		return &ValidationError{Name: "premium_id", err: errors.New(`ent: missing required field "Application.premium_id"`)}
	}
	if v, ok := ac.mutation.PremiumID(); ok {
		if err := application.PremiumIDValidator(v); err != nil {
			return &ValidationError{Name: "premium_id", err: fmt.Errorf(`ent: validator failed for field "Application.premium_id": %w`, err)}
		}
	}
	if _, ok := ac.mutation.BotDisplayName(); !ok {
		return &ValidationError{Name: "bot_display_name", err: errors.New(`ent: missing required field "Application.bot_display_name"`)}
	}
	if v, ok := ac.mutation.BotDisplayName(); ok {
		if err := application.BotDisplayNameValidator(v); err != nil {
			return &ValidationError{Name: "bot_display_name", err: fmt.Errorf(`ent: validator failed for field "Application.bot_display_name": %w`, err)}
		}
	}
	if _, ok := ac.mutation.BotMid(); !ok {
		return &ValidationError{Name: "bot_mid", err: errors.New(`ent: missing required field "Application.bot_mid"`)}
	}
	if v, ok := ac.mutation.BotMid(); ok {
		if err := application.BotMidValidator(v); err != nil {
			return &ValidationError{Name: "bot_mid", err: fmt.Errorf(`ent: validator failed for field "Application.bot_mid": %w`, err)}
		}
	}
	if _, ok := ac.mutation.BotActiveStatus(); !ok {
		return &ValidationError{Name: "bot_active_status", err: errors.New(`ent: missing required field "Application.bot_active_status"`)}
	}
	if v, ok := ac.mutation.BotActiveStatus(); ok {
		if err := application.BotActiveStatusValidator(v); err != nil {
			return &ValidationError{Name: "bot_active_status", err: fmt.Errorf(`ent: validator failed for field "Application.bot_active_status": %w`, err)}
		}
	}
	if _, ok := ac.mutation.BotSuspendReason(); !ok {
		return &ValidationError{Name: "bot_suspend_reason", err: errors.New(`ent: missing required field "Application.bot_suspend_reason"`)}
	}
	if v, ok := ac.mutation.BotSuspendReason(); ok {
		if err := application.BotSuspendReasonValidator(v); err != nil {
			return &ValidationError{Name: "bot_suspend_reason", err: fmt.Errorf(`ent: validator failed for field "Application.bot_suspend_reason": %w`, err)}
		}
	}
	if _, ok := ac.mutation.ApplicantName(); !ok {
		return &ValidationError{Name: "applicant_name", err: errors.New(`ent: missing required field "Application.applicant_name"`)}
	}
	if v, ok := ac.mutation.ApplicantName(); ok {
		if err := application.ApplicantNameValidator(v); err != nil {
			return &ValidationError{Name: "applicant_name", err: fmt.Errorf(`ent: validator failed for field "Application.applicant_name": %w`, err)}
		}
	}
	if _, ok := ac.mutation.ApplicantBizID(); !ok {
		return &ValidationError{Name: "applicant_biz_id", err: errors.New(`ent: missing required field "Application.applicant_biz_id"`)}
	}
	if v, ok := ac.mutation.ApplicantBizID(); ok {
		if err := application.ApplicantBizIDValidator(v); err != nil {
			return &ValidationError{Name: "applicant_biz_id", err: fmt.Errorf(`ent: validator failed for field "Application.applicant_biz_id": %w`, err)}
		}
	}
	if _, ok := ac.mutation.ApplicantMid(); !ok {
		return &ValidationError{Name: "applicant_mid", err: errors.New(`ent: missing required field "Application.applicant_mid"`)}
	}
	if v, ok := ac.mutation.ApplicantMid(); ok {
		if err := application.ApplicantMidValidator(v); err != nil {
			return &ValidationError{Name: "applicant_mid", err: fmt.Errorf(`ent: validator failed for field "Application.applicant_mid": %w`, err)}
		}
	}
	if _, ok := ac.mutation.ApplicantEmail(); !ok {
		return &ValidationError{Name: "applicant_email", err: errors.New(`ent: missing required field "Application.applicant_email"`)}
	}
	if v, ok := ac.mutation.ApplicantEmail(); ok {
		if err := application.ApplicantEmailValidator(v); err != nil {
			return &ValidationError{Name: "applicant_email", err: fmt.Errorf(`ent: validator failed for field "Application.applicant_email": %w`, err)}
		}
	}
	if _, ok := ac.mutation.Remark(); !ok {
		return &ValidationError{Name: "remark", err: errors.New(`ent: missing required field "Application.remark"`)}
	}
	if v, ok := ac.mutation.Remark(); ok {
		if err := application.RemarkValidator(v); err != nil {
			return &ValidationError{Name: "remark", err: fmt.Errorf(`ent: validator failed for field "Application.remark": %w`, err)}
		}
	}
	if _, ok := ac.mutation.StoreType(); !ok {
		return &ValidationError{Name: "store_type", err: errors.New(`ent: missing required field "Application.store_type"`)}
	}
	if v, ok := ac.mutation.StoreType(); ok {
		if err := application.StoreTypeValidator(v); err != nil {
			return &ValidationError{Name: "store_type", err: fmt.Errorf(`ent: validator failed for field "Application.store_type": %w`, err)}
		}
	}
	if _, ok := ac.mutation.WebsiteURL(); !ok {
		return &ValidationError{Name: "website_url", err: errors.New(`ent: missing required field "Application.website_url"`)}
	}
	if _, ok := ac.mutation.ApplicationStatus(); !ok {
		return &ValidationError{Name: "application_status", err: errors.New(`ent: missing required field "Application.application_status"`)}
	}
	if v, ok := ac.mutation.ApplicationStatus(); ok {
		if err := application.ApplicationStatusValidator(v); err != nil {
			return &ValidationError{Name: "application_status", err: fmt.Errorf(`ent: validator failed for field "Application.application_status": %w`, err)}
		}
	}
	if _, ok := ac.mutation.ReviewComment(); !ok {
		return &ValidationError{Name: "review_comment", err: errors.New(`ent: missing required field "Application.review_comment"`)}
	}
	if _, ok := ac.mutation.Assigner(); !ok {
		return &ValidationError{Name: "assigner", err: errors.New(`ent: missing required field "Application.assigner"`)}
	}
	if v, ok := ac.mutation.Assigner(); ok {
		if err := application.AssignerValidator(v); err != nil {
			return &ValidationError{Name: "assigner", err: fmt.Errorf(`ent: validator failed for field "Application.assigner": %w`, err)}
		}
	}
	if _, ok := ac.mutation.Assignee(); !ok {
		return &ValidationError{Name: "assignee", err: errors.New(`ent: missing required field "Application.assignee"`)}
	}
	if v, ok := ac.mutation.Assignee(); ok {
		if err := application.AssigneeValidator(v); err != nil {
			return &ValidationError{Name: "assignee", err: fmt.Errorf(`ent: validator failed for field "Application.assignee": %w`, err)}
		}
	}
	if _, ok := ac.mutation.CreatedDtime(); !ok {
		return &ValidationError{Name: "created_dtime", err: errors.New(`ent: missing required field "Application.created_dtime"`)}
	}
	if _, ok := ac.mutation.UpdatedDtime(); !ok {
		return &ValidationError{Name: "updated_dtime", err: errors.New(`ent: missing required field "Application.updated_dtime"`)}
	}
	return nil
}

func (ac *ApplicationCreate) sqlSave(ctx context.Context) (*Application, error) {
	_node, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	return _node, nil
}

func (ac *ApplicationCreate) createSpec() (*Application, *sqlgraph.CreateSpec) {
	var (
		_node = &Application{config: ac.config}
		_spec = &sqlgraph.CreateSpec{
			Table: application.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: application.FieldID,
			},
		}
	)
	if id, ok := ac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := ac.mutation.BasicID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: application.FieldBasicID,
		})
		_node.BasicID = value
	}
	if value, ok := ac.mutation.PremiumID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: application.FieldPremiumID,
		})
		_node.PremiumID = value
	}
	if value, ok := ac.mutation.BotDisplayName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: application.FieldBotDisplayName,
		})
		_node.BotDisplayName = value
	}
	if value, ok := ac.mutation.BotMid(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: application.FieldBotMid,
		})
		_node.BotMid = value
	}
	if value, ok := ac.mutation.BotActiveStatus(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: application.FieldBotActiveStatus,
		})
		_node.BotActiveStatus = value
	}
	if value, ok := ac.mutation.BotSuspendReason(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: application.FieldBotSuspendReason,
		})
		_node.BotSuspendReason = value
	}
	if value, ok := ac.mutation.ApplicantName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: application.FieldApplicantName,
		})
		_node.ApplicantName = value
	}
	if value, ok := ac.mutation.ApplicantBizID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: application.FieldApplicantBizID,
		})
		_node.ApplicantBizID = value
	}
	if value, ok := ac.mutation.ApplicantMid(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: application.FieldApplicantMid,
		})
		_node.ApplicantMid = value
	}
	if value, ok := ac.mutation.ApplicantEmail(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: application.FieldApplicantEmail,
		})
		_node.ApplicantEmail = value
	}
	if value, ok := ac.mutation.Remark(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: application.FieldRemark,
		})
		_node.Remark = value
	}
	if value, ok := ac.mutation.StoreType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: application.FieldStoreType,
		})
		_node.StoreType = value
	}
	if value, ok := ac.mutation.WebsiteURL(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: application.FieldWebsiteURL,
		})
		_node.WebsiteURL = value
	}
	if value, ok := ac.mutation.ApplicationStatus(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: application.FieldApplicationStatus,
		})
		_node.ApplicationStatus = value
	}
	if value, ok := ac.mutation.ReviewComment(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: application.FieldReviewComment,
		})
		_node.ReviewComment = value
	}
	if value, ok := ac.mutation.Assigner(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: application.FieldAssigner,
		})
		_node.Assigner = value
	}
	if value, ok := ac.mutation.Assignee(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: application.FieldAssignee,
		})
		_node.Assignee = value
	}
	if value, ok := ac.mutation.CreatedDtime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: application.FieldCreatedDtime,
		})
		_node.CreatedDtime = value
	}
	if value, ok := ac.mutation.UpdatedDtime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: application.FieldUpdatedDtime,
		})
		_node.UpdatedDtime = value
	}
	if nodes := ac.mutation.TicketsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   application.TicketsTable,
			Columns: []string{application.TicketsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: ticket.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ac.mutation.AssignmentHistoriesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   application.AssignmentHistoriesTable,
			Columns: []string{application.AssignmentHistoriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: applicationassignmenthistory.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ac.mutation.StatusHistoriesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   application.StatusHistoriesTable,
			Columns: []string{application.StatusHistoriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: applicationstatushistory.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ac.mutation.AttachmentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   application.AttachmentsTable,
			Columns: []string{application.AttachmentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: attachment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ApplicationCreateBulk is the builder for creating many Application entities in bulk.
type ApplicationCreateBulk struct {
	config
	builders []*ApplicationCreate
}

// Save creates the Application entities in the database.
func (acb *ApplicationCreateBulk) Save(ctx context.Context) ([]*Application, error) {
	specs := make([]*sqlgraph.CreateSpec, len(acb.builders))
	nodes := make([]*Application, len(acb.builders))
	mutators := make([]Mutator, len(acb.builders))
	for i := range acb.builders {
		func(i int, root context.Context) {
			builder := acb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ApplicationMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, acb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, acb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, acb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (acb *ApplicationCreateBulk) SaveX(ctx context.Context) []*Application {
	v, err := acb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acb *ApplicationCreateBulk) Exec(ctx context.Context) error {
	_, err := acb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acb *ApplicationCreateBulk) ExecX(ctx context.Context) {
	if err := acb.Exec(ctx); err != nil {
		panic(err)
	}
}
