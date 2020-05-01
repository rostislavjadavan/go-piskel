package piskel

import (
	"encoding/base64"
	"encoding/json"
	"image"
	_ "image/png" // png format decoder
	"io/ioutil"
	"strings"
)

// LoadFromString load Piksel format from string
func LoadFromString(content string) (*Piskel, error) {
	piskel := Piskel{}
	err := json.Unmarshal([]byte(content), &piskel)
	if err != nil {
		return nil, err
	}
	piskel.Content.Layers = parseLayers(piskel.Content.LayerString)
	return &piskel, nil
}

// LoadFromFile load Piksel format from file
func LoadFromFile(filename string) (*Piskel, error) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return LoadFromString(string(file))
}

func parseLayers(layerString []string) []PiskelLayer {
	if len(layerString) == 0 {
		return nil
	}
	var layers []PiskelLayer
	for _, l := range layerString {
		layer := PiskelLayer{}
		err := json.Unmarshal([]byte(l), &layer)
		if err == nil {
			for i := range layer.Chunks {
				processChunk(&layer.Chunks[i])
			}
			layers = append(layers, layer)
		}
	}
	return layers
}

func processChunk(chunk *PiskelChunk) {
	data := strings.TrimPrefix(chunk.Base64PNG, "data:image/png;base64,")
	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(data))
	image, _, err := image.Decode(reader)
	if err == nil {
		chunk.Image = image
	}
}
