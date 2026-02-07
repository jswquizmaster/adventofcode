#!/usr/bin/env python3

from collections import Counter

leftList = []
rightList = []

with open('input.txt', 'r') as file:
    for line in file:
        (left, right) = line.split()
        leftList.append(int(left))
        rightList.append(int(right))
    count = Counter(rightList)
    answer = 0
    for id in leftList:
        answer = answer + (id * count[id])
    
    print("Answer; ", answer)
    