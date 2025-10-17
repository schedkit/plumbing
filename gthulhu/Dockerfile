FROM ghcr.io/gthulhu/gthulhu:latest AS builder
FROM alpine:3.22.2
COPY --from=builder /gthulhu/main /main
CMD ["/main"]