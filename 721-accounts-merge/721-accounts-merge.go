

func accountsMerge(accounts [][]string) [][]string {
    emailToID := make(map[string]int)
    emailToName := make(map[string]string)
    for _, acct := range accounts {
        name := acct[0]
        for _, email := range acct[1:] {
            if _, exists := emailToID[email]; !exists {
                emailToID[email] = len(emailToID)
                emailToName[email] = name
            }
        }
    }
    //ufSize := make([]int, len(emailToID)) 
    ufParent := make([]int, len(emailToID)) 
    for idx := range ufParent {
        ufParent[idx] = idx
    }
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
        ufParent[findRoot(v1)] = findRoot(v2) 
    }
    
    for _, acct := range accounts {
        id1 := emailToID[acct[1]]
        for _, email := range acct[2:] {
            union(emailToID[email], id1)
        }
    }
    
    idToEmails := make(map[int][]string)
    for email, id := range emailToID {
        parId := findRoot(id)
        idToEmails[parId] = append(idToEmails[parId], email)
    }
    
    var res [][]string
    for _, emails := range idToEmails {
        sort.Strings(emails)
        acct := append([]string{emailToName[emails[0]]}, emails...)  
        res = append(res, acct)
    }
    
    return res
    
}