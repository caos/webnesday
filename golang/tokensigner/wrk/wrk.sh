#! /bin/bash

wrk -c 15 -d 10 -t 15 -s wrk/post.lua http://localhost:8080/verify
