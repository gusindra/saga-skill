components {
  id: "explosion"
  component: "/examples/ocean-commotion/assets/explode.particlefx"
}
components {
  id: "script"
  component: "/examples/ocean-commotion/fish.script"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"fish-blue\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/examples/ocean-commotion/assets/sprites.atlas\"\n"
  "}\n"
  ""
}
embedded_components {
  id: "sprite-eyes"
  type: "sprite"
  data: "default_animation: \"eyes-blue\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/examples/ocean-commotion/assets/sprites.atlas\"\n"
  "}\n"
  ""
  position {
    z: 0.1
  }
}
