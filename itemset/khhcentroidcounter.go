package itemset;

import (
    "math"
    "math/rand"
    "time"
    "github.com/wenkesj/rphash/types"
    "github.com/wenkesj/rphash/utils"
);

const (
    width = 200000; // int(math.Ceil(2/epsOfTotalCount))
    depth = 7; // int(math.Ceil(-math.Log(float64(1 - confidence)) / math.Log(2)))
);

type KHHCentroidCounter struct {
    depth int;
    width int;
    table [depth][width]int;
    hashA []int32;
    count int32;
    k int;
    origk int;
    frequentItems map[int32]types.Centroid;
    countlist map[int32]int32;
    priorityQueue *utils.PQueue;
    topcent []types.Centroid;
    counts []int32;
};

func NewKHHCentroidCounter(k int) *KHHCentroidCounter {
    newK := int(float64(k) * math.Log(float64(k))) * 4;
    seed := int64(time.Now().UnixNano() / int64(time.Millisecond));
    var countlist map[int32]int32;
    priorityQueue := utils.NewPQueue(countlist);
    var frequentItems map[int32]types.Centroid;
    var table [depth][width]int;
    hashA := make([]int32, depth);
    random := rand.New(rand.NewSource(seed));
    for i := 0; i < depth; i++ {
        hashA[i] = random.Int31();
    }
    var result = new(KHHCentroidCounter);
    result.depth = depth;
    result.width = width;
    result.table = table;
    result.hashA = hashA;
    result.k = newK;
    result.origk = k;
    result.countlist = countlist;
    result.priorityQueue = priorityQueue;
    result.frequentItems = frequentItems;
    return result;
};

func (this *KHHCentroidCounter) Add(c types.Centroid) {
    this.count++;
    count := this.AddLong(c.GetID(), 1);
    delete(this.frequentItems, c.GetID());
    probed := this.frequentItems[c.GetID()];

    for i := 0; i < c.GetIDs().Length(); i++ {
        if probed != nil {
            break;
        }
        if c.GetIDs().Get(int32(i)) {
            delete(this.frequentItems, int32(i));
            probed = this.frequentItems[int32(i)];
        }
    }

    if probed == nil {
        this.countlist[c.GetID()] = count;
        this.frequentItems[c.GetID()] = c;
        this.priorityQueue.Add(c);
    } else {
        this.priorityQueue.Remove(probed);
        probed.UpdateVector(c.Centroid());
        probed.GetIDs().AddAll(c.GetIDs());
        this.frequentItems[probed.GetID()] = probed;
        this.countlist[probed.GetID()] = count + 1;
        this.priorityQueue.Add(probed);
    }

    if this.priorityQueue.Size() > this.k {
        removed := this.priorityQueue.Poll();
        delete(this.frequentItems, removed.GetID());
        delete(this.countlist, removed.GetID());
    }
};

func (this *KHHCentroidCounter) Hash(item int32, i int) int {
    const PRIME_MODULUS = (1 << 31) - 1;
    hash := uint32(this.hashA[i] * item);
    hash += hash >> 32;
    hash &= PRIME_MODULUS;
    return int(hash % uint32(this.width));
};

/**
 * Add item hashed to a long value to count min sketch table add long comes
 * from streaminer documentation
 * @param item
 * @param count
 * @return size of min count bucket
 */
func (this *KHHCentroidCounter) AddLong(item, count int32) int32 {
    this.table[0][int(this.Hash(item, 0))] += int(count);
    min := this.table[0][int(this.Hash(item, 0))];
    for i := 1; i < depth; i++ {
        this.table[i][int(this.Hash(item, i))] += int(count);
        if this.table[i][int(this.Hash(item, i))] < min {
            min = this.table[i][int(this.Hash(item, i))];
        }
    }
    return int32(min);
};

func (this *KHHCentroidCounter) Count(item int32) int32 {
    min := this.table[0][int(this.Hash(item, 0))];
    for i := 1; i < this.depth; i++ {
        if this.table[i][int(this.Hash(item, i))] < min {
            min = this.table[i][int(this.Hash(item, i))];
        }
    }
    return int32(min);
};

func (this *KHHCentroidCounter) GetTop() []types.Centroid {
    if this.topcent != nil {
        return this.topcent;
    }
    this.topcent = []types.Centroid{};
    this.counts = []int32{};
    for !this.priorityQueue.IsEmpty() {
        tmp := this.priorityQueue.Poll();
        this.topcent = append(this.topcent, tmp);
        this.counts = append(this.counts, this.Count(tmp.GetID()));
    }
    return this.topcent;
};

func (this *KHHCentroidCounter) GetCounts() []int32 {
    if this.counts != nil {
        return this.counts;
    }
    this.GetTop();
    return this.counts;
};
