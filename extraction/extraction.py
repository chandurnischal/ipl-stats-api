import os
from time import perf_counter


def main():
    print("Extracting match URLS...")
    os.system("python extraction/schedule.py")

    print("Extracting match data...")
    os.system("python extraction/scraper.py")

    # print("Preprocessing data...")
    # os.system("python extraction/query.py")

if __name__ == '__main__':
    start = perf_counter()
    main()
    print('\n\nExecution time {} seconds'.format(round(perf_counter() - start, 2)))
