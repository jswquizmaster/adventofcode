from shapely.geometry import Polygon, box

with open("input.txt", 'r') as file:
    lines = file.readlines()

redTiles = []

for line in lines:
    x, y = line.split(',')
    redTiles.append( (int(x), int(y)) )

# Define polygon
polygon = Polygon(redTiles)

maxArea = 0
for i in range(0, len(redTiles)):
    for j in range(0, i):
        minx = min(redTiles[i][0], redTiles[j][0])
        miny = min(redTiles[i][1], redTiles[j][1])
        maxx = max(redTiles[i][0], redTiles[j][0])
        maxy = max(redTiles[i][1], redTiles[j][1])
        
        # Define rectangle
        rect_coords = (minx, miny, maxx, maxy)
        rectangle = box(*rect_coords)

        # Check if the rectangle is within the polygon
        if polygon.contains(rectangle):
            area = (maxx-minx+1) * (maxy-miny+1)
            if area > maxArea:
                maxArea = area

print("Answer: ", maxArea)
