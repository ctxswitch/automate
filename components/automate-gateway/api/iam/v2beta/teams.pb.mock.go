// Code generated by protoc-gen-grpc-mock. DO NOT EDIT.
// source: components/automate-gateway/api/iam/v2beta/teams.proto

package v2beta

import (
	"context"

	request "github.com/chef/automate/components/automate-gateway/api/iam/v2beta/request"
	response "github.com/chef/automate/components/automate-gateway/api/iam/v2beta/response"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// verify that the mock satisfies the TeamsServer interface (at compile time)
var _ TeamsServer = &TeamsServerMock{}

// NewTeamsServerMock gives you a fresh instance of TeamsServerMock.
func NewTeamsServerMock() *TeamsServerMock {
	return &TeamsServerMock{validateRequests: true}
}

// NewTeamsServerMockWithoutValidation gives you a fresh instance of
// TeamsServerMock which does not attempt to validate requests before passing
// them to their respective '*Func'.
func NewTeamsServerMockWithoutValidation() *TeamsServerMock {
	return &TeamsServerMock{}
}

// TeamsServerMock is the mock-what-you-want struct that stubs all not-overridden
// methods with "not implemented" returns
type TeamsServerMock struct {
	validateRequests          bool
	ListTeamsFunc             func(context.Context, *request.ListTeamsReq) (*response.ListTeamsResp, error)
	GetTeamFunc               func(context.Context, *request.GetTeamReq) (*response.GetTeamResp, error)
	CreateTeamFunc            func(context.Context, *request.CreateTeamReq) (*response.CreateTeamResp, error)
	UpdateTeamFunc            func(context.Context, *request.UpdateTeamReq) (*response.UpdateTeamResp, error)
	DeleteTeamFunc            func(context.Context, *request.DeleteTeamReq) (*response.DeleteTeamResp, error)
	GetTeamMembershipFunc     func(context.Context, *request.GetTeamMembershipReq) (*response.GetTeamMembershipResp, error)
	AddTeamMembersFunc        func(context.Context, *request.AddTeamMembersReq) (*response.AddTeamMembersResp, error)
	RemoveTeamMembersFunc     func(context.Context, *request.RemoveTeamMembersReq) (*response.RemoveTeamMembersResp, error)
	GetTeamsForMemberFunc     func(context.Context, *request.GetTeamsForMemberReq) (*response.GetTeamsForMemberResp, error)
	ApplyV2DataMigrationsFunc func(context.Context, *request.ApplyV2DataMigrationsReq) (*response.ApplyV2DataMigrationsResp, error)
}

func (m *TeamsServerMock) ListTeams(ctx context.Context, req *request.ListTeamsReq) (*response.ListTeamsResp, error) {
	if msg, ok := interface{}(req).(interface{ Validate() error }); m.validateRequests && ok {
		if err := msg.Validate(); err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if f := m.ListTeamsFunc; f != nil {
		return f(ctx, req)
	}
	return nil, status.Error(codes.Internal, "mock: 'ListTeams' not implemented")
}

func (m *TeamsServerMock) GetTeam(ctx context.Context, req *request.GetTeamReq) (*response.GetTeamResp, error) {
	if msg, ok := interface{}(req).(interface{ Validate() error }); m.validateRequests && ok {
		if err := msg.Validate(); err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if f := m.GetTeamFunc; f != nil {
		return f(ctx, req)
	}
	return nil, status.Error(codes.Internal, "mock: 'GetTeam' not implemented")
}

func (m *TeamsServerMock) CreateTeam(ctx context.Context, req *request.CreateTeamReq) (*response.CreateTeamResp, error) {
	if msg, ok := interface{}(req).(interface{ Validate() error }); m.validateRequests && ok {
		if err := msg.Validate(); err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if f := m.CreateTeamFunc; f != nil {
		return f(ctx, req)
	}
	return nil, status.Error(codes.Internal, "mock: 'CreateTeam' not implemented")
}

func (m *TeamsServerMock) UpdateTeam(ctx context.Context, req *request.UpdateTeamReq) (*response.UpdateTeamResp, error) {
	if msg, ok := interface{}(req).(interface{ Validate() error }); m.validateRequests && ok {
		if err := msg.Validate(); err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if f := m.UpdateTeamFunc; f != nil {
		return f(ctx, req)
	}
	return nil, status.Error(codes.Internal, "mock: 'UpdateTeam' not implemented")
}

func (m *TeamsServerMock) DeleteTeam(ctx context.Context, req *request.DeleteTeamReq) (*response.DeleteTeamResp, error) {
	if msg, ok := interface{}(req).(interface{ Validate() error }); m.validateRequests && ok {
		if err := msg.Validate(); err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if f := m.DeleteTeamFunc; f != nil {
		return f(ctx, req)
	}
	return nil, status.Error(codes.Internal, "mock: 'DeleteTeam' not implemented")
}

func (m *TeamsServerMock) GetTeamMembership(ctx context.Context, req *request.GetTeamMembershipReq) (*response.GetTeamMembershipResp, error) {
	if msg, ok := interface{}(req).(interface{ Validate() error }); m.validateRequests && ok {
		if err := msg.Validate(); err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if f := m.GetTeamMembershipFunc; f != nil {
		return f(ctx, req)
	}
	return nil, status.Error(codes.Internal, "mock: 'GetTeamMembership' not implemented")
}

func (m *TeamsServerMock) AddTeamMembers(ctx context.Context, req *request.AddTeamMembersReq) (*response.AddTeamMembersResp, error) {
	if msg, ok := interface{}(req).(interface{ Validate() error }); m.validateRequests && ok {
		if err := msg.Validate(); err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if f := m.AddTeamMembersFunc; f != nil {
		return f(ctx, req)
	}
	return nil, status.Error(codes.Internal, "mock: 'AddTeamMembers' not implemented")
}

func (m *TeamsServerMock) RemoveTeamMembers(ctx context.Context, req *request.RemoveTeamMembersReq) (*response.RemoveTeamMembersResp, error) {
	if msg, ok := interface{}(req).(interface{ Validate() error }); m.validateRequests && ok {
		if err := msg.Validate(); err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if f := m.RemoveTeamMembersFunc; f != nil {
		return f(ctx, req)
	}
	return nil, status.Error(codes.Internal, "mock: 'RemoveTeamMembers' not implemented")
}

func (m *TeamsServerMock) GetTeamsForMember(ctx context.Context, req *request.GetTeamsForMemberReq) (*response.GetTeamsForMemberResp, error) {
	if msg, ok := interface{}(req).(interface{ Validate() error }); m.validateRequests && ok {
		if err := msg.Validate(); err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if f := m.GetTeamsForMemberFunc; f != nil {
		return f(ctx, req)
	}
	return nil, status.Error(codes.Internal, "mock: 'GetTeamsForMember' not implemented")
}

func (m *TeamsServerMock) ApplyV2DataMigrations(ctx context.Context, req *request.ApplyV2DataMigrationsReq) (*response.ApplyV2DataMigrationsResp, error) {
	if msg, ok := interface{}(req).(interface{ Validate() error }); m.validateRequests && ok {
		if err := msg.Validate(); err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if f := m.ApplyV2DataMigrationsFunc; f != nil {
		return f(ctx, req)
	}
	return nil, status.Error(codes.Internal, "mock: 'ApplyV2DataMigrations' not implemented")
}

// Reset resets all overridden functions
func (m *TeamsServerMock) Reset() {
	m.ListTeamsFunc = nil
	m.GetTeamFunc = nil
	m.CreateTeamFunc = nil
	m.UpdateTeamFunc = nil
	m.DeleteTeamFunc = nil
	m.GetTeamMembershipFunc = nil
	m.AddTeamMembersFunc = nil
	m.RemoveTeamMembersFunc = nil
	m.GetTeamsForMemberFunc = nil
	m.ApplyV2DataMigrationsFunc = nil
}
