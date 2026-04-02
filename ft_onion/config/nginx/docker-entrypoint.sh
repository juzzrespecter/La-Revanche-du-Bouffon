#!/bin/bash
SSH_UNIX_SOCKET=/var/run/tor/sockets/onion-ssh.sock

if ! [ -f "$SSH_UNIX_SOCKET" ]; then
    socat TCP-LISTEN:4242,reuseaddr,fork TCP:127.0.0.1:4242 &
fi
systemctl sshd start
nginx -g "daemon off;"
