FROM  golang:1.14-buster as builder

WORKDIR /tmp

RUN git clone https://github.com/ilcm96/high-school-project-backend --depth 1

WORKDIR /tmp/programming-club-project-backend

RUN go get -u -d -v ./...

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags '-s' -o main main.go

FROM scratch
COPY --from=builder /tmp/high-school-project-backend/main /
CMD ["/main"]
