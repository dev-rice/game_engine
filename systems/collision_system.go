package systems

import (
	"fmt"
	"github.com/donutmonger/game_engine/components"
	"github.com/donutmonger/game_engine/world"
	"github.com/go-gl/mathgl/mgl32"
	"time"
)

type CollisionSystem struct {
}

func NewCollisionSystem() *CollisionSystem {
	return &CollisionSystem{}
}

func (c *CollisionSystem) Update(w *world.World) {
	mask := components.COMPONENT_POSITION | components.COMPONENT_SCALE | components.COMPONENT_COLLISION
	collisionEntities := make([]uint64, 0)
	for entity := uint64(0); entity < w.MaxEntities; entity++ {
		if w.EntitySatisfiesMask(entity, mask) {
			collisionEntities = append(collisionEntities, entity)
		}
	}

	for i := 0; i < len(collisionEntities); i++ {
		for j := 0; j < len(collisionEntities); j++ {
			if isColliding(w, uint64(i), uint64(j)) {
				fmt.Printf("%s Collision occurred between %d and %d\n", time.Now().Format(time.ANSIC), i, j)
			}
		}
	}

}

func isColliding(w *world.World, a uint64, b uint64) bool {
	if a == b {
		return false
	}

	aPositionComponent := w.PositionComponents[a]
	bPositionComponent := w.PositionComponents[b]

	aPosition := mgl32.Vec2{aPositionComponent.X, aPositionComponent.Y}
	bPosition := mgl32.Vec2{bPositionComponent.X, bPositionComponent.Y}

	distance := aPosition.Sub(bPosition).Len()

	return distance <= 0.1
}
