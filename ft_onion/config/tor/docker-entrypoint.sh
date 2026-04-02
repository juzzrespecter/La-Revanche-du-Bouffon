#!/bin/bash

systemctl tor start

# Test logs
ls -lah /var/log/tor
sleep infinity