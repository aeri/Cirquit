
<p align="center">
  <img width="350" src="https://i.imgur.com/31oY57g.png">
</p>

# Cirquit
> SNMP client for hub ports management

Cirquit is a small client written in Go for manage the ports of a hub using the SNMP protocol.

Abstraction of complex SNMP control commands and the ability to perform port management automatically and unattended.



## Executing

The management is very simple! The execution is automated and guided

Only build using the Makefile and run it:

```bash
make
make run
```

There is also a precompiled binary that can be used with the following command:
```bash
./Cirquit [ip] [community]
```
