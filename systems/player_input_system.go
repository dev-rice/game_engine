package systems

import (
	"github.com/donutmonger/game_engine/components"
	"github.com/donutmonger/game_engine/world"
	"github.com/go-gl/glfw/v3.2/glfw"
)

type PlayerInputSystem struct {
	window *glfw.Window
}

func NewPlayerInputSystem(w *glfw.Window) *PlayerInputSystem {
	return &PlayerInputSystem{
		window: w,
	}
}

func (p *PlayerInputSystem) Update(w *world.World) {
	playerMask := components.COMPONENT_VELOCITY | components.COMPONENT_PLAYER | components.COMPONENT_PARTICLE_EMITTER
	for entity := uint64(0); entity < w.MaxEntities; entity++ {
		if w.EntitySatisfiesMask(entity, playerMask) {
			speed := w.PlayerStatsComponents[entity].Speed

			w.VelocityComponents[entity].X = 0
			w.VelocityComponents[entity].Y = 0

			if p.window.GetKey(glfw.KeyA) == glfw.Press {
				w.VelocityComponents[entity].X = -speed
			}
			if p.window.GetKey(glfw.KeyD) == glfw.Press {
				w.VelocityComponents[entity].X = speed
			}
			if p.window.GetKey(glfw.KeyW) == glfw.Press {
				w.VelocityComponents[entity].Y = speed
			}
			if p.window.GetKey(glfw.KeyS) == glfw.Press {
				w.VelocityComponents[entity].Y = -speed
			}

			if p.window.GetKey(glfw.KeySpace) == glfw.Press {
				w.ParticleEmitterComponents[entity].FireRequested = true
			} else {
				w.ParticleEmitterComponents[entity].FireRequested = false
			}
		}
	}
}
