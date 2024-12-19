package parsing

import (
	"log"
	"net"
	"strconv"
	"strings"

	"github.com/cduffaut/Port-Scanner-Go/utils"
)

// return TRUE if ip is OK, FALSE if NOT
func ParseIp(ip string) bool {
	return net.ParseIP(ip) != nil
}

// return TRUE if hostname is OK, FALSE if NOT
func ParseHostname(hostname string) bool {
	// conversion of the domain name into IP format (DNS conversion
	if _, err := net.LookupHost(hostname); err == nil {
		return true
	}
	return false
}

// return true if range is in ascending order, false if not
func is_range_is_in_order(start, end int) bool {
	result := end - start
	return result >= 0
}

// create and filled a struct with starting end ending range values
func range_in_struct(tab_range []string) utils.PortRange {
	start, err := strconv.Atoi(tab_range[0])
	if err != nil {
		log.Fatal("error: range is not only made of digits: ", tab_range[0])
	}
	end, err := strconv.Atoi(tab_range[1])
	if err != nil {
		log.Fatal("error: range is not only made of digits: ", tab_range[1])
	}
	if start < 1 || end < 1 {
		log.Fatal("error: port scanning cannot be under 1.")
	}
	if is_range_is_in_order(start, end) {
		var struct_range utils.PortRange = utils.PortRange{
			Start: start,
			End:   end,
		}
		return struct_range
	}
	return utils.PortRange{}
}

func parse_range(port_range string) utils.PortRange {
	if port_range != "" {
		port_range = strings.TrimSpace(port_range)
		tab_range := strings.Split(port_range, "-")
		if len(tab_range) > 2 {
			log.Fatal("error: range should be only 2 numbers separate by one \"-\".")
		}
		struct_range := range_in_struct(tab_range)
		return struct_range
	} else {
		return utils.PortRange{} // empty struct given return if no range given
	}
}

// check if doublons are presents in the list given, true=yes, false=no
func is_nbr_already_in_list(port_list []int, nbr int) bool {
	for _, ele := range port_list {
		if ele == nbr {
			return true
		}
	}
	return false
}

// check if doublons are presents between the list and the range given, true=yes, false=no
func is_number_in_range(port_range utils.PortRange, nbr int) bool {
	for i := port_range.Start; i <= port_range.End; i++ {
		if i == nbr {
			return true
		}
	}
	return false
}

// check if each strings of the tabs is only digits made, true=yes, false=no
func only_digits(p_list []string, p_range utils.PortRange) []int {
	var digits_list []int

	for _, ele := range p_list {
		ele = strings.TrimSpace(ele)
		nbr, err := strconv.Atoi(ele)
		if err != nil {
			log.Fatal("error: list of ports is not only made of digits... : ", ele)
		}
		if nbr < 1 {
			log.Fatal("error: port number cannot be under 1: ", nbr)
		}
		if !is_number_in_range(p_range, nbr) && !is_nbr_already_in_list(digits_list, nbr) {
			digits_list = append(digits_list, nbr)
		}
	}

	return digits_list
}

// parse the port list to scan: no doublons...
func parse_ports(port_list string, port_range utils.PortRange) []int {
	if port_list != "" {
		tab_list := strings.Split(port_list, ",")
		digits_list := only_digits(tab_list, port_range)
		if digits_list != nil {
			return digits_list
		}
	}
	return nil
}

// stop the program if number of port to scan is > 1024
func max_port_to_scan(bag *utils.Bag) {
	var total_range int

	if bag.IsRange {
		total_range = (bag.PortRange.End - bag.PortRange.Start) + 1
	}
	total_to_scan := total_range + len(bag.PortList)

	if total_to_scan > 1024 {
		log.Fatal("error: cannot scan more than 1024 port in total\nNumber of port asked: ", total_to_scan)
	}
}

func ParseEnv(bag *utils.Bag) {
	bag.PortRange = parse_range(bag.VarEnv.FOR_RANGE)
	bag.PortList = parse_ports(bag.VarEnv.FOR_PORTS, bag.PortRange)
	if bag.PortRange.Start >= 1 {
		bag.IsRange = true
	}
	if len(bag.PortList) > 0 {
		bag.IsList = true
	}
	max_port_to_scan(bag)
	if !bag.IsList && !bag.IsRange {
		bag.PortRange = utils.PortRange{
			Start: 1,
			End:   1024,
		}
		bag.IsRange = true
	}
}
