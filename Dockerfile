FROM golang:1.9 as builder
WORKDIR /go/src/github.com/smarthut/agent
COPY . .
RUN make vendor
RUN make build

FROM scratch
COPY --from=builder /go/src/github.com/smarthut/agent/agent /
EXPOSE 8080
ENTRYPOINT ["/agent"]
