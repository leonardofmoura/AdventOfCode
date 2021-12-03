use lib::{self, read_strings};

pub fn main() {
    let file = "./src/day2/input.txt";

    p1(read_strings(file));
    p2(read_strings(file));
}

fn p1(lines: Vec<String>) {
    let mut x = 0;
    let mut y = 0;

    for line in lines {
        let l: Vec<&str> = line.split(" ").collect();

        match l[0] {
            "forward" => {
                x += l[1].parse::<i32>().unwrap();
            },
            "down" => {
                y += l[1].parse::<i32>().unwrap();
            },
            "up" => {
                y -= l[1].parse::<i32>().unwrap();
            },
            invalid => {
                println!("Invalid input: {}", invalid);
            }
        }
    }

    println!("\n=== PART 1 ===\nx: {}, y: {}\nmult: {}",x,y,x*y);
}

fn p2(lines: Vec<String>) {
    let mut x = 0;
    let mut y = 0;
    let mut aim = 0;

    for line in lines {
        let l: Vec<&str> = line.split(" ").collect();

        match l[0] {
            "forward" => {
                let delta = l[1].parse::<i32>().unwrap();
                x += delta;               
                y += aim*delta;
            },
            "down" => {
                aim += l[1].parse::<i32>().unwrap();
            },
            "up" => {
                aim -= l[1].parse::<i32>().unwrap();
            },
            invalid => {
                println!("Invalid input: {}", invalid);
            }
        }
    }

    println!("\n=== PART 2 ===\nx: {}, y: {}\nmult: {}",x,y,x*y);

}
