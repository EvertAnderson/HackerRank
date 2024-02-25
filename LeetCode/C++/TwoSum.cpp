#include <iostream>
#include <iostream>
#include <vector>
#include <string>
using namespace std;

vector<int> twoSum(vector<int>& nums, int target) {
    for (int i = 0; i < nums.size(); i++)
    {
        int tempTarget = target - nums[i];
        for (int j = i + 1; j < nums.size(); j++)
        {
            if (nums[j] == tempTarget)
            {
                std::cout << std::to_string(i) + " " + std::to_string(j);
                return vector<int> {i, j};
            }
        }
    }
    return vector<int> {0, 0};
}

int main()
{
    std::cout << "Hello World!\n";
    //char c;
    //cin >> c;
    //std::cout << c;
    vector<int> a = { 3, 2, 4 };
    twoSum(a, 6);
    cin.get();
}