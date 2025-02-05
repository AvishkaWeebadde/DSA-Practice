#include<iostream>
#include<string>

// the input will be the following
// (()) and ()() both result in floor 0.
// ((( and (()(()( both result in floor 3.
// ))((((( also results in floor 3.
// ()) and ))( both result in floor -1 (the first basement level).
// ))) and )())()) both result in floor -3.
// according to the input given if there is paranthesis pair, it will count as zero, if it is a left paranthesis it will count as positive and else negative
// write the c++ program to handle this logic

int calculateFloor(const std::string& input) {
    int floor = 0;
    int index = 0;
    for (char c : input) {
        if (c == '(') {
            floor++;
        } else if (c == ')') {
            floor--;
        }
        index++;
        if(floor == -1) {
            return index - 1;
        }
    }
    return floor;
}

int main() {

    std::string input;
    std::cin >> input;

    std::cout << calculateFloor(input) << std::endl;

    return 0;
}