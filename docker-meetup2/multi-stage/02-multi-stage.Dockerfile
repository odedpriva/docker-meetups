FROM golang:1.8-onbuild as builder

FROM alpine
COPY --from=builder /go/bin/app .
CMD ["./app"] 
