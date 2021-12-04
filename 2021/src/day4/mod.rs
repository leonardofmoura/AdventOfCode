use lib::{self, read_strings};

pub fn main() {
    let file = "src/day4/input.txt";

    p1(read_strings(file));
    p2(read_strings(file));
}

struct Board {
    matrix: Vec<Vec<i32>>,
    highlight: Vec<Vec<bool>>
}

impl Board {
    pub fn new() -> Board {
        Board {
            matrix: Vec::new(),
            highlight: vec![vec![false;5];5],
        }
    }

    pub fn add_line(&mut self, line: Vec<&str>) {
        self.matrix.push(line.iter().map(|n| n.parse::<i32>().unwrap()).collect());
    }

    pub fn mark(&mut self, num: i32) {
        for (line_index, line) in self.matrix.iter().enumerate() {
            for (col_index, number) in line.iter().enumerate() {
                if *number == num {
                    self.highlight[line_index][col_index] = true;
                }
            }
        }
    }

    pub fn verify(&self) -> bool {
        //verify lines
        let mut count = 0;
        for line in &self.highlight {
            for val in line {
                if *val == true {
                    count += 1;
                }
            }

            if count == 5 {
                return true;
            }
            else {
                count = 0;
            }
        }

        // verify columns
        count = 0;
        for col in 0..5 {
            for line in 0..5 {
                if self.highlight[line][col] == true {
                    count += 1;
                }
            }

            if count == 5 {
                return true;
            }
            else {
                count = 0;
            }
        }

        false
    }

    pub fn get_score(&self) -> i32{
        let mut sum = 0;

        for (line_index, line) in self.highlight.iter().enumerate() {
            for (col_index, col) in line.iter().enumerate() {
                if *col == false {
                    sum += self.matrix[line_index][col_index];
                }
            }
        }

        sum
    }

    pub fn reset(&mut self) {
        for i in 0..5 {
            for j in 0..5 {
                self.highlight[i][j] = false;
            }
        }
    }
}

fn p1(lines: Vec<String>) {
    //parse the order of the numbers
    let numbers: Vec<i32> = lines.get(0).unwrap().split(",").map(|n| n.parse::<i32>().unwrap()).collect();
    let mut boards: Vec<Board> = Vec::new();
    let mut board_index = 0;
    boards.push(Board::new());

    //parse the boards
    for (index, line) in lines.iter().enumerate() {
        if index > 1 {
            if line == "" {
                board_index += 1;
                boards.push(Board::new());
            }
            else {
                boards.get_mut(board_index).unwrap().add_line(line.trim().split_whitespace().collect());
            }
        }
    }

    //run the simulation
    for (index, n) in numbers.iter().enumerate() {
        for b in 0..boards.len() {
            boards.get_mut(b).unwrap().mark(*n); 
        }

        if index > 4 {
            for b in &boards {
                if b.verify() {
                    let score = b.get_score();
                    println!("\n=== PART 1 ===\nFound result -> score:{}, number: {}\nsol:{}",score,n,score*n);
                    return;
                }
            }
        }
    }
}

fn p2(lines: Vec<String>) {
    //parse the order of the numbers
    let numbers: Vec<i32> = lines.get(0).unwrap().split(",").map(|n| n.parse::<i32>().unwrap()).collect();
    let mut boards: Vec<Board> = Vec::new();
    let mut board_index = 0;
    boards.push(Board::new());

    //parse the boards
    for (index, line) in lines.iter().enumerate() {
        if index > 1 {
            if line == "" {
                board_index += 1;
                boards.push(Board::new());
            }
            else {
                boards.get_mut(board_index).unwrap().add_line(line.trim().split_whitespace().collect());
            }
        }
    }

    let mut stack: Vec<usize> = Vec::new();

    //run the simulation
    for (index, n) in numbers.iter().enumerate() {
        for b in 0..boards.len() {
            boards.get_mut(b).unwrap().mark(*n); 
        }

        if index > 4 {
            for (i, b) in boards.iter().enumerate() {
                if !stack.contains(&i) && b.verify() {
                    stack.push(i);
                }
            }
        }
    }

    //determine the last winner
    let winner = stack.pop().unwrap();

    //reset the boards
    for b in boards.iter_mut() {
        b.reset();
    }

    //run simulation until last winner
    for (index, n) in numbers.iter().enumerate() {
        for b in 0..boards.len() {
            boards.get_mut(b).unwrap().mark(*n); 
        }

        if index > 4 {
            for (i, b) in boards.iter().enumerate() {
                if b.verify() && i == winner {
                    let score = b.get_score();
                    println!("\n=== PART 2 ===\nFound result -> score:{}, number: {}\nsol:{}",score,n,score*n);
                    return;
                }
            }
        }
    }
}
