package systems

import (
	"github.com/donutmonger/game_engine/components"
	"github.com/donutmonger/game_engine/world"
	"time"
)

type ParticleEmitterSystem struct {
}

func NewParticleEmitterSystem() *ParticleEmitterSystem {
	return &ParticleEmitterSystem{}
}

func (p *ParticleEmitterSystem) Update(w *world.World) {
	mask := components.COMPONENT_POSITION | components.COMPONENT_PARTICLE_EMITTER
	for entity := uint64(0); entity < w.MaxEntities; entity++ {
		if w.EntitySatisfiesMask(entity, mask) {
			p := &w.ParticleEmitterComponents[entity]
			if p.Continuous || p.FireRequested {
				minTime := 1.0 / p.MaxFiresPerSecond
				timeSinceLast := float32(time.Since(p.LastFireTime).Seconds())
				if timeSinceLast >= minTime {
					w.CreateLaser(w.PositionComponents[entity], p.ParticleTexture)
					p.LastFireTime = time.Now()
				}
			}
		}
	}
}
