#!/usr/bin/env python

import time
import threading


def countdown(n):
    while n > 0:
        n -= 1


def main():
    threads = []
    for i in range(5):
        t = threading.Thread(target=countdown, args=(100000000,))
        t.daemon = True
        threads.append(t)

    start_time = time.time()

    [t.start() for t in threads]
    [t.join() for t in threads]

    end_time = time.time()
    print(end_time - start_time)


if __name__ == '__main__':
    main()
