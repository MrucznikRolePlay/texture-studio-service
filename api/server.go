package api

import (
	"context"
	"github.com/MruV-RP/mruv-pb-go/texturestudio"
	"github.com/MruV-RP/mruv-service-template/engine"
)

type Server struct {
	tss *engine.TextureStudioServer
}

func NewServer() *Server {
	return &Server{tss: &engine.TextureStudioServer{}}
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
	err := s.tss.Start()
	return &texturestudio.StartServerResponse{}, err
}

func (s Server) StopServer(ctx context.Context, request *texturestudio.StopServerRequest) (*texturestudio.StopServerResponse, error) {
	err := s.tss.Stop()
	return &texturestudio.StopServerResponse{}, err
}

func (s Server) RestartServer(ctx context.Context, request *texturestudio.RestartServerRequest) (*texturestudio.RestartServerResponse, error) {
	err := s.tss.Restart()
	return &texturestudio.RestartServerResponse{}, err
}

func (s Server) ServerStatus(ctx context.Context, request *texturestudio.ServerStatusRequest) (*texturestudio.ServerStatusResponse, error) {
	status := texturestudio.ServerStatus_UNKNOWN_STATUS
	if s.tss.IsServerRunning() {
		status = texturestudio.ServerStatus_ON
	} else {
		status = texturestudio.ServerStatus_OFF
	}
	return &texturestudio.ServerStatusResponse{
		Port:   0,
		Status: status,
	}, nil

}

func (s Server) UploadProject(ctx context.Context, request *texturestudio.UploadProjectRequest) (*texturestudio.UploadProjectResponse, error) {
	panic("implement me")
}

func (s Server) GetProject(ctx context.Context, request *texturestudio.GetProjectRequest) (*texturestudio.GetProjectResponse, error) {
	code, err := s.tss.GetProjectCode(request.Name)
	if err != nil {
		return nil, err
	}
	response := &texturestudio.GetProjectResponse{Code: code}
	return response, nil
}

func (s Server) GetProjects(ctx context.Context, request *texturestudio.GetProjectsRequest) (*texturestudio.GetProjectsResponse, error) {
	projects := s.tss.GetSavedProjects()
	return &texturestudio.GetProjectsResponse{
		Names: projects,
	}, nil
}

func (s Server) SubscribeToProjectsChanges(request *texturestudio.SubscribeToProjectsChangesRequest, server texturestudio.TextureStudioService_SubscribeToProjectsChangesServer) error {
	panic("implement me")
}
