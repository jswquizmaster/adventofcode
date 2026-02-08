#!/usr/bin/env python3

area = []
with open("input.txt", 'r') as file:
    for line in file:
        area.append(list(line.strip()))

# Start position and direction of the guard
x = 74
y = 69
dir = (-1, 0)
answer = 0       

def turn(dir):
    match dir:
        case (-1, 0):
            return (0, 1)
        case (0, 1):
            return (1, 0)
        case (1, 0):
            return (0, -1)
        case (0, -1):
            return (-1, 0)
    
while True:
    try:
        area[y][x] = 'X'
        # Since in python negative indices do not raise an exception
        # we need to check for negative values here explicitly
        if y+dir[0] < 0 or x+dir[1] < 0:
            break
        if area[y+dir[0]][x+dir[1]] == '#':
            dir = turn(dir)
            continue
        y = y + dir[0]
        x = x + dir[1]
    except IndexError:
       break

for a in area:
    answer = answer + a.count('X')

print("Answer; ", answer)
