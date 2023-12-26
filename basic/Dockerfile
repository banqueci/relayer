FROM golang:1.18

WORKDIR /go/src/app
COPY ./ .

RUN go env -w GOPROXY=https://goproxy.oneitfarm.com,https://goproxy.cn,direct
RUN go get -d -v ./...
RUN go install -v ./...
RUN cd basic && make
