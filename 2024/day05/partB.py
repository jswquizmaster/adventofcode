#!/usr/bin/env python3

from functools import cmp_to_key

ordering = []
answer = 0

def check(update):
    def cmp(a, b):
        # Either (a, b) or (b, a) is in the ordering list
        # That's why it is sufficient to only check for (a, b) here 
        if (a, b) in ordering:
            return -1
        else:
            return 1 

    update_sorted = sorted(update, key=cmp_to_key(cmp))
    if update != update_sorted:
        return(update_sorted[len(update) // 2])
    else:
        return 0
        
with open('input.txt', 'r') as file:
    for line in file:
        if line == "\n":
            break
        a, b = line.split('|')
        ordering.append((int(a), int(b)))

    for line in file:
        update = line.split(',')
        update = [int(x) for x in update]
        answer = answer + check(update)

print("Answer; ", answer)