with open("input.txt", 'r') as file:
    inputs = file.readlines()

with open("output.txt", 'r') as file:
    results = file.readlines()

for line in range(len(results)):
    if results[line][0] == "A":
        continue
    elements = inputs[line].split(' ')
    target = [int(item) for item in elements[-1][1:-2].split(',')]
    buttons = elements[1:-1]
    presses = [int(item) for item in results[line].split(' ')[3:-1]]
    for j in range(len(buttons)):
        button = [int(item) for item in buttons[j][1:-1].split(',')]
        for b in button:
            target[b] = target[b] - presses[j]
    if not all(x == 0 for x in target):
        print("Error in line ", line, ": ", target)
        
    

    