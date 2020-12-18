import collections
import math
import re
import sys

lines = [l.rstrip('\n') for l in sys.stdin]

grid = {}

for rows, ln in enumerate(lines):
    for cols, c in enumerate(ln):
        grid[(rows,cols,0,0)] = bool(c == '#')

def mink(grid, d):
    return min(k[d] for k in grid.keys())-1

def maxk(grid, d):
    return max(k[d] for k in grid.keys())+1

def cycle(grid):
    new_grid = {}
    for x in range(mink(grid, 0), maxk(grid, 0)+1):
        for y in range(mink(grid, 1), maxk(grid, 1)+1):
            for z in range(mink(grid, 2), maxk(grid, 2)+1):
                for w in range(mink(grid, 3), maxk(grid, 3)+1):
                    count = 0
                    cube = grid.get((x,y,z,w), False)
                    for dx in (-1, 0, 1):
                        for dy in (-1, 0, 1):
                            for dz in (-1, 0, 1):
                                for dw in (-1, 0, 1):
                                    if dz == dy == dx == dw == 0:
                                        continue
                                    target_cube = grid.get((x+dx,y+dy,z+dz,dw+w), False)
                                    if target_cube:
                                        count+=1
                    new_grid[(x,y,z,w)] = (cube and count in (2,3)) or (not cube and count == 3)
    return new_grid

def score(grid):
    return sum(grid.values())
    
for _ in range(6):
    grid = cycle(grid)

print(score(grid))


