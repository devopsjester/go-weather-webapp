# Language: dockerfile
FROM golang:1.18

# Install additional packages if needed.
RUN apt-get update && apt-get install -y git

# Create a non-root user for Codespaces.
RUN useradd -m vscode
USER vscode

# Set the working directory.
WORKDIR /workspace