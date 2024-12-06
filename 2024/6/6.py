directions = {'up': 0, 'right': 1, 'down': 2, 'left': 3}

def solvePart1(fileName):
  file = readFile(fileName).split('\n')
  grid = []
  for line in file:
    grid.append(list(line))
  visited = patrole(grid)
  return len(visited)

def patrole(grid, returnDirection = False):
  rowLength = len(grid)
  colLength = len(grid[0])

  visited = set()
  direction = directions['up']
  row, col = guardPosition(grid)
  while row >= 0 and row <= rowLength - 1 and col >= 0 and col <= colLength - 1:
    current = str(row) + "|" + str(col)
    if returnDirection:
      current + str(direction)
    visited.add(current)
    if direction == directions['up']:
      if row > 0 and grid[row - 1][col] == "#":
        direction += 1
        continue
      row -= 1
    elif direction == directions['right']:
      if col < colLength - 1 and grid[row][col + 1] == "#":
        direction += 1
        continue
      col += 1
    elif direction == directions['down']:
      if row < rowLength - 1 and grid[row + 1][col] == "#":
        direction += 1
        continue
      row += 1
    elif direction == directions['left']:
      if col > 0 and grid[row][col - 1] == "#":
        direction = directions['up']
        continue
      col -= 1
  return visited

def guardPosition(grid):
  for r, row in enumerate(grid):
    for c, col in enumerate(row):
      if col == "^":
        return r, c

def solvePart2(fileName):
  file = readFile(fileName).split('\n')
  grid = []
  for line in file:
    grid.append(list(line))
  visited = patrole(grid)
  return len(visited)

def readFile(fileName):
   with open(fileName) as f:
      return f.read()

if __name__ == "__main__":
    print('Part 1:', solvePart1('input.txt'))
    print('Part 2:', solvePart2('input.txt'))