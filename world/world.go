package world

import (
	"github.com/donutmonger/game_engine/texture"

	"github.com/donutmonger/game_engine/components"
	"math/rand"
	"time"
)

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

type PlayerStatsComponent struct {
	Speed float32
}

type ParticleEmitterComponent struct {
	Continuous        bool
	MaxFiresPerSecond float32
	ParticleTexture   texture.Texture
	LastFireTime      time.Time
	FireRequested     bool
}

type ParticleComponent struct {
	DestroyWhenOffScreen bool
}

type World struct {
	MaxEntities               uint64
	PositionComponents        []PositionComponent
	VelocityComponents        []VelocityComponent
	ScaleComponents           []ScaleComponent
	SpriteComponents          []SpriteComponent
	PlayerStatsComponents     []PlayerStatsComponent
	ParticleEmitterComponents []ParticleEmitterComponent
	ParticleComponents        []ParticleComponent

	mask []uint64
}

func NewWorld(maxEntities uint64) World {
	return World{
		MaxEntities:               maxEntities,
		PositionComponents:        make([]PositionComponent, maxEntities),
		VelocityComponents:        make([]VelocityComponent, maxEntities),
		ScaleComponents:           make([]ScaleComponent, maxEntities),
		SpriteComponents:          make([]SpriteComponent, maxEntities),
		PlayerStatsComponents:     make([]PlayerStatsComponent, maxEntities),
		ParticleEmitterComponents: make([]ParticleEmitterComponent, maxEntities),
		ParticleComponents:        make([]ParticleComponent, maxEntities),
		mask:                      make([]uint64, maxEntities),
	}
}

func (w World) CreateEntity() uint64 {
	for entity := uint64(0); entity < w.MaxEntities; entity++ {
		if w.mask[entity] == components.COMPONENT_NONE {
			return entity
		}
	}
	return 0
}

func (w *World) DestroyEntity(entity uint64) {
	w.mask[entity] = components.COMPONENT_NONE
}

func (w *World) EntitySatisfiesMask(entity uint64, mask uint64) bool {
	return (w.mask[entity] & mask) == mask
}

func (w *World) CreateLaser(position PositionComponent, texture texture.Texture) uint64 {
	entity := w.CreateEntity()
	w.mask[entity] = components.COMPONENT_POSITION | components.COMPONENT_SCALE | components.COMPONENT_VELOCITY | components.COMPONENT_SPRITE | components.COMPONENT_PARTICLE
	w.PositionComponents[entity] = position
	w.VelocityComponents[entity] = VelocityComponent{X: 0, Y: 2}
	w.ScaleComponents[entity] = ScaleComponent{X: 0.0025, Y: 0.0075}
	w.SpriteComponents[entity] = SpriteComponent{Texture: texture}
	w.ParticleComponents[entity] = ParticleComponent{DestroyWhenOffScreen: true}

	return entity
}

func (w *World) CreatePlayerSpaceship(t *texture.Texture, laserTexture *texture.Texture) uint64 {
	entity := w.CreateEntity()
	w.mask[entity] = components.COMPONENT_POSITION | components.COMPONENT_VELOCITY | components.COMPONENT_SCALE | components.COMPONENT_SPRITE | components.COMPONENT_PLAYER | components.COMPONENT_PARTICLE_EMITTER
	w.PositionComponents[entity] = PositionComponent{X: 0, Y: -0.9}
	w.VelocityComponents[entity] = VelocityComponent{X: 0, Y: 0}
	w.ScaleComponents[entity] = ScaleComponent{X: 0.05, Y: 0.0525}
	w.SpriteComponents[entity] = SpriteComponent{Texture: *t}
	w.PlayerStatsComponents[entity] = PlayerStatsComponent{Speed: 1}
	w.ParticleEmitterComponents[entity] = ParticleEmitterComponent{Continuous: false, MaxFiresPerSecond: 5, ParticleTexture: *laserTexture, LastFireTime: time.Now(), FireRequested: false}

	return entity
}

func (w *World) CreateEnemyFighter(t *texture.Texture) uint64 {
	entity := w.CreateEntity()
	w.mask[entity] = components.COMPONENT_POSITION | components.COMPONENT_VELOCITY | components.COMPONENT_SCALE | components.COMPONENT_SPRITE
	w.PositionComponents[entity] = PositionComponent{X: getRandomFloatInGLScreenSpace(), Y: getRandomFloatInGLScreenSpace()}
	w.VelocityComponents[entity] = VelocityComponent{X: 0, Y: 0}
	w.ScaleComponents[entity] = ScaleComponent{X: 0.0375, Y: 0.05}
	w.SpriteComponents[entity] = SpriteComponent{Texture: *t}

	return entity
}

func getRandomFloatInGLScreenSpace() float32 {
	return (rand.Float32() * 2) - 1.0
}
