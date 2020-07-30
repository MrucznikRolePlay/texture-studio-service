package engine

import (
	"context"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
)

type TextureStudioServer struct {
	ctx      context.Context
	stopFunc context.CancelFunc
}

func (s *TextureStudioServer) Start() error {
	if s.IsServerRunning() {
		return errors.New("server is already running")
	}

	s.ctx, s.stopFunc = context.WithCancel(context.Background())
	cmd := exec.CommandContext(s.ctx, "sampctl", "package", "run")
	cmd.Dir = viper.GetString("texture_studio.path")
	cmd.Stdout = logrus.StandardLogger().Out
	cmd.Stderr = logrus.StandardLogger().Out

	go func() {
		logrus.Infoln("Starting texture studio server.")
		if err := cmd.Run(); err != nil {
			logrus.Error("Texture studio server has shut down with error: ")
			logrus.Errorln(err)
			return
		}
		logrus.Infoln("Texture studio server has shut down successfully.")
	}()
	return nil
}

func (s *TextureStudioServer) Stop() error {
	if !s.IsServerRunning() {
		return errors.New("server is not running")
	}

	s.stopFunc()
	return nil
}

func (s *TextureStudioServer) Restart() error {
	if err := s.Stop(); err != nil {
		logrus.Infoln("Triggered a server restart, but server is not running.")
	}
	return s.Start()
}

func (s *TextureStudioServer) IsServerRunning() bool {
	if s.ctx == nil {
		return false
	}

	select {
	case <-s.ctx.Done():
		return false
	default:
		return true
	}
}

func (s *TextureStudioServer) GetProjectCode(name string) (string, error) {
	path := fmt.Sprintf("%s/scriptfiles/tstudio/ExportMaps/%s.txt", viper.GetString("texture_studio.path"), name)
	content, err := ioutil.ReadFile(path)
	return string(content), err
}

func (s *TextureStudioServer) GetSavedProjects() []string {
	var files []string
	err := filepath.Walk(fmt.Sprintf("%s/scriptfiles/tstudio/ExportMaps", viper.GetString("texture_studio.path")),
		func(path string, info os.FileInfo, err error) error {
			files = append(files, path)
			return nil
		})
	if err != nil {
		panic(err)
	}
	return files
}

func (s *TextureStudioServer) UploadProject(name string, data []byte) error {
	err := ioutil.WriteFile(fmt.Sprintf("%s/scriptfiles/tstudio/ImportMaps/%s.txt", viper.GetString("texture_studio.path"), name), data, 0755)
	if err != nil {
		return err
	}

	
}
