#include <string>
#include <vector>
#include <algorithm>
#include <iostream>
#include <set>
#include <map>
#include "../lib/file_read.hpp"
#include "../lib/string_manipulation.hpp"

int count_answers(std::string group_answers) {
    std::vector<std::string> answers = tokenize_string_multiple_char_delim(group_answers,"\n");
    std::set<char> res;

    for (auto ans: answers) {
        for (auto question : ans) {
            res.insert(question);
        }
    }

    return res.size();
}

int count_answers_all(std::string group_answers) {
    std::vector<std::string> answers = tokenize_string_multiple_char_delim(group_answers,"\n");
    std::map<char,int> res;

    for (auto ans: answers) {
        for (auto question : ans) {
            if (res.find(question) == res.end()) {
                res[question] = 1;
            }
            else {
                res[question] += 1;
            }
        }
    }

    int count = 0;

    for (auto q: res) {
        if (q.second == answers.size()) {
            count++;
        }
    }

    return count;
}

int main() {
    std::string str = read_to_string("answers.txt");
    int count = 0;

    std::vector<std::string> group_answers = tokenize_string_multiple_char_delim(str,"\n\n");

    for (auto group: group_answers) {
        count += count_answers(group);
    }

    std::cout << "Solution: " << count << std::endl;

    int count2 = 0;

    for (auto group: group_answers) {
        count2 += count_answers_all(group);
    }

    std::cout << "Solution 2 " << count2 << std::endl;

    return 0;
}