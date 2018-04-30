#!/usr/bin/env python
import time
import threading


def fib(n):
    if n <= 1:
        return n
    return fib(n - 1) + fib(n - 2)


num_of_threads = 600

#print(fib(35))


def main():
    for fib_num in range(1, 20):
    #for num_of_threads in range(1, 20):
        # Thread Creation #
        threads = [threading.Thread(target=fib, args=(fib_num,)) for _ in range(num_of_threads)]
        #threads = [threading.Thread(target=fib, args=(35,)) for _ in range(num_of_threads)]

        # Timer Start #
        start_time = time.time()

        # Thread Start #
        [thread.start() for thread in threads]

        # Thread Wait#
        [thread.join() for thread in threads]

        end_time = time.time()

        print('Number Of Threads: ' + str(num_of_threads) + ' Fib Number: ' + str(fib_num) + ' Time: ' + str(end_time - start_time))
        #print('Number Of Threads: ' + str(num_of_threads) + ' Fib Number: 35' + ' Time: ' + str(end_time - start_time))


if __name__ == '__main__':
    main()
