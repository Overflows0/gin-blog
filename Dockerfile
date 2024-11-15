FROM scratch

WORKDIR $GOPATH/src/gin-blog
COPY . $GOPATH/src/gin-blog


ENTRYPOINT [ "./gin-blog" ]