import os
from time import perf_counter
import json


def main():

    os.system("python extraction/schedule.py")
    os.system("python extraction/scraper.py")

if __name__ == '__main__':
    start = perf_counter()
    main()
    print('\n\nExecution time {} seconds'.format(round(perf_counter() - start, 2)))
