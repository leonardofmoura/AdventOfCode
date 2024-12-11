import fs from 'node:fs'

function part1(filename) {
    let memory = fs.readFileSync(filename,'utf-8');

    const mulre = /mul\([0-9]{1,3},[0-9]{1,3}\)/gm;

    let match;
    let sum = 0;

    do {
        match = mulre.exec(memory);
        if (match) {
            const resre = /[0-9]{1,3}/gm

            let i = parseInt(resre.exec(match[0])[0]);
            let j = parseInt(resre.exec(match[0])[0]);

            sum += i*j;
        }
    } while (match);

    console.log(`part 1: ${sum}`);
}


function part2(filename) {
    let memory = fs.readFileSync(filename,'utf-8');

    const mulre = /mul\([0-9]{1,3},[0-9]{1,3}\)|do\(\)|don\'t\(\)/gm 

    let match;
    let sum = 0;
    let enable = true;

    do {
        match = mulre.exec(memory);
        if (match) {
            if (match == 'do()') {
                enable = true;
                continue;
            }
            else if (match == 'don\'t()') {
                enable = false;
                continue;
            }
            
            if (enable) {
                const resre = /[0-9]{1,3}/gm

                let i = parseInt(resre.exec(match[0])[0]);
                let j = parseInt(resre.exec(match[0])[0]);

                sum += i*j;
            }
        }
    } while (match);

    console.log(`part 1: ${sum}`);
}

part1("inputs/day3.txt");
part2("inputs/day3.txt")
