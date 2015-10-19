package itemset;
//
// import (
//     "github.com/wenkesj/rphash/types"
// );
//
// type KHHCountMinSketch struct {
//     depth int;
//     width int;
//     table [][]int32;
//     hashA []int32;
//     size int32;
//     p types.PQueue;
//     k int;
//     items map[int32]interface{};
//     countlist map[int32]int32;
//     topcent []interface{};
//     counts []int32;
// };
//
// func NewKHHCountMinSketch(m int) *KHHCountMinSketch {
//     newK := int(m * math.Log(float64(k))) * 4;
//     epsOfTotalCount := 0.00001;
//     confidence := 0.99;
//     seed := int(time.Now().UnixNano() / int64(time.Millisecond));
//     countlist := make(map[int32]int32);
//
//     priorityQueue := utils.NewPQueue(countlist);
//     items := make(map[int32]interface{});
//     width := int(math.Ceil(2 / epsOfTotalCount));
//     depth := int(math.Ceil(-math.Log(float64(1 - confidence)) / math.Log(2)));
//     table := make([depth][width]int);
//     hashA := make([]int32, depth);
//     random := rand.New(rand.NewSource(seed));
//     for i := 0; i < depth; i++ {
//         hashA[i] = random.Int32();
//     }
//     return &KHHCountMinSketch{
//         k: newK,
//         count: 0,
//         topcent: nil,
//         counts: nil,
//         countlist: countlist,
//         priorityQueue: priorityQueue,
//         items: items,
//         width: width,
//         depth: depth,
//         table: table,
//         hashA: hashA,
//     };
// };
//
// func (this *KHHCountMinSketch) Add(e interface{}) bool {
//     count := this.AddLong(c.HashCode(), 1);
//     if e.(type) == types.Centroid {
//         c := e.(types.Centroid);
//         delete(this.items, c.id);
//         probed := this.items[c.id];
//
//         if probed == nil {
//             this.countlist[c.id] = count;
//             this.priorityQueue.Add(e);
//             this.items[int32(c.id)] = e;
//         } else {
//             this.priorityQueue.Remove(probed);
//             this.items[(probed.(types.Centroid)).id] = probed;
//             this.countlist[(probed.(types.Centroid)).id] = count;
//             this.priorityQueue.Add(probed);
//         }
//     } else {
//         if val, ok := items[e.HashCode().(int32)]; !ok {
//             this.countlist[e.HashCode().(int32)] = count;
//             this.priorityQueue.Add(e);
//             this.items[e.HashCode().(int32)] = e;
//         } else {
//             this.priorityQueue.Remove(e);
//             this.items[e.HashCode().(int32)] = e;
//             this.countlist[e.HashCode().(int32)] = count;
//             this.priorityQueue.Add(e);
//         }
//     }
//     if this.priorityQueue.Size() > this.k {
//         removed := this.priorityQueue.Poll();
//         delete(this.items, removed);
//     }
//     return false;
// };
//
// func (this *KHHCountMinSketch) Hash(item int32, int i) int {
//     const PRIME_MODULUS = (1 << 31) - 1;
//     var hash uint32 = this.hashA[i] * item;
//     hash += hash >> 32;
//     hash &= PRIME_MODULUS;
//     return int(hash % this.width);
// };
//
// func (this *KHHCountMinSketch) AddLong(item, count int32) int32 {
//     this.table[0][this.Hash(item, 0)] += count;
//     min := int(this.table[0][this.Hash(item, 0)]);
//     for i := 1; i < depth; i++ {
//         this.table[i][this.Hash(item, i)] += count;
//         if this.table[i][this.Hash(item, i)] < min {
//             min = int(this.table[i][this.Hash(item, i)]);
//         }
//     }
//     this.size += count;
//     return min;
// };
//
// func (this *KHHCountMinSketch) GetTop() []interface{} {
//     if this.topcent != nil {
//         return this.topcent;
//     }
//     this.topcent = make([]interface{});
//     this.counts = make([]interface{});
//     for !this.priorityQueue.IsEmpty() {
//         tmp := this.priorityQueue.Poll();
//         append(this.topcent, tmp);
//         append(this.countlist, countlist[tmp.HashCode().(int32)]);
//     }
//     return this.topcent;
// };
//
// func (this *KHHCountMinSketch) GetCounts() []int32 {
//     if this.counts != nil {
//         return this.counts;
//     }
//     this.GetTop();
//     return this.counts;
// };
