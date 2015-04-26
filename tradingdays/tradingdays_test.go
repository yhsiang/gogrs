package tradingdays

import (
	"fmt"
	"testing"

	"github.com/toomore/gogrs/utils"
)

func ExampleIsOpen() {
	fmt.Println(IsOpen(2015, 4, 17, utils.TaipeiTimeZone))
	fmt.Println(IsOpen(2015, 5, 1, utils.TaipeiTimeZone))
	// output:
	// true
	// false
}

func TestIsOpen(t *testing.T) {
	DownloadCSV(true)
	if IsOpen(2015, 4, 17, utils.TaipeiTimeZone) != true {
		t.Error("Should be `true`")
	}
	if IsOpen(2015, 4, 18, utils.TaipeiTimeZone) != false {
		t.Error("Should be `false`")
	}
	if IsOpen(2015, 4, 20, utils.TaipeiTimeZone) != true {
		t.Error("Should be `true`")
	}
	if IsOpen(2015, 5, 1, utils.TaipeiTimeZone) != false {
		t.Error("Should be `false`")
	}
}

func BenchmarkIsOpen(t *testing.B) {
	for i := 0; i < t.N; i++ {
		IsOpen(2015, 4, 19, utils.TaipeiTimeZone)
	}
}

func BenchmarkDownloadCSV(t *testing.B) {
	for i := 0; i < t.N; i++ {
		DownloadCSV(true)
	}
}
