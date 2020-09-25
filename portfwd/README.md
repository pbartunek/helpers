## PortFwd

Simple TCP port forwarder


### Usage example

Forward port local 8888 to port 8080 on 192.168.1.1:

```
portfwd -l 127.0.0.1:8888 -f 192.168.1.1:8080
```

Parameters:
```
  -f string
    	forward ip:port (default "127.0.0.1:8080")
  -l string
    	listen ip:port (default "127.0.0.1:1337")
```
