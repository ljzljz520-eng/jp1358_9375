package report

import (
	"encoding/json"
	"io"
	"time"
	"yoga.example/studio/internal/model"
)

type Export struct {
	Generated time.Time
	Items     []model.Audit
}

func BuildExport(items []model.Audit) Export {
	return Export{Generated: time.Now().UTC(), Items: SortAudits(items)}
}
func WriteJSON(w io.Writer, e Export) error {
	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	return enc.Encode(e)
}
func CountActions(items []model.Audit) map[string]int {
	out := map[string]int{}
	for _, i := range items {
		out[i.Action]++
	}
	return out
}
func Latest(items []model.Audit) model.Audit {
	if len(items) == 0 {
		return model.Audit{}
	}
	return SortAudits(items)[len(items)-1]
}
