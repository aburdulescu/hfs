# hfs
HTTP file server

## Install

Install golang and then run:

```
go get -u github.com/aburdulescu/hfs
```

If you want a smaller binary:

```
go get -u -ldflags="-s -w" github.com/aburdulescu/hfs
```

## Usage

```
Usage of hfs:
  -addr string
        listen address(address:port) (default "localhost:12345")
  -dir string
        path to directory which contains files (default ".")
```
