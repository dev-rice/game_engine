package systems

import "github.com/donutmonger/game_engine/world"

type System interface {
	Update(world *world.World)
}
