#version 330

in vec2 Texcoord;

out vec4 outputColor;

uniform sampler2D base_texture;

void main() {
    vec4 texel = textureOffset(base_texture, Texcoord, ivec2(-0.5 , -0.5));

    outputColor = vec4(texel.rgb * vec3(0, 1, 0), texel.a);
}