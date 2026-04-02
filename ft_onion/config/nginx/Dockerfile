FROM nginx:1.29.7

RUN apt update && \
    apt install sshd socat -y

COPY ./config/nginx/sshd_config /etc/sshd/sshd_config
COPY ./config/nginx/docker-entrypoint.sh /docker-entrypoint.sh

ENTRYPOINT ["/docker-entrypoint.sh"]

