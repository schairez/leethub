class Solution {
public:
    vector<vector<string>> res;
    vector<vector<string>> solveNQueens(int n) {
        //gen board
        vector<string> chessBoard(n, string(n, '.')); // O(n*n) time to create
        dfs(chessBoard, 0, n);
        return res;
    }
    void dfs(vector<string> &board, int x, int queensLeft) {
        if (queensLeft == 0) {
            res.push_back(board);
            return;
        } 
        for (int y = 0; y < board.size(); ++y) {
            if (!isValid(board, x, y)) {
                continue;
            }
            board[x][y] = 'Q';
            dfs(board, x+1, queensLeft-1);
            board[x][y] = '.';
        } 
    }
    bool isValid(vector<string> &board, int rowLoc, int colLoc) {
        int n = board.size();
        // check row and col vals
        for (int x = 0; x < n; x++) {
            if (board[x][colLoc] == 'Q') {
                return false;
            } 
        }
        for (int y = 0; y < n; y++) {
            if (board[rowLoc][y] == 'Q') {
                return false;
            }
        }
        int offset = 1;
        // check upper diag
        for (int x = rowLoc-1; x >= 0; x--) {
            if (colLoc -offset >= 0) {
                if (board[x][colLoc-offset] == 'Q') {
                    return false;
                }
            }
            if (colLoc + offset < n) {
                if (board[x][colLoc+offset] == 'Q') {
                    return false;
                }
            }
            offset++;
        }
        offset = 1;
        // check lower diag
        for (int x = rowLoc+1; x < n; x++) {
            if (colLoc -offset >= 0 && colLoc -offset < n) {
                if (board[x][colLoc-offset] == 'Q') {
                    return false;
                }
            }
            if (colLoc + offset < n && colLoc+offset >= 0) {
                if (board[x][colLoc+offset] == 'Q') {
                    return false;
                }
            }
            offset++;
        }
        return true;
    }
};

