use std::fs;

fn main() {
  let input = fetch_input();
  let lines: Vec<_> = input.lines().collect();
  let part1 = part_one(lines.clone());
  let part2 = part_two(lines.clone());
  println!("Part 1: {part1}");
  println!("Part 2: {part2}");
}

fn part_one(lines: Vec<&str>) -> i32 {
  let (mut a, mut b) = get_lists(lines);
  a.sort();
  b.sort();
  
  let mut total = 0;
  for (i, num) in a.iter().enumerate() {
    total += i32::abs(num - b[i])
  }
  return total;
}

fn part_two(lines: Vec<&str>) -> i32 {
  let (a, b) = get_lists(lines);

  let mut similarity = 0;
  for num in a {
    similarity += num * times_in_list(num, b.clone());
  }
  return similarity;
}

fn get_lists(lines: Vec<&str>) -> (Vec<i32>, Vec<i32>) {
  let mut a = Vec::new();
  let mut b = Vec::new();

  for line in lines.into_iter() {
    let args: Vec<_> = line.split("   ").collect();
    a.push(args[0].parse::<i32>().unwrap());
    b.push(args[1].parse::<i32>().unwrap());
  }

  return (a, b);
}

fn times_in_list(num: i32, list: Vec<i32>) -> i32 {
  let mut total = 0;
  for x in list {
    if x == num {
      total += 1;
    }
  }
  return total;
}

fn fetch_input() -> String {
  let contents = fs::read_to_string("./input.txt")
    .expect("Input file not found.");
  return contents
}