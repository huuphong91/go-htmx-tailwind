# Base image
FROM mcr.microsoft.com/devcontainers/base:jammy

# Install dependencies
RUN apt-get update && apt-get install -y \
    curl \
    git \
    build-essential \
    zlib1g-dev \
    libffi-dev \
    libssl-dev \
    libsqlite3-dev \
    libbz2-dev \
    libreadline-dev \
    libncurses5-dev \
    libgdbm-dev \
    liblzma-dev \
    && apt-get clean && rm -rf /var/lib/apt/lists/*

# Define build arguments
ARG VARIANT=1.23.3
ARG NODE_VERSION=22.11.0

# Install NVM
ENV NVM_DIR=/usr/local/nvm
RUN mkdir -p $NVM_DIR \
    && curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.40.0/install.sh | bash \
    && echo 'export NVM_DIR="$NVM_DIR"' >> /etc/bash.bashrc \
    && echo '[ -s "$NVM_DIR/nvm.sh" ] && \. "$NVM_DIR/nvm.sh"' >> /etc/bash.bashrc \
    && echo '[ -s "$NVM_DIR/bash_completion" ] && \. "$NVM_DIR/bash_completion"' >> /etc/bash.bashrc

# Install Node.js using NVM
RUN bash -c "source $NVM_DIR/nvm.sh && nvm install 22 && nvm alias default 22 && nvm use default"

# Ensure Node.js and npm are available globally
ENV PATH=$NVM_DIR/versions/node/v$NODE_VERSION/bin:$PATH

# Verify Node.js and npm installation
RUN node -v && npm -v

# Install Bun
ENV BUN_INSTALL=/home/vscode/.bun
RUN curl -fsSL https://bun.sh/install | bash \
    && echo 'export BUN_INSTALL="$HOME/.bun"' >> /etc/bash.bashrc \
    && echo 'export PATH="$BUN_INSTALL/bin:$PATH"' >> /etc/bash.bashrc

# Ensure Bun is available globally
ENV PATH=/home/vscode/.bun/bin:$PATH

# Verify Bun installation
RUN bun -v

# Install Go
RUN curl -fsSL https://go.dev/dl/go$VARIANT.linux-amd64.tar.gz -o go.tar.gz \
    && tar -C /usr/local -xzf go.tar.gz \
    && rm go.tar.gz \
    && ln -s /usr/local/go/bin/* /usr/local/bin/

# Set Go environment variables
ENV PATH="/usr/local/go/bin:${PATH}"
ENV GOPATH="/go"

RUN mkdir -p /go && chown -R vscode:vscode /go

RUN curl -fsSL https://www.python.org/ftp/python/3.13.1/Python-3.13.1.tgz -o Python-3.13.1.tgz \
    && tar -xzf Python-3.13.1.tgz \
    && cd Python-3.13.1 \
    && ./configure --enable-optimizations --prefix=/usr/local/python3.13 \
    && make -j$(nproc) \
    && make install \
    && cd .. \
    && rm -rf Python-3.13.1 Python-3.13.1.tgz

ENV PATH="/usr/local/python3.13/bin:$PATH"

RUN python3.13 --version && pip3 --version

USER vscode

WORKDIR /workspace