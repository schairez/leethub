class Solution {
public:
    int racecar(long int target) {
        queue<pair<int, int>> q;
        q.push({0, 1});
        int depth = 0;
        while (true) {
            //vector<pair<long int, long int>> next_q;
            int n = q.size(); 
            for (int idx = 0; idx < n; idx++) {
                pair<int, int> node = q.front();
                q.pop();
                int pos = node.first;
                int vel = node.second;
                if (pos == target) {
                    return depth;
                }
                // 1 <= target <= 104 
                if (pos + vel <= 10000 && pos + vel > 0) {
                    q.push({ pos + vel, 2 * vel });
                }
                
                if (pos + vel > target && vel > 0) {
                    q.push({ pos, -1 });
                } else if (pos + vel < target && vel < 0) {
                    q.push({ pos, 1 });
                }
            }
            ++depth;
        }
        return -1;
    }
};