package viewmodels

import(
  "github.com/purple-polar-bear/go-ogc-api/core/viewmodels"
)

type Collections struct {
  Collections []*Collection `json:"collections"`
  Links []*viewmodels.Link `json:"links"`
}

func NewCollections() *Collections {
  return &Collections{
    Collections: []*Collection{},
    Links: []*viewmodels.Link{},
  }
}
