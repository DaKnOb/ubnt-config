package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	/* Command Line Arguments needed to generate output */
	netname := flag.String("d", "Office", "The DHCP Name")
	subname := flag.String("s", "10.0.0.0/24", "The IPv4 Subnet")
	sfre := flag.Int("f", 10, "Amount of entries between sleeps")
	sltim := flag.Float64("t", 1.0, "Time to sleep")

	/* Parse flags */
	flag.Parse()

	/* Save flag values to strings / int / float */
	dhcpname := *netname
	subnet := *subname
	entriesSleep := *sfre
	sleepTime := *sltim

	/* Get the CSV from stdin */
	r := bufio.NewScanner(os.Stdin)

	/* Used to add sleep every n entries */
	entries := 0

	/* Output configure just in case the user didn't */
	fmt.Println("configure")

	/* Process every CSV line from stdin */
	for r.Scan() {
		line := r.Text()

		/* CSV is formatted MAC, IP, Hostname */
		mac := strings.Split(line, ",")[0]
		sip := strings.Split(line, ",")[1]
		hname := strings.Split(line, ",")[2]

		/* If hostname isn't present, generate one based on the IP */
		if hname == "" {
			hname = dhcpname + "-"
			hname += strings.Split(sip, ".")[2]
			hname += "-"
			hname += strings.Split(sip, ".")[3]
		}

		/* Ensure hostname charset is valid */
		hname = sanitizeHostname(hname)

		/* Output valid EdgeMAX (Vyatta) commands to stdout */
		fmt.Println("set service dhcp-server shared-network-name", dhcpname, "subnet", subnet, "static-mapping", hname, "ip-address", sip)
		fmt.Println("set service dhcp-server shared-network-name", dhcpname, "subnet", subnet, "static-mapping", hname, "mac-address", mac)

		entries++

		/* Sleep every n entries */
		if entries%entriesSleep == 0 {
			fmt.Println("sleep", sleepTime)
		}
	}

}

/*
sanitizeHostname gets the desired hostname and then strips it of
characters that may cause problems..
*/
func sanitizeHostname(hname string) string {
	ret := ""
	for c := range hname {
		if (hname[c] < 48) || (hname[c] > 57 && hname[c] < 65) || (hname[c] > 90 && hname[c] < 97) || (hname[c] > 122) {
			ret += "-"
		} else {
			ret += string(hname[c])
		}
	}
	return ret
}
