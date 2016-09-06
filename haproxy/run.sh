#!/bin/bash

docker run --net=host -v $(pwd):/usr/local/etc/haproxy haproxy
