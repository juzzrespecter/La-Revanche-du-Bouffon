FROM golang:bookworm AS build

WORKDIR /opt
RUN apt update && \
    apt install -y make && \
    rm -rf /var/lib/apt/lists/*

COPY . .

RUN make build

FROM golang:bookworm

COPY --from=build /opt/bin/ft_otp /opt/ft_otp
ENTRYPOINT ["/opt/ft_otp"]
