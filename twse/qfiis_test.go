package twse

import (
	"testing"
	"time"

	"github.com/toomore/gogrs/utils"
)

func TestQFIISTOP20_Get(t *testing.T) {
	qf := &QFIISTOP20{Date: time.Date(2015, 5, 25, 0, 0, 0, 0, utils.TaipeiTimeZone)}
	t.Log(qf.URL())
	t.Log(qf.Get())
}

func TestBFI82U_Get(t *testing.T) {
	bfi := &BFI82U{
		Begin: time.Date(2015, 5, 25, 0, 0, 0, 0, utils.TaipeiTimeZone),
		End:   time.Date(2015, 5, 26, 0, 0, 0, 0, utils.TaipeiTimeZone),
	}
	t.Log(bfi.URL())
	t.Log(bfi.Get())
}

func TestT86_Get(t *testing.T) {
	t86 := &T86{Date: time.Date(2015, 5, 25, 0, 0, 0, 0, utils.TaipeiTimeZone)}
	t.Log(t86.URL("01"))
	data, _ := t86.Get("ALLBUT0999")
	t.Log(data, len(data), data[:5])
}

func TestTWTXXU_Get(t *testing.T) {
	date := time.Date(2015, 5, 26, 0, 0, 0, 0, utils.TaipeiTimeZone)
	twt43u := NewTWT43U(date)
	t.Log(twt43u.URL())
	t.Log(twt43u.Get())
}