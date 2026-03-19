# This Dockerfile will: 
# set up a Rust environment(using Rust Docker Image)
# and install the Hermes Relayer

# Use Rust Docker Image
FROM rust:1.70.0

# Set working dir
WORKDIR /hermes_relayer_de_hub

#RUN rustup default nightly && rustup update
RUN rustup update

# Install Hermes
RUN cargo install ibc-relayer-cli@1.5.1 --bin hermes --locked

# Copy config file
COPY ./config.toml /root/.hermes/config.toml

# Validate the config file during image building
RUN hermes config validate

# Copy mnemonic files
# COPY ./university_chain_de_keys.json /home/hermes/university_chain_de_keys.json
# COPY ./hub_keys.json /home/hermes/hub_keys.json

# Copy entrypoint script
# COPY ./run-relayer.sh /home/hermes/

# Set permission
# RUN chmod +x ./run-relayer.sh

# Run entrypoint script
# ENTRYPOINT ["/bin/bash", "-c", "/home/hermes/run-relayer.sh"]

# build with
# docker build -f Dockerfile . -t relayer_hermes
