package unit

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/donutmonger/game_engine/texture"
	"github.com/go-gl/gl/v4.1-core/gl"

	"github.com/donutmonger/game_engine/shader"
)

type Unit struct {
	Position mgl32.Vec2
	Scale mgl32.Vec2
	Sprite *texture.Texture
	ShaderProgram *shader.ShaderProgram
}

func (u Unit) Draw() {
	model2 := mgl32.Mat3FromCols(
		mgl32.Vec3{u.Scale.X(), 0          , u.Position.X()},
		mgl32.Vec3{0          , u.Scale.Y(), u.Position.Y()},
		mgl32.Vec3{0          , 0          , 1},
	)

	modelUniform := u.ShaderProgram.GetUniformLocation("transformation")
	gl.UniformMatrix3fv(modelUniform, 1, false, &model2[0])

	gl.ActiveTexture(gl.TEXTURE0)
	u.Sprite.Bind2D()

	gl.DrawArrays(gl.TRIANGLES, 0, 6)
}