# sfs
Static files HTTP server

## Install

Install golang and then run:

```
go get -u github.com/aburdulescu/sfs
```

If you want a smaller binary:

```
go get -u -ldflags="-s -w" github.com/aburdulescu/sfs
```

## Usage

```
Usage of sfs:
  -addr string
        listen address(address:port) (default "localhost:12345")
  -dir string
        path to directory which contains files (default ".")
```
