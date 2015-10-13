package itemset;

import (
    "math"
    "math/rand"
    "time"
    "github.com/wenkesj/rphash/types"
    "github.com/wenkesj/rphash/utils"
);

const PRIME_MODULUS = (1 << 31) - 1;

type KHHCentroidCounter struct {
    depth int;
    width int;
    table [][]int;
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
    newK := int(k * math.Log(float64(k))) * 4;
    epsOfTotalCount := 0.00001;
    confidence := 0.99;
    seed := int(time.Now().UnixNano() / int64(time.Millisecond));
    var countlist map[int32]int32;

     /* TODO: Need Comparator in Heap. */
    // cmp := new Comparator<Centroid>() {
    //     @Override
    //     public int compare(Centroid n1, Centroid n2) {
    //         long cn1 = countlist.get(n1.id);// count(n1.id);
    //         long cn2 = countlist.get(n2.id);// count(n2.id);
    //         if (cn1 > cn2)
    //             return +1;
    //         else if (cn1 < cn2)
    //             return -1;
    //         return 0;
    //     }
    // };

    priorityQueue := utils.NewPQueue(/*cmp*/);
    var frequentItems map[int32]types.Centroid;
    width := int(math.Ceil(2 / epsOfTotalCount));
    depth := int(math.Ceil(-math.Log(float64(1 - confidence)) / math.Log(2)));
    var table [depth][width]int;
    hashA := make([]int32, depth);
    random := rand.New(rand.NewSource(seed));
    for i := 0; i < depth; i++ {
        hashA[i] = random.Int32();
    }
    return &KHHCentroidCounter{
        origk: k,
        k: newK,
        count: 0,
        topcent: nil,
        counts: nil,
        countlist: countlist,
        priorityQueue: priorityQueue,
        frequentItems: frequentItems,
        width: width,
        depth: depth,
        table: table,
        hashA: hashA,
    };
};

func (this *KHHCentroidCounter) Add(c types.Centroid) {
	this.count++;
	count := this.AddLong(c.id, 1);
    delete(frequentItems, c.id);
    probed = frequentItems[c.id];

	for _, h := range c.ids {
		if probed != nil {
			break;
		}
        delete(frequentItems, h);
		probed = frequentItems[h];
	}

	if probed == nil {
		countlist[c.id] = count;
		frequentItems[c.id] = c;
		priorityQueue.Add(c);
	} else {
		priorityQueue.Remove(probed);
		probed.UpdateVector(c.Centroid());
		probed.ids.AddAll(c.ids);
		frequentItems[probed.id] = probed;
		countlist[probed.id] = count + 1;
		priorityQueue.Add(probed);
	}

	if priorityQueue.Size() > this.k {
		removed := priorityQueue.Poll();
		delete(frequentItems, removed.id);
		delete(countlist, removed.id);
	}
};

func (this *KHHCentroidCounter) Hash(item int32, int i) int {
	var hash uint32 = this.hashA[i] * item;
	hash += hash >> 32;
	hash &= PRIME_MODULUS;
	return int(hash % this.width);
};

/**
 * Add item hashed to a long value to count min sketch table add long comes
 * from streaminer documentation
 * @param item
 * @param count
 * @return size of min count bucket
 */
func (this *KHHCentroidCounter) AddLong(item, count int32) int32 {
	this.table[0][this.Hash(item, 0)] += count;
    min := int(this.table[0][this.Hash(item, 0)]);
	for i := 1; i < depth; i++ {
		this.table[i][this.Hash(item, i)] += count;
		if this.table[i][this.Hash(item, i)] < min {
            min = int(this.table[i][this.Hash(item, i)]);
        }
	}
	return min;
};

func (this *KHHCentroidCounter) Count(item int32) int32 {
    min := int(this.table[0][this.Hash(item, 0)]);
	for i := 1; i < this.depth; i++ {
		if this.table[i][this.Hash(item, i)] < min {
            min = int(this.table[i][this.Hash(item, i)]);
        }
	}
	return min;
};

func (this *KHHCentroidCounter) GetTop() []types.Centroid {
	if this.topcent != nil {
        return this.topcent;
    }
    this.topcent = []types.Centroid{};
    this.counts = []int32{};
	for !this.priorityQueue.IsEmpty() {
        tmp := this.priorityQueue.Poll();
		append(this.topcent, tmp);
		append(this.counts, this.Count(tmp.id));
	}
	return this.topcent;
};

func (this *KHHCentroidCounter) GetCounts() [][]float64 {
	if this.counts != nil {
        return this.counts;
    }
	this.GetTop();
	return this.counts;
};
