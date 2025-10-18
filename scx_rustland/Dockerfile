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

FROM gcr.io/distroless/cc

ARG TARGETPLATFORM

RUN if [ "$TARGETPLATFORM" = "linux/arm64" ]; then \
      echo "Setting $COPY_LIBS_DIR to aarch64-linux-gnu" && \
      export COPY_LIBS_DIR="aarch64-linux-gnu"; \
    else \
      echo "Setting default COPY_LIBS_DIR x86_64-linux-gnu" && \
      export COPY_LIBS_DIR="x86_64-linux-gnu"; \
    fi

COPY --from=builder /scx/target/release/scx_rustland /scx_rustland
COPY --from=builder /usr/lib/$COPY_LIBS_DIR/libelf.so.1 /usr/lib/libelf.so.1
COPY --from=builder /usr/lib/$COPY_LIBS_DIR/libz.so.1 /usr/lib/libz.so.1
COPY --from=builder /usr/lib/$COPY_LIBS_DIR/libseccomp.so.2 /usr/lib/libseccomp.so.2

CMD ["/scx_rustland"]
