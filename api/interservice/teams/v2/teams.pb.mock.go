// Code generated by protoc-gen-grpc-mock. DO NOT EDIT.
// source: api/interservice/teams/v2/teams.proto

package v2

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// verify that the mock satisfies the TeamsV2Server interface (at compile time)
var _ TeamsV2Server = &TeamsV2ServerMock{}

// NewTeamsV2ServerMock gives you a fresh instance of TeamsV2ServerMock.
func NewTeamsV2ServerMock() *TeamsV2ServerMock {
	return &TeamsV2ServerMock{validateRequests: true}
}

// NewTeamsV2ServerMockWithoutValidation gives you a fresh instance of
// TeamsV2ServerMock which does not attempt to validate requests before passing
// them to their respective '*Func'.
func NewTeamsV2ServerMockWithoutValidation() *TeamsV2ServerMock {
	return &TeamsV2ServerMock{}
}

// TeamsV2ServerMock is the mock-what-you-want struct that stubs all not-overridden
// methods with "not implemented" returns
type TeamsV2ServerMock struct {
	validateRequests      bool
	GetTeamFunc           func(context.Context, *GetTeamReq) (*GetTeamResp, error)
	ListTeamsFunc         func(context.Context, *ListTeamsReq) (*ListTeamsResp, error)
	CreateTeamFunc        func(context.Context, *CreateTeamReq) (*CreateTeamResp, error)
	UpdateTeamFunc        func(context.Context, *UpdateTeamReq) (*UpdateTeamResp, error)
	DeleteTeamFunc        func(context.Context, *DeleteTeamReq) (*DeleteTeamResp, error)
	AddTeamMembersFunc    func(context.Context, *AddTeamMembersReq) (*AddTeamMembersResp, error)
	RemoveTeamMembersFunc func(context.Context, *RemoveTeamMembersReq) (*RemoveTeamMembersResp, error)
	GetTeamsForMemberFunc func(context.Context, *GetTeamsForMemberReq) (*GetTeamsForMemberResp, error)
	GetTeamMembershipFunc func(context.Context, *GetTeamMembershipReq) (*GetTeamMembershipResp, error)
	UpgradeToV2Func       func(context.Context, *UpgradeToV2Req) (*UpgradeToV2Resp, error)
}

func (m *TeamsV2ServerMock) GetTeam(ctx context.Context, req *GetTeamReq) (*GetTeamResp, error) {
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

func (m *TeamsV2ServerMock) ListTeams(ctx context.Context, req *ListTeamsReq) (*ListTeamsResp, error) {
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

func (m *TeamsV2ServerMock) CreateTeam(ctx context.Context, req *CreateTeamReq) (*CreateTeamResp, error) {
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

func (m *TeamsV2ServerMock) UpdateTeam(ctx context.Context, req *UpdateTeamReq) (*UpdateTeamResp, error) {
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

func (m *TeamsV2ServerMock) DeleteTeam(ctx context.Context, req *DeleteTeamReq) (*DeleteTeamResp, error) {
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

func (m *TeamsV2ServerMock) AddTeamMembers(ctx context.Context, req *AddTeamMembersReq) (*AddTeamMembersResp, error) {
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

func (m *TeamsV2ServerMock) RemoveTeamMembers(ctx context.Context, req *RemoveTeamMembersReq) (*RemoveTeamMembersResp, error) {
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

func (m *TeamsV2ServerMock) GetTeamsForMember(ctx context.Context, req *GetTeamsForMemberReq) (*GetTeamsForMemberResp, error) {
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

func (m *TeamsV2ServerMock) GetTeamMembership(ctx context.Context, req *GetTeamMembershipReq) (*GetTeamMembershipResp, error) {
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

func (m *TeamsV2ServerMock) UpgradeToV2(ctx context.Context, req *UpgradeToV2Req) (*UpgradeToV2Resp, error) {
	if msg, ok := interface{}(req).(interface{ Validate() error }); m.validateRequests && ok {
		if err := msg.Validate(); err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if f := m.UpgradeToV2Func; f != nil {
		return f(ctx, req)
	}
	return nil, status.Error(codes.Internal, "mock: 'UpgradeToV2' not implemented")
}

// Reset resets all overridden functions
func (m *TeamsV2ServerMock) Reset() {
	m.GetTeamFunc = nil
	m.ListTeamsFunc = nil
	m.CreateTeamFunc = nil
	m.UpdateTeamFunc = nil
	m.DeleteTeamFunc = nil
	m.AddTeamMembersFunc = nil
	m.RemoveTeamMembersFunc = nil
	m.GetTeamsForMemberFunc = nil
	m.GetTeamMembershipFunc = nil
	m.UpgradeToV2Func = nil
}
