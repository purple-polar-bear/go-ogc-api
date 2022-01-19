package htmltemplates

import(
  "html/template"
  "github.com/purple-polar-bear/go-ogc-api/core/models"
  "github.com/purple-polar-bear/go-ogc-api/core/templates/html"
  "github.com/purple-polar-bear/go-ogc-api/features/models"
  "github.com/purple-polar-bear/go-ogc-api/features/viewmodels"
)

// Transforms a renderlandingpage function into a renderlandingpage object
func NewFeatureRenderer() *FeatureRenderer {
  templates := htmltemplates.NewTemplate("features", []string{
    "collections.html",
    "collection.html",
    "features.html",
    "feature.html",
  })

  return &FeatureRenderer{
    Templates: templates,
  }
}

// Internal
type FeatureRenderer struct {
  Templates *template.Template
}

func (renderer *FeatureRenderer) RenderCollections(context *coremodels.Webcontext, collections *viewmodels.Collections) {
  renderer.Templates.ExecuteTemplate(context.W, "collections.html", collections)
}

func (renderer *FeatureRenderer) RenderCollection(context *coremodels.Webcontext, collection *viewmodels.Collection) {
  renderer.Templates.ExecuteTemplate(context.W, "collection.html", collection)
}

func (renderer *FeatureRenderer) RenderItems(context *coremodels.Webcontext, items *viewmodels.Features) {
  renderer.Templates.ExecuteTemplate(context.W, "features.html", items)
}

func (renderer *FeatureRenderer) RenderItem(context *coremodels.Webcontext, item *featuremodels.Feature) {
  renderer.Templates.ExecuteTemplate(context.W, "feature.html", item)
}
