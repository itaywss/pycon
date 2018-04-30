#!/usr/bin/env bash
echo pypy
pypy3 ./2.py

echo grumpy
cd ../grumpy
cat ../grumpy_example/2.py | make run
