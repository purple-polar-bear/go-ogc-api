package featurescontrollers

import(
  "net/http"

  "github.com/purple-polar-bear/go-ogc-api/core/services"
  "github.com/purple-polar-bear/go-ogc-api/core/models"
  coreviewmodels "github.com/purple-polar-bear/go-ogc-api/core/viewmodels"
  "github.com/purple-polar-bear/go-ogc-api/features/models"
  "github.com/purple-polar-bear/go-ogc-api/features/services"
  "github.com/purple-polar-bear/go-ogc-api/features/templates/core"

  // "github.com/go-spatial/geom/encoding/geojson"
)

type FeatureController struct {

}

func (controller *FeatureController) HandleFunc(app coremodels.Application, r interface{}) coremodels.ControllerFunc {
  renderer := r.(coretemplates.RenderFeaturesType)

  return func(handler coremodels.Handler, w http.ResponseWriter, r *http.Request, routeParameters coremodels.MatchedRouteParameters) {
    templates := app.Templates("feature", "")

    featureService, ok := app.GetService("features").(featureservices.FeatureService)
    if !ok {
      panic("Cannot find featureservice")
    }

    coreservice, ok := app.GetService("core").(coreservices.CoreService)
    if !ok {
      panic("Cannot find coreservice")
    }

    encoding := coreservice.ContentTypeUrlEncoder()
    collectionId := routeParameters.Get("collection_id")
    featureId := routeParameters.Get("item_id")
    feature := featureService.Feature(collectionId, featureId)
    baseUrl := app.Config().FullUri()
    hrefParams := make(map[string]string)
    hrefParams["collection_id"] = collectionId
    hrefParams["item_id"] = featureId
    links := []*coreviewmodels.Link{}
    // current link
    for _, template := range templates {
  		baseHref := template.Href(baseUrl, hrefParams, encoding)
      link := &coreviewmodels.Link{
        Title: template.Title(),
        Rel: template.Rel(handler.Type()),
        Type: template.Type(),
        Href: baseHref,
      }

  		links = append(links, link)
  	}

    resource := &featuremodels.Feature{
      Feature: feature,
      Links: links,
    }
    renderer.RenderItem(coremodels.NewWebcontext(w, r), resource)
  }
}
