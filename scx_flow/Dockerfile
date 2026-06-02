ARG BUILDARCH
FROM rust:1.87-slim AS builder
ARG SCX_VERSION

ADD https://github.com/sched-ext/scx.git#${SCX_VERSION} /scx

RUN apt-get update && apt-get install -y --no-install-recommends \
    build-essential \
    libc6-dev \
    libncurses5-dev \
    libz-dev \
    libelf1 \
    libelf-dev \
    libz1 \
    pkg-config \
    curl \
    gnupg2 \
    lsb-release \
    wget \
    software-properties-common \
    && curl -fsSL https://apt.llvm.org/llvm.sh | bash -s -- 20 \
    && rm -rf /var/lib/apt/lists/*

RUN update-alternatives --install /usr/bin/clang clang /usr/bin/clang-20 100 \
    && update-alternatives --install /usr/bin/clang++ clang++ /usr/bin/clang++-20 100

WORKDIR /scx/scheds/experimental/scx_flow
COPY . .

RUN rustup component add rustfmt

RUN cargo build -p scx_flow --release

FROM alpine AS libs-amd64
COPY --from=builder /usr/lib/x86_64-linux-gnu/libelf.so.1 /usr/lib/libelf.so.1
COPY --from=builder /usr/lib/x86_64-linux-gnu/libz.so.1 /usr/lib/libz.so.1

FROM alpine AS libs-arm64
COPY --from=builder /usr/lib/aarch64-linux-gnu/libelf.so.1 /usr/lib/libelf.so.1
COPY --from=builder /usr/lib/aarch64-linux-gnu/libz.so.1 /usr/lib/libz.so.1

FROM libs-${BUILDARCH} AS libs

FROM gcr.io/distroless/cc
ARG SCX_VERSION
COPY --from=builder /scx/target/release/scx_flow /scx_flow
COPY --from=libs /usr/lib/libelf.so.1 /usr/lib/libelf.so.1
COPY --from=libs /usr/lib/libz.so.1 /usr/lib/libz.so.1

LABEL org.opencontainers.image.title="scx_flow"
LABEL org.opencontainers.image.description="Budget-based sched_ext scheduler with bounded service lanes balancing interactivity and throughput."
LABEL org.opencontainers.image.documentation="https://github.com/sched-ext/scx/tree/main/scheds/experimental/scx_flow"
LABEL org.opencontainers.image.source="https://github.com/schedkit/plumbing/tree/main/scx_flow"
LABEL org.opencontainers.image.version="${SCX_VERSION}"
LABEL org.schedkit.scheduler.kernel.min-version="6.12"

ENTRYPOINT ["/scx_flow"]
