FROM golang:1.8-onbuild as builder

FROM scratch
COPY --from=builder /go/bin/app .
CMD ["./app"] 
