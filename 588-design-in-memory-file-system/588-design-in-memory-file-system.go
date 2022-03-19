

type fsNode struct {
    key string
    isDir bool
    children map[string]*fsNode
    contents *strings.Builder 
}

func (fs *fsNode) addChildNode(fsPath []string) *fsNode {
    for _, key := range fsPath {
        if _, exists := fs.children[key]; !exists {
            fs.children[key] = &fsNode{key:key,
                                       isDir: true, 
                                       children: make(map[string]*fsNode),
                                      }
        }
        fs = fs.children[key]
        
    }
    return fs
}


func (fs *fsNode) FindNode(path []string) *fsNode {
    //condition: no derefenrce nil ptr
    for i := 0; i < len(path); i++ {
        fs = fs.children[path[i]]
    }
    return fs
}


type FileSystem struct {
    root *fsNode
}


func Constructor() FileSystem {
    return FileSystem{root: &fsNode{key:"/", isDir: true,
                                    children: make(map[string]*fsNode)}}
}


func (this *FileSystem) Ls(path string) []string {
    var lvlNodes []string 
    
    getNodes := func(currFS *fsNode) []string {
        if currFS.isDir {
            for nodeNames := range currFS.children {
                lvlNodes = append(lvlNodes, nodeNames)
            }
        } else {
            return []string{currFS.key}
        }
        if len(lvlNodes) == 0 {
            return lvlNodes
        }
        sort.Strings(lvlNodes) // O(nlogn); size of keys
        return  lvlNodes
    }
    // condition: no err in path
    pathVals := splitFilePath(path)
    if len(pathVals) == 0 { // if rootfs
        lvlNodes = getNodes(this.root)
    } else {
        currNode := this.root.FindNode(pathVals)
        lvlNodes = getNodes(currNode)
    }
    
    return lvlNodes
}


func (this *FileSystem) Mkdir(path string)  {
    pathVals := splitFilePath(path)
    this.root.addChildNode(pathVals)
    
}

func (this *FileSystem) AddContentToFile(filePath string, content string)  {
    pathVals := splitFilePath(filePath)
    n := len(pathVals)
    currNode := this.root.addChildNode(pathVals[:n-1])
    fileName := pathVals[n-1] 
    fileNode, exists := currNode.children[fileName]
    if exists {
        fileNode.contents.WriteString(content)
        return
    }
    
    var sb strings.Builder
    sb.WriteString(content)
    currNode.children[pathVals[n-1]] = &fsNode{key: fileName, isDir:false, contents: &sb}
    
}

func splitFilePath(path string) []string {
    if path == "/" {
        return []string{}
    }
    pathVals := strings.Split(path, "/")
    return pathVals[1:]
}


func (this *FileSystem) ReadContentFromFile(filePath string) string {
    pathVals := splitFilePath(filePath)
    currNode := this.root.addChildNode(pathVals)
    //condition: path is valid and last val is file
    
    return currNode.contents.String()
}


/**
 * Your FileSystem object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ls(path);
 * obj.Mkfs(path);
 * obj.AddContentToFile(filePath,content);
 * param_4 := obj.ReadContentFromFile(filePath);
 */