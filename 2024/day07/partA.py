#!/usr/bin/env python3

from itertools import product
from operator import add, mul

operators = [add, mul]
answer = 0

with open("input.txt", 'r') as file:
    for line in file:
        t_value, equation = line.split(':')
        t_value = int(t_value)
        eq = [int(x) for x in equation.split()]
        first = eq.pop(0)
        for ops in product(operators, repeat=len(eq)):
            r = first
            for op, a in zip(ops, eq):
                r = op(r, a)
            if r == t_value:
                answer = answer + t_value
                break

print("Answer; ", answer)