package jsontemplates

import(
  "github.com/purple-polar-bear/go-ogc-api/core/models"
  "github.com/purple-polar-bear/go-ogc-api/core/viewmodels"
  "github.com/purple-polar-bear/go-ogc-api/core/templates"
)

type CoreRenderer struct {
}

func NewCoreRenderer() coretemplates.RenderCoreType {
  return &CoreRenderer{}
}

func (renderer *CoreRenderer) RenderLandingpage(context *coremodels.Webcontext, landingpage *viewmodels.Landingpage) {
  RenderPage(context, landingpage)
}

func (renderer *CoreRenderer) RenderConformance(context *coremodels.Webcontext, conformanceClasses *viewmodels.Conformanceclasses) {
  RenderPage(context, conformanceClasses)
}

func (renderer *CoreRenderer) RenderAPI(context *coremodels.Webcontext, api interface{}) {
  RenderPage(context, api)
}
