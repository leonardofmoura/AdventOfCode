#include <fstream>
#include <iostream>
#include <vector>
#include <string>

int findTrees(int right, int down, std::vector<std::string> map) {
    int y = 0, x = 0, trees = 0;
    while (y < map.size()) {
        if (map[y][x] == '#') {
            trees++;
        }

        y += down;
        x = (x + right) % map[0].length();
    }

    return trees;
}

int main() {
    std::fstream mapFile;
    mapFile.open("map.txt",std::ios::in);

    if (!mapFile.is_open()) {
        return 1;
    }

    std::vector<std::string> map;
    std::string line;

    while (std::getline(mapFile,line)) {
        map.push_back(line);
    }

    int trees1 = findTrees(1,1,map), trees2 = findTrees(3,1,map), 
        trees3 = findTrees(5,1,map), trees4 = findTrees(7,1,map), 
        trees5 = findTrees(1,2,map);

    std::cout << "Part1: There are " << trees2 << " trees in the way!\n\n";

    std::cout << "Part 2:\n> " << trees1 << "\n> " << trees2 << "\n> " << trees3 
        << "\n> " << trees4 << "\n> " << trees5 << "\nTheir product is " 
        << trees1 * trees2 * trees3 * trees4 * trees5 << std::endl; 


    mapFile.close();

    return 0;
}