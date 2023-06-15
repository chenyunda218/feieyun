package feieyun

import (
	"os"
	"testing"
)

func TestPrint(t *testing.T) {
	t.Log("Print testing")
	ukey := os.Getenv("UKEY")
	endpoint := os.Getenv("ENDPOINT")
	sn := os.Getenv("SN")
	user := os.Getenv("USER")
	printer := &Printer{endpoint, ukey, sn, user}
	t.Log(printer.print("<C><B>放大一倍</B></C>"))
}
