# Cast Control

Cast control is a project to expose control over Google Cast devices to other networks, and to perform this control with very low latency.
A Python server daemon lives in a Docker container on the same network as the Cast device (for example, on a Raspberry Pi). 
The Golang client binary connects to the daemon and executes commands against it. The commands are forwarded to the cast device.

## Quickstart

Run the server daemon on localhost:
```
~/server$ docker build -t xairos/cast-control .
~/server$ docker run -d --rm -p 7004:7004 xairos/cast-control
```

Now you are ready to configure and use the client!

[![asciicast](https://asciinema.org/a/ZLdZEVABaGUzbL0oR0POiROEu.svg)](https://asciinema.org/a/ZLdZEVABaGUzbL0oR0POiROEu)
