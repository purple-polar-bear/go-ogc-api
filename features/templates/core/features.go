package coretemplates

import(
  "github.com/purple-polar-bear/go-ogc-api/core/models"
  "github.com/purple-polar-bear/go-ogc-api/features/models"
  "github.com/purple-polar-bear/go-ogc-api/features/viewmodels"
)

// Interface definition for features renderer
type RenderFeaturesType interface {
  RenderCollections(context *coremodels.Webcontext, collections *viewmodels.Collections)
  RenderCollection(context *coremodels.Webcontext, collection *viewmodels.Collection)
  RenderItems(context *coremodels.Webcontext, items *viewmodels.Features)
  RenderItem(context *coremodels.Webcontext, item *featuremodels.Feature)
}
