directions = {'up': 0, 'right': 1, 'down': 2, 'left': 3}

def solvePart1(fileName):
  file = readFile(fileName).split('\n')
  grid = []
  for line in file:
    grid.append(list(line))
  visited, _ = patrole(grid)
  return len(set(visited))

def solvePart2(fileName):
  file = readFile(fileName).split('\n')
  grid = []
  for line in file:
    grid.append(list(line))
  
  count = 0
  total = 0
  for r, row in enumerate(grid):
    for c, col in enumerate(row):
      count += 1
      if grid[r][c] == "#" or grid[r][c] == "^":
        continue
      grid[r][c] = "#"
      _, loop = patrole(grid, True)
      if loop:
        total += 1
      grid[r][c] = col
  return total

def patrole(grid, checkLoop = False):
  rowLength = len(grid)
  colLength = len(grid[0])

  visited = []
  loop = False
  direction = directions['up']
  row, col = guardPosition(grid)

  while row >= 0 and row <= rowLength - 1 and col >= 0 and col <= colLength - 1:
    current = str(row) + str(col)
    if checkLoop:
      current += str(direction)
      if hasVisited(current, visited):
        loop = True
        break

    visited.append(current)

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

  return visited, loop

def guardPosition(grid):
  for r, row in enumerate(grid):
    for c, col in enumerate(row):
      if col == "^":
        return r, c

def hasVisited(current, visited):
  for v in visited:
    if v == current:
      return True
  return False

def readFile(fileName):
   with open(fileName) as f:
      return f.read()

if __name__ == "__main__":
    print('Part 1:', solvePart1('input.txt'))
    print('Part 2:', solvePart2('input.txt'))