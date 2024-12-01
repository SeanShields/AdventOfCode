# Read the input data and parse it into a list of tuples
with open("input.txt") as f:
    data = f.read().strip().split("\n")
instructions = [(x.split()[0], int(x.split()[1])) for x in data]

# Initialize the register X and the cycle counter
x = 1
cycle = 1

# Iterate through the instructions
signal_strengths = []
for instruction, value in instructions:
    if instruction == "addx":
        # Execute the addx instruction and update the value of the register X
        x += value
    else:
        # Execute the noop instruction
        pass

    # Check if the current cycle is one of the interesting signal strength cycles
    if cycle % 40 == 20:
        signal_strengths.append(cycle * x)

    # Increment the cycle counter
    cycle += 1

# Print the result
print(f"The interesting signal strengths are: {signal_strengths}")
