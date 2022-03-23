class Solution {
public:
    
    int racecar(long int target) {
        queue<pair<long int, long int>> q;
        q.push({0, 1});
        long int depth = 0;
        
        while (true) {
            //vector<pair<long int, long int>> next_q;
            int n = q.size(); 
            for (long int idx = 0; idx < n; idx++) {
                pair<long int, long int> node = q.front();
                q.pop();
                long int pos = node.first;
                long int speed = node.second;
                if (pos == target) {
                    return depth;
                }
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