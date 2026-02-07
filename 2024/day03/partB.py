#!/usr/bin/env python3

import re

with open("input.txt", 'r') as file:
    content = file.read()

lines = re.split(r"(?=do(?:n't)?\(\))", content)

answer = 0
for line in lines:
    if line.startswith("don't()"):
        continue
    matches = re.findall(r"mul\((\d{1,3}),(\d{1,3})\)", line)
    int_matches = [(int(x), int(y)) for x, y in matches]
    answer = answer + sum(x * y for x, y in int_matches)

print("Answer; ", answer)
