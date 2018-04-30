#!/usr/bin/env bash

echo pypy
pypy3 ./1.py

echo grumpy
cd ../grumpy
cat ../fib_example/1.py | make run
