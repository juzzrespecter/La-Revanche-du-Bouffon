ARG DISTRO=bullseye
FROM debian:$DISTRO-slim

ARG DISTRO
ENV DISTRO=$DISTRO
RUN <<EOF
#!/bin/bash
SUPPORTED_ARCHS=(amd64 arm64)

ARCH=$(dpkg --print-architecture)
if ! [[ ${SUPPORTED_ARCHS[*]} =~ "$ARCH" ]]
then
    echo "Not supported: $ARCH"
    exit 1
fi

apt update 
apt install apt-transport-https gnupg wget -y
EOF

COPY <<EOF /etc/apt/sources.list.d/tor.sources
Types: deb deb-src
URIs: https://deb.torproject.org/torproject.org/
Suites: ${DISTRO}
Components: main
Signed-By: /usr/share/keyrings/deb.torproject.org-keyring.gpg
EOF

RUN wget -qO- https://deb.torproject.org/torproject.org/A3C4F0F979CAA22CDBA8F512EE8CBC9E886DDD89.asc |\
    gpg --dearmor |\
    tee /usr/share/keyrings/deb.torproject.org-keyring.gpg >/dev/null
RUN apt update && \
    apt install tor deb.torproject.org-keyring

COPY ./config/tor/docker-entrypoint.sh /

CMD ["/docker-entrypoint.sh"]
