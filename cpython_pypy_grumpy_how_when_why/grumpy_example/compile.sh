#!/usr/bin/env bash

cd ../grumpy
make
export PATH=$PWD/build/bin:$PATH
export GOPATH=$PWD/build
export PYTHONPATH=$PWD/build/lib/python2.7/site-packages
grumpc ../grumpy_example/1.py 
