#!/usr/bin/env python

import time


def countdown(n):
    while n > 0:
        n -= 1


def main():
    start_time = time.time()

    # 100 Million
    countdown(100000000)
    countdown(100000000)
    countdown(100000000)
    countdown(100000000)
    countdown(100000000)

    end_time = time.time()

    print(end_time - start_time)


if __name__ == '__main__':
    main()
