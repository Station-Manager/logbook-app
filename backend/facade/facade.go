package facade

import (
	"github.com/Station-Manager/database/sqlite/meta"
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

func (s *Service) GetDatabaseMetadata() ([]meta.SqliteMeta, error) {
	const op errors.Op = "facade.Service.GetDatabaseFileList"

	if !s.initialized.Load() {
		err := errors.New(op).Msg(errMsgServiceNotInit)
		s.LoggerService.ErrorWith().Err(err).Msg(errMsgServiceNotInit)
		return nil, err
	}

	//dbDir := filepath.Dir(s.DatabaseService.DatabaseConfig.Path)
	//
	//metaList, err := meta.LoadSqliteDatabaseList(dbDir)
	//if err != nil {
	//	return nil, err
	//}

	//	fmt.Printf("Loaded metadata: %+v\n", metaList)

	//entries, err := os.ReadDir(dbDir)
	//if err != nil {
	//	err = errors.New(op).Err(err)
	//	s.LoggerService.ErrorWith().Err(err).Msg("Failed to read database directory.")
	//	return nil, err
	//}
	//
	//var slice []string
	//for _, entry := range entries {
	//	if entry.IsDir() {
	//		continue
	//	}
	//	matched, merr := filepath.Match("*.db", entry.Name())
	//	if merr != nil {
	//		merr = errors.New(op).Err(merr)
	//		s.LoggerService.ErrorWith().Err(merr).Msg("Failed to match database file.")
	//		return nil, merr
	//	}
	//	if matched {
	//		slice = append(slice, strings.ReplaceAll(entry.Name(), ".db", ""))
	//	}
	//}

	return nil, nil
}

func (s *Service) GetLogbookList() ([]types.Logbook, error) {
	const op errors.Op = "facade.Service.GetLogbookList"

	if !s.initialized.Load() {
		err := errors.New(op).Msg(errMsgServiceNotInit)
		s.LoggerService.ErrorWith().Err(err).Msg(errMsgServiceNotInit)
		return nil, err
	}

	slice, err := s.DatabaseService.FetchAllLogbooks()
	if err != nil {
		err = errors.New(op).Err(err)
		s.LoggerService.ErrorWith().Err(err).Msg("Failed to fetch logbook list.")
		return nil, err
	}

	return slice, nil
}

func (s *Service) NewLogbook(logbook types.Logbook) error {
	const op errors.Op = "facade.Service.NewLogbook"

	if !s.initialized.Load() {
		err := errors.New(op).Msg(errMsgServiceNotInit)
		s.LoggerService.ErrorWith().Err(err).Msg(errMsgServiceNotInit)
		return err
	}

	id, err := s.DatabaseService.InsertLogbook(logbook)
	if err != nil {
		err = errors.New(op).Err(err)
		s.LoggerService.ErrorWith().Err(err).Msgf("Failed to insert logbook: %+v", logbook)
		return err
	}

	logbook.ID = id

	return nil
}

func (s *Service) UpdateLogbook(logbook types.Logbook) error {
	return nil
}

func (s *Service) DeleteLogbook(id int64) error {
	const op errors.Op = "facade.Service.DeleteLogbook"

	if !s.initialized.Load() {
		err := errors.New(op).Msg(errMsgServiceNotInit)
		s.LoggerService.ErrorWith().Err(err).Msg(errMsgServiceNotInit)
		return err
	}

	logbook, err := s.DatabaseService.FetchLogbookByID(id)
	if err != nil {
		err = errors.New(op).Err(err)
		s.LoggerService.ErrorWith().Err(err).Msgf("Failed to fetch logbook with id: %d", id)
		return err
	}

	qsoSlice, err := s.DatabaseService.FetchQsoSliceByLogbookId(logbook.ID)
	if err != nil {
		err = errors.New(op).Err(err)
		s.LoggerService.ErrorWith().Err(err).Msgf("Failed to fetch QSO count for logbook with id: %d", logbook.ID)
		return err
	}

	if len(qsoSlice) > 0 {
		return errors.New(op).Msgf("Logbook with id: %d has QSOs. Cannot delete.", logbook.ID)
	}

	return nil
}
