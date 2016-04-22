FROM elixir

ENV UPDATED_ON 2016-04-08

RUN apt-get update

RUN apt-get install -y file curl netcat git

# Install RUST
# RUN curl -sSf https://static.rust-lang.org/rustup.sh | sh

# Install Multirust
RUN git clone --recursive https://github.com/brson/multirust 
WORKDIR multirust
RUN git submodule update --init
RUN ./build.sh
RUN ./install.sh
RUN multirust default 1.8

RUN which cargo

# Install Elixir deps for running Hound
RUN mix do local.rebar --force
RUN mix local.hex --force
