package types

import (
	"encoding/json"
	"io"
	"text/template"
)

// Anime is a kitsu api payload
// type Anime = kitsuclient.Anime
type Anime struct {
	Attributes map[string]interface{} `json:"attributes"`
	ID         string                 `json:"id"`
}

var (
	animeMutation = template.Must(template.New("").Funcs(template.FuncMap{
		"strCast": func(v interface{}) (string, error) {
			switch v.(type) {
			case string:
				return v.(string), nil
			default:
				b, err := json.Marshal(v)
				return string(b), err
			}
		},
	}).Parse(`
{{$id := .ID}}
mutation {
	set {
		_:a{{$id}} <id> "{{$id}}" .
		_:a{{$id}} <type> "anime" .
		_:a{{$id}} <identity> "a{{$id}}"
		{{range $k, $v := .Attributes}}
		_:a{{$id}} <{{$k}}> "{{ strCast $v | js }}" .
		{{end}}
	}
}
`))

	animeSchema = `
mutation {
	schema {
		id: int @index(int) .
		type: string @index(exact) . 
		identity: string @index(exact) .
	}
}
`
)

// WriteMutation to the io.Writer.
func (a *Anime) WriteMutation(w io.Writer) error {
	return animeMutation.Execute(w, a)
}
