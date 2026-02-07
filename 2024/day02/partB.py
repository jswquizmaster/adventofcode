#!/usr/bin/env python3

from itertools import pairwise

def check(levels):
    if check.asc == None:
        check.asc = (levels[0] < levels[1])
    return levels[0] != levels[1] and (levels[0] < levels[1]) == check.asc and abs(levels[0] - levels[1]) <= 3
    
with open("input.txt", 'r') as file:
    lines = file.readlines()

answer = 0
for line in lines:
    report = [int(id) for id in line.split()]
    for i in range(len(report)+1):
        check.asc = None
        if all(check(levels) for levels in pairwise(report[:i] + report[i+1:])):
            answer = answer + 1
            break
    
print("Answer; ", answer)
