FROM golang:alpine as builder
RUN apk add git
RUN mkdir /build
ADD . /build/
WORKDIR /build
RUN go get -u github.com/gorilla/mux
RUN go get -u gopkg.in/ini.v1
RUN go get -u github.com/brotherpowers/ipsubnet
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o main .
FROM scratch
COPY --from=builder /build/main /app/
WORKDIR /app
CMD ["./main"]
