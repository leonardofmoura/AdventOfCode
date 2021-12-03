#include "./string_manipulation.hpp"
#include <vector>
#include <string>
#include <cstring>

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