# Base image
FROM mcr.microsoft.com/devcontainers/base:jammy

# Install dependencies
RUN apt-get update && apt-get install -y \
    curl \
    git \
    build-essential \
    && apt-get clean && rm -rf /var/lib/apt/lists/*

# Define build arguments
ARG VARIANT
ARG NODE_VERSION

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

# Install Go
RUN curl -fsSL https://go.dev/dl/go$VARIANT.linux-amd64.tar.gz -o go.tar.gz \
    && tar -C /usr/local -xzf go.tar.gz \
    && rm go.tar.gz \
    && ln -s /usr/local/go/bin/* /usr/local/bin/

# Set Go environment variables
ENV PATH="/usr/local/go/bin:${PATH}"
ENV GOPATH="/go"

RUN mkdir -p /go && chown -R vscode:vscode /go

USER vscode

WORKDIR /workspace