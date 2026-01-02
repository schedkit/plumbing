ARG BUILDARCH
FROM rust:1.87-slim AS builder

ADD https://github.com/sched-ext/scx.git /scx

RUN apt-get update && apt-get install -y --no-install-recommends \
    build-essential \
    libseccomp-dev \
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

WORKDIR /scx
COPY . .

RUN rustup component add rustfmt

RUN cargo build --release -p scx_rustland

FROM alpine AS libs-amd64
COPY --from=builder /usr/lib/x86_64-linux-gnu/libelf.so.1 /usr/lib/libelf.so.1
COPY --from=builder /usr/lib/x86_64-linux-gnu/libz.so.1 /usr/lib/libz.so.1
COPY --from=builder /usr/lib/x86_64-linux-gnu/libseccomp.so.2 /usr/lib/libseccomp.so.2

FROM alpine AS libs-arm64
COPY --from=builder /usr/lib/aarch64-linux-gnu/libelf.so.1 /usr/lib/libelf.so.1
COPY --from=builder /usr/lib/aarch64-linux-gnu/libz.so.1 /usr/lib/libz.so.1
COPY --from=builder /usr/lib/aarch64-linux-gnu/libseccomp.so.2 /usr/lib/libseccomp.so.2

FROM libs-${BUILDARCH} AS libs

FROM gcr.io/distroless/cc
COPY --from=builder /scx/target/release/scx_rustland /scx_rustland
COPY --from=libs /usr/lib/libelf.so.1 /usr/lib/libelf.so.1
COPY --from=libs /usr/lib/libz.so.1 /usr/lib/libz.so.1
COPY --from=libs /usr/lib/libseccomp.so.2 /usr/lib/libseccomp.so.2

ENTRYPOINT ["/scx_rustland"]
