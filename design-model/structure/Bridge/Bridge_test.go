package Bridge

import (
	"fmt"
	"testing"
)

func TestBridge(t *testing.T) {
	black := Black{}
	gold := Gold{}

	p800 := P800{}
	p1000 := P1000{}

	xiaomi := NewXiaomi(black, p800)
	xiaomi.color.showColor()
	xiaomi.pixel.showPixel()
	xiaomi.call()

	fmt.Println()
	
	huawei := NewHuawei(gold, p1000)
	huawei.color.showColor()
	huawei.pixel.showPixel()
	huawei.call()

}
