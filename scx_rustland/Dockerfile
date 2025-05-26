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
    && curl -fsSL https://apt.llvm.org/llvm.sh | bash -s -- 18 \
    && rm -rf /var/lib/apt/lists/*

RUN update-alternatives --install /usr/bin/clang clang /usr/bin/clang-18 100 \
    && update-alternatives --install /usr/bin/clang++ clang++ /usr/bin/clang++-18 100

WORKDIR /scx/scheds/rust/scx_rustland
COPY . .

# No idea why building scx_rustland requires rustfmt
RUN rustup component add rustfmt

# Build the project using the default glibc target
RUN cargo build -p scx_rustland --release

# Final distroless image
FROM gcr.io/distroless/cc
COPY --from=builder /scx/target/release/scx_rustland /scx_rustland
COPY --from=builder /usr/lib/x86_64-linux-gnu/libelf.so.1 /usr/lib/libelf.so.1
COPY --from=builder /usr/lib/x86_64-linux-gnu/libz.so.1 /usr/lib/libz.so.1
COPY --from=builder /usr/lib/x86_64-linux-gnu/libseccomp.so.2 /usr/lib/libseccomp.so.2


CMD ["/scx_rustland"]
