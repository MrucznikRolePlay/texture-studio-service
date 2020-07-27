package engine

import (
	"context"
	"errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os/exec"
)

type TextureStudioServer struct {
	ctx      context.Context
	stopFunc context.CancelFunc
}

func (s *TextureStudioServer) Start() error {
	s.ctx, s.stopFunc = context.WithCancel(context.Background())
	cmd := exec.CommandContext(s.ctx, "sampctl", "package run")
	cmd.Dir = viper.GetString("texture_studio.path")
	successfullStart := true

	go func() {
		logrus.Infoln("Starting texture studio server.")
		if err := cmd.Run(); err != nil {
			logrus.Error("Texture studio server has shut down with error: ")
			logrus.Errorln(err)
		}
		logrus.Infoln("Texture studio server has shut down successfully.")
		successfullStart = false
	}()

	if successfullStart {
		return nil
	} else {
		return errors.New("server did not started successfully")
	}
}

func (s *TextureStudioServer) Stop() error {
	select {
	case <-s.ctx.Done():
		return errors.New("server is not running")
	default:
		s.stopFunc()
		return nil
	}
}

func (s *TextureStudioServer) Restart() error {
	if err := s.Stop(); err != nil {
		logrus.Infoln("Triggered a server restart, but server is not running.")
	}
	return s.Start()
}

func (s *TextureStudioServer) IsServerRunning() bool {
	select {
	case <-s.ctx.Done():
		return true
	default:
		return false
	}
}
