use lib;

pub fn main() {
    let filename = "src/day7/input.txt";
    let v = lib::read_line_ints(filename); 

    p1(&v);
    p2(&v)
}

fn sim(nums: &Vec<i32>, start: usize) -> i32 {
    let mut cost = 0;

    for n in nums {
        let fuel = *n - start as i32;
        cost += fuel.abs();
    }

    cost
}

fn sim_new(nums: &Vec<i32>, start: usize) -> i32 {
    let mut cost = 0;

    for n in nums {
        let s = (*n - start as i32).abs();
        let fuel = s*(s+1)/2; 
        cost += fuel;
    }

    cost
}

fn p1(nums: &Vec<i32>) {
    let mut min_cost = i32::MAX;

    // brute force because its good enough
    for n in 0..nums.len() {
        let n_cost = sim(&nums,n);

        if n_cost < min_cost {
            min_cost = n_cost
        }
    }

    println!("sol: {}",min_cost);
}

fn p2(nums: &Vec<i32>) {
    let mut min_cost = i32::MAX;

    // brute force because its good enough
    for n in 0..nums.len() {
        let n_cost = sim_new(&nums,n);

        if n_cost < min_cost {
            min_cost = n_cost
        }
    }

    println!("sol: {}",min_cost);

}

