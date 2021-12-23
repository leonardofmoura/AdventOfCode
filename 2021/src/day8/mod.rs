use std::collections::HashMap;

use lib;

pub fn main() {
    let filename = "src/day8/input.txt";
    let lines = lib::read_strings(filename);

    p1(&lines);
    p2(&lines);
}

fn p1(lines: &Vec<String>) {
    let mut c = 0;

    for line in lines {
        let l: Vec<&str> = line.split(" | ").collect();
        let digits = l[1];

        let n = digits.split(" ").collect::<Vec<&str>>();

        for d in n {
            let l = d.len();
            if l == 2 || l == 3 || l == 4 || l == 7 {
                c += 1;
            }
        }
    }

    println!("sol: {}",c);
}


fn p2(lines: &Vec<String>) {
    let mut s = 0;

    for line in lines {
        let mut table = HashMap::new();
        let mut inv_table = HashMap::new();

        let l: Vec<&str> = line.split(" | ").collect();
        let digits = l[0];

        let n = digits.split(" ").collect::<Vec<&str>>();

        // find the digits with different length
        for i in &n {
            let d = lib::sort_str(i.to_string());

            if d.len() == 2 { // 1
                table.insert(d.clone(),1);
                inv_table.insert(1,d);
            }
            else if d.len() == 3 { // 7
                table.insert(d.clone(),7);
                inv_table.insert(7,d);
            }
            else if d.len() == 4 { // 4
                table.insert(d.clone(),4);
                inv_table.insert(4,d);
            }
            else if d.len() == 7 { // 8
                table.insert(d.clone(),8);
                inv_table.insert(8,d);
            }       
        }

        // find digits of length 5
        for i in &n {
            let d = lib::sort_str(i.to_string());
            
            if d.len() == 5 {
                if lib::common_chars(&d, inv_table.get(&7).unwrap()) == 3 { // 3 overlaps 7 in 3 chars
                    table.insert(d.clone(),3);
                    inv_table.insert(3,d);
                }
                else if lib::common_chars(&d, inv_table.get(&4).unwrap()) == 2 { // 2 overlaps 4 in 2 chars
                    table.insert(d.clone(),2);
                    inv_table.insert(2,d);
                }
                else if lib::common_chars(&d, inv_table.get(&4).unwrap()) == 3 { // 5 overlaps 4 in 3 chars
                    table.insert(d.clone(),5);
                    inv_table.insert(5,d);
                }
            }
        }

        // find digits of length 6
        for i in &n {
            let d = lib::sort_str(i.to_string());
            
            if d.len() == 6 {
                if lib::common_chars(&d, inv_table.get(&3).unwrap()) == 5 { // 9 overlaps 3 in 5 chars
                    table.insert(d.clone(),9);
                    inv_table.insert(9,d);
                }
                else if lib::common_chars(&d, inv_table.get(&1).unwrap()) == 2 { // 0 overlaps 1 in 2 chars
                    table.insert(d.clone(),0);
                    inv_table.insert(0,d);
                }
                else if lib::common_chars(&d, inv_table.get(&5).unwrap()) == 5 { // 6 overlaps 5 in 5 chars
                    table.insert(d.clone(),6);
                    inv_table.insert(6,d);
                }
            }
        }

        // now that we have the table, we can decode the digits
        let decode = l[1].trim().split(" ").collect::<Vec<&str>>();

        for i in 0..decode.len() {
            let search = lib::sort_str(decode[i].to_string());
            
            let n = table.get(&search).unwrap();
            s += n * i32::pow(10, 3 - i as u32);
        }
    }

    println!("sol: {}",s);
}

