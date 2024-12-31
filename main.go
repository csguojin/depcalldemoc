package depcalldemoc

import (
	"fmt"

	"github.com/csguojin/depcalldemod"
)

func MyFunc1() {
	fmt.Println("depcalldemoC.MyFunc1")
	depcalldemod.MyFunc3()
}
