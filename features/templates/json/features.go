package jsontemplates

import(
  "github.com/purple-polar-bear/go-ogc-api/core/models"
  "github.com/purple-polar-bear/go-ogc-api/core/templates/json"
  "github.com/purple-polar-bear/go-ogc-api/features/models"
  "github.com/purple-polar-bear/go-ogc-api/features/viewmodels"
  "github.com/purple-polar-bear/go-ogc-api/features/templates/core"
)

type FeatureRenderer struct {
}

func NewFeatureRenderer() coretemplates.RenderFeaturesType {
  return &FeatureRenderer{}
}

func (renderer *FeatureRenderer) RenderCollections(context *coremodels.Webcontext, collections *viewmodels.Collections) {
  jsontemplates.RenderPage(context, collections)
}

func (renderer *FeatureRenderer) RenderCollection(context *coremodels.Webcontext, collection *viewmodels.Collection) {
  jsontemplates.RenderPage(context, collection)
}

func (renderer *FeatureRenderer) RenderItems(context *coremodels.Webcontext, items *viewmodels.Features) {
  jsontemplates.RenderPage(context, items)
}

func (renderer *FeatureRenderer) RenderItem(context *coremodels.Webcontext, item *featuremodels.Feature) {
  jsontemplates.RenderPage(context, item)
}
