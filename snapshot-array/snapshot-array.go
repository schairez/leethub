type data struct {
    id, val int
}

type snapPairs []data

type SnapshotArray struct {
    snapStore []snapPairs
    snapId int
}


func Constructor(length int) SnapshotArray {
    snapStore := make([]snapPairs, length)
    return SnapshotArray{snapStore, 0}
}


func (this *SnapshotArray) Set(index int, val int)  {
    if index >= len(this.snapStore) {
        return
    }
    snapValsAtIdx := this.snapStore[index]
    if size := len(snapValsAtIdx); size == 0 || 
    snapValsAtIdx[size-1].id != this.snapId {
        dataPair := data{this.snapId, val}
        this.snapStore[index] = append(this.snapStore[index], dataPair)
    } else {
        this.snapStore[index][len(snapValsAtIdx)-1].val = val
    }
}


func (this *SnapshotArray) Snap() int {
    this.snapId++
    return this.snapId-1
}

func (this *SnapshotArray) Get(index int, snap_id int) int {
    //idxSnapIdValPairs := this.snapStore[index]
    snapId := snap_id //idiomatic go
    size := len(this.snapStore[index])
    if size == 0 { 
        return 0 
    }
    if this.snapStore[index][0].id > snapId {
        return 0
    } 
    if this.snapStore[index][size-1].id <= snapId {
        return this.snapStore[index][size-1].val
    }
    lo, hi := 0, size-1
    for lo <= hi {
        mid := (lo+hi) >> 1
        midId := this.snapStore[index][mid].id
        if midId == snapId {
            return this.snapStore[index][mid].val
        } else if midId < snapId {
            lo = mid+1
        } else {
            hi = mid-1
        }
    }
    return this.snapStore[index][hi].val
}


/**
 * Your SnapshotArray object will be instantiated and called as such:
 * obj := Constructor(length);
 * obj.Set(index,val);
 * param_2 := obj.Snap();
 * param_3 := obj.Get(index,snap_id);
 */