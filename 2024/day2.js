import fs from 'node:fs'

function getMatrix(filename) {
    let matrix = fs.readFileSync(filename,'utf8').split('\n').map((elem) => elem.split(" ").map((num) => parseInt(num)));
    matrix.pop();

    //console.log(matrix);
    return matrix;
}

function testLine(line, index, increase) {
    // verify ending condition
    if (index == line.length-1) {
        return true;
    }

    // verify diff
    let diff = Math.abs(line[index+1] - line[index]);
    //console.log(`${line[index]} -> ${line[index+1]}: ${diff}`);
    if (diff < 1 || diff > 3) {
        return false;
    }

    //console.log(increase);
    //console.log(`${line[index]} < ${line[index+1]} -> ${line[index] < line[index+1]}`);
    if ((increase && (line[index] < line[index+1])) || (!increase && (line[index] > line[index+1]))) {
        return testLine(line,index+1,increase)
    }

    return false;
}

function part1(filename) {
    let matrix = getMatrix(filename);

    let sum = 0;

    for (let line of matrix) {
        if (testLine(line,0,line[0] < line[1])) {
            //console.log(`${line}: true`);
            sum++;
        }
        else {
            //console.log(`${line}: false`);
        }
    }

    console.log("part 1: " + sum);
}

function part2(filename) {
    let matrix = getMatrix(filename);

    let sum = 0;

    for (let line of matrix) {
        if (testLine(line,0,line[0] < line[1])) {
            console.log(`${line}: true`);
            sum++;
        }
        else {
            // the line failed -> test all combinations of removing an elem
            for (let i = 0; i <= line.length-1; i++) {
                let newline = [...line]; // copy line 
                newline.splice(i,1);
                if (testLine(newline,0,newline[0] < newline[1])) {
                    console.log(`${line}: true; remove ${line[i]}`);
                    sum++;
                    break;
                }

                if (i == line.length-1) {
                    console.log(`${line}: false`);
                }
            }
        }
    }

    console.log("part 2: " + sum);
}

part1('inputs/day2.txt');
part2('inputs/day2.txt');
