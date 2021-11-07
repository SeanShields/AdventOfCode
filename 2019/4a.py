import re

inputStart = 372037
inputEnd = 905157

def solve():
  passwords = []

  for number in range(inputStart, inputEnd):
    # validate repeating adjacent digit
    isSixDigitsWithRepeating = re.match(r'(\d){0,6}\1', str(number))
    if not isSixDigitsWithRepeating:
      continue

    # validate decreasing/same digits
    decreases = False
    prevDigit = None
    for digit in map(int, str(number)):
      if prevDigit is None:
        prevDigit = digit
        continue
      else:
        if prevDigit > digit:
          decreases = True
          break
      prevDigit = digit

    if not decreases:
      passwords.append(number)
    
  return passwords

combinations = solve()
print(combinations)
print(f'valid passwords: {len(combinations)}')
