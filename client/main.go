package main

import (
	"flag"
	"github.com/gaspardpeduzzi/spring_block/data"
	"github.com/gaspardpeduzzi/spring_block/display"
)

func main(){


	var addr = flag.String("addr", "s1.ripple.com:51233", "http service address")
	var verb = flag.Bool("verb", false, "Display more information")
	flag.Parse()
	display.VERBOSE = *verb

	//c := make(chan int)
	lastIndex := data.GetLastLedgerSeq(addr)
	display.DisplayVerbose("last index", lastIndex)


}