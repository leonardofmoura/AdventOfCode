use lib;

pub fn main() {
    let filename = "src/day9/input.txt";

    p1(lib::read_matrix(filename));
}

fn p1(matrix: Vec<Vec<i32>>) {
    let mut low_points: Vec<i32> = vec![];
    let mut low_coords: Vec<(i32,i32)> = vec![];

    for line in 0..matrix.len() {
        for i in 0..matrix[line].len() {
            let elem = matrix[line][i];
            
            if line > 0  && matrix[line-1][i] <= elem {
                continue;
            }
            else if line < matrix.len() - 1 && matrix[line+1][i] <= elem {
                continue;
            }
            else if i > 0 && matrix[line][i-1] <= elem {
                continue;
            }
            else if i < matrix[line].len() - 1 && matrix[line][i+1] <= elem {
                continue;
            }

            low_points.push(elem);
            low_coords.push((line as i32,i as i32));
        }
    }

    let mut s = 0;

    for elem in &low_points {
        s += 1 + elem;
    }

    println!("Sol1: {}", s);

    let mut explored = vec![vec![false; matrix[0].len()]; matrix.len()];

    let mut basins = vec![];

    for elem in low_coords {
        basins.push(fill(elem.1, elem.0, &matrix, &mut explored));
    }

    basins.sort_by(|a,b| b.cmp(a));

    println!("Sol2: {}", basins[0] * basins[1] * basins[2]);
}

fn fill(x: i32, y: i32, matrix: &Vec<Vec<i32>>, explored: &mut Vec<Vec<bool>>) -> i32 {
    if y >= matrix.len() as i32 || y < 0 {
        return 0;
    } 
    else if x >= matrix[y as usize].len() as i32 || x < 0 {
        return 0;
    } 
    else if  matrix[y as usize][x as usize] == 9 {
        return 0;
    }
    else if explored[y as usize][x as usize] {
        return 0;
    }

    explored[y as usize][x as usize] = true;

    return 1 + fill(x+1, y, matrix, explored) + fill(x-1, y, matrix, explored) + fill(x, y+1, matrix, explored) + fill(x, y-1, matrix, explored);
}
