#!/usr/bin/env python3

ordering = []
answer = 0

def check(update):
    for order in ordering:
        try:
            if update.index(order[0]) > update.index(order[1]):
                return 0 
        except ValueError:
            continue
    return(update[len(update) // 2])
        
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