#!/usr/bin/env python3

area = []
with open("input.txt", 'r') as file:
    for line in file:
        area.append(list(line.strip()))

max_y = 130
max_x = 130
# Since the puzzle input is relatively small we use an upper bound of iterations 
# instead of a accurate loop detection in this case
# The area is 130x130 blocks. Each block can be reached from 4 different directions
# So 130*130*4 is an upper limit of iterations before the guard is stuck in a loop 
max_i = max_x*max_y*4
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

for oy in range(max_y):
    for ox in range (max_x):
        x = 74
        y = 69
        dir = (-1, 0)
        if area[oy][ox] == '^' or area[oy][ox] == '#':
            continue
        area[oy][ox] = '#'
        for i in range(max_i):
            try:
                area[y][x] = 'X'
                if y+dir[0] < 0 or x+dir[1] < 0:
                    break
                if area[y+dir[0]][x+dir[1]] == '#':
                    dir = turn(dir)
                    continue
                y = y + dir[0]
                x = x + dir[1]
            except IndexError:
                break
        # Check if've gone through the hole for loop
        # which means the guard is stuck in a loop
        if i == max_i-1:
            answer = answer + 1
        area[oy][ox] = '.'
    # Since this can take a while, print out status message
    print("%d percent finished, %d positions found so far..." % (oy*100//max_y, answer))


print("Answer; ", answer)
