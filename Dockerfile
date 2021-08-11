FROM golang:latest

ENV GOPROXY https://goproxy.cn,direct
WORKDIR $GOPATH/src/github.com/iamxuwenjin/blog
COPY . $GOPATH/src/github.com/iamxuwenjin/blog
RUN go build .

EXPOSE 8000
ENTRYPOINT ["./blog"]