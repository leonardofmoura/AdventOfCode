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
    lines.map(|l| l.unwrap().parse::<i32>().unwrap()).collect()
}

pub fn read_strings<P>(file: P) -> Vec<String> 
where P: AsRef<Path>, {
    let lines = read_lines(file).unwrap();
    lines.map(|l| l.unwrap()).collect()
}

pub fn read_line_ints<P>(file: P) -> Vec<i32>
where P: AsRef<Path>, {
    let lines = read_strings(file);
    lines[0].split(",").map(|l| l.parse::<i32>().unwrap()).collect()
}

pub fn read_matrix<P>(file: P) -> Vec<Vec<i32>>
where P: AsRef<Path>, {
    let lines = read_strings(file);
    let mut vec: Vec<Vec<i32>> = vec![];
    for l in lines {
        let sp = l.chars();
        vec.push(sp.map(|l| l.to_digit(10).unwrap() as i32).collect());
    }
     vec
}

pub fn calc_binary(num: &String) -> i32 {
    let mut res = 0;
    let inv = num.chars().rev();

    for (index, c) in inv.enumerate() {
        if c == '1' {
            res += 1 << index;
        } 
    }

    res
}

pub fn sort_str(s : String) -> String {
    let mut l: Vec<char> = s.to_owned().chars().collect();
    l.sort_unstable();
    l.into_iter().collect()
}

pub fn common_chars(s1: &String, s2: &String) -> i32 {
    let mut counter = 0;

    for c in s1.chars() {
        if s2.find(c) != None {
            counter += 1;
        }
    }

    counter
}
