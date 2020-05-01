package piskel

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var testOk = `{"modelVersion":2,"piskel":{"name":"Sheep_1","description":"","fps":3,"height":16,"width":16,"layers":["{\"name\":\"Layer 1\",\"opacity\":1,\"frameCount\":2,\"chunks\":[{\"layout\":[[0],[1]],\"base64PNG\":\"data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAACAAAAAQCAYAAAB3AH1ZAAAAsElEQVRIS82VWQ6AMAhEy/0PjRGLoQgUXFL9cxl4TMoIbfEFi/s3DwADsAx0Wm8VQ0RfD0CSCKKk14VOcW80GMFgAUTYnItJvQQgsWzM995z5cTQPBqALOxOXgAUpemALCA+uDXA6wDVAf4LIGymofRmGAfx8RnofaIVPgzOboE+0LMt4PelPTYCq6R3k9AKo0QIhUNY+mmi7SKRA5kYHiBm+hDAsLcEkNFXCn7y49wAICOSEeflEdkAAAAASUVORK5CYII=\"}]}"],"hiddenFrames":[]}}`
var testEmptyLayers = `{"modelVersion":2,"piskel":{"name":"Sheep_1","description":"","fps":3,"height":16,"width":16,"layers":[]}}`
var testNoLayers = `{"modelVersion":2,"piskel":{"name":"Sheep_1","description":"","fps":3,"height":16,"width":16}}`
var testMultipleLayers = `{"modelVersion":2,"piskel":{"name":"Layers","description":"","fps":12,"height":16,"width":16,"layers":["{\"name\":\"Layer 1\",\"opacity\":1,\"frameCount\":6,\"chunks\":[{\"layout\":[[0],[1],[2],[3],[4],[5]],\"base64PNG\":\"data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAGAAAAAQCAYAAADpunr5AAAAYUlEQVRYR+2VsQ0AIAzD0v+PBmXggg6WkLtTSzGBiYMmMChdeBQAXwIFKABOAMbbgA8EnGT9l2x3bM9XA7KjDSjYgRLwCYKCf1gFKABOAMbbAAXACcB4G6AAOAEYbwNgARei+gYQ5y+M5AAAAABJRU5ErkJggg==\"}]}","{\"name\":\"Layer 2\",\"opacity\":1,\"frameCount\":6,\"chunks\":[{\"layout\":[[0],[1],[2],[3],[4],[5]],\"base64PNG\":\"data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAGAAAAAQCAYAAADpunr5AAAAZElEQVRYR+2TMQ4AIBDC5P+PPlfjKkkH6wNAW8nyoASCtlu+FAB/AgUoACYA17sABcAE4HoXQAuYmUnyJKKR0eDQvEcz637bmf0EvgHt9wwFwD9AAQqACcD1LkABMAG43gXAAjbvhhQRz1ldQgAAAABJRU5ErkJggg==\"}]}"],"hiddenFrames":[]}}`

func TestLoadFromFile(t *testing.T) {
	p, err := LoadFromFile("test_example.piskel")
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	assert.NotNil(t, p)
	assert.Equal(t, 2, p.ModelVersion)
	assert.NotNil(t, p.Content)
	assert.Equal(t, "Sheep_1", p.Content.Name)
	assert.Equal(t, 3, p.Content.Fps)
}

func TestLoadFromFileNotFound(t *testing.T) {
	_, err := LoadFromFile("file_not_found.piskel")
	if err == nil {
		t.Error(err)
		t.Fail()
		return
	}
}

func TestLoadFromString(t *testing.T) {
	p, err := LoadFromString(testOk)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	assert.NotNil(t, p)
	assert.Equal(t, 2, p.ModelVersion)
	assert.NotNil(t, p.Content)
	assert.Equal(t, "Sheep_1", p.Content.Name)
	assert.Equal(t, 3, p.Content.Fps)
	assert.NotNil(t, p.Content.Layers)
}

func TestLoadFromStringInvalidJson(t *testing.T) {
	_, err := LoadFromString(":-(")
	if err == nil {
		t.Error(err)
		t.Fail()
		return
	}
}

func TestEmptyLayers(t *testing.T) {
	p, err := LoadFromString(testEmptyLayers)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	assert.NotNil(t, p)
	assert.Equal(t, 2, p.ModelVersion)
	assert.NotNil(t, p.Content)
	assert.Equal(t, "Sheep_1", p.Content.Name)
	assert.Equal(t, 3, p.Content.Fps)
	assert.Nil(t, p.Content.Layers)
}

func TestNoLayers(t *testing.T) {
	p, err := LoadFromString(testNoLayers)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	assert.NotNil(t, p)
	assert.Equal(t, 2, p.ModelVersion)
	assert.NotNil(t, p.Content)
	assert.Equal(t, "Sheep_1", p.Content.Name)
	assert.Equal(t, 3, p.Content.Fps)
	assert.Nil(t, p.Content.Layers)
}

func TestMultipleLayers(t *testing.T) {
	p, err := LoadFromString(testMultipleLayers)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	assert.NotNil(t, p)
	assert.Equal(t, 2, p.ModelVersion)
	assert.NotNil(t, p.Content)
	assert.Equal(t, "Layers", p.Content.Name)
	assert.Equal(t, 12, p.Content.Fps)
	assert.NotNil(t, p.Content.Layers)
	assert.Equal(t, 2, len(p.Content.Layers))
	assert.Equal(t, 1, len(p.Content.Layers[0].Chunks))
	assert.Equal(t, 1, len(p.Content.Layers[1].Chunks))
}

func TestImageDecoding(t *testing.T) {
	p, err := LoadFromString(testMultipleLayers)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	assert.NotNil(t, p)
	assert.NotNil(t, p.Content)
	assert.NotNil(t, p.Content.Layers)
	assert.Equal(t, 2, len(p.Content.Layers))
	assert.Equal(t, 1, len(p.Content.Layers[0].Chunks))

	ch1 := p.Content.Layers[0].Chunks[0]
	assert.NotEmpty(t, ch1.Base64PNG)
	assert.NotNil(t, ch1.Image)
	assert.Equal(t, 96, ch1.Image.Bounds().Max.X)
	assert.Equal(t, 16, ch1.Image.Bounds().Max.Y)
}
