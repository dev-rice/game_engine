package shader

import (
	"fmt"
	"github.com/go-gl/gl/v4.1-core/gl"
	"strings"
)

type ShaderProgram struct {
	GLid uint32
}

func NewShaderProgram(vertexShader *VertexShader, fragmentShader *FragmentShader) (*ShaderProgram, error) {

	program_id := gl.CreateProgram()
	gl.AttachShader(program_id, vertexShader.GLid)
	gl.AttachShader(program_id, fragmentShader.GLid)
	gl.LinkProgram(program_id)

	var status int32
	gl.GetProgramiv(program_id, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(program_id, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetProgramInfoLog(program_id, logLength, nil, gl.Str(log))

		return nil, fmt.Errorf("failed to link program: %v", log)
	}

	gl.DeleteShader(vertexShader.GLid)
	gl.DeleteShader(fragmentShader.GLid)

	shaderProgram := &ShaderProgram{
		GLid: program_id,
	}
	return shaderProgram, nil
}

func (s ShaderProgram) Use() {
	gl.UseProgram(s.GLid)
}

func (s ShaderProgram) GetUniformLocation(uniform_name string) int32 {
	return gl.GetUniformLocation(s.GLid, gl.Str(uniform_name+"\x00"))
}

func (s ShaderProgram) GetAttribLocation(attrib_name string) int32 {
	return gl.GetAttribLocation(s.GLid, gl.Str(attrib_name+"\x00"))
}

type VertexShader struct {
	GLid uint32
}

func NewVertexShader(source string) (*VertexShader, error) {
	id, err := compileShader(source, gl.VERTEX_SHADER)
	if err != nil {
		return nil, err
	}

	return &VertexShader{GLid: id}, nil
}

type FragmentShader struct {
	GLid uint32
}

func NewFragmentShader(source string) (*FragmentShader, error) {
	id, err := compileShader(source, gl.FRAGMENT_SHADER)
	if err != nil {
		return nil, err
	}

	return &FragmentShader{GLid: id}, nil
}

func compileShader(source string, shaderType uint32) (uint32, error) {
	shader := gl.CreateShader(shaderType)

	csources, free := gl.Strs(source)
	gl.ShaderSource(shader, 1, csources, nil)
	free()
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))

		return 0, fmt.Errorf("failed to compile %v: %v", source, log)
	}

	return shader, nil
}
