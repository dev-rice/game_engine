package systems

import (
	"github.com/donutmonger/game_engine/components"
	"github.com/donutmonger/game_engine/shader"
	"github.com/donutmonger/game_engine/world"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

type RenderSystem struct {
	shader shader.ShaderProgram
	vao    uint32
}

var flatMeshVertices = []float32{
	-1.0, 1.0, 0.0, 1.0,
	1.0, 1.0, 1.0, 1.0,
	1.0, -1.0, 1.0, 0.0,

	1.0, -1.0, 1.0, 0.0,
	-1.0, -1.0, 0.0, 0.0,
	-1.0, 1.0, 0.0, 1.0,
}

func NewRenderSystem(s shader.ShaderProgram) *RenderSystem {
	// Configure the vertex data
	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)

	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(flatMeshVertices)*4, gl.Ptr(flatMeshVertices), gl.STATIC_DRAW)

	posAttrib := uint32(s.GetAttribLocation("position"))
	gl.EnableVertexAttribArray(posAttrib)
	gl.VertexAttribPointer(posAttrib, 2, gl.FLOAT, false, 4*4, gl.PtrOffset(0))

	texcoordAttrib := uint32(s.GetAttribLocation("texcoord"))
	gl.EnableVertexAttribArray(texcoordAttrib)
	gl.VertexAttribPointer(texcoordAttrib, 2, gl.FLOAT, false, 4*4, gl.PtrOffset(2*4))

	return &RenderSystem{
		shader: s,
		vao:    vao,
	}
}

// The Draw System
// A system is just a function and a mask on which it operates
func (r *RenderSystem) Update(w *world.World) {
	// Render
	r.shader.Use()
	gl.BindVertexArray(r.vao)

	drawableMask := components.COMPONENT_POSITION | components.COMPONENT_SCALE | components.COMPONENT_SPRITE
	for entity := uint64(0); entity < w.MaxEntities; entity++ {
		if w.EntitySatisfiesMask(entity, drawableMask) {
			position := w.PositionComponents[entity]
			scale := w.ScaleComponents[entity]

			model := mgl32.Mat3FromCols(
				mgl32.Vec3{scale.X, 0, position.X},
				mgl32.Vec3{0, scale.Y, position.Y},
				mgl32.Vec3{0, 0, 1},
			)

			modelUniform := r.shader.GetUniformLocation("transformation")
			gl.UniformMatrix3fv(modelUniform, 1, false, &model[0])

			gl.ActiveTexture(gl.TEXTURE0)
			sprite := w.SpriteComponents[entity]
			sprite.Texture.Bind2D()

			gl.DrawArrays(gl.TRIANGLES, 0, 6)
		}
	}
}
