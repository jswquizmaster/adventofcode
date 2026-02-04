import math
from igraph import *

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

for vertexes in dict(sorted(distances.items(), key=lambda item: item[0])).values():
  graph.add_edge(vertexes[0], vertexes[1])
  if len(graph.connected_components()) == 1:
    answer = vertexes[0]["coords"][0] * vertexes[1]["coords"][0]
    break

print("Answer; ", answer)
