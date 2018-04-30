#!/usr/bin/env python
import time
import threading


def fib(n):
    if n <= 1:
        return n
    return fib(n - 1) + fib(n - 2)


print(fib(35))


def main():
    for num_of_threads in range(1, 5):
        # Thread Creation #
        threads = [threading.Thread(target=fib, args=(35,)) for _ in range(num_of_threads)]

        # Timer Start #
        start_time = time.time()

        # Thread Start #
        [thread.start() for thread in threads]

        # Thread Start #
        [thread.join() for thread in threads]

        end_time = time.time()

        print('Number Of Threads: ' + str(num_of_threads), ' Time: ' + str(end_time - start_time))


if __name__ == '__main__':
    main()
