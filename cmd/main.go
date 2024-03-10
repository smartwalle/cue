package main

import (
	"bufio"
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

	var f, _ = os.Create("./xx.txt")
	sheet.WriteTo(bufio.NewWriter(f))
	f.Close()
}
