def solvePart1(fileName):
  lines = readFile(fileName).split('\n')
  for line in lines:
    args = line.split(': ')
    value = args[0]
    nums = args[1].split(' ')
    test(value, nums)
  return 0

def solvePart2(fileName):
  file = readFile(fileName)
  return file

def test(value, nums):
   operators = 2
   possibilities = pow(len(nums) - 1, 1) * operators
   print(possibilities)
  #  for num in range(possibilities):
      

def readFile(fileName):
   with open(fileName) as f:
      return f.read()

if __name__ == "__main__":
    print('Part 1:', solvePart1('test.txt'))
    # print('Part 2:', solvePart2('input.txt'))