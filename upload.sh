#!/usr/bin/env bash
scp -r DevUtils swagger static conf restart.sh root@139.199.39.76:~/dev/devUtils
scp -r DevUtils swagger static restart.sh root@192.168.1.201:/home/devUtils