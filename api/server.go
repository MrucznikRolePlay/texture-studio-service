package api

import (
	"context"
	"github.com/MruV-RP/mruv-pb-go/texturestudio"
)

type Server struct {
}

func NewServer() *Server {
	return &Server{}
}

func (s Server) CreateServer(ctx context.Context, request *texturestudio.CreateServerRequest) (*texturestudio.CreateServerResponse, error) {
	panic("implement me")
}

func (s Server) MyServer(ctx context.Context, request *texturestudio.MyServerRequest) (*texturestudio.MyServerResponse, error) {
	panic("implement me")
}

func (s Server) DeleteServer(ctx context.Context, request *texturestudio.DeleteServerRequest) (*texturestudio.DeleteServerResponse, error) {
	panic("implement me")
}

func (s Server) StartServer(ctx context.Context, request *texturestudio.StartServerRequest) (*texturestudio.StartServerResponse, error) {
	panic("implement me")
}

func (s Server) StopServer(ctx context.Context, request *texturestudio.StopServerRequest) (*texturestudio.StopServerResponse, error) {
	panic("implement me")
}

func (s Server) RestartServer(ctx context.Context, request *texturestudio.RestartServerRequest) (*texturestudio.RestartServerResponse, error) {
	panic("implement me")
}

func (s Server) ServerStatus(ctx context.Context, request *texturestudio.ServerStatusRequest) (*texturestudio.ServerStatusResponse, error) {
	panic("implement me")
}

func (s Server) UploadProject(ctx context.Context, request *texturestudio.UploadProjectRequest) (*texturestudio.UploadProjectResponse, error) {
	panic("implement me")
}

func (s Server) GetProject(ctx context.Context, request *texturestudio.GetProjectRequest) (*texturestudio.GetProjectResponse, error) {
	panic("implement me")
}

func (s Server) GetProjects(ctx context.Context, request *texturestudio.GetProjectsRequest) (*texturestudio.GetProjectsResponse, error) {
	panic("implement me")
}

func (s Server) SubscribeToProjectsChanges(request *texturestudio.SubscribeToProjectsChangesRequest, server texturestudio.TextureStudioService_SubscribeToProjectsChangesServer) error {
	panic("implement me")
}
