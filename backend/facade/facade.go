package facade

import (
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

	//requiredCfg, err := s.ConfigService.RequiredConfigs()
	//if err != nil {
	//	err = errors.New(op).Err(err)
	//	s.LoggerService.ErrorWith().Err(err).Msg("Failed to fetch required configs.")
	//	return nil, err
	//}
	//
	//loggingStation, err := s.ConfigService.LoggingStationConfigs()
	//if err != nil {
	//	err = errors.New(op).Err(err)
	//	s.LoggerService.ErrorWith().Err(err).Msg("Failed to fetch logging station configs.")
	//	return nil, err
	//}

	return &types.UiConfig{
		Logbook: s.currentLogbook,
	}, nil
}
