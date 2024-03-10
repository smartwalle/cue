package main

import (
	"fmt"
	"github.com/smartwalle/cue"
	"os"
)

func main() {
	var sheet, err = cue.Decode("./CDImage.cue")
	if err != nil {
		fmt.Println(err)
		return
	}

	var f, _ = os.Create("./out.cue")
	sheet.WriteTo(cue.NewWriter(f))
	f.Close()
}
