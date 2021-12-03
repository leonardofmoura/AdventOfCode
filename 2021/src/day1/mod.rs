use lib;

pub fn main() {
    let file = "./src/day1/input.txt";

    p1(lib::read_ints(file));
    p2(lib::read_ints(file));
}

fn p1(lines: Vec<i32>) {
    let mut last: i32 = 0;
    let mut counter: i32 = -1;

    for i in lines {
        if i > last {
            counter += 1;
        }

        last = i;
    } 

    println!("Result 1: {}",counter);
}

fn p2(lines: Vec<i32>) {
    let mut last: i32 = 0;
    let mut counter: i32 = -1;

    for (index, elem) in lines.iter().enumerate() {
        if index >= lines.len() - 2 {
            break;
        }

        let sum = elem + lines[index+1] + lines[index+2];

        if sum > last {
            counter += 1;
        }

        last = sum;
    }

    println!("Result 2: {}",counter);
}


