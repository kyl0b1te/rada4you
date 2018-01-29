package rada4you

import (
	"time"
)

type request interface {
	Values() map[string]string
}

type GetAllDivisionsRequest struct {
	Start time.Time
	End   time.Time
	House string
}

func (r *GetAllDivisionsRequest) Values() map[string]string {
	val := make(map[string]string)
	if !r.Start.IsZero() {
		val["start_date"] = r.Start.Format("2006-01-02")
	}
	if !r.End.IsZero() {
		val["end_date"] = r.End.Format("2006-01-02")
	}
	if r.House != "" {
		val["house"] = r.House
	}
	// 2 dates should be set
	if val["start_date"] == "" || val["end_date"] == "" {
		delete(val, "start_date")
		delete(val, "end_date")
	}
	return val
}
