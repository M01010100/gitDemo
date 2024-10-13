#include <iostream>
#include <string>
using namespace std;
#
string print_following(string word_1, string word_2, int n){
	string out;
	for(int i = 0; i < n; i++){
		out.append(word_1 + ' ');
		out.append(word_2 + '\n');
	}
	return out;
}

int main(){
	string one = "Pesto is";
	string two = "a Penguin";

	string display = print_following(one, two, 2);
	cout << display;
	return 0;
}
