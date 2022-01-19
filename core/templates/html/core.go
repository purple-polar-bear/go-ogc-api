package htmltemplates

import(
  "embed"
  "html/template"
  "github.com/purple-polar-bear/go-ogc-api/core/models"
  "github.com/purple-polar-bear/go-ogc-api/core/viewmodels"
)

//go:embed templates/*
var templateFiles embed.FS

// Transforms a renderlandingpage function into a renderlandingpage object
func NewCoreRenderer() *CoreRenderer {
  templates := NewTemplate("core", templateFiles)

  return &CoreRenderer{
    Templates: templates,
  }
}

// Internal
type CoreRenderer struct {
  Templates *template.Template
}

func (renderer *CoreRenderer) RenderLandingpage(context *coremodels.Webcontext, landingpageClasses *viewmodels.Landingpage) {
  renderer.Templates.ExecuteTemplate(context.W, "landingpage.html", landingpageClasses)
}

func (renderer *CoreRenderer) RenderConformance(context *coremodels.Webcontext, conformanceClasses *viewmodels.Conformanceclasses) {
  renderer.Templates.ExecuteTemplate(context.W, "conformance.html", conformanceClasses)
}

func (renderer *CoreRenderer) RenderAPI(context *coremodels.Webcontext, api interface{}) {
  renderer.Templates.ExecuteTemplate(context.W, "api.html", api)
}
