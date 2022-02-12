// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/Xanonymous-GitHub/carnival/db/pkg/ent/migrate"
	"github.com/google/uuid"

	"github.com/Xanonymous-GitHub/carnival/db/pkg/ent/application"
	"github.com/Xanonymous-GitHub/carnival/db/pkg/ent/applicationassignmenthistory"
	"github.com/Xanonymous-GitHub/carnival/db/pkg/ent/applicationstatushistory"
	"github.com/Xanonymous-GitHub/carnival/db/pkg/ent/attachment"
	"github.com/Xanonymous-GitHub/carnival/db/pkg/ent/reviewer"
	"github.com/Xanonymous-GitHub/carnival/db/pkg/ent/ticket"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Application is the client for interacting with the Application builders.
	Application *ApplicationClient
	// ApplicationAssignmentHistory is the client for interacting with the ApplicationAssignmentHistory builders.
	ApplicationAssignmentHistory *ApplicationAssignmentHistoryClient
	// ApplicationStatusHistory is the client for interacting with the ApplicationStatusHistory builders.
	ApplicationStatusHistory *ApplicationStatusHistoryClient
	// Attachment is the client for interacting with the Attachment builders.
	Attachment *AttachmentClient
	// Reviewer is the client for interacting with the Reviewer builders.
	Reviewer *ReviewerClient
	// Ticket is the client for interacting with the Ticket builders.
	Ticket *TicketClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Application = NewApplicationClient(c.config)
	c.ApplicationAssignmentHistory = NewApplicationAssignmentHistoryClient(c.config)
	c.ApplicationStatusHistory = NewApplicationStatusHistoryClient(c.config)
	c.Attachment = NewAttachmentClient(c.config)
	c.Reviewer = NewReviewerClient(c.config)
	c.Ticket = NewTicketClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:                          ctx,
		config:                       cfg,
		Application:                  NewApplicationClient(cfg),
		ApplicationAssignmentHistory: NewApplicationAssignmentHistoryClient(cfg),
		ApplicationStatusHistory:     NewApplicationStatusHistoryClient(cfg),
		Attachment:                   NewAttachmentClient(cfg),
		Reviewer:                     NewReviewerClient(cfg),
		Ticket:                       NewTicketClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:                          ctx,
		config:                       cfg,
		Application:                  NewApplicationClient(cfg),
		ApplicationAssignmentHistory: NewApplicationAssignmentHistoryClient(cfg),
		ApplicationStatusHistory:     NewApplicationStatusHistoryClient(cfg),
		Attachment:                   NewAttachmentClient(cfg),
		Reviewer:                     NewReviewerClient(cfg),
		Ticket:                       NewTicketClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Application.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Application.Use(hooks...)
	c.ApplicationAssignmentHistory.Use(hooks...)
	c.ApplicationStatusHistory.Use(hooks...)
	c.Attachment.Use(hooks...)
	c.Reviewer.Use(hooks...)
	c.Ticket.Use(hooks...)
}

// ApplicationClient is a client for the Application schema.
type ApplicationClient struct {
	config
}

// NewApplicationClient returns a client for the Application from the given config.
func NewApplicationClient(c config) *ApplicationClient {
	return &ApplicationClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `application.Hooks(f(g(h())))`.
func (c *ApplicationClient) Use(hooks ...Hook) {
	c.hooks.Application = append(c.hooks.Application, hooks...)
}

// Create returns a create builder for Application.
func (c *ApplicationClient) Create() *ApplicationCreate {
	mutation := newApplicationMutation(c.config, OpCreate)
	return &ApplicationCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Application entities.
func (c *ApplicationClient) CreateBulk(builders ...*ApplicationCreate) *ApplicationCreateBulk {
	return &ApplicationCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Application.
func (c *ApplicationClient) Update() *ApplicationUpdate {
	mutation := newApplicationMutation(c.config, OpUpdate)
	return &ApplicationUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ApplicationClient) UpdateOne(a *Application) *ApplicationUpdateOne {
	mutation := newApplicationMutation(c.config, OpUpdateOne, withApplication(a))
	return &ApplicationUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ApplicationClient) UpdateOneID(id uuid.UUID) *ApplicationUpdateOne {
	mutation := newApplicationMutation(c.config, OpUpdateOne, withApplicationID(id))
	return &ApplicationUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Application.
func (c *ApplicationClient) Delete() *ApplicationDelete {
	mutation := newApplicationMutation(c.config, OpDelete)
	return &ApplicationDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *ApplicationClient) DeleteOne(a *Application) *ApplicationDeleteOne {
	return c.DeleteOneID(a.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *ApplicationClient) DeleteOneID(id uuid.UUID) *ApplicationDeleteOne {
	builder := c.Delete().Where(application.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ApplicationDeleteOne{builder}
}

// Query returns a query builder for Application.
func (c *ApplicationClient) Query() *ApplicationQuery {
	return &ApplicationQuery{
		config: c.config,
	}
}

// Get returns a Application entity by its id.
func (c *ApplicationClient) Get(ctx context.Context, id uuid.UUID) (*Application, error) {
	return c.Query().Where(application.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ApplicationClient) GetX(ctx context.Context, id uuid.UUID) *Application {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryTickets queries the tickets edge of a Application.
func (c *ApplicationClient) QueryTickets(a *Application) *TicketQuery {
	query := &TicketQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := a.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(application.Table, application.FieldID, id),
			sqlgraph.To(ticket.Table, ticket.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, application.TicketsTable, application.TicketsColumn),
		)
		fromV = sqlgraph.Neighbors(a.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryAssignmentHistories queries the assignment_histories edge of a Application.
func (c *ApplicationClient) QueryAssignmentHistories(a *Application) *ApplicationAssignmentHistoryQuery {
	query := &ApplicationAssignmentHistoryQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := a.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(application.Table, application.FieldID, id),
			sqlgraph.To(applicationassignmenthistory.Table, applicationassignmenthistory.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, application.AssignmentHistoriesTable, application.AssignmentHistoriesColumn),
		)
		fromV = sqlgraph.Neighbors(a.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryStatusHistories queries the status_histories edge of a Application.
func (c *ApplicationClient) QueryStatusHistories(a *Application) *ApplicationStatusHistoryQuery {
	query := &ApplicationStatusHistoryQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := a.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(application.Table, application.FieldID, id),
			sqlgraph.To(applicationstatushistory.Table, applicationstatushistory.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, application.StatusHistoriesTable, application.StatusHistoriesColumn),
		)
		fromV = sqlgraph.Neighbors(a.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryAttachments queries the attachments edge of a Application.
func (c *ApplicationClient) QueryAttachments(a *Application) *AttachmentQuery {
	query := &AttachmentQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := a.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(application.Table, application.FieldID, id),
			sqlgraph.To(attachment.Table, attachment.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, application.AttachmentsTable, application.AttachmentsColumn),
		)
		fromV = sqlgraph.Neighbors(a.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ApplicationClient) Hooks() []Hook {
	return c.hooks.Application
}

// ApplicationAssignmentHistoryClient is a client for the ApplicationAssignmentHistory schema.
type ApplicationAssignmentHistoryClient struct {
	config
}

// NewApplicationAssignmentHistoryClient returns a client for the ApplicationAssignmentHistory from the given config.
func NewApplicationAssignmentHistoryClient(c config) *ApplicationAssignmentHistoryClient {
	return &ApplicationAssignmentHistoryClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `applicationassignmenthistory.Hooks(f(g(h())))`.
func (c *ApplicationAssignmentHistoryClient) Use(hooks ...Hook) {
	c.hooks.ApplicationAssignmentHistory = append(c.hooks.ApplicationAssignmentHistory, hooks...)
}

// Create returns a create builder for ApplicationAssignmentHistory.
func (c *ApplicationAssignmentHistoryClient) Create() *ApplicationAssignmentHistoryCreate {
	mutation := newApplicationAssignmentHistoryMutation(c.config, OpCreate)
	return &ApplicationAssignmentHistoryCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of ApplicationAssignmentHistory entities.
func (c *ApplicationAssignmentHistoryClient) CreateBulk(builders ...*ApplicationAssignmentHistoryCreate) *ApplicationAssignmentHistoryCreateBulk {
	return &ApplicationAssignmentHistoryCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for ApplicationAssignmentHistory.
func (c *ApplicationAssignmentHistoryClient) Update() *ApplicationAssignmentHistoryUpdate {
	mutation := newApplicationAssignmentHistoryMutation(c.config, OpUpdate)
	return &ApplicationAssignmentHistoryUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ApplicationAssignmentHistoryClient) UpdateOne(aah *ApplicationAssignmentHistory) *ApplicationAssignmentHistoryUpdateOne {
	mutation := newApplicationAssignmentHistoryMutation(c.config, OpUpdateOne, withApplicationAssignmentHistory(aah))
	return &ApplicationAssignmentHistoryUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ApplicationAssignmentHistoryClient) UpdateOneID(id int) *ApplicationAssignmentHistoryUpdateOne {
	mutation := newApplicationAssignmentHistoryMutation(c.config, OpUpdateOne, withApplicationAssignmentHistoryID(id))
	return &ApplicationAssignmentHistoryUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for ApplicationAssignmentHistory.
func (c *ApplicationAssignmentHistoryClient) Delete() *ApplicationAssignmentHistoryDelete {
	mutation := newApplicationAssignmentHistoryMutation(c.config, OpDelete)
	return &ApplicationAssignmentHistoryDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *ApplicationAssignmentHistoryClient) DeleteOne(aah *ApplicationAssignmentHistory) *ApplicationAssignmentHistoryDeleteOne {
	return c.DeleteOneID(aah.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *ApplicationAssignmentHistoryClient) DeleteOneID(id int) *ApplicationAssignmentHistoryDeleteOne {
	builder := c.Delete().Where(applicationassignmenthistory.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ApplicationAssignmentHistoryDeleteOne{builder}
}

// Query returns a query builder for ApplicationAssignmentHistory.
func (c *ApplicationAssignmentHistoryClient) Query() *ApplicationAssignmentHistoryQuery {
	return &ApplicationAssignmentHistoryQuery{
		config: c.config,
	}
}

// Get returns a ApplicationAssignmentHistory entity by its id.
func (c *ApplicationAssignmentHistoryClient) Get(ctx context.Context, id int) (*ApplicationAssignmentHistory, error) {
	return c.Query().Where(applicationassignmenthistory.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ApplicationAssignmentHistoryClient) GetX(ctx context.Context, id int) *ApplicationAssignmentHistory {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryApplications queries the applications edge of a ApplicationAssignmentHistory.
func (c *ApplicationAssignmentHistoryClient) QueryApplications(aah *ApplicationAssignmentHistory) *ApplicationQuery {
	query := &ApplicationQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := aah.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(applicationassignmenthistory.Table, applicationassignmenthistory.FieldID, id),
			sqlgraph.To(application.Table, application.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, applicationassignmenthistory.ApplicationsTable, applicationassignmenthistory.ApplicationsColumn),
		)
		fromV = sqlgraph.Neighbors(aah.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ApplicationAssignmentHistoryClient) Hooks() []Hook {
	return c.hooks.ApplicationAssignmentHistory
}

// ApplicationStatusHistoryClient is a client for the ApplicationStatusHistory schema.
type ApplicationStatusHistoryClient struct {
	config
}

// NewApplicationStatusHistoryClient returns a client for the ApplicationStatusHistory from the given config.
func NewApplicationStatusHistoryClient(c config) *ApplicationStatusHistoryClient {
	return &ApplicationStatusHistoryClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `applicationstatushistory.Hooks(f(g(h())))`.
func (c *ApplicationStatusHistoryClient) Use(hooks ...Hook) {
	c.hooks.ApplicationStatusHistory = append(c.hooks.ApplicationStatusHistory, hooks...)
}

// Create returns a create builder for ApplicationStatusHistory.
func (c *ApplicationStatusHistoryClient) Create() *ApplicationStatusHistoryCreate {
	mutation := newApplicationStatusHistoryMutation(c.config, OpCreate)
	return &ApplicationStatusHistoryCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of ApplicationStatusHistory entities.
func (c *ApplicationStatusHistoryClient) CreateBulk(builders ...*ApplicationStatusHistoryCreate) *ApplicationStatusHistoryCreateBulk {
	return &ApplicationStatusHistoryCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for ApplicationStatusHistory.
func (c *ApplicationStatusHistoryClient) Update() *ApplicationStatusHistoryUpdate {
	mutation := newApplicationStatusHistoryMutation(c.config, OpUpdate)
	return &ApplicationStatusHistoryUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ApplicationStatusHistoryClient) UpdateOne(ash *ApplicationStatusHistory) *ApplicationStatusHistoryUpdateOne {
	mutation := newApplicationStatusHistoryMutation(c.config, OpUpdateOne, withApplicationStatusHistory(ash))
	return &ApplicationStatusHistoryUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ApplicationStatusHistoryClient) UpdateOneID(id int) *ApplicationStatusHistoryUpdateOne {
	mutation := newApplicationStatusHistoryMutation(c.config, OpUpdateOne, withApplicationStatusHistoryID(id))
	return &ApplicationStatusHistoryUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for ApplicationStatusHistory.
func (c *ApplicationStatusHistoryClient) Delete() *ApplicationStatusHistoryDelete {
	mutation := newApplicationStatusHistoryMutation(c.config, OpDelete)
	return &ApplicationStatusHistoryDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *ApplicationStatusHistoryClient) DeleteOne(ash *ApplicationStatusHistory) *ApplicationStatusHistoryDeleteOne {
	return c.DeleteOneID(ash.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *ApplicationStatusHistoryClient) DeleteOneID(id int) *ApplicationStatusHistoryDeleteOne {
	builder := c.Delete().Where(applicationstatushistory.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ApplicationStatusHistoryDeleteOne{builder}
}

// Query returns a query builder for ApplicationStatusHistory.
func (c *ApplicationStatusHistoryClient) Query() *ApplicationStatusHistoryQuery {
	return &ApplicationStatusHistoryQuery{
		config: c.config,
	}
}

// Get returns a ApplicationStatusHistory entity by its id.
func (c *ApplicationStatusHistoryClient) Get(ctx context.Context, id int) (*ApplicationStatusHistory, error) {
	return c.Query().Where(applicationstatushistory.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ApplicationStatusHistoryClient) GetX(ctx context.Context, id int) *ApplicationStatusHistory {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryApplications queries the applications edge of a ApplicationStatusHistory.
func (c *ApplicationStatusHistoryClient) QueryApplications(ash *ApplicationStatusHistory) *ApplicationQuery {
	query := &ApplicationQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := ash.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(applicationstatushistory.Table, applicationstatushistory.FieldID, id),
			sqlgraph.To(application.Table, application.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, applicationstatushistory.ApplicationsTable, applicationstatushistory.ApplicationsColumn),
		)
		fromV = sqlgraph.Neighbors(ash.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ApplicationStatusHistoryClient) Hooks() []Hook {
	return c.hooks.ApplicationStatusHistory
}

// AttachmentClient is a client for the Attachment schema.
type AttachmentClient struct {
	config
}

// NewAttachmentClient returns a client for the Attachment from the given config.
func NewAttachmentClient(c config) *AttachmentClient {
	return &AttachmentClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `attachment.Hooks(f(g(h())))`.
func (c *AttachmentClient) Use(hooks ...Hook) {
	c.hooks.Attachment = append(c.hooks.Attachment, hooks...)
}

// Create returns a create builder for Attachment.
func (c *AttachmentClient) Create() *AttachmentCreate {
	mutation := newAttachmentMutation(c.config, OpCreate)
	return &AttachmentCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Attachment entities.
func (c *AttachmentClient) CreateBulk(builders ...*AttachmentCreate) *AttachmentCreateBulk {
	return &AttachmentCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Attachment.
func (c *AttachmentClient) Update() *AttachmentUpdate {
	mutation := newAttachmentMutation(c.config, OpUpdate)
	return &AttachmentUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *AttachmentClient) UpdateOne(a *Attachment) *AttachmentUpdateOne {
	mutation := newAttachmentMutation(c.config, OpUpdateOne, withAttachment(a))
	return &AttachmentUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *AttachmentClient) UpdateOneID(id int) *AttachmentUpdateOne {
	mutation := newAttachmentMutation(c.config, OpUpdateOne, withAttachmentID(id))
	return &AttachmentUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Attachment.
func (c *AttachmentClient) Delete() *AttachmentDelete {
	mutation := newAttachmentMutation(c.config, OpDelete)
	return &AttachmentDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *AttachmentClient) DeleteOne(a *Attachment) *AttachmentDeleteOne {
	return c.DeleteOneID(a.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *AttachmentClient) DeleteOneID(id int) *AttachmentDeleteOne {
	builder := c.Delete().Where(attachment.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &AttachmentDeleteOne{builder}
}

// Query returns a query builder for Attachment.
func (c *AttachmentClient) Query() *AttachmentQuery {
	return &AttachmentQuery{
		config: c.config,
	}
}

// Get returns a Attachment entity by its id.
func (c *AttachmentClient) Get(ctx context.Context, id int) (*Attachment, error) {
	return c.Query().Where(attachment.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *AttachmentClient) GetX(ctx context.Context, id int) *Attachment {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryApplications queries the applications edge of a Attachment.
func (c *AttachmentClient) QueryApplications(a *Attachment) *ApplicationQuery {
	query := &ApplicationQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := a.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(attachment.Table, attachment.FieldID, id),
			sqlgraph.To(application.Table, application.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, attachment.ApplicationsTable, attachment.ApplicationsColumn),
		)
		fromV = sqlgraph.Neighbors(a.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryTickets queries the tickets edge of a Attachment.
func (c *AttachmentClient) QueryTickets(a *Attachment) *TicketQuery {
	query := &TicketQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := a.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(attachment.Table, attachment.FieldID, id),
			sqlgraph.To(ticket.Table, ticket.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, attachment.TicketsTable, attachment.TicketsColumn),
		)
		fromV = sqlgraph.Neighbors(a.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *AttachmentClient) Hooks() []Hook {
	return c.hooks.Attachment
}

// ReviewerClient is a client for the Reviewer schema.
type ReviewerClient struct {
	config
}

// NewReviewerClient returns a client for the Reviewer from the given config.
func NewReviewerClient(c config) *ReviewerClient {
	return &ReviewerClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `reviewer.Hooks(f(g(h())))`.
func (c *ReviewerClient) Use(hooks ...Hook) {
	c.hooks.Reviewer = append(c.hooks.Reviewer, hooks...)
}

// Create returns a create builder for Reviewer.
func (c *ReviewerClient) Create() *ReviewerCreate {
	mutation := newReviewerMutation(c.config, OpCreate)
	return &ReviewerCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Reviewer entities.
func (c *ReviewerClient) CreateBulk(builders ...*ReviewerCreate) *ReviewerCreateBulk {
	return &ReviewerCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Reviewer.
func (c *ReviewerClient) Update() *ReviewerUpdate {
	mutation := newReviewerMutation(c.config, OpUpdate)
	return &ReviewerUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ReviewerClient) UpdateOne(r *Reviewer) *ReviewerUpdateOne {
	mutation := newReviewerMutation(c.config, OpUpdateOne, withReviewer(r))
	return &ReviewerUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ReviewerClient) UpdateOneID(id string) *ReviewerUpdateOne {
	mutation := newReviewerMutation(c.config, OpUpdateOne, withReviewerID(id))
	return &ReviewerUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Reviewer.
func (c *ReviewerClient) Delete() *ReviewerDelete {
	mutation := newReviewerMutation(c.config, OpDelete)
	return &ReviewerDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *ReviewerClient) DeleteOne(r *Reviewer) *ReviewerDeleteOne {
	return c.DeleteOneID(r.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *ReviewerClient) DeleteOneID(id string) *ReviewerDeleteOne {
	builder := c.Delete().Where(reviewer.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ReviewerDeleteOne{builder}
}

// Query returns a query builder for Reviewer.
func (c *ReviewerClient) Query() *ReviewerQuery {
	return &ReviewerQuery{
		config: c.config,
	}
}

// Get returns a Reviewer entity by its id.
func (c *ReviewerClient) Get(ctx context.Context, id string) (*Reviewer, error) {
	return c.Query().Where(reviewer.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ReviewerClient) GetX(ctx context.Context, id string) *Reviewer {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *ReviewerClient) Hooks() []Hook {
	return c.hooks.Reviewer
}

// TicketClient is a client for the Ticket schema.
type TicketClient struct {
	config
}

// NewTicketClient returns a client for the Ticket from the given config.
func NewTicketClient(c config) *TicketClient {
	return &TicketClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `ticket.Hooks(f(g(h())))`.
func (c *TicketClient) Use(hooks ...Hook) {
	c.hooks.Ticket = append(c.hooks.Ticket, hooks...)
}

// Create returns a create builder for Ticket.
func (c *TicketClient) Create() *TicketCreate {
	mutation := newTicketMutation(c.config, OpCreate)
	return &TicketCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Ticket entities.
func (c *TicketClient) CreateBulk(builders ...*TicketCreate) *TicketCreateBulk {
	return &TicketCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Ticket.
func (c *TicketClient) Update() *TicketUpdate {
	mutation := newTicketMutation(c.config, OpUpdate)
	return &TicketUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *TicketClient) UpdateOne(t *Ticket) *TicketUpdateOne {
	mutation := newTicketMutation(c.config, OpUpdateOne, withTicket(t))
	return &TicketUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *TicketClient) UpdateOneID(id int) *TicketUpdateOne {
	mutation := newTicketMutation(c.config, OpUpdateOne, withTicketID(id))
	return &TicketUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Ticket.
func (c *TicketClient) Delete() *TicketDelete {
	mutation := newTicketMutation(c.config, OpDelete)
	return &TicketDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *TicketClient) DeleteOne(t *Ticket) *TicketDeleteOne {
	return c.DeleteOneID(t.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *TicketClient) DeleteOneID(id int) *TicketDeleteOne {
	builder := c.Delete().Where(ticket.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &TicketDeleteOne{builder}
}

// Query returns a query builder for Ticket.
func (c *TicketClient) Query() *TicketQuery {
	return &TicketQuery{
		config: c.config,
	}
}

// Get returns a Ticket entity by its id.
func (c *TicketClient) Get(ctx context.Context, id int) (*Ticket, error) {
	return c.Query().Where(ticket.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *TicketClient) GetX(ctx context.Context, id int) *Ticket {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryApplications queries the applications edge of a Ticket.
func (c *TicketClient) QueryApplications(t *Ticket) *ApplicationQuery {
	query := &ApplicationQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := t.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(ticket.Table, ticket.FieldID, id),
			sqlgraph.To(application.Table, application.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, ticket.ApplicationsTable, ticket.ApplicationsColumn),
		)
		fromV = sqlgraph.Neighbors(t.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryAttachments queries the attachments edge of a Ticket.
func (c *TicketClient) QueryAttachments(t *Ticket) *AttachmentQuery {
	query := &AttachmentQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := t.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(ticket.Table, ticket.FieldID, id),
			sqlgraph.To(attachment.Table, attachment.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ticket.AttachmentsTable, ticket.AttachmentsColumn),
		)
		fromV = sqlgraph.Neighbors(t.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *TicketClient) Hooks() []Hook {
	return c.hooks.Ticket
}