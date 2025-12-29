package facade

import (
	"os"
	"path/filepath"

	"github.com/Station-Manager/errors"
	"github.com/Station-Manager/types"
)

// FetchUiConfig retrieves the UI configuration object. It returns an error if the service is not initialized, or the underlying
// ConfigService returns an error.
func (s *Service) FetchUiConfig() (*types.UiConfig, error) {
	const op errors.Op = "facade.Service.UiConfig"

	if !s.initialized.Load() {
		err := errors.New(op).Msg(errMsgServiceNotInit)
		s.LoggerService.ErrorWith().Err(err).Msg(errMsgServiceNotInit)
		return nil, err
	}

	return &types.UiConfig{
		Logbook: s.currentLogbook,
	}, nil
}

// GetDatabaseFileList retrieves a list of database file paths from the configured database directory.
// It returns an error if the service is not initialized or the file glob operation fails.
func (s *Service) GetDatabaseFileList() ([]string, error) {
	const op errors.Op = "facade.Service.GetDatabaseFileList"

	if !s.initialized.Load() {
		err := errors.New(op).Msg(errMsgServiceNotInit)
		s.LoggerService.ErrorWith().Err(err).Msg(errMsgServiceNotInit)
		return nil, err
	}

	dbDir := filepath.Dir(s.DatabaseService.DatabaseConfig.Path)
	entries, err := os.ReadDir(dbDir)
	if err != nil {
		err = errors.New(op).Err(err)
		s.LoggerService.ErrorWith().Err(err).Msg("Failed to read database directory.")
		return nil, err
	}

	var slice []string
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		matched, merr := filepath.Match("*.db", entry.Name())
		if merr != nil {
			merr = errors.New(op).Err(merr)
			s.LoggerService.ErrorWith().Err(merr).Msg("Failed to match database file.")
			return nil, merr
		}
		if matched {
			slice = append(slice, entry.Name())
		}
	}

	return slice, nil
}
