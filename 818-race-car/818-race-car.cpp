// naive BFS w/o pruning would lead to O(2^n) runtime and space;
// but here we're not enumerating the accelerate and reverse option for
// each node; with the pruning step our bound reduces dramatically
// as it takes some logN steps to overshoot our target and some 

class Solution {
public:
    int racecar(long int target) {
        queue<pair<int, int>> q;
        q.push({0, 1});
        int lvl = 0;
        while (!q.empty()) {
            int n = q.size(); 
            cout << n << '\n';
            for (int idx = 0; idx < n; idx++) {
                pair<int, int> node = q.front();
                q.pop();
                int pos = node.first;
                int vel = node.second;
                if (pos == target) {
                    return lvl;
                }
                // 1 <= target <= 10^4 ; avoids using long type for vel
                if (pos + vel <= 2 * target && pos + vel > 0) {
                    q.push({ pos + vel, 2 * vel });
                }
                if (pos + vel > target && vel > 0) {
                    q.push({ pos, -1 });
                } else if (pos + vel < target && vel < 0) {
                    q.push({ pos, 1 });
                }
            }
            ++lvl;
        }
        return -1;
    }
};