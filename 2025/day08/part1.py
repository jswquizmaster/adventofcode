import math
from igraph import *

num_connection = 1000
answer = 1
graph = Graph()

with open("input.txt", 'r') as file:
    lines = file.readlines()

for line in lines:
  x, y, z = line.split(',')
  graph.add_vertex(coords=(int(x),int(y),int(z)))

distances = {}
for i in range(len(graph.vs)):
  for j in range(i):
    dist = math.dist(graph.vs[i]["coords"], graph.vs[j]["coords"])
    distances[dist] = (graph.vs[i], graph.vs[j])

count = 0
for vertexes in dict(sorted(distances.items(), key=lambda item: item[0])).values():
  graph.add_edge(vertexes[0], vertexes[1])
  count = count + 1
  if count == num_connection:
    break

components = sorted(graph.connected_components(), key=len, reverse=True)
for i in range(3):
  answer = answer * len(components[i])

print("Answer; ", answer)
  