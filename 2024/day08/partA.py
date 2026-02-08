#!/usr/bin/env python3

from collections import defaultdict
from itertools import permutations

def antinode(a1, a2):
    x1, y1 = a1
    xm, ym = a2
    x2 = 2 * xm - x1
    y2 = 2 * ym - y1
    
    return (x2, y2)

antennas = defaultdict(list)
antinodes = []
with open('input.txt', 'r') as file:
    for y, line in enumerate(file):
        for x, ch in enumerate(line.strip()):
            if ch !='.':
                antennas[ch].append((x, y))

for ant in antennas.values():
    for points in permutations(ant, 2):
        an = antinode(points[0], points[1])
        if an[0]>=0 and an[0]<=x and an[1]>=0 and an[1]<=y and not an in antinodes:
            antinodes.append(an)

print("Answer; ", len(antinodes))
