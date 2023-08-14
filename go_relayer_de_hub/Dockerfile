# This Dockerfile will: 
# set up a Golang environment(using Golang Docker Image)
# and install Go Relayer

# Use Golang Docker Image
#FROM golang:1.18
FROM golang:1.20.5

# Set working dir
WORKDIR /go_relayer_de_hub

# Install git
RUN apt-get install git

# Clone relayer repository and install relayer
# RUN git clone https://github.com/cosmos/relayer.git && cd relayer && git checkout v2.0.0-rc3 && make install
RUN git clone https://github.com/cosmos/relayer.git && cd relayer && make install

# Copy configs
# ADD configs /home/rly/configs

# Copy pahts
# ADD paths /home/rly/paths

# Copy entrypoint script
# COPY ./run-relayer_de_hub.sh /home/rly/

# Set permission
# RUN chmod +x ./run-relayer_de_hub.sh

