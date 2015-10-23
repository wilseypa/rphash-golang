package itemset;

type KHHCountMinSketch struct {
    count int64;
};

func NewKHHCountMinSketch(k int) *KHHCountMinSketch {
    return &KHHCountMinSketch{};
};

func (this *KHHCountMinSketch) Add(c int64) { };

func (this *KHHCountMinSketch) GetCount() int64 {
    return this.count;
};

func (this *KHHCountMinSketch) GetCounts() (a []int64) {
    return a;
};

func (this *KHHCountMinSketch) GetTop() (a []int64) {
    return a;
};
