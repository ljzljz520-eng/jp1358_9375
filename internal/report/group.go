package report

import (
	"strings"
	"yoga.example/studio/internal/model"
)

func GroupByEntity(items []model.Audit) map[string][]model.Audit {
	out := map[string][]model.Audit{}
	for _, i := range items {
		out[i.Entity] = append(out[i.Entity], i)
	}
	return out
}
func Actions(items []model.Audit) []string {
	seen := map[string]bool{}
	var out []string
	for _, i := range items {
		if !seen[i.Action] {
			seen[i.Action] = true
			out = append(out, strings.ToLower(i.Action))
		}
	}
	return out
}
