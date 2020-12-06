#include <iostream>
#include <vector>
#include <fstream>
#include <sstream>
#include <cstring>
#include <algorithm>
#include <unordered_map>
#include <regex>

//sorted array of valid fields
std::vector<std::string> VALID_FIELDS = {"byr","ecl","eyr","hcl","hgt","iyr","pid"};
std::vector<std::string> VALID_FIELDS2 = {"byr","cid","ecl","eyr","hcl","hgt","iyr","pid"};
std::vector<std::string> VALID_EYES = {"amb", "blu", "brn", "gry", "grn", "hzl", "oth"};

std::vector<std::string> tokenize_string_multiple_char_delim(std::string str, std::string delim) {
    std::size_t pos = str.find(delim);
    std::size_t start_pos = 0;
    std::vector<std::string> out;

    while (pos != std::string::npos) {
        out.push_back(str.substr(start_pos,pos-start_pos));
        start_pos = pos + delim.length();
        pos = str.find(delim,start_pos);
    }

    out.push_back(str.substr(start_pos));

    return out;
}

std::vector<std::string> tokenize_string_multiple_delims(std::string const& str, std::string delim_str) {
    char* cstr = new char[str.length()+1];
    strcpy(cstr,str.c_str());

    char* delim = new char[delim_str.length()+1];
    strcpy(delim,delim_str.c_str());

    char* token = strtok(cstr,delim);

    std::vector<std::string> result;
    while (token != NULL) {
        std::string str_token (token);
        result.push_back(str_token);
        token = strtok(NULL,delim);
    }

    delete[] cstr;
    delete[] delim;

    return result;
}

bool verify_height(std::string height) {
    //detect inch
    std::size_t pos = height.find("in");
    bool cm = false;

    if (pos == std::string::npos) {
        pos = height.find("cm");
        cm = true;
    } 

    if (pos == std::string::npos) return false;

    int num = std::stoi(height.substr(0,pos));

    if (cm) return num >= 150 && num <= 193;  
    else return num >= 59 && num <= 76;
}

bool analyze_passport(std::string const& passport) {
    //divide the passport in parameters
    std::vector<std::string> parameters = tokenize_string_multiple_delims(passport," \n");

    std::vector<std::vector<std::string>> passport_info;
    std::vector<std::string> params;

    for (auto i : parameters) {
        passport_info.push_back(tokenize_string_multiple_char_delim(i,":"));
    }

    for (int i = 0; i < passport_info.size(); i++) {
        params.push_back(passport_info[i][0]);
    }

    std::sort(params.begin(),params.end());

    if (params != VALID_FIELDS && params != VALID_FIELDS2) {
        return false;
    }

    //Convert passport info to unordered map
    std::unordered_map<std::string,std::string> info;

    for (auto i : passport_info) {
        info.insert(std::make_pair(i[0],i[1]));
    }

    //Verify information
    if (std::stoi(info["byr"]) > 2002 || std::stoi(info["byr"]) < 1920) return false;
    else if (std::stoi(info["iyr"]) > 2020 || std::stoi(info["iyr"]) < 2010) return false;
    else if (std::stoi(info["eyr"]) > 2030 || std::stoi(info["eyr"]) < 2020) return false;
    else if (!verify_height(info["hgt"])) return false;
    else if (!std::regex_match(info["hcl"],std::regex("#[a-f0-9]{6}"))) return false;
    else if (std::find(VALID_EYES.begin(),VALID_EYES.end(),info["ecl"]) == VALID_EYES.end()) return false;
    else if (info["pid"].length() != 9 || !std::all_of(info["pid"].begin(),info["pid"].end(),::isdigit)) return false;

    return true;
}

int main() {
    std::fstream passports;
    passports.open("passports.txt",std::ios::in);

    if (!passports.is_open()) {
        return 1;
    }

    std::ostringstream out;
    out << passports.rdbuf();
    
    std::vector<std::string> passport_tokens = tokenize_string_multiple_char_delim(out.str(),"\n\n");
    int count = 0;

    for (auto i : passport_tokens) {
        if (analyze_passport(i)) count++;
    }

    std::cout << "There are " << count << " valid passports\n";

    passports.close();

    return 0;
}

