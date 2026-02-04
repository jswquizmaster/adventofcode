def is_safe(levels, **kwargs):
    if len(levels) <= 1:
        return True
    asc = kwargs.get('asc', levels[0] < levels[1])
    id = levels.pop(0)
    if id == levels[0] or (id < levels[0]) != asc or abs(id - levels[0]) > 3:
        return False
    return is_safe(levels, asc=asc)
    
with open("input.txt", 'r') as file:
    lines = file.readlines()

answer = 0
for line in lines:
    report = list(map(int, line.split()))
    if is_safe(report):
        answer = answer + 1
    
print("Answer; ", answer)
