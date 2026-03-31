FROM golang:bookworm AS build

WORKDIR /opt
RUN apt update && \
    apt install -y make && \
    rm -rf /var/lib/apt/lists/*

COPY . .

RUN make build

FROM debian:bullseye-slim

RUN apt update && \
    apt install -y oathtool xxd && \
    rm -rf /var/lib/apt/lists/*

COPY --from=build /opt/bin/ft_otp /usr/bin/ft_otp
ENTRYPOINT ["ft_otp"]
