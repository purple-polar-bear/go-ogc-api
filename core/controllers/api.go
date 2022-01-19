package corecontrollers

import(
  "net/http"

  "github.com/purple-polar-bear/go-ogc-api/core/models"
  "github.com/purple-polar-bear/go-ogc-api/core/services"
  "github.com/purple-polar-bear/go-ogc-api/core/templates"
)

type APIController struct {
}

func (controller *APIController) HandleFunc(app coremodels.Application, r interface{}) coremodels.ControllerFunc {
  renderer := r.(coretemplates.RenderCoreType)

  return func(handler coremodels.Handler, w http.ResponseWriter, r *http.Request, routeParameters coremodels.MatchedRouteParameters) {
    apiService, ok := app.GetService("core").(coreservices.CoreService)
    if !ok {
      panic("Cannot find featureservice")
    }

    resource := apiService.OpenAPI()
    renderer.RenderAPI(coremodels.NewWebcontext(w, r), resource)
  }
}
