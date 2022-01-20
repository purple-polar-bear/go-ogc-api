# Readme

## Engine & Core

### Enable engine

```go
func EnableEngine() {
  mountingPath := "/package"
  engine := apicore.NewSimpleEngine(mountingPath)
  apicore.AddBaseJSONTemplates(engine)
  apicore.AddBaseHTMLTemplates(engine)
}
```

### Configuring the engine

```go
func ConfigureEngine() {
  config := engine.Config()
  config.SetTitle("Demo application - running latest GitHub version")
  config.SetDescription("DemoApp provides an API to geospatial data")
  config.SetProtocol("http")
  config.SetHost("172.17.0.1")
  config.SetPort(8080)
}
```

### Enabling Core

```go
func EnableOGCAPICore() {
  apiService := engine.GetService("core").(coreservices.CoreService)
  apiService.SetContact(&coreservices.ContactInfo{Name: "PDOK", Url: "https://pdok.nl/contact"})
  apiService.SetLicense(&coreservices.LicenseInfo{Name: "CC-BY 4.0 license", Url: "https://creativecommons.org/licenses/by/4.0/"})
  ctUrlEncoder := coremodels.NewContentTypeUrlEncoding("f")
  ctUrlEncoder.AddContentType("json", "application/vnd.oai.openapi+json;version=3.0")
  ctUrlEncoder.AddContentType("json", "application/json")
  ctUrlEncoder.AddContentType("html", "text/html")
  apiService.SetContentTypeUrlEncoder(ctUrlEncoder)
}
```

The following code fragment shows the minimum requirements for creating the engine and mounting it inside an existing router.

## Features

### Enabling Features
```go
func EnableOGCAPIFeatures() {
  engine := apif.NewSimpleEngine(mountingPath)
  apif.AddBaseJSONTemplates(engine)
  apif.AddBaseHTMLTemplates(engine)

  featuredatasource := myDatasourceImplementation.NewImplementation()
  apif.EnableFeatures(engine, featuredatasource)
  apif.AddFeaturesJSONTemplates(engine)
  apif.AddFeaturesHTMLTemplates(engine)

  engine.RebuildOpenAPI()
  router.HandleFunc(regexp.MustCompile("^"+mountingPath), engine.HTTPHandler)
}
```

### Implementing a FeatureService
In order to provide a datasource implementation for features, the following methods must be implemented:

```go
type FeatureService interface {
  Collections() []featuremodels.Collection
  Collection(string) featuremodels.Collection
  Features(*http.Request, *featuremodels.FeaturesParams) featuremodels.Features
  Feature(string, string) *featuremodels.Feature
  BuildOpenAPISpecification(builder coreservices.OpenAPIBuilder)
}
```
