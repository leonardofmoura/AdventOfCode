use std::env;

mod day1;
mod day2;
mod day3;
mod day4;
mod day5;
mod day6;
mod day7;
mod day8;
mod day9;

fn main() {
    let days: Vec<fn()> = vec![
        day1::main,
        day2::main,
        day3::main,
        day4::main,
        day5::main,
        day6::main,
        day7::main,
        day8::main,
        day9::main,
    ];

    let args: Vec<String> = env::args().collect();
    let day: usize = args[1].parse().unwrap();

    days[day-1]()
}
