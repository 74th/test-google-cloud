FROM "mcr.microsoft.com/devcontainers/base:bullseye"

RUN apt-get update && apt-get install -y \
    git \
    curl \
    unzip \
    && rm -rf /var/lib/apt/lists/*

ENV PATH="/home/vscode/.tfenv/bin:$PATH"

USER vscode

RUN git clone --depth=1 https://github.com/tfutils/tfenv.git /home/vscode/.tfenv
