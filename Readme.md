# Golang Mock Server
This utility can be used to mock apis and respond with xml/json response.

## Setup
This service runs on go.

- Checkout the code and build the project:

```

git clone https://github.com/bhambri94/mockit.git

cd mockit/

docker build -t mockit:latest .

docker images

docker run -it --rm -p 8010:8010 -v $PWD/src:/go/src/mockit mockit

```
