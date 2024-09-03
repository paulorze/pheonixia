package engine

import "phoenixia/internal/ports"

var _ ports.EngineService = &Service{}

type Service struct {
	Repository ports.EngineRepository
}
