package scalar

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"time"
)

const layoutISO8601 = "2006-01-02T15:04:05.000Z07:00"

type DateTime time.Time

func (dt DateTime) MarshalJSON() ([]byte, error) {
	str := time.Time(dt).Format(layoutISO8601)
	return json.Marshal(str)
}

func (dt DateTime) MarshalGQL(w io.Writer) {
	buf, err := dt.MarshalJSON()
	if err != nil {
		log.Printf("DateTime#MarshalGQL error %v", err)
	}
	_, _ = w.Write(buf)
}

func (dt *DateTime) UnmarshalGQL(v any) (err error) {
	switch v := v.(type) {
	case string:
		t, err := time.Parse(layoutISO8601, v)
		if err == nil {
			*dt = DateTime(t)
		}
		return err
	default:
		return fmt.Errorf("%T is not a time", v)
	}
}
