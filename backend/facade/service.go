package facade

import (
	"github.com/Station-Manager/config"
	"github.com/Station-Manager/database/sqlite"
	"github.com/Station-Manager/errors"
	"github.com/Station-Manager/logging"
	"github.com/Station-Manager/types"

	"sync"
	"sync/atomic"
)

const (
	ServiceName = "logbook-app-facade"
)

type Service struct {
	ConfigService   *config.Service  `di.inject:"configservice"`
	LoggerService   *logging.Service `di.inject:"loggingservice"`
	DatabaseService *sqlite.Service  `di.inject:"sqliteservice"`

	initialized atomic.Bool
	started     atomic.Bool // guarded via atomic operations; Start/Stop also hold mu for broader state

	initOnce sync.Once
	mu       sync.Mutex

	currentLogbook types.Logbook
	requiredCfgs   *types.RequiredConfigs
}

func (s *Service) Initialize() error {
	const op errors.Op = "facade.Service.Initialize"

	var initErr error
	s.initOnce.Do(func() {
		if s.ConfigService == nil {
			initErr = errors.New(op).Msg(errMsgNilConfigService)
			return
		}

		if s.LoggerService == nil {
			initErr = errors.New(op).Msg(errMsgNilLoggerService)
			return
		}

		if s.DatabaseService == nil {
			initErr = errors.New(op).Msg(errMsgNilDatabaseService)
			return
		}

		reqCfg, err := s.ConfigService.RequiredConfigs()
		if err != nil {
			initErr = errors.New(op).Err(err)
			return
		}
		s.requiredCfgs = &reqCfg

		if err = s.openAndLoadFromDatabase(); err != nil {
			initErr = errors.New(op).Err(err)
			return
		}

		s.initialized.Store(true)
	})

	return initErr
}
