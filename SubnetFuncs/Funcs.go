package SubnetFuncs

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

const (
	AddrBytelen  = 35 //32 bit + 3 .
	AddrBitLen   = 32 //32 Bit
	AddrOctetlen = 9  //8 bit + 1.
)

func SubnetMaskAddrMaker(Subnetvalue int) ([]int, string) {
	var (
		SubnetMaskAddr  string
		bit1            = "1"
		bit0            = "0"
		d               = "."
		intSubnetOctets []int
		tempOctet       int
	)

	for i := 1; i <= AddrBytelen; i++ {
		if Subnetvalue != 0 && i%AddrOctetlen != 0 {
			SubnetMaskAddr += bit1
			Subnetvalue--
		} else if i%AddrOctetlen == 0 {
			SubnetMaskAddr += d

		} else {
			SubnetMaskAddr += bit0

		}
	}

	subnetOctets := strings.Split(SubnetMaskAddr, ".")
	for _, t := range subnetOctets {
		fmt.Sscanf(t, "%b", &tempOctet)
		intSubnetOctets = append(intSubnetOctets, tempOctet)
	}

	return intSubnetOctets, SubnetMaskAddr
}

func ConvertToInt(addr []string) []int {
	var (
		intdataArray []int
		intValue     int
	)

	for _, t := range addr {
		intValue, _ = strconv.Atoi(t)
		intdataArray = append(intdataArray, intValue)

	}
	return intdataArray
}

func AndNow(subnetmaskaddr, ipaddr []int) []int {

	var NetworkAdrr []int
	for index := range ipaddr {
		NetworkAdrr = append(NetworkAdrr, ipaddr[index]&subnetmaskaddr[index])
	}
	return NetworkAdrr
}

func Not(Addr []int) []int {
	var NotAdrr []int
	var byteOctets byte
	for _, data := range Addr {
		byteOctets = byte(^data)
		NotAdrr = append(NotAdrr, int(byteOctets))
	}
	return NotAdrr
}

func BroadCastAddrMaker(IpAddr, WisCardAddr []int) []int {

	var BroadCastAddr []int
	for i, _ := range IpAddr {
		BroadCastAddr = append(BroadCastAddr, IpAddr[i]+WisCardAddr[i])
	}
	return BroadCastAddr
}

func Hosts(SubnetValue int) float64 {

	Hosts := math.Pow(2, AddrBitLen-float64(SubnetValue)) - 2

	return Hosts
}

func SliceTOAddr(AddrSlice []int) string {
	var (
		tempString string
		AddrString string
	)
	builder := strings.Builder{}
	for i, data := range AddrSlice {
		tempString = strconv.Itoa(data)
		builder.WriteString(tempString)
		if i < 3 {
			builder.WriteString(".")
		}

	}
	AddrString = builder.String()
	return AddrString
}
