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
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	strconv "strconv"
	strings "strings"
)

// AttachmentService implements AttachmentServiceServer
type AttachmentService struct {
	client *ent.Client
	UnimplementedAttachmentServiceServer
}

// NewAttachmentService returns a new AttachmentService
func NewAttachmentService(client *ent.Client) *AttachmentService {
	return &AttachmentService{
		client: client,
	}
}

func toProtoAttachment_AType(e attachment.AType) Attachment_AType {
	if v, ok := Attachment_AType_value[strings.ToUpper(string(e))]; ok {
		return Attachment_AType(v)
	}
	return Attachment_AType(0)
}

func toEntAttachment_AType(e Attachment_AType) attachment.AType {
	if v, ok := Attachment_AType_name[int32(e)]; ok {
		return attachment.AType(strings.ToLower(v))
	}
	return ""
}

// toProtoAttachment transforms the ent type to the pb type
func toProtoAttachment(e *ent.Attachment) (*Attachment, error) {
	v := &Attachment{}
	atype := toProtoAttachment_AType(e.AType)
	v.AType = atype
	applicationid, err := e.ApplicationID.MarshalBinary()
	if err != nil {
		return nil, err
	}
	v.ApplicationId = applicationid
	createddtime := timestamppb.New(e.CreatedDtime)
	v.CreatedDtime = createddtime
	id := int32(e.ID)
	v.Id = id
	obshash := e.ObsHash
	v.ObsHash = obshash
	obsoid := e.ObsOid
	v.ObsOid = obsoid
	if e.TicketID != nil {
		ticketid := wrapperspb.Int32(*e.TicketID)
		v.TicketId = ticketid
	}
	if edg := e.Edges.Applications; edg != nil {
		id, err := edg.ID.MarshalBinary()
		if err != nil {
			return nil, err
		}
		v.Applications = &Application{
			Id: id,
		}
	}
	if edg := e.Edges.Tickets; edg != nil {
		id := int32(edg.ID)
		v.Tickets = &Ticket{
			Id: id,
		}
	}
	return v, nil
}

// Create implements AttachmentServiceServer.Create
func (svc *AttachmentService) Create(ctx context.Context, req *CreateAttachmentRequest) (*Attachment, error) {
	attachment := req.GetAttachment()
	m := svc.client.Attachment.Create()
	attachmentAType := toEntAttachment_AType(attachment.GetAType())
	m.SetAType(attachmentAType)
	var attachmentApplicationID uuid.UUID
	if err := (&attachmentApplicationID).UnmarshalBinary(attachment.GetApplicationId()); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	}
	m.SetApplicationID(attachmentApplicationID)
	attachmentCreatedDtime := runtime.ExtractTime(attachment.GetCreatedDtime())
	m.SetCreatedDtime(attachmentCreatedDtime)
	attachmentObsHash := attachment.GetObsHash()
	m.SetObsHash(attachmentObsHash)
	attachmentObsOid := attachment.GetObsOid()
	m.SetObsOid(attachmentObsOid)
	if attachment.GetTicketId() != nil {
		attachmentTicketID := int32(attachment.GetTicketId().GetValue())
		m.SetTicketID(attachmentTicketID)
	}
	var attachmentApplications uuid.UUID
	if err := (&attachmentApplications).UnmarshalBinary(attachment.GetApplications().GetId()); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	}
	m.SetApplicationsID(attachmentApplications)
	attachmentTickets := int(attachment.GetTickets().GetId())
	m.SetTicketsID(attachmentTickets)
	res, err := m.Save(ctx)
	switch {
	case err == nil:
		proto, err := toProtoAttachment(res)
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

// Get implements AttachmentServiceServer.Get
func (svc *AttachmentService) Get(ctx context.Context, req *GetAttachmentRequest) (*Attachment, error) {
	var (
		err error
		get *ent.Attachment
	)
	id := int(req.GetId())
	switch req.GetView() {
	case GetAttachmentRequest_VIEW_UNSPECIFIED, GetAttachmentRequest_BASIC:
		get, err = svc.client.Attachment.Get(ctx, id)
	case GetAttachmentRequest_WITH_EDGE_IDS:
		get, err = svc.client.Attachment.Query().
			Where(attachment.ID(id)).
			WithApplications(func(query *ent.ApplicationQuery) {
				query.Select(application.FieldID)
			}).
			WithTickets(func(query *ent.TicketQuery) {
				query.Select(ticket.FieldID)
			}).
			Only(ctx)
	default:
		return nil, status.Error(codes.InvalidArgument, "invalid argument: unknown view")
	}
	switch {
	case err == nil:
		return toProtoAttachment(get)
	case ent.IsNotFound(err):
		return nil, status.Errorf(codes.NotFound, "not found: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}
	return nil, nil

}

// Update implements AttachmentServiceServer.Update
func (svc *AttachmentService) Update(ctx context.Context, req *UpdateAttachmentRequest) (*Attachment, error) {
	attachment := req.GetAttachment()
	attachmentID := int(attachment.GetId())
	m := svc.client.Attachment.UpdateOneID(attachmentID)
	attachmentAType := toEntAttachment_AType(attachment.GetAType())
	m.SetAType(attachmentAType)
	var attachmentApplicationID uuid.UUID
	if err := (&attachmentApplicationID).UnmarshalBinary(attachment.GetApplicationId()); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	}
	m.SetApplicationID(attachmentApplicationID)
	attachmentObsHash := attachment.GetObsHash()
	m.SetObsHash(attachmentObsHash)
	attachmentObsOid := attachment.GetObsOid()
	m.SetObsOid(attachmentObsOid)
	if attachment.GetTicketId() != nil {
		attachmentTicketID := int32(attachment.GetTicketId().GetValue())
		m.SetTicketID(attachmentTicketID)
	}
	var attachmentApplications uuid.UUID
	if err := (&attachmentApplications).UnmarshalBinary(attachment.GetApplications().GetId()); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	}
	m.SetApplicationsID(attachmentApplications)
	attachmentTickets := int(attachment.GetTickets().GetId())
	m.SetTicketsID(attachmentTickets)
	res, err := m.Save(ctx)
	switch {
	case err == nil:
		proto, err := toProtoAttachment(res)
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

// Delete implements AttachmentServiceServer.Delete
func (svc *AttachmentService) Delete(ctx context.Context, req *DeleteAttachmentRequest) (*emptypb.Empty, error) {
	var err error
	id := int(req.GetId())
	err = svc.client.Attachment.DeleteOneID(id).Exec(ctx)
	switch {
	case err == nil:
		return &emptypb.Empty{}, nil
	case ent.IsNotFound(err):
		return nil, status.Errorf(codes.NotFound, "not found: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

// List implements AttachmentServiceServer.List
func (svc *AttachmentService) List(ctx context.Context, req *ListAttachmentRequest) (*ListAttachmentResponse, error) {
	var (
		err      error
		entList  []*ent.Attachment
		pageSize int
	)
	pageSize = int(req.GetPageSize())
	switch {
	case pageSize < 0:
		return nil, status.Errorf(codes.InvalidArgument, "page size cannot be less than zero")
	case pageSize == 0 || pageSize > entproto.MaxPageSize:
		pageSize = entproto.MaxPageSize
	}
	listQuery := svc.client.Attachment.Query().
		Order(ent.Desc(attachment.FieldID)).
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
			Where(attachment.IDLTE(pageToken))
	}
	switch req.GetView() {
	case ListAttachmentRequest_VIEW_UNSPECIFIED, ListAttachmentRequest_BASIC:
		entList, err = listQuery.All(ctx)
	case ListAttachmentRequest_WITH_EDGE_IDS:
		entList, err = listQuery.
			WithApplications(func(query *ent.ApplicationQuery) {
				query.Select(application.FieldID)
			}).
			WithTickets(func(query *ent.TicketQuery) {
				query.Select(ticket.FieldID)
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
		var pbList []*Attachment
		for _, entEntity := range entList {
			pbEntity, err := toProtoAttachment(entEntity)
			if err != nil {
				return nil, status.Errorf(codes.Internal, "internal error: %s", err)
			}
			pbList = append(pbList, pbEntity)
		}
		return &ListAttachmentResponse{
			AttachmentList: pbList,
			NextPageToken:  nextPageToken,
		}, nil
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}