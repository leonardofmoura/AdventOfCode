use std::env;

mod day1;
mod day2;
mod day3;

fn main() {
    let days: Vec<fn()> = vec![
        day1::main,
        day2::main,
        day3::main,
    ];

    let args: Vec<String> = env::args().collect();
    let day: usize = args[1].parse().unwrap();

    days[day-1]()
}