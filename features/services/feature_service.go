package featureservices

import (
  "net/http"
  "github.com/purple-polar-bear/go-ogc-api/core/services"
  "github.com/purple-polar-bear/go-ogc-api/features/models"
)

type FeatureService interface {
  Collections() []featuremodels.Collection
  Collection(string) featuremodels.Collection
  Features(*http.Request, *featuremodels.FeaturesParams) featuremodels.Features
  Feature(string, string) *featuremodels.Feature
  BuildOpenAPISpecification(builder coreservices.OpenAPIBuilder)
}
