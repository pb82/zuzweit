package scenes

import (
	"github.com/joelschutz/stagehand"
	"time"
	"zuzweit/api"
	"zuzweit/ecs"
)

type BaseScene struct {
	entityManager *ecs.EntityManager
	context       *api.GameContext
	state         api.GameState
	milliseconds  int64

	sm *stagehand.SceneManager[api.GameState]
}

func (s *BaseScene) Load(state api.GameState, sm stagehand.SceneController[api.GameState]) {
	s.sm = sm.(*stagehand.SceneManager[api.GameState])
	s.state = state
}

func (s *BaseScene) Unload() api.GameState {
	return s.state
}

func (s *BaseScene) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}

func NewBaseScene(em *ecs.EntityManager, context *api.GameContext) BaseScene {
	return BaseScene{
		entityManager: em,
		context:       context,
		milliseconds:  time.Now().UnixMilli(),
	}
}
