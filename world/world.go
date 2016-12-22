package world

import (
	"github.com/donutmonger/game_engine/texture"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/donutmonger/game_engine/shader"
	"github.com/go-gl/gl/v4.1-core/gl"

	"math/rand"
	"github.com/go-gl/glfw/v3.2/glfw"

)

var COMPONENT_NONE uint64 = 0
var COMPONENT_POSITION uint64 = 1 << 0
var COMPONENT_SCALE uint64 = 1 << 1
var COMPONENT_SPRITE uint64 = 1 << 2
var COMPONENT_INPUT uint64 = 1 << 3

type PositionComponent struct {
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

type InputComponent struct {
	OnPress func()
}

type World struct {
	mask []uint64
	positionComponents []PositionComponent
	scaleComponents []ScaleComponent
	spriteComponents []SpriteComponent
	inputComponents []InputComponent

	maxEntities uint64
}

func NewWorld(maxEntities uint64) World {
	return World {
		mask: make([]uint64, maxEntities),
		positionComponents: make([]PositionComponent, maxEntities),
		scaleComponents: make([]ScaleComponent, maxEntities),
		spriteComponents: make([]SpriteComponent, maxEntities),
		inputComponents: make([]InputComponent, maxEntities),
		maxEntities: maxEntities,
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

// The Draw System
// A system is just a function and a mask on which it operates
func (w World) Draw(shader *shader.ShaderProgram) {
	drawableMask := COMPONENT_POSITION | COMPONENT_SCALE | COMPONENT_SPRITE
	for entity := uint64(0); entity < w.maxEntities; entity++ {
		if (w.mask[entity] & drawableMask) == drawableMask {
			position := w.positionComponents[entity]
			scale := w.scaleComponents[entity]

			model := mgl32.Mat3FromCols(
				mgl32.Vec3{scale.X, 0      , position.X},
				mgl32.Vec3{0      , scale.Y, position.Y},
				mgl32.Vec3{0      , 0      , 1},
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

func (w *World) InputSystem(window *glfw.Window) {
	inputMask := COMPONENT_INPUT
	for entity := uint64(0); entity < w.maxEntities; entity++ {
		if (w.mask[entity] & inputMask) == inputMask {
			if (window.GetKey(glfw.KeyB) == glfw.Press) {
				w.inputComponents[entity].OnPress()
			}
		}
	}
}

func (w *World) CreateBarracks(t *texture.Texture) uint64 {
	barracks := w.CreateEntity()
	w.mask[barracks] = COMPONENT_INPUT
	w.inputComponents[barracks] = InputComponent{OnPress: func() {
		w.CreateEnemy(t)
	}}

	return barracks
}

func (w *World) CreateEnemy(t *texture.Texture) uint64 {
	enemy := w.CreateEntity()
	w.mask[enemy] = COMPONENT_POSITION | COMPONENT_SCALE | COMPONENT_SPRITE
	w.positionComponents[enemy] = PositionComponent{randomFloatInGLScreenSpace(), randomFloatInGLScreenSpace()}
	w.scaleComponents[enemy] = ScaleComponent{0.1, 0.15}
	w.spriteComponents[enemy] = SpriteComponent{*t}

	return enemy
}


func randomFloatBetween(min float32, max float32) float32 {
	return rand.Float32()
}

func randomFloatInGLScreenSpace() float32 {
	return (rand.Float32() * 2) - 1
}
