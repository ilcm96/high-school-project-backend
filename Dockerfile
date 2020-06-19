FROM  golang:1.14-buster as builder

WORKDIR /tmp

RUN git clone https://github.com/ilcm96/high-school-project-backend --depth 1

WORKDIR /tmp/high-school-project-backend
RUN go get -u -d -v ./...
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags '-s -w' -o main main.go \
    && chmod +x upx \
    && ./upx --lzma main

FROM scratch
COPY --from=builder /tmp/high-school-project-backend /
CMD ["/main"]
