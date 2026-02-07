#!/usr/bin/env python3

with open("input.txt", 'r') as file:
    p = file.readlines()

max_y = len(p)
max_x = len(p[0])
answer = 0

for y in range(1, max_y-1):
    for x in range(1, max_x-1):
        if p[y][x] != 'A':
            continue
        if p[y-1][x-1] == 'M' and p[y-1][x+1] == 'M' and p[y+1][x-1] == 'S' and p[y+1][x+1] == 'S':
            answer = answer + 1
        if p[y-1][x-1] == 'M' and p[y-1][x+1] == 'S' and p[y+1][x-1] == 'M' and p[y+1][x+1] == 'S':
            answer = answer + 1
        if p[y-1][x-1] == 'S' and p[y-1][x+1] == 'M' and p[y+1][x-1] == 'S' and p[y+1][x+1] == 'M':
            answer = answer + 1
        if p[y-1][x-1] == 'S' and p[y-1][x+1] == 'S' and p[y+1][x-1] == 'M' and p[y+1][x+1] == 'M':
            answer = answer + 1

print("Answer; ", answer)
