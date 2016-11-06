package sprite

import (
	"fmt"
	"github.com/donutmonger/game_engine/texture"
)

type Sprite struct {
	texture texture.Texture
}

func NewSprite(tex texture.Texture) Sprite {
	return Sprite{
		texture: tex,
	}
}

func (sprite Sprite) Draw() {

}

// A sprite animation holds several textures that can be animated
type SpriteAnimation struct {
	textures          []texture.Texture
	timeBetweenFrames float64
	currentFrame      int
	timeSinceLast     float64
}

func NewSpriteAnimation(textures []texture.Texture, timeBetweenFrames float64) SpriteAnimation {

	return SpriteAnimation{
		textures:          textures,
		timeBetweenFrames: timeBetweenFrames,
		currentFrame:      0,
		timeSinceLast:     0,
	}
}

func (sa *SpriteAnimation) Animate(elapsed float64) {
	sa.timeSinceLast += elapsed

	if sa.timeSinceLast >= sa.timeBetweenFrames {
		sa.timeSinceLast = 0

		sa.currentFrame += 1
		if sa.currentFrame >= len(sa.textures) {
			sa.currentFrame = 0
		}
		fmt.Printf("Current Frame: %d\n", sa.currentFrame)

		frame := sa.textures[sa.currentFrame]
		fmt.Printf("Frame with ID: %v\n", frame.GLid)
		sa.textures[sa.currentFrame].Bind2D()
	}

}

func (sa *SpriteAnimation) Nothing() {

}
