package itemset;

type KHHCountMinSketch struct {
    count int32;
};

func NewKHHCountMinSketch(k int) *KHHCountMinSketch {
    return &KHHCountMinSketch{};
};

func (this *KHHCountMinSketch) Add(c int32) { };

func (this *KHHCountMinSketch) GetCount() int32 {
    return this.count;
};

func (this *KHHCountMinSketch) GetCounts() (a []int32) {
    return a;
};

func (this *KHHCountMinSketch) GetTop() (a []int32) {
    return a;
};
