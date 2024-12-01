package scanner_cache

import (
	"testing"

	"github.com/photoview/photoview/api/scanner/media_type"
	"github.com/stretchr/testify/assert"
)

func TestMakeAlbumCache(t *testing.T) {
	cache := MakeAlbumCache()
	assert.NotNil(t, cache)
	assert.NotNil(t, cache.path_contains_photos)
	assert.NotNil(t, cache.photo_types)
	assert.NotNil(t, cache.ignore_data)
}

func TestInsertAlbumPath(t *testing.T) {
	cache := MakeAlbumCache()
	cache.InsertAlbumPath("/test/path", true)

	containsPhoto := cache.AlbumContainsPhotos("/test/path")
	assert.NotNil(t, containsPhoto)
	assert.True(t, *containsPhoto)
}

func TestGetMediaType(t *testing.T) {
	cache := MakeAlbumCache()
	mediaType := media_type.MediaType("image/jpeg")
	cache.photo_types["/test/image.jpg"] = mediaType

	result, err := cache.GetMediaType("/test/image.jpg")
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, mediaType, *result)
}

func TestInsertAlbumIgnore(t *testing.T) {
	cache := MakeAlbumCache()
	ignoreData := []string{"ignore1", "ignore2"}
	cache.InsertAlbumIgnore("/test/path", ignoreData)

	result := cache.GetAlbumIgnore("/test/path")
	assert.NotNil(t, result)
	assert.Equal(t, ignoreData, *result)
}
