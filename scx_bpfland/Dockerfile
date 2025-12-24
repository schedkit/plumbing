ARG BUILDARCH
# Build stage using Debian-based Rust image
FROM rust:1.87-slim AS builder

ADD https://github.com/sched-ext/scx.git /scx

# Install necessary build dependencies, including the missing tools for LLVM installation
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

WORKDIR /scx/scheds/rust/scx_bpfland
COPY . .

# No idea why building scx_bpfland requires rustfmt
RUN rustup component add rustfmt

# Build the project using the default glibc target
RUN cargo build -p scx_bpfland --release

# Architecture-specific library stages
FROM alpine AS libs-amd64
COPY --from=builder /usr/lib/x86_64-linux-gnu/libelf.so.1 /usr/lib/libelf.so.1
COPY --from=builder /usr/lib/x86_64-linux-gnu/libz.so.1 /usr/lib/libz.so.1
#COPY --from=builder /usr/lib/x86_64-linux-gnu/libseccomp.so.2 /usr/lib/libseccomp.so.2

FROM alpine AS libs-arm64
COPY --from=builder /usr/lib/aarch64-linux-gnu/libelf.so.1 /usr/lib/libelf.so.1
COPY --from=builder /usr/lib/aarch64-linux-gnu/libz.so.1 /usr/lib/libz.so.1
#COPY --from=builder /usr/lib/aarch64-linux-gnu/libseccomp.so.2 /usr/lib/libseccomp.so.2

FROM libs-${BUILDARCH} AS libs

# Final distroless image
FROM gcr.io/distroless/cc
COPY --from=builder /scx/target/release/scx_bpfland /scx_bpfland
COPY --from=libs /usr/lib/libelf.so.1 /usr/lib/libelf.so.1
COPY --from=libs /usr/lib/libz.so.1 /usr/lib/libz.so.1

CMD ["/scx_bpfland"]

