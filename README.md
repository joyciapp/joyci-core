# joyci-core
Core of JoyCI

# Tests

To run the test suite runs:
```
$ go test ./... -v
```

# Docker

## Build a new image
```
$ docker build -t joyciapp/joyci-core:0.0.1 .
```

## Run the built image
```
$ docker run --rm -it joyciapp/joyci-core:0.0.1
```

# Releases

To release a new version:
```
$ git tag -a vx.x.x
$ git push origin vx.x.x
```