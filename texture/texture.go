package texture

import (
	"fmt"
	"github.com/go-gl/gl/v4.1-core/gl"
	"image"
	"image/draw"
	_ "image/png"
	"os"
)

type Texture struct {
	GLid uint32
}

func NewTextureFromFile(filename string) (*Texture, error) {
	imgFile, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("texture file %q not found on disk: %v", filename, err)
	}
	img, _, err := image.Decode(imgFile)
	if err != nil {
		return nil, err
	}

	rgba := image.NewRGBA(img.Bounds())
	if rgba.Stride != rgba.Rect.Size().X*4 {
		return nil, fmt.Errorf("unsupported stride")
	}
	draw.Draw(rgba, rgba.Bounds(), img, image.Point{0, 0}, draw.Src)

	var texture_id uint32
	gl.GenTextures(1, &texture_id)
	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, texture_id)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.NEAREST)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
	gl.TexImage2D(
		gl.TEXTURE_2D,
		0,
		gl.RGBA,
		int32(rgba.Rect.Size().X),
		int32(rgba.Rect.Size().Y),
		0,
		gl.RGBA,
		gl.UNSIGNED_BYTE,
		gl.Ptr(rgba.Pix))

	return &Texture{GLid: texture_id}, nil
}

func (t Texture) Bind2D() {
	gl.BindTexture(gl.TEXTURE_2D, t.GLid)
}
