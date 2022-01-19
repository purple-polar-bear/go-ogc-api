package htmltemplates

import(
  "io/fs"
  "html/template"
  "strings"
)

func NewTemplate(klass string, pages fs.FS) *template.Template {
  /*
  path := "./" + klass + "/templates/html/templates/"
  pagesWithPath := []string{}
  for _, page := range pages {
    pagesWithPath = append(pagesWithPath, path + page)
  }
  */

  return template.Must(template.New("templates").Funcs(
   template.FuncMap{
     "isOdd":       func(i int) bool { return i%2 != 0 },
     "titleize":  func(title string) string { return strings.Title(title) },
   },
  ).ParseFS(pages, "**/*"))
}
