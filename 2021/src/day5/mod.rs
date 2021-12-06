use lib;

pub fn main() {
    let file = "src/day5/input.txt";

    solve(lib::read_strings(file)); 
}

struct Vents {
    matrix: Vec<Vec<i32>>,
}

impl Vents {
    pub fn new(x: usize, y: usize) -> Vents {
        Vents {
            matrix: vec![vec![0; x]; y],
        }
    }

    pub fn populate(&mut self, coords: &Vec<((usize,usize),(usize,usize))>) {
        for c in coords {
            let (x1, y1) = c.0;
            let (x2, y2) = c.1;
            
            if x1 == x2 { //vertical
                for i in y1.min(y2)..=y1.max(y2) {
                    self.matrix[i][x1] += 1;
                }
            }
            else if y1 == y2 { //horizontal
                for i in x1.min(x2)..=x1.max(x2) {
                    self.matrix[y1][i] += 1;
                }
            }
        }
    }

    pub fn populate_diag(&mut self, coords: &Vec<((usize,usize),(usize,usize))>) {
        for c in coords {
            let (x1, y1) = c.0;
            let (x2, y2) = c.1;
        
            if x1 == x2 { //vertical
                for i in y1.min(y2)..=y1.max(y2) {
                    self.matrix[i][x1] += 1;
                }
            }
            else if y1 == y2 { //horizontal
                for i in x1.min(x2)..=x1.max(x2) {
                    self.matrix[y1][i] += 1;
                }
            }
            else { //diagonal
                let (mut init, fin) = if x1 < x2 {
                    ((x1,y1),(x2,y2))
                } 
                else {
                    ((x2,y2),(x1,y1))
                };

                let neg_slope = if fin.1 > init.1 {
                    false
                } else {
                    true
                }; 

                while init != fin {
                    self.matrix[init.1][init.0] += 1;
                    match neg_slope {
                        true => init.1 -= 1,
                        false => init.1 += 1,
                    }
                    init.0 += 1;
                }

                self.matrix[init.1][init.0] += 1; //set fin
                
            }
        }
    }

    pub fn overlaps(&self) -> i32 {
        let mut o = 0;

        for line in &self.matrix {
            for elem in line {
                if *elem > 1 {
                    o += 1;
                } 
            }
        }

        o
    }
}

fn solve(lines: Vec<String>) {
    let mut coords: Vec<((usize,usize),(usize,usize))> = Vec::new();
    let mut max_x = 0;
    let mut max_y = 0;
    
    // first pass to find out the max in each line to construct the vector
    // it also constructs the sequence of vents 
    for line in lines {
        let mut coord: Vec<usize> = Vec::new();
        line.split(" -> ").for_each(|l| l.split(",").for_each(|c| coord.push(c.parse().unwrap())));
        coords.push(((coord[0],coord[1]),(coord[2],coord[3])));

        max_x = max_x.max(coord[0].max(coord[2]));
        max_y = max_y.max(coord[1].max(coord[3]));
    }
    
    let mut vents = Vents::new(max_x+1, max_y+1);
    let mut vents_2 = Vents::new(max_x+1, max_y+1);

    vents.populate(&coords);

    println!("=== PART 1 ===\nOverlaps: {}",vents.overlaps());

    vents_2.populate_diag(&coords);
    println!("=== PART 2 ===\nOverlaps: {}",vents_2.overlaps());
}
