package coincap

import "time"

//
// Interval
//

type Interval int

const (
	M1  Interval = iota // max range: 1day
	M5                  // max range: 5day
	M15                 // max range: 7day
	M30                 // max range: 14day
	H1                  // max range: 30day
	H2                  // max range: 61day
	H6                  // max range: 183day
	H12                 // max range: 365day
	D1                  // max range: 7305day
)

func (i Interval) String() string {
	switch i {
	case M1:
		return "m1"
	case M5:
		return "m5"
	case M15:
		return "m15"
	case M30:
		return "m30"
	case H1:
		return "h1"
	case H2:
		return "h2"
	case H6:
		return "h6"
	case H12:
		return "h12"
	case D1:
		return "d1"
	default:
		return ""
	}
}

func (i Interval) Value() time.Duration {
	switch i {
	case M1:
		return time.Minute * 1
	case M5:
		return time.Minute * 5
	case M15:
		return time.Minute * 15
	case M30:
		return time.Minute * 30
	case H1:
		return time.Hour * 1
	case H2:
		return time.Hour * 2
	case H6:
		return time.Hour * 6
	case H12:
		return time.Hour * 12
	case D1:
		return time.Hour * 24
	default:
		return time.Duration(0)
	}

}
