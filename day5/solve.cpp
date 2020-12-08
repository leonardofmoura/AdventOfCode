#include <iostream>
#include <fstream>
#include <vector>
#include <algorithm>
#include <cmath>
#include "../lib/file_read.hpp"

int get_seat_id(const std::string seat) {
    if (seat.length() != 10) {
        return -1;
    }

    double low = 0, high = 127;
    double row, column;

    for (int i = 0; i < 6; i++) {
        if (seat[i] == 'F') {
            high = floor(high - (high - low)/2);
        }
        else if (seat[i] == 'B') {
            low = ceil(low + (high - low)/2);
        }
    }

    if (seat[6] == 'F') {
        row = low;
    }
    else if (seat[6] == 'B') {
        row = high;
    }

    low = 0, high = 7;
    for (int i = 7; i < 9; i++) {
        if (seat[i] == 'L') {
            high = floor(high - (high - low)/2);
        }
        else if (seat[i] == 'R') {
            low = ceil(low + (high - low)/2);
        }
    }

    if (seat[9] == 'L') {
        column = low;
    }
    else if (seat[9] == 'R') {
        column = high;
    }

    return row * 8 + column;
}

int main() {
    std::vector<std::string> seats = read_strings("seats.txt");
    std::vector<int> ids;

    for (auto seat : seats) {
        ids.push_back(get_seat_id(seat));
    }

    std::sort(ids.begin(),ids.end());

    std::cout << "Highest id: " << ids[ids.size()-1] << std::endl;

    for (int i = 0; i < ids.size()-1; i++) {
        if (ids[i]+1 != ids[i+1]) {
            std::cout << "Seat found! ID: " << ids[i]+1 << std::endl;
        }
    }

    return 0;
}