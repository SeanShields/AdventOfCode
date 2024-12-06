def solvePart1(fileName):
  file = readFile(fileName).split('\n')
  grid = []
  for line in file:
    grid.append(list(line))

  total = 0
  for rowIndex, row in enumerate(grid):
    for columnIndex, _ in enumerate(row):
      total += check(rowIndex, columnIndex, "XMAS", grid)
  return total

def solvePart2(fileName):
  file = readFile(fileName)
  return file

def resetCache(rowIndex, colIndex):
  global textCache, rowCache, colCache
  textCache = ""
  rowCache = rowIndex
  colCache = colIndex

def check(rowIndex, colIndex, text, grid):
  global textCache, rowCache, colCache
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

def readFile(fileName):
   with open(fileName) as f:
      return f.read()

if __name__ == "__main__":
    print('Part 1:', solvePart1('input.txt'))
    # print('Part 2:', solvePart2('test.txt'))