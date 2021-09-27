FROM golang:latest AS builder

RUN mkdir /build
WORKDIR /build

RUN GO111MODULE=on
#RUN apt-get update
#ENV GO111MODULE=on \
#    CGO_ENABLED=0 \
#    GOOS=linux \
#    GOARCH=amd64
#WORKDIR /go/src/mutants
#COPY go.mod .
#RUN go mod download
#COPY . .
#RUN go install


#FROM scratch
#COPY --from=builder /go/bin/mutants .
#ENTRYPOINT ["./main"]
