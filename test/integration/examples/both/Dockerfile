FROM elixir

ENV UPDATED_ON 2016-04-08

RUN apt-get update

RUN apt-get install -y file curl netcat git

# Install RUST
RUN curl -sSf https://static.rust-lang.org/rustup.sh | sh
RUN which cargo

# Install Elixir deps for running Hound
RUN mix do local.rebar --force
RUN mix local.hex --force
