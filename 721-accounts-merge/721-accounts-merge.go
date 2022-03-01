
//
//
//
//

func accountsMerge(accounts [][]string) [][]string {
    // space: O(n)
    emailToID := make(map[string]int)
    emailToName := make(map[string]string)
    
    //time: O(len(accounts) * len(acct))
    for _, acct := range accounts {
        name := acct[0]
        for _, email := range acct[1:] {
            if _, exists := emailToID[email]; !exists {
                emailToID[email] = len(emailToID)
                emailToName[email] = name
            }
        }
    }
    ufSize := make([]int, len(emailToID))
    ufParent := make([]int, len(emailToID)) 
    for idx := range ufParent {
        ufParent[idx] = idx
        ufSize[idx] = 1
    }
    
    // inv Ackerman O(a(n)) <= 4; quasi-constant
    findRoot := func(id int) int {
        for id != ufParent[ufParent[id]] {
            ufParent[id] = ufParent[ufParent[id]]
            id = ufParent[id]
        }
        return id
    }
    union := func(v1, v2 int) {
        if ufParent[v1] == ufParent[v2] {
            return
        }
        v1Root, v2Root := findRoot(v1), findRoot(v2)
        if ufSize[v1Root] <= ufSize[v2Root] {
            ufParent[v1Root] = v2Root
            ufSize[v2Root] += ufSize[v1Root]
        } else {
            ufParent[v2Root] = v1Root
            ufSize[v1Root] += ufSize[v2Root]
        }
        //ufParent[findRoot(v1)] = findRoot(v2) 
    }
    
    for _, acct := range accounts {
        id1 := emailToID[acct[1]]
        for _, email := range acct[2:] {
            union(emailToID[email], id1)
        }
    }
    
    //time: O(len(accounts) * len(acct))
    idToEmails := make(map[int][]string)
    for email, id := range emailToID {
        parId := findRoot(id)
        idToEmails[parId] = append(idToEmails[parId], email)
    }
    
    var res [][]string
    
    // e = len(emails)
    // time: O(e * nlogn)
    for _, emails := range idToEmails {
        // intro sort go std lib sort avg & worst case time: O(nlogn)
        sort.Strings(emails)
        acct := append([]string{emailToName[emails[0]]}, emails...)  
        res = append(res, acct)
    }
    
    return res
    
}