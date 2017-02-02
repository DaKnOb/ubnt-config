# ubnt-config
A collection of command line tools to quickly configure EdgeMAX devices

## What is this?
The `ubnt-config` project is a collection of command line tools I have created
to interact with EdgeMAX devices, especially EdgeRouters. Although configuring
one is easy and straight forward (for a pro-grade router, not a home router),
some tasks just seem to take a lot of time and can be easily automated.

Luckily, EdgeMAX allows full `ssh` access, which means we can write scripts
that directly interact with the device. However, both for simplicity as well
as safety, the tools in this repository output to the `stdout`, the terminal
they are executed in, and the user can then copy the output and paste it to
the shell of the device either directly or through pipes and/or `pbcopy`.

## What's available?
The following tools are currently available. This list will grow as new tools
are added. For more information about a particular tool, click on it to read
its documentation.

* [DHCP Static Maps Importer](docs/dhcp-static-maps.md)

## License
Apache 2
