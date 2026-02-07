#!/usr/bin/env python3

import re

with open("input.txt", 'r') as file:
    content = file.read()

matches = re.findall(r"mul\((\d{1,3}),(\d{1,3})\)", content)
int_matches = [(int(x), int(y)) for x, y in matches]

answer = sum(x * y for x, y in int_matches)

print(answer)
