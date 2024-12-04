def solvePart1(fileName):
  file = readFile(fileName).split('\n')
  grid = []
  for line in file:
    grid.append(list(line))

  total = 0
  for rowIndex, row in enumerate(grid):
    for columnIndex, _ in enumerate(row):
      if check(rowIndex, columnIndex, "XMAS", grid):
        total += 1
  return total

def solvePart2(fileName):
  file = readFile(fileName)
  return file

global textCache
global rowCache
global colCache

def check(rowIndex, colIndex, text, grid):
  textLength = len(text)
  colLength = len(grid[0])
  rowLength = len(grid)

  textCache = ""
  rowCache = rowIndex
  colCache = colIndex

  # up
  for _ in range(textLength):
    textCache += grid[rowCache][colCache]
    if textCache == text:
      return True
    elif rowCache == 0:
      break
    rowCache -= 1

    # up-right
    textCache += grid[rowCache][colCache]
    if textCache == text:
      return True
    elif rowCache == 0 or colCache == colLength - 1:
      break
    rowCache -= 1
    colCache += 1

     # right
    textCache += grid[rowCache][colCache]
    if textCache == text:
      return True
    elif colCache == colLength - 1:
      break
    colCache += 1

    # down-right
    textCache += grid[rowCache][colCache]
    if textCache == text:
      return True
    elif rowCache == rowLength - 1 or colCache == colLength - 1:
      break
    rowCache += 1
    colCache += 1

    # down
    textCache += grid[rowCache][colCache]
    if textCache == text:
      return True
    elif rowCache == rowLength - 1:
      break
    rowCache += 1

    # down-left
    textCache += grid[rowCache][colCache]
    if textCache == text:
      return True
    elif rowCache == rowLength - 1 or colCache == 0:
      break
    rowCache += 1
    colCache -= 1

    # left
    textCache += grid[rowCache][colCache]
    if textCache == text:
      return True
    elif colCache == 0:
      break
    colCache -= 1

    # up-left
    textCache += grid[rowCache][colCache]
    if textCache == text:
      return True
    elif rowCache == rowLength - 1 or colCache == colLength - 1:
      break
    rowCache -= 1
    colCache -= 1

  return False

def readFile(fileName):
   with open(fileName) as f:
      return f.read()

if __name__ == "__main__":
    print('Part 1:', solvePart1('test.txt'))
    # print('Part 2:', solvePart2('test.txt'))