use std::io::{self, BufRead};
use std::fs::File;
use std::path::Path;

// from https://doc.rust-lang.org/stable/rust-by-example/std_misc/file/read_lines.html
pub fn read_lines<P>(file: P) -> io::Result<io::Lines<io::BufReader<File>>>
where P: AsRef<Path>, {
    let file = File::open(file)?;
    Ok(io::BufReader::new(file).lines())
}

pub fn read_ints<P>(file: P) -> Vec<i32>
where P: AsRef<Path>, {
    let lines = read_lines(file).unwrap();
    let mut res: Vec<i32> = Vec::new();

    for line in lines {
        if let Ok(number) = line {
            res.push(number.parse::<i32>().unwrap());
        } 
    }

    res
}
