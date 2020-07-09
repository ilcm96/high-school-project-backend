FROM  golang:1.14-buster as builder

WORKDIR /tmp/high-school-auth-backend
COPY . .

RUN go mod tidy \
    && go get -u -d -v ./...
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags '-s -w' -o main cmd/main.go \
    && chmod +x upx \
    && ./upx --lzma main

FROM scratch
COPY --from=builder /tmp/high-school-auth-backend /
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
ENV TZ=Asia/Seoul
CMD ["/main"]
