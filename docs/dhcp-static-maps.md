# DHCP Static Maps Importer
A tool to import DHCP Address Reservations from a CSV file

## What is this?
This is a tool that will read a CSV file with MAC Addresses, IP Addresses,
as well as optionally a hostname, and will add DHCP Address reservations to
the EdgeRouter. This tool was created in order to quickly add over 300 hosts
to an EdgeRouter Pro.

## Anything I should know?
Yes. Copying the output of this tool and pasting to the EdgeRouter shell
for a large amount of hosts (>200) caused some errors that I was not willing
to debug. Pasting it in smaller quantities, say, two hosts at a time, was
perfectly fine. But do not worry, I have added a `sleep` functionality which
will add a `sleep` command every now and then and will remove issues.

## Where is the code?
You can find the code in the [src/dhcp-static-maps](../src/dhcp-static-maps)
folder. I will try to make it available in binary format as well as source
in case you do not want to compile it yourself.

## How do I use it?
By running `./dhcp-static-maps -h` you get a help message. Do you need more?

Just kidding.. Here is a short / full tutorial on how to use the tool in its
current state:

First, you need to create the DHCP Server. Make sure its name does not contain
any spaces since this tool will not work. Use a `-` instead of spaces. Add
all the information in that such as subnet, starting and ending IPs, etc.

When you are done, find a CSV file you have all the reservations you want to
add. Currently only one format is supported and that is
`MAC Address, IP Address, Hostname`. The hostname is optional and may be blank,
however if your CSV only has two rows, the tool will not work. It is okay if
some hosts have a hostname and some do not, or if all do, or none does. In
the future it is planned to support multiple formats, defined by the user.
If you use Google Sheets / Numbers / Excel for this you should be able to
export the list to CSV pretty easily.

Now you can execute the program. There are some command line arguments you
need to know of:

* `-d`: This is **required** and must contain the name of the DHCP Server,
such as `Office`.
* `-s`: This is **required** and must contain the IPv4 subnet, such as
`192.0.2.0/24`.
* `-f`: This is an integer, which is every how many entries the tool should
call `sleep` in the EdgeRouter. The default is set to `10`.
* `-t`: This is a number (with decimals) that defines the length of the pause,
the default being one second (1).

You can now execute the program like so:

```bash
cat /path/to/reservations.csv | ./dhcp-static-maps -d "Home" -s "192.168.1.0/24"
```
