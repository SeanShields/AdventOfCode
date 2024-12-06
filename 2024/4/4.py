def solvePart1(fileName):
  file = readFile(fileName).split('\n')
  grid = []
  for line in file:
    grid.append(list(line))

  total = 0
  for rowIndex, row in enumerate(grid):
    for columnIndex, _ in enumerate(row):
      total += checkXMAS(rowIndex, columnIndex, grid)
  return total

def solvePart2(fileName):
  file = readFile(fileName).split('\n')
  grid = []
  for line in file:
    grid.append(list(line))

  total = 0
  for rowIndex, row in enumerate(grid):
    for columnIndex, _ in enumerate(row):
      if checkMASX(rowIndex, columnIndex, grid):
        total += 1
  return total

def checkMASX(rowIndex, colIndex, grid):
  colLength = len(grid[0])
  rowLength = len(grid)

  if grid[rowIndex][colIndex] != "A" \
    or \
      (rowIndex == 0 \
      or colIndex == 0 \
      or rowIndex == rowLength - 1 or colIndex == colLength - 1):
    return False

  topRight = grid[rowIndex - 1][colIndex + 1]
  topLeft = grid[rowIndex - 1][colIndex - 1]
  bottomRight = grid[rowIndex + 1][colIndex + 1]
  bottomLeft = grid[rowIndex + 1][colIndex - 1]

  firstHalf = topLeft == "M" and bottomRight == "S" \
    or topLeft == "S" and bottomRight == "M"
  
  secondHalf = topRight == "M" and bottomLeft == "S" \
    or topRight == "S" and bottomLeft == "M"

  return firstHalf and secondHalf

def checkXMAS(rowIndex, colIndex, grid):
  global textCache, rowCache, colCache
  text = "XMAS"
  textLength = len(text)
  colLength = len(grid[0])
  rowLength = len(grid)
  total = 0

  resetCache(rowIndex, colIndex)

  # up
  for _ in range(textLength):
    textCache += grid[rowCache][colCache]
    if textCache == text:
      total += 1
    elif rowCache == 0:
      break
    rowCache -= 1

  resetCache(rowIndex, colIndex)
  
  # up-right
  for _ in range(textLength):
    textCache += grid[rowCache][colCache]
    if textCache == text:
      total += 1
    elif rowCache == 0 or colCache == colLength - 1:
      break
    rowCache -= 1
    colCache += 1

  resetCache(rowIndex, colIndex)

    # right
  for _ in range(textLength):
    textCache += grid[rowCache][colCache]
    if textCache == text:
      total += 1
    elif colCache == colLength - 1:
      break
    colCache += 1

  resetCache(rowIndex, colIndex)

  # down-right
  for _ in range(textLength):
    textCache += grid[rowCache][colCache]
    if textCache == text:
      total += 1
    elif rowCache == rowLength - 1 or colCache == colLength - 1:
      break
    rowCache += 1
    colCache += 1

  resetCache(rowIndex, colIndex)

  # down
  for _ in range(textLength):
    textCache += grid[rowCache][colCache]
    if textCache == text:
      total += 1
    elif rowCache == rowLength - 1:
      break
    rowCache += 1

  resetCache(rowIndex, colIndex)

  # down-left
  for _ in range(textLength):
    textCache += grid[rowCache][colCache]
    if textCache == text:
      total += 1
    elif rowCache == rowLength - 1 or colCache == 0:
      break
    rowCache += 1
    colCache -= 1

  resetCache(rowIndex, colIndex)

  # left
  for _ in range(textLength):
    textCache += grid[rowCache][colCache]
    if textCache == text:
      total += 1
    elif colCache == 0:
      break
    colCache -= 1

  resetCache(rowIndex, colIndex)

  # up-left
  for _ in range(textLength):
    textCache += grid[rowCache][colCache]
    if textCache == text:
      total += 1
    elif rowCache == 0 or colCache == 0:
      break
    rowCache -= 1
    colCache -= 1

  return total

def resetCache(rowIndex, colIndex):
  global textCache, rowCache, colCache
  textCache = ""
  rowCache = rowIndex
  colCache = colIndex

def readFile(fileName):
   with open(fileName) as f:
      return f.read()

if __name__ == "__main__":
    print('Part 1:', solvePart1('input.txt'))
    print('Part 2:', solvePart2('input.txt'))