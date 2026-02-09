#!/usr/bin/env python3

from collections import defaultdict
from itertools import combinations

antennas = defaultdict(list)

def is_antinode(p):
    for ant in antennas.values():
        for a, b in combinations(ant, 2):
            if (b[1] - a[1]) * (p[0] - a[0]) == (b[0] - a[0]) * (p[1] - a[1]):
                return True
    return False

answer = 0
with open('input.txt', 'r') as file:
    for max_y, line in enumerate(file):
        for max_x, ch in enumerate(line.strip()):
            if ch !='.':
                antennas[ch].append((max_x, max_y))

answer = sum([is_antinode((ax, ay)) for ay in range(max_y+1) for ax in range(max_x+1)])

print("Answer; ", answer)
