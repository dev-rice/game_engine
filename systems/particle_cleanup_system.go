package systems

import (
	"github.com/donutmonger/game_engine/components"
	"github.com/donutmonger/game_engine/world"
)

type ParticleCleanupSystem struct {
}

func NewParticleCleanupSystem() *ParticleCleanupSystem {
	return &ParticleCleanupSystem{}
}

func (p *ParticleCleanupSystem) Update(w *world.World) {
	mask := components.COMPONENT_PARTICLE | components.COMPONENT_POSITION
	for entity := uint64(0); entity < w.MaxEntities; entity++ {
		if w.EntitySatisfiesMask(entity, mask) {
			if w.ParticleComponents[entity].DestroyWhenOffScreen {
				position := w.PositionComponents[entity]
				if isOffScreen(position.X, position.Y) {
					w.DestroyEntity(entity)
				}
			}
		}
	}
}

func isOffScreen(x float32, y float32) bool {
	return x > 1.0 || x < -1.0 || y > 1.0 || y < -1.0
}
