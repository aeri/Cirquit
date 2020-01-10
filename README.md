
<p align="center">
  <img width="350" src="https://i.imgur.com/31oY57g.png">
</p>

# Cirquit
> SNMP client for convenient hub ports management

Cirquit is a small client written in Go for manage the ports of a hub using the SNMP protocol.

Abstraction of complex SNMP control commands and the ability to perform port management automatically and unattended.

## Libraries

This program needs to build this external libraries with:

* **soniah** - *gosnmp* - [GitHub](https://github.com/soniah/gosnmp)
* **fatih** - *color* - [GitHub](https://github.com/fatih/color)

Those libraries can be installed by executing the next commands:

```bash
go get github.com/fatih/color
go get github.com/soniah/gosnmp
```




## Executing

The management is very simple! The execution is automated and guided

There are two ways to execute this program:

* Interactive mode

Just build using the Makefile and run it:

```bash
make
make run
```

* Unattended mode

In order to do an automated and quiet management of your hub, there is a possibility to do this with the following sequence:

```bash
./Cirquit ip community optionMenu [gateway]
```

The second menu option can be used by loading the port configuration from a file ```ports.cfg```, the included file by default has an example of the syntax.

## License
This project is licensed under the GNU General Public License v3.0 - see the [LICENSE](LICENSE) file for details
