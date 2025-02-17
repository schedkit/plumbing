# Build stage using Debian-based Rust image
FROM rust:1.84-slim AS builder

# Install necessary build dependencies, including the missing tools for LLVM installation
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
    && curl -fsSL https://apt.llvm.org/llvm.sh | bash - \
    && apt-get install -y --no-install-recommends \
    clang-16 \
    llvm-16 \
    llvm-16-dev \
    && rm -rf /var/lib/apt/lists/*

# Set Clang version 16 as default
RUN update-alternatives --install /usr/bin/clang clang /usr/bin/clang-16 100 \
    && update-alternatives --install /usr/bin/clang++ clang++ /usr/bin/clang++-16 100

WORKDIR /scx_rusty
COPY . .

# Build the project using the default glibc target
RUN cargo build -p scx_rusty --release

# Final distroless image
FROM gcr.io/distroless/cc
COPY --from=builder /scx_rusty/target/release/scx_rusty /scx_rusty
COPY --from=builder /usr/lib/x86_64-linux-gnu/libelf.so.1 /usr/lib/libelf.so.1
COPY --from=builder /usr/lib/x86_64-linux-gnu/libz.so.1 /usr/lib/libz.so.1


CMD ["/scx_rusty"]
