package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.55

import (
	"context"

	api "github.com/photoview/photoview/api/graphql"
	"github.com/photoview/photoview/api/graphql/models"
	"github.com/photoview/photoview/api/scanner/face_detection"
)

// SiteInfo is the resolver for the siteInfo field.
func (r *queryResolver) SiteInfo(ctx context.Context) (*models.SiteInfo, error) {
	return models.GetSiteInfo(r.DB(ctx))
}

// FaceDetectionEnabled is the resolver for the faceDetectionEnabled field.
func (r *siteInfoResolver) FaceDetectionEnabled(ctx context.Context, obj *models.SiteInfo) (bool, error) {
	return face_detection.GlobalFaceDetector != nil, nil
}

// SiteInfo returns api.SiteInfoResolver implementation.
func (r *Resolver) SiteInfo() api.SiteInfoResolver { return &siteInfoResolver{r} }

type siteInfoResolver struct{ *Resolver }
