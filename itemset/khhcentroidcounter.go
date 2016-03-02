package itemset;

import (
    "math"
    "math/rand"
    "time"
    "github.com/wenkesj/rphash/types"
    "github.com/wenkesj/rphash/utils"
);

const (
    width = 200000;
    depth = 7;
);

type KHHCentroidCounter struct {
    depth int;
    width int;
    sketchTable [depth][width]int;
    hashVector []int64;
    count int64;
    k int;
    frequentItems map[int64]types.Centroid;
    countlist map[int64]int64;
    priorityQueue *utils.CentriodPriorityQueue;
    topCentroid []types.Centroid;
    counts []int64;
};

func NewKHHCentroidCounter(k int) *KHHCentroidCounter {
    newK := int(float64(k) * math.Log(float64(k))) * 4;
    seed := int64(time.Now().UnixNano() / int64(time.Millisecond));
    priorityQueue := utils.NewCentroidPriorityQueue();
    frequentItems := make(map[int64]types.Centroid);
    countlist := make(map[int64]int64);
    var sketchTable [depth][width]int;
    hashVector := make([]int64, depth);
    random := rand.New(rand.NewSource(seed));
    for i := 0; i < depth; i++ {
        hashVector[i] = random.Int63();
    }
    var result = &KHHCentroidCounter{
        depth: depth,
        width: width,
        sketchTable: sketchTable,
    };
    result.hashVector = hashVector;
    result.k = newK;
    result.countlist = countlist;
    result.priorityQueue = priorityQueue;
    result.frequentItems = frequentItems;
    return result;
};

func (this *KHHCentroidCounter) Add(centroid types.Centroid) {
    this.count++;
    count := this.AddLong(centroid.GetID(), 1);
    probed := this.frequentItems[centroid.GetID()];
    delete(this.frequentItems, centroid.GetID());

    /*for i := 0; i < centroid.GetIDs().Length(); i++ {
        if probed != nil {
            break;
        }
        if centroid.GetIDs().Get(int64(i)) {
            delete(this.frequentItems, int64(i));
            probed = this.frequentItems[int64(i)];
        }
    }*/

    if probed == nil {
        this.countlist[centroid.GetID()] = count;
        this.frequentItems[centroid.GetID()] = centroid;
        this.priorityQueue.Enqueue(centroid);
    } else {
        this.priorityQueue.Remove(centroid.GetID());
        probed.UpdateVector(centroid.Centroid());
        probed.GetIDs().AddAll(centroid.GetIDs());
        this.frequentItems[probed.GetID()] = probed;
        this.countlist[probed.GetID()] = count;
        this.priorityQueue.Enqueue(probed);
    }

    if this.priorityQueue.Size() > this.k {
        removed := this.priorityQueue.Poll();
        delete(this.frequentItems, removed.GetID());
        delete(this.countlist, removed.GetID());
    }
};

func (this *KHHCentroidCounter) Hash(item int64, i int) int {
    const PRIME_MODULUS = uint64((int64(1) << 31) - 1);
    hash := uint64(this.hashVector[i] * item);
    hash += hash >> 32;
    hash &= PRIME_MODULUS;
    return int(hash % uint64(this.width));
};

/**
 * Add item hashed to a long value to count min sketch table add long comes
 * from streaminer documentation
 * @param item
 * @param count
 * @return size of min count bucket
 */
func (this *KHHCentroidCounter) AddLong(item, count int64) int64 {
    this.sketchTable[0][int(this.Hash(item, 0))] += int(count);
    min := this.sketchTable[0][int(this.Hash(item, 0))];
    for i := 1; i < depth; i++ {
        this.sketchTable[i][int(this.Hash(item, i))] += int(count);
        if this.sketchTable[i][int(this.Hash(item, i))] < min {
            min = this.sketchTable[i][int(this.Hash(item, i))];
        }
    }
    return int64(min);
};

func (this *KHHCentroidCounter) Count(item int64) int64 {
    min := this.sketchTable[0][int(this.Hash(item, 0))];
    for i := 1; i < this.depth; i++ {
        if this.sketchTable[i][int(this.Hash(item, i))] < min {
            min = this.sketchTable[i][int(this.Hash(item, i))];
        }
    }
    return int64(min);
};

func (this *KHHCentroidCounter) GetTop() []types.Centroid {
    if this.topCentroid != nil {
        return this.topCentroid;
    }
    this.topCentroid = []types.Centroid{};
    this.counts = []int64{};
    for !this.priorityQueue.IsEmpty() {
        tmp := this.priorityQueue.Poll();
        this.topCentroid = append(this.topCentroid, tmp);
        this.counts = append(this.counts, this.Count(tmp.GetID()));
    }
    return this.topCentroid;
};

func (this *KHHCentroidCounter) GetCount() int64 {
    return this.count;
};

func (this *KHHCentroidCounter) GetCounts() []int64 {
    if this.counts != nil {
        return this.counts;
    }
    this.GetTop();
    return this.counts;
};
