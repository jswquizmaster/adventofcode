#!/usr/bin/env python3

with open("input.txt", 'r') as file:
    p = file.readlines()

max_y = len(p)
max_x = len(p[0])
answer = 0

for y in range(max_y):
    for x in range(max_x):
        if p[y][x] != 'X':
            continue
        if x<max_x-3 and p[y][x+1] == 'M' and p[y][x+2] == 'A' and p[y][x+3] == 'S':
            answer = answer + 1
        if x<max_x-3 and y<max_y-3 and p[y+1][x+1] == 'M' and p[y+2][x+2] == 'A' and p[y+3][x+3] == 'S':
            answer = answer + 1
        if y<max_y-3 and p[y+1][x] == 'M' and p[y+2][x] == 'A' and p[y+3][x] == 'S':
            answer = answer + 1
        if x>=3 and y<max_y-3 and p[y+1][x-1] == 'M' and p[y+2][x-2] == 'A' and p[y+3][x-3] == 'S':
            answer = answer + 1
        if x>=3 and p[y][x-1] == 'M' and p[y][x-2] == 'A' and p[y][x-3] == 'S':
            answer = answer + 1
        if x>=3 and y>=3 and p[y-1][x-1] == 'M' and p[y-2][x-2] == 'A' and p[y-3][x-3] == 'S':
            answer = answer + 1
        if y>=3 and p[y-1][x] == 'M' and p[y-2][x] == 'A' and p[y-3][x] == 'S':
            answer = answer + 1
        if x<max_x-3 and y>=3 and p[y-1][x+1] == 'M' and p[y-2][x+2] == 'A' and p[y-3][x+3] == 'S':
            answer = answer + 1

print("Answer; ", answer)
