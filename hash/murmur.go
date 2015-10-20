package hash;

type Murmur struct { hashMod int32 };

func NewMurmur(hashMod int32) *Murmur {
    return &Murmur{hashMod};
};

func (this *Murmur) Hash(key []int32) (hash int32) {
    return hash;
};
