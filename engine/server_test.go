package engine

import (
	"context"
	"testing"
)

func TestTextureStudioServer_Restart(t *testing.T) {
	type fields struct {
		ctx      context.Context
		stopFunc context.CancelFunc
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &TextureStudioServer{
				ctx:      tt.fields.ctx,
				stopFunc: tt.fields.stopFunc,
			}
			if err := s.Restart(); (err != nil) != tt.wantErr {
				t.Errorf("Restart() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTextureStudioServer_Start(t *testing.T) {
	type fields struct {
		ctx      context.Context
		stopFunc context.CancelFunc
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &TextureStudioServer{
				ctx:      tt.fields.ctx,
				stopFunc: tt.fields.stopFunc,
			}
			if err := s.Start(); (err != nil) != tt.wantErr {
				t.Errorf("Start() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTextureStudioServer_Stop(t *testing.T) {
	type fields struct {
		ctx      context.Context
		stopFunc context.CancelFunc
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &TextureStudioServer{
				ctx:      tt.fields.ctx,
				stopFunc: tt.fields.stopFunc,
			}
			if err := s.Stop(); (err != nil) != tt.wantErr {
				t.Errorf("Stop() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
