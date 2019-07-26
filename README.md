# Install

## i2c

```
$ sudo apt-get install i2c-tools
$ sudo raspi-config # Advanced Options -> I2C
$ sudo i2cdetect -y 1 # try sometimes if not detected
$ sudo gpasswd -a $USERNAME i2c
```


## go

go(1.11~)

``` .zshrc etc...
GO111MODULE=on
```

```
$ go mod init
$ go mod tidy
$ make build
```

