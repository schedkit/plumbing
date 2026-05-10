FROM alpine
ARG SCX_VERSION

LABEL org.opencontainers.image.title="scx_test"
LABEL org.opencontainers.image.description="No-op container used to exercise schedctl image plumbing without loading a real scheduler."
LABEL org.opencontainers.image.documentation="https://github.com/schedkit/plumbing/tree/main/scx_test"
LABEL org.opencontainers.image.source="https://github.com/schedkit/plumbing/tree/main/scx_test"
LABEL org.opencontainers.image.version="${SCX_VERSION}"
LABEL org.schedkit.scheduler.kernel.min-version="6.12"

CMD ["sleep", "infinity"]
