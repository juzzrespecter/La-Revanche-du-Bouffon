FROM golang:1.26 AS build

WORKDIR /app

RUN apt update && apt install make -y
COPY . .

RUN make build

FROM alpine:latest

RUN apt update && apt install -y git perl && \
    

COPY --from=build /app/bin/spider /usr/local/bin/spider
COPY --from=build /app/bin/scorpion /usr/local/bin/scorpion

CMD ["spider"]