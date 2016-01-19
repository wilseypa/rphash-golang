package itemset;

import (
    "math"
    "math/rand"
    "time"
    "github.com/wenkesj/rphash/utils"
);

type KHHCountMinSketch struct {
    depth int;
    width int;
    sketchTable [depth][width]int64;
    hashVector []int64;
    size int64;
    priorityQueue *utils.Int64PriorityQueue;
    k int;
    items map[int64]int64;
    countlist map[int64]int64;
    count int64;
    counts []int64;
    topCentroid []int64;
};

func NewKHHCountMinSketch(m int) *KHHCountMinSketch {
    k := int(float64(m) * math.Log(float64(m)));
    seed := int64(time.Now().UnixNano() / int64(time.Millisecond));
    countlist := make(map[int64]int64);
    items := make(map[int64]int64);
    var sketchTable [depth][width]int64;
    hashVector := make([]int64, depth);
    random := rand.New(rand.NewSource(seed));
    for i := 0; i < depth; i++ {
        hashVector[i] = random.Int63n(2147483647);
    }
    result := new(KHHCountMinSketch);
    result.k = k;
    result.countlist = countlist;
    result.items = items;
    result.sketchTable = sketchTable;
    result.width = width;
    result.depth = depth;
    result.size = 0;
    result.hashVector = hashVector;
    result.priorityQueue = utils.NewInt64PriorityQueue();
    result.topCentroid = nil;
    return result;
};

func (this *KHHCountMinSketch) Hash(item int64, i int) int {
    PRIME_MODULUS := int64(1 << 63 - 1);
    hash := this.hashVector[i] * item;
    hash += hash >> 64;
    hash &= PRIME_MODULUS;
    return int(hash) % this.width;
};

func (this *KHHCountMinSketch) Add(e int64) {
    count := this.AddLong(utils.HashCode(e), 1);
    if _, ok := this.items[utils.HashCode(e)]; !ok {
        this.countlist[utils.HashCode(e)] = count;
        this.priorityQueue.Enqueue(e);
        this.items[utils.HashCode(e)] = e;
    } else {
        this.priorityQueue.Dequeue();
        this.items[utils.HashCode(e)] = e;
        this.countlist[utils.HashCode(e)] = count;
        this.priorityQueue.Enqueue(e);
    }

    if this.priorityQueue.Size() > this.k {
        removed := this.priorityQueue.Poll();
        delete(this.items, removed);
    }
};

func (this *KHHCountMinSketch) AddLong(item, count int64) int64 {
    this.sketchTable[0][this.Hash(item, 0)] += count;
    min := int64(this.sketchTable[0][this.Hash(item, 0)]);
    for i := 1; i < this.depth; i++ {
        this.sketchTable[i][this.Hash(item, i)] += count;
        if this.sketchTable[i][this.Hash(item, i)] < min {
            min = int64(this.sketchTable[i][this.Hash(item, i)]);
        }
    }
    this.size += count;
    return min;
};

func (this *KHHCountMinSketch) GetCount() int64 {
    return this.count;
};

func (this *KHHCountMinSketch) GetCounts() []int64 {
    if this.counts != nil {
        return this.counts;
    }
    this.GetTop();
    return this.counts;
};

func (this *KHHCountMinSketch) GetTop() []int64 {
    if this.topCentroid != nil {
        return this.topCentroid;
    }
    this.topCentroid = []int64{};
    this.counts = []int64{};
    for !this.priorityQueue.IsEmpty() {
        tmp := this.priorityQueue.Poll();
        this.topCentroid = append(this.topCentroid, tmp);
        this.counts = append(this.counts, this.countlist[utils.HashCode(tmp)]);
    }
    return this.topCentroid;
};
