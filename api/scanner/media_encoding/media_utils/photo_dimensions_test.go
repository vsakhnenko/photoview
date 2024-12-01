package media_utils

import (
	"image"
	"image/jpeg"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPhotoDimensions(t *testing.T) {
	// Create a temporary image file for testing
	imagePath := "test_image.jpg"
	img := image.NewRGBA(image.Rect(0, 0, 100, 200))
	file, err := os.Create(imagePath)
	assert.NoError(t, err)
	defer os.Remove(imagePath)
	defer file.Close()

	err = jpeg.Encode(file, img, nil)
	assert.NoError(t, err)

	dimensions, err := GetPhotoDimensions(imagePath)
	assert.NoError(t, err)
	assert.NotNil(t, dimensions)
	assert.Equal(t, 100, dimensions.Width)
	assert.Equal(t, 200, dimensions.Height)
}

func TestPhotoDimensionsFromRect(t *testing.T) {
	rect := image.Rect(0, 0, 100, 200)
	dimensions := PhotoDimensionsFromRect(rect)
	assert.Equal(t, 100, dimensions.Width)
	assert.Equal(t, 200, dimensions.Height)
}

func TestThumbnailScale(t *testing.T) {
	dimensions := &PhotoDimensions{Width: 2000, Height: 1000}
	thumbnail := dimensions.ThumbnailScale()
	assert.Equal(t, 1024, thumbnail.Width)
	assert.Equal(t, 512, thumbnail.Height)

	dimensions = &PhotoDimensions{Width: 1000, Height: 2000}
	thumbnail = dimensions.ThumbnailScale()
	assert.Equal(t, 512, thumbnail.Width)
	assert.Equal(t, 1024, thumbnail.Height)

	dimensions = &PhotoDimensions{Width: 800, Height: 600}
	thumbnail = dimensions.ThumbnailScale()
	assert.Equal(t, 800, thumbnail.Width)
	assert.Equal(t, 600, thumbnail.Height)
}
