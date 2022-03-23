
class Solution {
public:
    int racecar(long int target) {
        queue<pair<int, int>> q;
        q.push({0, 1});
        int depth = 0;
        while (!q.empty()) {
            int n = q.size(); 
            for (int idx = 0; idx < n; idx++) {
                pair<int, int> node = q.front();
                q.pop();
                int pos = node.first;
                int speed = node.second;
                if (pos == target) {
                    return depth;
                }
                if (abs(pos) > 2 * target) continue;
                
                q.push({ pos + speed, 2 * speed });
                if (pos + speed > target && speed > 0) {
                    q.push({ pos, -1 });
                } else if (pos + speed < target && speed < 0) {
                    q.push({ pos, 1 });
                }
            }
            ++depth;
        }
        return -1;
    }
};