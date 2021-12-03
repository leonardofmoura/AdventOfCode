#include "file_read.hpp"
#include <fstream>
#include <vector>
#include <string>
#include <iostream>
#include <sstream>

std::vector<int> read_integers(std::string file_name) {
    std::fstream file;
    std::string str;
    int temp;
    std::vector<int> res;

    file.open(file_name,std::ios::in);

    if (!file.is_open()) {
        std::cerr << "Error reading file\n"; 
        return res;
    }

    //read the file
    while (std::getline(file, str)) {
        temp = std::stoi(str);
        res.push_back(temp);
    }

    file.close();
    return res;
} 


std::vector<std::string> read_strings(std::string file_name) {
    std::fstream file;
    std::vector<std::string> res;
    std::string str;


    file.open(file_name,std::ios::in);

    if (!file.is_open()) {
        std::cerr << "Error reading file\n"; 
        return res;
    }

    while (std::getline(file,str)) {
        res.push_back(str);
    }

    file.close();
    return res;
}


std::string read_to_string(std::string file_name) {
    std::fstream file;
    file.open(file_name,std::ios::in);

    if (!file.is_open()) {
        std::cerr << "Error reading file\n"; 
        return "";
    }

    std::ostringstream out;
    out << file.rdbuf();

    file.close();
    return out.str();
}