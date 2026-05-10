FROM ghcr.io/gthulhu/gthulhu:latest AS builder
FROM alpine:3.22.2
ARG SCX_VERSION
COPY --from=builder /gthulhu/main /main

LABEL org.opencontainers.image.title="gthulhu"
LABEL org.opencontainers.image.description="gthulhu sched_ext scheduler."
LABEL org.opencontainers.image.documentation="https://github.com/gthulhu/gthulhu"
LABEL org.opencontainers.image.source="https://github.com/schedkit/plumbing/tree/main/gthulhu"
LABEL org.opencontainers.image.version="${SCX_VERSION}"
LABEL org.schedkit.scheduler.kernel.min-version="6.12"

ENTRYPOINT ["/main"]
