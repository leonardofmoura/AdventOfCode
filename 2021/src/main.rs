use std::env;

mod day1;
mod day2;
mod day3;
mod day4;
mod day5;
mod day6;

fn main() {
    let days: Vec<fn()> = vec![
        day1::main,
        day2::main,
        day3::main,
        day4::main,
        day5::main,
        day6::main,
    ];

    let args: Vec<String> = env::args().collect();
    let day: usize = args[1].parse().unwrap();

    days[day-1]()
}
