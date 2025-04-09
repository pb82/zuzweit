package scenes

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/joelschutz/stagehand"
	"log"
	"time"
	"zuzweit/api"
	"zuzweit/ecs"
)

type BaseScene struct {
	entityManager *ecs.EntityManager
	context       *api.GameContext
	state         api.GameState
	milliseconds  int64
	keys          []ebiten.Key

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

func (s *BaseScene) Update() error {
	milliseconds := time.Now().UnixMilli()
	delta := milliseconds - s.milliseconds
	s.milliseconds = milliseconds

	s.keys = inpututil.AppendPressedKeys(s.keys[:0])
	if len(s.keys) > 0 {
		log.Println(s.keys)
	}

	s.entityManager.Collect()
	s.entityManager.Update(float64(delta), s.context)

	return nil
}

func NewBaseScene(em *ecs.EntityManager, context *api.GameContext) BaseScene {
	return BaseScene{
		entityManager: em,
		context:       context,
		milliseconds:  time.Now().UnixMilli(),
	}
}
