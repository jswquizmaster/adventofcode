#!/usr/bin/env python3

disk = []

def get_next():
    global disk
    while disk != []:
        c = disk.pop()
        if c != None:
            return c
    return None
        
with open('input.txt', 'r') as file:
    for line in file:
        for i, c in enumerate(line.strip()):
            if i % 2 == 0:
                disk.extend([i//2] * int(c))
            else:
                disk.extend([None] * int(c))
            
answer = 0
i = 0
while disk != []:
    c = disk.pop(0)
    if c == None:
        c = get_next()
    if c == None:
        continue
    answer = answer + (i * c)
    i = i + 1
    
print("Answer; ", answer)
