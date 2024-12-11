import fs from 'node:fs';

// input parsing

function parseInput(input) {
    const data = fs.readFileSync(input,'utf-8').split('\n');
    data.pop(); //remove last element

    let first = []
    let second = []

    data.forEach((line) => {
        let nums = line.split("   ");
        first.push(parseInt(nums[0]));
        second.push(parseInt(nums[1]));
    })

    //console.log(first);
    //console.log(second);

    return [first,second]
}


// PART 1

function getMin(list) {
    let min = list[0];
    let ret_index = 0;

    for (let index in list) {
        if (list[index] < min) {
            min = list[index];
            ret_index = index;
        }
    }

    return ret_index;
}

function part1(input) {
    let parsed = parseInput(input);
    let first = parsed[0];
    let second = parsed[1];

    let sum = 0;

    for (let i = 0; i < first.length; i++) {
        let first_index = getMin(first);
        let second_index = getMin(second);

        //console.log('first: ' + first[first_index] + ' second: ' + second[second_index]);
        let add = Math.abs(first[first_index] - second[second_index]);
        sum += Math.abs( first[first_index] - second[second_index]);
        
        // 'remove the elements'
        first[first_index] = 99999999;
        second[second_index] = 99999999;
    }

    return sum;
}


function part2(input) {
    let [first,second] = parseInput(input);
    
    let sum = 0;

    for (let elem of first) {
        let times = second.filter((x) => x == elem).length;
        sum += elem * times;
    }

    return sum;
}

console.log("part1: " + part1("inputs/day1.txt"));
console.log("part2: " + part2("inputs/day1.txt"));
