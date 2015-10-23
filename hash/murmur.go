package hash;

type Murmur struct {
    tablesize int64;
};

func NewMurmur(tablesize int64) *Murmur {
    return &Murmur{
        tablesize: tablesize,
    };
};

func (this *Murmur) Hash(data []int64) (hash int64) {
    return hash;
};
