# Cast Control

Cast control is a small command line app to expose volume control of Google Cast devices to other networks.
The Python server runs on the same network as the Cast device (for example, on a Raspberry Pi). The Golang client sends requests to the server over gRPC to raise, lower, and set the volume. The server receives these commands and forwards them to the Cast device.

## Quickstart

Run the server daemon on localhost:
```
~/server$ docker build -t xairos/cast-control .
~/server$ docker run -it --rm -network host xairos/cast-control
```

Now you are ready to configure and use the client!

[![asciicast](https://asciinema.org/a/ZLdZEVABaGUzbL0oR0POiROEu.svg)](https://asciinema.org/a/ZLdZEVABaGUzbL0oR0POiROEu)

### **Note:** Running the server on non-Linux platforms
The server uses [zeroconf](https://en.wikipedia.org/wiki/Zero-configuration_networking) to discover Cast devices on the local network. A Dockerfile is provided for deploying the server as a container, with a caveat: zeroconf cannot operate inside Docker, unless the container is on [Docker's `host` network driver](https://docs.docker.com/network/host/) (no network isolation). This only works on Docker for Linux. To run the server in Windows or MacOS, run the Python server directly.
