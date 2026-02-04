leftList = []
rightList = []

with open('input.txt', 'r') as file:
    for line in file:
        (left, right) = line.split()
        leftList.append(int(left))
        rightList.append(int(right))
    leftList.sort()
    rightList.sort()
    answer = sum(map(lambda x, y: abs(x - y), leftList, rightList))
    
    print("Answer; ", answer)
