#include <cmath>
#include <cstdio>
#include <vector>
#include <iostream>
#include <algorithm>
#include <cstdlib>
#include <cstdio>

using namespace std;

int solve_me_first (int v1, int v2)
{
    return v1 + v2;
}

int main() {
    /* Enter your code here. Read input from STDIN. Print output to STDOUT */  
    int a, b; 
    cin >> a;
    cin >> b;
    
    // string buffer
    char str[1000];
 
    // sprintf() to print num to str buffer
    sprintf(str, "%d", solve_me_first(a, b));

    cout << str;
    
    return 0;
}
