package coretemplates

import(
  "github.com/purple-polar-bear/go-ogc-api/core/models"
  "github.com/purple-polar-bear/go-ogc-api/core/viewmodels"
)

// Interface definition for features renderer
type RenderCoreType interface {
  RenderLandingpage(context *coremodels.Webcontext, landingpage *viewmodels.Landingpage)
  RenderConformance(context *coremodels.Webcontext, conformanceClasses *viewmodels.Conformanceclasses)
  RenderAPI(context *coremodels.Webcontext, api interface{})
}
