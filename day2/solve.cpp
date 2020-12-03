#include <iostream>
#include <string>
#include <sstream>
#include <fstream>

bool validPassword(char letter,int least,int most, std::string password) {
    int count = 0;

    for (int i = 0; i < password.length(); i++) {
        if (password[i] == letter) {
            count ++;
        }
    }

    return count >= least && count <= most;
}


//we can use the least and most values directly because password includes the ' ' in the beginning
bool validPassword2(char letter,int least,int most, std::string password) {
    return (password[least] == letter) != (password[most] == letter);
}

int main() {
    std::fstream passwords;
    passwords.open("passwords.txt",std::ios::in);

    if (!passwords.is_open()) {
        return 1;
    }

    int validPasswords = 0;
    int validPass2 = 0;
    std::string str;

    while (std::getline(passwords,str)) {
        std::stringstream sstream(str);

        std::string least;
        std::string most;
        std::string letter;
        std::string password;

        getline(sstream,least,'-');
        getline(sstream,most,' ');
        getline(sstream,letter,':');
        getline(sstream,password,'\n');

        if (validPassword(letter[0],std::stoi(least),std::stoi(most),password)) validPasswords++;
        if (validPassword2(letter[0],std::stoi(least),std::stoi(most),password)) validPass2++;
    }

    std::cout << "There are " << validPasswords << " valid passwords according to rule 1\n";
    std::cout << "There are " << validPass2 << " valid passwords according to rule 2\n";

    passwords.close();
    return 0;
}