package window

import "github.com/go-gl/glfw/v3.2/glfw"

// Window that holds the opengl context
type Window struct {
	height     int
	width      int
	GlfwWindow *glfw.Window
}

// NewWindow creates a new window
func NewWindow(width int, height int) Window {
	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	glfwWindow, err := glfw.CreateWindow(width, height, "Space Fight", nil, nil)
	if err != nil {
		panic(err)
	}

	window := Window{
		height:     height,
		width:      width,
		GlfwWindow: glfwWindow,
	}

	return window

}

func (w *Window) AspectRatio() float32 {
	return float32(w.width) / float32(w.height)
}
