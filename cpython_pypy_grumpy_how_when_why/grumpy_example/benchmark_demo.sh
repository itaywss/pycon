#!/usr/bin/env bash

echo pypy2
pypy ./dict.py

echo pypy3
pypy3 ./dict3.py

echo python2.7
python2.7 ./dict.py

echo python3.6
python3.6 ./dict3.py

echo grumpy
cd ../grumpy
cat benchmarks/dict.py | make run
