// Code generated by protoc-gen-entgrpc. DO NOT EDIT.
package entpb

import (
	context "context"
	base64 "encoding/base64"
	entproto "entgo.io/contrib/entproto"
	runtime "entgo.io/contrib/entproto/runtime"
	sqlgraph "entgo.io/ent/dialect/sql/sqlgraph"
	fmt "fmt"
	ent "github.com/Xanonymous-GitHub/carnival/db/pkg/ent"
	application "github.com/Xanonymous-GitHub/carnival/db/pkg/ent/application"
	attachment "github.com/Xanonymous-GitHub/carnival/db/pkg/ent/attachment"
	ticket "github.com/Xanonymous-GitHub/carnival/db/pkg/ent/ticket"
	uuid "github.com/google/uuid"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	strconv "strconv"
	strings "strings"
)

// TicketService implements TicketServiceServer
type TicketService struct {
	client *ent.Client
	UnimplementedTicketServiceServer
}

// NewTicketService returns a new TicketService
func NewTicketService(client *ent.Client) *TicketService {
	return &TicketService{
		client: client,
	}
}

func toProtoTicket_Status(e ticket.Status) Ticket_Status {
	if v, ok := Ticket_Status_value[strings.ToUpper(string(e))]; ok {
		return Ticket_Status(v)
	}
	return Ticket_Status(0)
}

func toEntTicket_Status(e Ticket_Status) ticket.Status {
	if v, ok := Ticket_Status_name[int32(e)]; ok {
		return ticket.Status(strings.ToLower(v))
	}
	return ""
}

// toProtoTicket transforms the ent type to the pb type
func toProtoTicket(e *ent.Ticket) (*Ticket, error) {
	v := &Ticket{}
	applicationid, err := e.ApplicationID.MarshalBinary()
	if err != nil {
		return nil, err
	}
	v.ApplicationId = applicationid
	content := e.Content
	v.Content = content
	createddtime := timestamppb.New(e.CreatedDtime)
	v.CreatedDtime = createddtime
	creator := e.Creator
	v.Creator = creator
	id := int32(e.ID)
	v.Id = id
	replieddtime := timestamppb.New(e.RepliedDtime)
	v.RepliedDtime = replieddtime
	replier := e.Replier
	v.Replier = replier
	reply := e.Reply
	v.Reply = reply
	reviewremark := e.ReviewRemark
	v.ReviewRemark = reviewremark
	revieweddtime := timestamppb.New(e.ReviewedDtime)
	v.ReviewedDtime = revieweddtime
	reviewer := e.Reviewer
	v.Reviewer = reviewer
	status := toProtoTicket_Status(e.Status)
	v.Status = status
	updateddtime := timestamppb.New(e.UpdatedDtime)
	v.UpdatedDtime = updateddtime
	if edg := e.Edges.Applications; edg != nil {
		id, err := edg.ID.MarshalBinary()
		if err != nil {
			return nil, err
		}
		v.Applications = &Application{
			Id: id,
		}
	}
	for _, edg := range e.Edges.Attachments {
		id := int32(edg.ID)
		v.Attachments = append(v.Attachments, &Attachment{
			Id: id,
		})
	}
	return v, nil
}

// Create implements TicketServiceServer.Create
func (svc *TicketService) Create(ctx context.Context, req *CreateTicketRequest) (*Ticket, error) {
	ticket := req.GetTicket()
	m := svc.client.Ticket.Create()
	var ticketApplicationID uuid.UUID
	if err := (&ticketApplicationID).UnmarshalBinary(ticket.GetApplicationId()); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	}
	m.SetApplicationID(ticketApplicationID)
	ticketContent := ticket.GetContent()
	m.SetContent(ticketContent)
	ticketCreatedDtime := runtime.ExtractTime(ticket.GetCreatedDtime())
	m.SetCreatedDtime(ticketCreatedDtime)
	ticketCreator := ticket.GetCreator()
	m.SetCreator(ticketCreator)
	ticketRepliedDtime := runtime.ExtractTime(ticket.GetRepliedDtime())
	m.SetRepliedDtime(ticketRepliedDtime)
	ticketReplier := ticket.GetReplier()
	m.SetReplier(ticketReplier)
	ticketReply := ticket.GetReply()
	m.SetReply(ticketReply)
	ticketReviewRemark := ticket.GetReviewRemark()
	m.SetReviewRemark(ticketReviewRemark)
	ticketReviewedDtime := runtime.ExtractTime(ticket.GetReviewedDtime())
	m.SetReviewedDtime(ticketReviewedDtime)
	ticketReviewer := ticket.GetReviewer()
	m.SetReviewer(ticketReviewer)
	ticketStatus := toEntTicket_Status(ticket.GetStatus())
	m.SetStatus(ticketStatus)
	ticketUpdatedDtime := runtime.ExtractTime(ticket.GetUpdatedDtime())
	m.SetUpdatedDtime(ticketUpdatedDtime)
	var ticketApplications uuid.UUID
	if err := (&ticketApplications).UnmarshalBinary(ticket.GetApplications().GetId()); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	}
	m.SetApplicationsID(ticketApplications)
	for _, item := range ticket.GetAttachments() {
		attachments := int(item.GetId())
		m.AddAttachmentIDs(attachments)
	}
	res, err := m.Save(ctx)
	switch {
	case err == nil:
		proto, err := toProtoTicket(res)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "internal error: %s", err)
		}
		return proto, nil
	case sqlgraph.IsUniqueConstraintError(err):
		return nil, status.Errorf(codes.AlreadyExists, "already exists: %s", err)
	case ent.IsConstraintError(err):
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

// Get implements TicketServiceServer.Get
func (svc *TicketService) Get(ctx context.Context, req *GetTicketRequest) (*Ticket, error) {
	var (
		err error
		get *ent.Ticket
	)
	id := int(req.GetId())
	switch req.GetView() {
	case GetTicketRequest_VIEW_UNSPECIFIED, GetTicketRequest_BASIC:
		get, err = svc.client.Ticket.Get(ctx, id)
	case GetTicketRequest_WITH_EDGE_IDS:
		get, err = svc.client.Ticket.Query().
			Where(ticket.ID(id)).
			WithApplications(func(query *ent.ApplicationQuery) {
				query.Select(application.FieldID)
			}).
			WithAttachments(func(query *ent.AttachmentQuery) {
				query.Select(attachment.FieldID)
			}).
			Only(ctx)
	default:
		return nil, status.Error(codes.InvalidArgument, "invalid argument: unknown view")
	}
	switch {
	case err == nil:
		return toProtoTicket(get)
	case ent.IsNotFound(err):
		return nil, status.Errorf(codes.NotFound, "not found: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}
	return nil, nil

}

// Update implements TicketServiceServer.Update
func (svc *TicketService) Update(ctx context.Context, req *UpdateTicketRequest) (*Ticket, error) {
	ticket := req.GetTicket()
	ticketID := int(ticket.GetId())
	m := svc.client.Ticket.UpdateOneID(ticketID)
	var ticketApplicationID uuid.UUID
	if err := (&ticketApplicationID).UnmarshalBinary(ticket.GetApplicationId()); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	}
	m.SetApplicationID(ticketApplicationID)
	ticketContent := ticket.GetContent()
	m.SetContent(ticketContent)
	ticketCreator := ticket.GetCreator()
	m.SetCreator(ticketCreator)
	ticketRepliedDtime := runtime.ExtractTime(ticket.GetRepliedDtime())
	m.SetRepliedDtime(ticketRepliedDtime)
	ticketReplier := ticket.GetReplier()
	m.SetReplier(ticketReplier)
	ticketReply := ticket.GetReply()
	m.SetReply(ticketReply)
	ticketReviewRemark := ticket.GetReviewRemark()
	m.SetReviewRemark(ticketReviewRemark)
	ticketReviewedDtime := runtime.ExtractTime(ticket.GetReviewedDtime())
	m.SetReviewedDtime(ticketReviewedDtime)
	ticketReviewer := ticket.GetReviewer()
	m.SetReviewer(ticketReviewer)
	ticketStatus := toEntTicket_Status(ticket.GetStatus())
	m.SetStatus(ticketStatus)
	ticketUpdatedDtime := runtime.ExtractTime(ticket.GetUpdatedDtime())
	m.SetUpdatedDtime(ticketUpdatedDtime)
	var ticketApplications uuid.UUID
	if err := (&ticketApplications).UnmarshalBinary(ticket.GetApplications().GetId()); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	}
	m.SetApplicationsID(ticketApplications)
	for _, item := range ticket.GetAttachments() {
		attachments := int(item.GetId())
		m.AddAttachmentIDs(attachments)
	}
	res, err := m.Save(ctx)
	switch {
	case err == nil:
		proto, err := toProtoTicket(res)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "internal error: %s", err)
		}
		return proto, nil
	case sqlgraph.IsUniqueConstraintError(err):
		return nil, status.Errorf(codes.AlreadyExists, "already exists: %s", err)
	case ent.IsConstraintError(err):
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

// Delete implements TicketServiceServer.Delete
func (svc *TicketService) Delete(ctx context.Context, req *DeleteTicketRequest) (*emptypb.Empty, error) {
	var err error
	id := int(req.GetId())
	err = svc.client.Ticket.DeleteOneID(id).Exec(ctx)
	switch {
	case err == nil:
		return &emptypb.Empty{}, nil
	case ent.IsNotFound(err):
		return nil, status.Errorf(codes.NotFound, "not found: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

// List implements TicketServiceServer.List
func (svc *TicketService) List(ctx context.Context, req *ListTicketRequest) (*ListTicketResponse, error) {
	var (
		err      error
		entList  []*ent.Ticket
		pageSize int
	)
	pageSize = int(req.GetPageSize())
	switch {
	case pageSize < 0:
		return nil, status.Errorf(codes.InvalidArgument, "page size cannot be less than zero")
	case pageSize == 0 || pageSize > entproto.MaxPageSize:
		pageSize = entproto.MaxPageSize
	}
	listQuery := svc.client.Ticket.Query().
		Order(ent.Desc(ticket.FieldID)).
		Limit(pageSize + 1)
	if req.GetPageToken() != "" {
		bytes, err := base64.StdEncoding.DecodeString(req.PageToken)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "page token is invalid")
		}
		token, err := strconv.ParseInt(string(bytes), 10, 32)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "page token is invalid")
		}
		pageToken := int(token)
		listQuery = listQuery.
			Where(ticket.IDLTE(pageToken))
	}
	switch req.GetView() {
	case ListTicketRequest_VIEW_UNSPECIFIED, ListTicketRequest_BASIC:
		entList, err = listQuery.All(ctx)
	case ListTicketRequest_WITH_EDGE_IDS:
		entList, err = listQuery.
			WithApplications(func(query *ent.ApplicationQuery) {
				query.Select(application.FieldID)
			}).
			WithAttachments(func(query *ent.AttachmentQuery) {
				query.Select(attachment.FieldID)
			}).
			All(ctx)
	}
	switch {
	case err == nil:
		var nextPageToken string
		if len(entList) == pageSize+1 {
			nextPageToken = base64.StdEncoding.EncodeToString(
				[]byte(fmt.Sprintf("%v", entList[len(entList)-1].ID)))
			entList = entList[:len(entList)-1]
		}
		var pbList []*Ticket
		for _, entEntity := range entList {
			pbEntity, err := toProtoTicket(entEntity)
			if err != nil {
				return nil, status.Errorf(codes.Internal, "internal error: %s", err)
			}
			pbList = append(pbList, pbEntity)
		}
		return &ListTicketResponse{
			TicketList:    pbList,
			NextPageToken: nextPageToken,
		}, nil
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}
