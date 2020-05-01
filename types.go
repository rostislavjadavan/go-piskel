package piskel

import "image"

type Piskel struct {
	ModelVersion int        `json:"modelVersion"`
	Content      PiskelContent `json:"piskel"`
}

type PiskelContent struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Fps         int      `json"fps"`
	Width       int      `json:"width"`
	Height      int      `json:"height"`
	LayerString []string `json:"layers"`
	Layers      []PiskelLayer
}

type PiskelLayer struct {
	Name       string        `json:"name"`
	Opacity    float32       `json:"opacity"`
	FrameCount int           `json:"frameCount"`
	Chunks     []PiskelChunk `json:"chunks"`
}

type PiskelChunk struct {
	Layout    [][]int `json:"layout"`
	Base64PNG string  `json:"base64PNG"`
	Image     image.Image
}
