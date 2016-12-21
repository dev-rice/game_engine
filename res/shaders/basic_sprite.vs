#version 330

in vec2 position;
in vec2 texcoord;

out vec2 Texcoord;

uniform mat3 transformation;

void main() {
    mat3 transform = mat3(0.5, 0.0 , 0.2,
                          0.0, 0.75, 0.5,
                          0.0, 0.0 , 1.0);


    Texcoord = vec2(texcoord.x, 1 - texcoord.y);
    vec3 position_temp = vec3(position, 1.0) * transformation;
    gl_Position = vec4(position_temp.xy, 0.0, 1.0);
}