package world

import (
	"github.com/donutmonger/game_engine/shader"
	"github.com/donutmonger/game_engine/texture"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"

	"github.com/go-gl/glfw/v3.2/glfw"
)

var COMPONENT_NONE uint64 = 0
var COMPONENT_POSITION uint64 = 1 << 0
var COMPONENT_VELOCITY uint64 = 1 << 1
var COMPONENT_SCALE uint64 = 1 << 2
var COMPONENT_SPRITE uint64 = 1 << 3
var COMPONENT_PLAYER uint64 = 1 << 4

type PositionComponent struct {
	X float32
	Y float32
}

type VelocityComponent struct {
	X float32
	Y float32
}

type ScaleComponent struct {
	X float32
	Y float32
}

type SpriteComponent struct {
	Texture texture.Texture
}

type PlayerComponent struct {
	Speed float32
}

type World struct {
	mask               []uint64
	positionComponents []PositionComponent
	velocityComponents []VelocityComponent
	scaleComponents    []ScaleComponent
	spriteComponents   []SpriteComponent
	playerComponents   []PlayerComponent

	maxEntities uint64
}

func NewWorld(maxEntities uint64) World {
	return World{
		mask:               make([]uint64, maxEntities),
		positionComponents: make([]PositionComponent, maxEntities),
		velocityComponents: make([]VelocityComponent, maxEntities),
		scaleComponents:    make([]ScaleComponent, maxEntities),
		spriteComponents:   make([]SpriteComponent, maxEntities),
		playerComponents:   make([]PlayerComponent, maxEntities),
		maxEntities:        maxEntities,
	}
}

func (w World) CreateEntity() uint64 {
	for entity := uint64(0); entity < w.maxEntities; entity++ {
		if w.mask[entity] == COMPONENT_NONE {
			return entity
		}
	}

	return 0
}

func (w *World) DestroyEntity(entity uint64) {
	w.mask[entity] = COMPONENT_NONE
}

func (w *World) entitySatisfiesMask(entity uint64, mask uint64) bool {
	return (w.mask[entity] & mask) == mask
}

// The Draw System
// A system is just a function and a mask on which it operates
func (w World) Draw(shader *shader.ShaderProgram) {
	drawableMask := COMPONENT_POSITION | COMPONENT_SCALE | COMPONENT_SPRITE
	for entity := uint64(0); entity < w.maxEntities; entity++ {
		if w.entitySatisfiesMask(entity, drawableMask) {
			position := w.positionComponents[entity]
			scale := w.scaleComponents[entity]

			model := mgl32.Mat3FromCols(
				mgl32.Vec3{scale.X, 0, position.X},
				mgl32.Vec3{0, scale.Y, position.Y},
				mgl32.Vec3{0, 0, 1},
			)

			modelUniform := shader.GetUniformLocation("transformation")
			gl.UniformMatrix3fv(modelUniform, 1, false, &model[0])

			gl.ActiveTexture(gl.TEXTURE0)
			sprite := w.spriteComponents[entity]
			sprite.Texture.Bind2D()

			gl.DrawArrays(gl.TRIANGLES, 0, 6)
		}
	}
}

func (w *World) PhysicsSystem(dt float32) {
	physicsMask := COMPONENT_POSITION | COMPONENT_VELOCITY
	for entity := uint64(0); entity < w.maxEntities; entity++ {
		if w.entitySatisfiesMask(entity, physicsMask) {
			velocity := w.velocityComponents[entity]
			dx := velocity.X * dt
			dy := velocity.Y * dt

			w.positionComponents[entity].X += dx
			w.positionComponents[entity].Y += dy
		}
	}
}

func (w *World) PlayerInputSystem(window *glfw.Window) {
	playerMask := COMPONENT_VELOCITY | COMPONENT_PLAYER
	for entity := uint64(0); entity < w.maxEntities; entity++ {
		if w.entitySatisfiesMask(entity, playerMask) {
			speed := w.playerComponents[entity].Speed

			w.velocityComponents[entity].X = 0
			w.velocityComponents[entity].Y = 0

			if window.GetKey(glfw.KeyA) == glfw.Press {
				w.velocityComponents[entity].X = -speed
			}
			if window.GetKey(glfw.KeyD) == glfw.Press {
				w.velocityComponents[entity].X = speed
			}
			if window.GetKey(glfw.KeyW) == glfw.Press {
				w.velocityComponents[entity].Y = speed
			}
			if window.GetKey(glfw.KeyS) == glfw.Press {
				w.velocityComponents[entity].Y = -speed
			}
			if window.GetKey(glfw.KeySpace) == glfw.Press {
				// Fire the guns
				w.CreateBullet(w.positionComponents[entity], w.spriteComponents[entity])
			}
		}
	}
}

func (w *World) CreateBullet(position PositionComponent, sprite SpriteComponent) uint64 {
	entity := w.CreateEntity()
	w.mask[entity] = COMPONENT_POSITION | COMPONENT_SCALE | COMPONENT_VELOCITY | COMPONENT_SPRITE
	w.positionComponents[entity] = position
	w.velocityComponents[entity] = VelocityComponent{X: 0, Y: 2}
	w.scaleComponents[entity] = ScaleComponent{X: 0.005, Y: 0.015}
	w.spriteComponents[entity] = sprite

	return entity
}

func (w *World) CreatePlayerSpaceship(t *texture.Texture) uint64 {
	entity := w.CreateEntity()
	w.mask[entity] = COMPONENT_POSITION | COMPONENT_VELOCITY | COMPONENT_SCALE | COMPONENT_SPRITE | COMPONENT_PLAYER
	w.positionComponents[entity] = PositionComponent{X: 0, Y: -0.9}
	w.velocityComponents[entity] = VelocityComponent{X: 0, Y: 0}
	w.scaleComponents[entity] = ScaleComponent{X: 0.05, Y: 0.0525}
	w.spriteComponents[entity] = SpriteComponent{Texture: *t}
	w.playerComponents[entity] = PlayerComponent{Speed: 1}

	return entity
}
