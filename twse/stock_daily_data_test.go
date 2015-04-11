package twse

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/toomore/gogrs/utils"
)

func assertType(t *testing.T, t1 interface{}, t2 interface{}) {
	if reflect.TypeOf(t1) != reflect.TypeOf(t2) {
		t.Errorf("Diff type t1(%s), t2(%s)", reflect.TypeOf(t1), reflect.TypeOf(t2))
	}
}

func TestURL(t *testing.T) {
	var d = NewTWSE("2618", time.Now())
	assertType(t, d, &Data{})
}

func ExampleData() {
	var d = NewTWSE("2618", time.Date(2014, 12, 26, 0, 0, 0, 0, time.Local))

	stockData, _ := d.Get()
	fmt.Println(stockData[0])
	// output:
	// [103/12/01 64,418,143 1,350,179,448 20.20 21.40 20.20 21.35 +1.35 13,249]
}

func TestData_Get(*testing.T) {
	var d = NewTWSE("2618", time.Date(2014, 12, 26, 0, 0, 0, 0, time.Local))
	fmt.Println(d)

	stockData, _ := d.Get()
	d.Get() // Test Cache.
	fmt.Println(stockData)
}

var twse = NewTWSE("2329", time.Date(2015, 03, 20, 0, 0, 0, 0, time.Local))
var otc = NewOTC("8446", time.Date(2015, 03, 20, 0, 0, 0, 0, time.Local))

func TestGetList(*testing.T) {
	for _, stock := range []*Data{twse, otc} {

		fmt.Println(stock)
		fmt.Println(stock.URL())
		stock.Get()
		fmt.Println(stock.RawData)
		fmt.Println(stock.MA(6))
		fmt.Println(stock.MAV(6))
		fmt.Println(stock.GetPriceList())
		fmt.Println(utils.ThanPastFloat64(stock.GetPriceList(), 3, true))
		fmt.Println(utils.ThanPastFloat64(stock.GetPriceList(), 3, false))
		fmt.Println(stock.GetVolumeList())
		fmt.Println(utils.ThanPastUint64(stock.GetVolumeList(), 3, true))
		fmt.Println(utils.ThanPastUint64(stock.GetVolumeList(), 3, false))
		fmt.Println(stock.GetRangeList())
		fmt.Println(stock.GetOpenList())
		fmt.Println(stock.IsRed())
	}
}

func BenchmarkGet(b *testing.B) {
	var d = NewTWSE("2618", time.Date(2014, 12, 26, 0, 0, 0, 0, time.Local))
	for i := 0; i <= b.N; i++ {
		d.Get()
	}
}

func BenchmarkGetVolumeList(b *testing.B) {
	var d = NewTWSE("2618", time.Date(2015, 3, 27, 0, 0, 0, 0, time.Local))
	d.Get()
	for i := 0; i <= b.N; i++ {
		d.GetVolumeList()
	}
}

func BenchmarkGetPriceList(b *testing.B) {
	var d = NewTWSE("2618", time.Date(2015, 3, 27, 0, 0, 0, 0, time.Local))
	d.Get()
	for i := 0; i <= b.N; i++ {
		d.GetPriceList()
	}
}

// 新增一個 TWSE 上市股票
func Example_newTWSE() {
	var stock = NewTWSE("2618", time.Date(2015, 3, 27, 0, 0, 0, 0, time.Local))
	stock.Get()
	fmt.Println(stock.RawData[0])
	// output:
	// [104/03/02 13,384,378 305,046,992 23.00 23.05 22.50 22.90 -0.10 3,793]
}

// 新增一個 OTC 上櫃股票
func Example_newOTC() {
	var stock = NewOTC("8446", time.Date(2015, 3, 27, 0, 0, 0, 0, time.Local))
	stock.Get()
	fmt.Println(stock.RawData[0])
	// output:
	// [104/03/02 354 33,018 92.00 94.90 90.80 92.60 3.50 299]
}

func ExampleData_Get_notEnoughData() {
	year, month, _ := time.Now().Date()
	var d = NewTWSE("2618", time.Date(year, month+1, 1, 0, 0, 0, 0, time.Local))

	stockData, err := d.Get()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(stockData)
	}
	// output:
	// Not enough data.
}

func ExampleData_PlusData() {
	var stock = NewTWSE("2618", time.Date(2015, 3, 27, 0, 0, 0, 0, time.Local))
	stock.Get() // 2015/3
	fmt.Println(stock.Date)
	stock.PlusData() // 2015/2
	fmt.Println(stock.Date)
	// output:
	// 2015-03-27 00:00:00 +0800 CST
	// 2015-02-01 00:00:00 +0800 CST
}

func ExampleData_Round() {
	var d = NewTWSE("2618", time.Date(2014, 12, 26, 0, 0, 0, 0, time.Local))

	fmt.Println(d.Date) // 2014/12

	d.Round()
	fmt.Println(d.Date) // 2014/11

	d.Round()
	fmt.Println(d.Date) // 2014/10
	// output:
	// 2014-12-26 00:00:00 +0800 CST
	// 2014-11-01 00:00:00 +0800 CST
	// 2014-10-01 00:00:00 +0800 CST
}

func TestData_Round(t *testing.T) {
	var now = time.Date(2015, 3, 27, 0, 0, 0, 0, time.Local)
	var past = time.Date(2015, 2, 1, 0, 0, 0, 0, time.Local)
	var d = NewTWSE("2618", now)

	t.Log(d.Date)
	if d.Date == past {
		t.Fatal(d.Date, past)
	}
	d.Round()
	t.Log(d.Date)
	if d.Date != past {
		t.Fatal(d.Date, past)
	}
}

func TestData_PlusData(t *testing.T) {
	var now = time.Date(2015, 3, 27, 0, 0, 0, 0, time.Local)
	var d = NewTWSE("2618", now)
	d.PlusData()
	var d2 = NewTWSE("2618", time.Date(2015, 2, 1, 0, 0, 0, 0, time.Local))
	d2.Get()
	for i := range d2.RawData {
		if d.RawData[i][0] != d2.RawData[i][0] {
			t.Fatal("Data not difference.")
			t.Log(d.RawData, d2.RawData)
		}
	}
}

func TestData_GetByTimeMap(*testing.T) {
	var d = NewTWSE("2618", time.Date(2014, 12, 26, 0, 0, 0, 0, time.Local))
	fmt.Println(d.GetByTimeMap())
}

func TestData_FormatData(*testing.T) {
	var d = NewTWSE("2618", time.Date(2014, 12, 26, 0, 0, 0, 0, time.Local))
	d.Get()
	fmt.Println(d.FormatData())
}
