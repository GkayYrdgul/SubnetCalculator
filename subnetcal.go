package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	. "subnetcal/JsonModel"
	. "subnetcal/SubnetFuncs"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage:\n  %s [ipAddr] [SubnetValue]", os.Args[0])
		flag.PrintDefaults()
		os.Exit(2)
	}
	flag.Parse()
	args := flag.Args()
	if len(args) != 2 {
		fmt.Println("Usage:\n./subnetcal [ipAddr] [SubnetValue]")
		os.Exit(1)
	}

	IpaddrString := os.Args[1]
	MaskValueStr := os.Args[2]
	IpOctetsStr := strings.Split(IpaddrString, ".")
	IpOCtetsInt := ConvertToInt(IpOctetsStr)
	MaskValueInt, _ := strconv.Atoi(MaskValueStr)
	SubnetAddrOctetsInt, _ := SubnetMaskAddrMaker(MaskValueInt)
	SubnetWildcard := Not(SubnetAddrOctetsInt)
	NetworkAddrOctets := AndNow(IpOCtetsInt, SubnetAddrOctetsInt)
	BroadCastAddr := BroadCastAddrMaker(NetworkAddrOctets, SubnetWildcard)
	Hostlen := Hosts(MaskValueInt)

	JsonFile := Jsonfile{}
	JsonFile.IpAddr = IpaddrString
	JsonFile.SubnetValue = MaskValueInt
	JsonFile.SubnetMaskAddr = SliceTOAddr(SubnetAddrOctetsInt)
	JsonFile.SubnetWildcard = SliceTOAddr(SubnetWildcard)
	JsonFile.NetworkAddr = SliceTOAddr(NetworkAddrOctets)
	JsonFile.BroadCastAdrr = SliceTOAddr(BroadCastAddr)
	JsonFile.Hosts = int(Hostlen)

	jsondata, _ := json.Marshal(JsonFile)
	fmt.Printf("%s", jsondata)
}
