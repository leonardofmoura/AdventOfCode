use std::collections::VecDeque;

use lib;

pub fn main() {
    let filename = "src/day6/input.txt";
    let vec = lib::read_line_ints(filename);

    solve(&vec,80);
    opt(&vec, 256);
}

// naive solution
fn solve(line: &Vec<i32>, days: i32) {
    let mut fishes = line.to_owned();

    for _ in 0..days {
        let mut count = 0;

        for i in 0..fishes.len() {
            if fishes[i] == 0 {
                fishes[i] = 6;
                count += 1;
            }
            else {
                fishes[i] -= 1;
            } 
        }

        for _ in 0..count {
            fishes.push(8);
        }
    }   


    println!("=== PART 1 ===\nsol: {}",fishes.len());
}

// optimized solution
fn opt(line: &Vec<i32>, days: i32) {
    let fishes = line.to_owned();

    let mut queue: VecDeque<usize> = VecDeque::from(vec![0;9]);

    for f in fishes {
        queue[f as usize] += 1;
    }

    let mut len = line.len();
    for _ in 0..days {
        let to_add = queue.pop_front().unwrap();
        queue[6] += to_add;
        queue.push_back(to_add);
        len += to_add;
    }

    println!("=== PART 2 ===\nsol: {}",len);
}
