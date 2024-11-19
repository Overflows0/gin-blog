FROM golang:latest

ENV GOPROXY=https://goproxy.cn,direct
WORKDIR $GOPATH/gin-blog
COPY . $GOPATH/gin-blog
RUN go build .

EXPOSE 8000
ENTRYPOINT [ "./gin-blog" ]