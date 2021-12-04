use lib::read_strings;

pub fn main() {
    let file = "src/day3/input.txt";

    p1(read_strings(file));
    p2(read_strings(file));
}

fn p1(lines: Vec<String>) {
    let order = lines[0].len();

    let mut indexes = vec![0; order];
    let mut index; //lsb will be in index 0 of the array

    for l in &lines {
        index = order;
        for i in l.chars() {
            index -= 1;
            match i {
            '1' => {
                indexes[index] += 1;
            },
            '0' => {},
                _ => {
                    println!("Invalid input");
                    break;
                }
            }
        }
    }

    let mut gamma = 0;
    let mut epsilon = 0;

    //println!("{:?}",indexes);

    for (i, count) in indexes.iter().enumerate() {
        if *count > lines.len()/2 {
            gamma += 1 << i;
        }
        else {
            epsilon += 1 << i;
        }
    }

    println!("=== PART 1 ===\ngamma: {}, epsilon: {}\nmul: {}",gamma,epsilon,gamma*epsilon);
}

fn generate_vecs(lines: &Vec<String>, index: usize) -> (Vec<String>,Vec<String>) {
    let mut one_vec: Vec<String> = Vec::new();
    let mut zero_vec: Vec<String> = Vec::new();

    for line in lines {
        match line.chars().nth(index).unwrap() {
            '1' => {
                one_vec.push(line.clone());
            },
            '0' => {
                zero_vec.push(line.clone());
            }
            _ => {
                println!("Invalid bit");
                break;
            }
        }      
    }

    (zero_vec,one_vec)
}

fn find_oxy(lines: &Vec<String>, index: usize) -> Vec<String> {
   let vecs = generate_vecs(lines, index);

    if vecs.1.len() >= vecs.0.len() {
        vecs.1
    }
    else {
        vecs.0
    }
}

fn find_co2(lines: &Vec<String>, index: usize) -> Vec<String> {
    let vecs = generate_vecs(lines, index);

    if vecs.0.len() <= vecs.1.len() {
        vecs.0
    }
    else {
        vecs.1
    }
}


fn p2(lines: Vec<String>) {
    let mut index = 0;
    let mut oxy = find_oxy(&lines, index);

    while oxy.len() != 1 {
        index += 1;
        oxy = find_oxy(&oxy, index);
    }
    
    let oxygen = lib::calc_binary(&oxy[0]);

    index = 0;
    let mut co2 = find_co2(&lines, index);

    while co2.len() != 1 {
        index += 1;
        co2 = find_co2(&co2, index);
    }

    let scrubber = lib::calc_binary(&co2[0]);

    println!("=== PART 2 ===\noxy: {}, co2: {}\nmul: {}",oxygen,scrubber,oxygen*scrubber);
}


