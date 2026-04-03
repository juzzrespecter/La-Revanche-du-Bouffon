FROM golang:1.26 AS build

WORKDIR /app

RUN apt update && apt install make -y
COPY . .

RUN make build

FROM debian:bookworm-slim

RUN apt update && \
    apt install -y git perl && \
    rm -rf /var/lib/apt/lists/*

WORKDIR /opt

RUN git clone https://github.com/exiftool/exiftool.git && \
    ln -s /opt/exiftool/exiftool /usr/bin/exiftool

COPY --from=build /app/bin/spider /usr/local/bin/spider
COPY --from=build /app/bin/scorpion /usr/local/bin/scorpion

CMD ["spider"]