package face_detection

import (
	"github.com/photoview/photoview/api/graphql/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
)

func setupTestDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&models.ImageFace{}, &models.FaceGroup{}, &models.Media{}, &models.MediaURL{}, &models.User{}, &models.Album{})
	return db, nil
}

func TestReloadFacesFromDatabase(t *testing.T) {
	db, err := setupTestDB()
	assert.NoError(t, err)

	fd := &faceDetector{}
	err = fd.ReloadFacesFromDatabase(db)
	assert.NoError(t, err)
	assert.Empty(t, fd.faceDescriptors)
}

func TestDetectFaces_NoThumbnail(t *testing.T) {
	db, err := setupTestDB()
	assert.NoError(t, err)

	// Create a media record without a thumbnail URL
	media := &models.Media{}
	err = db.Create(media).Error
	assert.NoError(t, err)

	fd := &faceDetector{}
	err = fd.DetectFaces(db, media, false)
	assert.Error(t, err)
	assert.Equal(t, "thumbnail url is missing", err.Error())
}

func TestMergeCategories(t *testing.T) {
	fd := &faceDetector{
		faceGroupIDs: []int32{1, 2, 3},
	}

	fd.MergeCategories(2, 4)
	assert.Equal(t, []int32{1, 4, 3}, fd.faceGroupIDs)
}

func TestMergeImageFaces(t *testing.T) {
	fd := &faceDetector{
		faceGroupIDs: []int32{1, 2, 3},
		imageFaceIDs: []int{10, 20, 30},
	}

	fd.MergeImageFaces([]int{20}, 4)
	assert.Equal(t, []int32{1, 4, 3}, fd.faceGroupIDs)
}
