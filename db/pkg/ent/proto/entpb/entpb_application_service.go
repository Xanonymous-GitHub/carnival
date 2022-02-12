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
	applicationassignmenthistory "github.com/Xanonymous-GitHub/carnival/db/pkg/ent/applicationassignmenthistory"
	applicationstatushistory "github.com/Xanonymous-GitHub/carnival/db/pkg/ent/applicationstatushistory"
	attachment "github.com/Xanonymous-GitHub/carnival/db/pkg/ent/attachment"
	ticket "github.com/Xanonymous-GitHub/carnival/db/pkg/ent/ticket"
	uuid "github.com/google/uuid"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	strings "strings"
)

// ApplicationService implements ApplicationServiceServer
type ApplicationService struct {
	client *ent.Client
	UnimplementedApplicationServiceServer
}

// NewApplicationService returns a new ApplicationService
func NewApplicationService(client *ent.Client) *ApplicationService {
	return &ApplicationService{
		client: client,
	}
}

func toProtoApplication_ApplicationStatus(e application.ApplicationStatus) Application_ApplicationStatus {
	if v, ok := Application_ApplicationStatus_value[strings.ToUpper(string(e))]; ok {
		return Application_ApplicationStatus(v)
	}
	return Application_ApplicationStatus(0)
}

func toEntApplication_ApplicationStatus(e Application_ApplicationStatus) application.ApplicationStatus {
	if v, ok := Application_ApplicationStatus_name[int32(e)]; ok {
		return application.ApplicationStatus(strings.ToLower(v))
	}
	return ""
}

func toProtoApplication_BotActiveStatus(e application.BotActiveStatus) Application_BotActiveStatus {
	if v, ok := Application_BotActiveStatus_value[strings.ToUpper(string(e))]; ok {
		return Application_BotActiveStatus(v)
	}
	return Application_BotActiveStatus(0)
}

func toEntApplication_BotActiveStatus(e Application_BotActiveStatus) application.BotActiveStatus {
	if v, ok := Application_BotActiveStatus_name[int32(e)]; ok {
		return application.BotActiveStatus(strings.ToLower(v))
	}
	return ""
}

func toProtoApplication_StoreType(e application.StoreType) Application_StoreType {
	if v, ok := Application_StoreType_value[strings.ToUpper(string(e))]; ok {
		return Application_StoreType(v)
	}
	return Application_StoreType(0)
}

func toEntApplication_StoreType(e Application_StoreType) application.StoreType {
	if v, ok := Application_StoreType_name[int32(e)]; ok {
		return application.StoreType(strings.ToLower(v))
	}
	return ""
}

// toProtoApplication transforms the ent type to the pb type
func toProtoApplication(e *ent.Application) (*Application, error) {
	v := &Application{}
	applicantemail := e.ApplicantEmail
	v.ApplicantEmail = applicantemail
	applicantname := e.ApplicantName
	v.ApplicantName = applicantname
	applicationmid := e.ApplicationMid
	v.ApplicationMid = applicationmid
	applicationstatus := toProtoApplication_ApplicationStatus(e.ApplicationStatus)
	v.ApplicationStatus = applicationstatus
	assignee := e.Assignee
	v.Assignee = assignee
	assigner := e.Assigner
	v.Assigner = assigner
	basicid := e.BasicID
	v.BasicId = basicid
	botactivestatus := toProtoApplication_BotActiveStatus(e.BotActiveStatus)
	v.BotActiveStatus = botactivestatus
	botdisplayname := e.BotDisplayName
	v.BotDisplayName = botdisplayname
	botmid := e.BotMid
	v.BotMid = botmid
	createddtime := timestamppb.New(e.CreatedDtime)
	v.CreatedDtime = createddtime
	id, err := e.ID.MarshalBinary()
	if err != nil {
		return nil, err
	}
	v.Id = id
	premiumid := e.PremiumID
	v.PremiumId = premiumid
	remark := e.Remark
	v.Remark = remark
	reviewcomment := e.ReviewComment
	v.ReviewComment = reviewcomment
	storetype := toProtoApplication_StoreType(e.StoreType)
	v.StoreType = storetype
	updatedtime := timestamppb.New(e.UpdateDtime)
	v.UpdateDtime = updatedtime
	websiteurl := e.WebsiteURL
	v.WebsiteUrl = websiteurl
	for _, edg := range e.Edges.AssignmentHistories {
		id := int32(edg.ID)
		v.AssignmentHistories = append(v.AssignmentHistories, &ApplicationAssignmentHistory{
			Id: id,
		})
	}
	for _, edg := range e.Edges.Attachments {
		id := int32(edg.ID)
		v.Attachments = append(v.Attachments, &Attachment{
			Id: id,
		})
	}
	for _, edg := range e.Edges.StatusHistories {
		id := int32(edg.ID)
		v.StatusHistories = append(v.StatusHistories, &ApplicationStatusHistory{
			Id: id,
		})
	}
	if edg := e.Edges.Tickets; edg != nil {
		id := int32(edg.ID)
		v.Tickets = &Ticket{
			Id: id,
		}
	}
	return v, nil
}

// Create implements ApplicationServiceServer.Create
func (svc *ApplicationService) Create(ctx context.Context, req *CreateApplicationRequest) (*Application, error) {
	application := req.GetApplication()
	m := svc.client.Application.Create()
	applicationApplicantEmail := application.GetApplicantEmail()
	m.SetApplicantEmail(applicationApplicantEmail)
	applicationApplicantName := application.GetApplicantName()
	m.SetApplicantName(applicationApplicantName)
	applicationApplicationMid := application.GetApplicationMid()
	m.SetApplicationMid(applicationApplicationMid)
	applicationApplicationStatus := toEntApplication_ApplicationStatus(application.GetApplicationStatus())
	m.SetApplicationStatus(applicationApplicationStatus)
	applicationAssignee := application.GetAssignee()
	m.SetAssignee(applicationAssignee)
	applicationAssigner := application.GetAssigner()
	m.SetAssigner(applicationAssigner)
	applicationBasicID := application.GetBasicId()
	m.SetBasicID(applicationBasicID)
	applicationBotActiveStatus := toEntApplication_BotActiveStatus(application.GetBotActiveStatus())
	m.SetBotActiveStatus(applicationBotActiveStatus)
	applicationBotDisplayName := application.GetBotDisplayName()
	m.SetBotDisplayName(applicationBotDisplayName)
	applicationBotMid := application.GetBotMid()
	m.SetBotMid(applicationBotMid)
	applicationCreatedDtime := runtime.ExtractTime(application.GetCreatedDtime())
	m.SetCreatedDtime(applicationCreatedDtime)
	applicationPremiumID := application.GetPremiumId()
	m.SetPremiumID(applicationPremiumID)
	applicationRemark := application.GetRemark()
	m.SetRemark(applicationRemark)
	applicationReviewComment := application.GetReviewComment()
	m.SetReviewComment(applicationReviewComment)
	applicationStoreType := toEntApplication_StoreType(application.GetStoreType())
	m.SetStoreType(applicationStoreType)
	applicationUpdateDtime := runtime.ExtractTime(application.GetUpdateDtime())
	m.SetUpdateDtime(applicationUpdateDtime)
	applicationWebsiteURL := application.GetWebsiteUrl()
	m.SetWebsiteURL(applicationWebsiteURL)
	for _, item := range application.GetAssignmentHistories() {
		assignmenthistories := int(item.GetId())
		m.AddAssignmentHistoryIDs(assignmenthistories)
	}
	for _, item := range application.GetAttachments() {
		attachments := int(item.GetId())
		m.AddAttachmentIDs(attachments)
	}
	for _, item := range application.GetStatusHistories() {
		statushistories := int(item.GetId())
		m.AddStatusHistoryIDs(statushistories)
	}
	applicationTickets := int(application.GetTickets().GetId())
	m.SetTicketsID(applicationTickets)
	res, err := m.Save(ctx)
	switch {
	case err == nil:
		proto, err := toProtoApplication(res)
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

// Get implements ApplicationServiceServer.Get
func (svc *ApplicationService) Get(ctx context.Context, req *GetApplicationRequest) (*Application, error) {
	var (
		err error
		get *ent.Application
	)
	var id uuid.UUID
	if err := (&id).UnmarshalBinary(req.GetId()); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	}
	switch req.GetView() {
	case GetApplicationRequest_VIEW_UNSPECIFIED, GetApplicationRequest_BASIC:
		get, err = svc.client.Application.Get(ctx, id)
	case GetApplicationRequest_WITH_EDGE_IDS:
		get, err = svc.client.Application.Query().
			Where(application.ID(id)).
			WithAssignmentHistories(func(query *ent.ApplicationAssignmentHistoryQuery) {
				query.Select(applicationassignmenthistory.FieldID)
			}).
			WithAttachments(func(query *ent.AttachmentQuery) {
				query.Select(attachment.FieldID)
			}).
			WithStatusHistories(func(query *ent.ApplicationStatusHistoryQuery) {
				query.Select(applicationstatushistory.FieldID)
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
		return toProtoApplication(get)
	case ent.IsNotFound(err):
		return nil, status.Errorf(codes.NotFound, "not found: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}
	return nil, nil

}

// Update implements ApplicationServiceServer.Update
func (svc *ApplicationService) Update(ctx context.Context, req *UpdateApplicationRequest) (*Application, error) {
	application := req.GetApplication()
	var applicationID uuid.UUID
	if err := (&applicationID).UnmarshalBinary(application.GetId()); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	}
	m := svc.client.Application.UpdateOneID(applicationID)
	applicationApplicantEmail := application.GetApplicantEmail()
	m.SetApplicantEmail(applicationApplicantEmail)
	applicationApplicantName := application.GetApplicantName()
	m.SetApplicantName(applicationApplicantName)
	applicationApplicationMid := application.GetApplicationMid()
	m.SetApplicationMid(applicationApplicationMid)
	applicationApplicationStatus := toEntApplication_ApplicationStatus(application.GetApplicationStatus())
	m.SetApplicationStatus(applicationApplicationStatus)
	applicationAssignee := application.GetAssignee()
	m.SetAssignee(applicationAssignee)
	applicationAssigner := application.GetAssigner()
	m.SetAssigner(applicationAssigner)
	applicationBasicID := application.GetBasicId()
	m.SetBasicID(applicationBasicID)
	applicationBotActiveStatus := toEntApplication_BotActiveStatus(application.GetBotActiveStatus())
	m.SetBotActiveStatus(applicationBotActiveStatus)
	applicationBotDisplayName := application.GetBotDisplayName()
	m.SetBotDisplayName(applicationBotDisplayName)
	applicationBotMid := application.GetBotMid()
	m.SetBotMid(applicationBotMid)
	applicationPremiumID := application.GetPremiumId()
	m.SetPremiumID(applicationPremiumID)
	applicationRemark := application.GetRemark()
	m.SetRemark(applicationRemark)
	applicationReviewComment := application.GetReviewComment()
	m.SetReviewComment(applicationReviewComment)
	applicationStoreType := toEntApplication_StoreType(application.GetStoreType())
	m.SetStoreType(applicationStoreType)
	applicationUpdateDtime := runtime.ExtractTime(application.GetUpdateDtime())
	m.SetUpdateDtime(applicationUpdateDtime)
	applicationWebsiteURL := application.GetWebsiteUrl()
	m.SetWebsiteURL(applicationWebsiteURL)
	for _, item := range application.GetAssignmentHistories() {
		assignmenthistories := int(item.GetId())
		m.AddAssignmentHistoryIDs(assignmenthistories)
	}
	for _, item := range application.GetAttachments() {
		attachments := int(item.GetId())
		m.AddAttachmentIDs(attachments)
	}
	for _, item := range application.GetStatusHistories() {
		statushistories := int(item.GetId())
		m.AddStatusHistoryIDs(statushistories)
	}
	applicationTickets := int(application.GetTickets().GetId())
	m.SetTicketsID(applicationTickets)
	res, err := m.Save(ctx)
	switch {
	case err == nil:
		proto, err := toProtoApplication(res)
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

// Delete implements ApplicationServiceServer.Delete
func (svc *ApplicationService) Delete(ctx context.Context, req *DeleteApplicationRequest) (*emptypb.Empty, error) {
	var err error
	var id uuid.UUID
	if err := (&id).UnmarshalBinary(req.GetId()); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	}
	err = svc.client.Application.DeleteOneID(id).Exec(ctx)
	switch {
	case err == nil:
		return &emptypb.Empty{}, nil
	case ent.IsNotFound(err):
		return nil, status.Errorf(codes.NotFound, "not found: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

// List implements ApplicationServiceServer.List
func (svc *ApplicationService) List(ctx context.Context, req *ListApplicationRequest) (*ListApplicationResponse, error) {
	var (
		err      error
		entList  []*ent.Application
		pageSize int
	)
	pageSize = int(req.GetPageSize())
	switch {
	case pageSize < 0:
		return nil, status.Errorf(codes.InvalidArgument, "page size cannot be less than zero")
	case pageSize == 0 || pageSize > entproto.MaxPageSize:
		pageSize = entproto.MaxPageSize
	}
	listQuery := svc.client.Application.Query().
		Order(ent.Desc(application.FieldID)).
		Limit(pageSize + 1)
	if req.GetPageToken() != "" {
		bytes, err := base64.StdEncoding.DecodeString(req.PageToken)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "page token is invalid")
		}
		pageToken, err := uuid.ParseBytes(bytes)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "page token is invalid")
		}
		listQuery = listQuery.
			Where(application.IDLTE(pageToken))
	}
	switch req.GetView() {
	case ListApplicationRequest_VIEW_UNSPECIFIED, ListApplicationRequest_BASIC:
		entList, err = listQuery.All(ctx)
	case ListApplicationRequest_WITH_EDGE_IDS:
		entList, err = listQuery.
			WithAssignmentHistories(func(query *ent.ApplicationAssignmentHistoryQuery) {
				query.Select(applicationassignmenthistory.FieldID)
			}).
			WithAttachments(func(query *ent.AttachmentQuery) {
				query.Select(attachment.FieldID)
			}).
			WithStatusHistories(func(query *ent.ApplicationStatusHistoryQuery) {
				query.Select(applicationstatushistory.FieldID)
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
		var pbList []*Application
		for _, entEntity := range entList {
			pbEntity, err := toProtoApplication(entEntity)
			if err != nil {
				return nil, status.Errorf(codes.Internal, "internal error: %s", err)
			}
			pbList = append(pbList, pbEntity)
		}
		return &ListApplicationResponse{
			ApplicationList: pbList,
			NextPageToken:   nextPageToken,
		}, nil
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}