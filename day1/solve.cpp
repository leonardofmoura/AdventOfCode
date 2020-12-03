#include <iostream>
#include <fstream>
#include <string>
#include <vector>
#include <algorithm>
#include <unordered_map>

int main() {
    std::vector<int> entries;
    std::string entryString;
    int entry;
    
    std::fstream reportFile;
    reportFile.open("report.txt",std::ios::in);

    if (!reportFile.is_open()) {
        return 1;
    }

    //read the file
    while (std::getline(reportFile, entryString)) {
        entry = std::stoi(entryString);
        entries.push_back(entry);
    }

    //The algorithm itself
    std::sort(entries.begin(),entries.end()); //sort the array
    std::vector<int>::iterator start = entries.begin();
    std::vector<int>::iterator end = std::prev(entries.end());

    int sum = 0;
    while (start != entries.end() || end != entries.begin()) {
        sum = *start + *end;

        if (sum == 2020) {
            std::cout << "Sum found! " << *start << "+" << *end << "= 2020" << std::endl;
            std::cout << "The solution is: " << *start * *end << std::endl;
            break;
        }
        else if (sum > 2020) {
            end--;
        }
        else {
            start++;
        }
    }

    std::cout << "Algorithm 1 ended\n";

    //PART TWO
    // reset iterators
    std::unordered_map<int,std::vector<int>> map;
    std::vector<int> values;

    for (int i = 0; i < entries.size(); i++) {
        for (int j = 0; j < entries.size(); j++) {
            if (i == j) {
                continue;
            }

            values = {entries[i],entries[j]};
            sum = entries[i] + entries[j];

            map.insert(std::make_pair(sum,values));
        }
    }

    for (int i = 0; i < entries.size(); i++) {
        std::unordered_map<int,std::vector<int>>::iterator val = map.find(2020-entries[i]);

        if (val != map.end()) {
            std::cout << "Sum found! " << val->second[0] << "+" << val->second[1] << "+" << entries[i] << "= 2020" << std::endl;
            std::cout << "The solution is: " << val->second[0] * val->second[1] * entries[i] << std::endl;
            break;
        }
    } 


    std::cout << "Algorithm 2 ended\n";
    reportFile.close();
}