package systems

import (
	"github.com/donutmonger/game_engine/components"
	"github.com/donutmonger/game_engine/world"
	"time"
)

type MovementSystem struct {
	lastUpdateTime time.Time
}

func NewMovementSystem() *MovementSystem {
	return &MovementSystem{
		lastUpdateTime: time.Now(),
	}
}

func (m *MovementSystem) Update(w *world.World) {
	dt := float32(time.Since(m.lastUpdateTime).Seconds())
	physicsMask := components.COMPONENT_POSITION | components.COMPONENT_VELOCITY
	for entity := uint64(0); entity < w.MaxEntities; entity++ {
		if w.EntitySatisfiesMask(entity, physicsMask) {
			velocity := w.VelocityComponents[entity]
			dx := velocity.X * dt
			dy := velocity.Y * dt

			w.PositionComponents[entity].X += dx
			w.PositionComponents[entity].Y += dy
		}
	}

	m.lastUpdateTime = time.Now()
}
