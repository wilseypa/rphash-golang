package hash;

const (
    seed =  216613626;
);

type Murmur struct {
    tablesize int64;
};

func NewMurmur(tablesize int64) *Murmur {
    return &Murmur{
        tablesize: tablesize,
    };
};

func (this *Murmur) Hash(data1 []int64) int64 {
    data := make([]byte, len(data1));
    ct := 0;
    for _, d := range data1 {
        data[ct] = byte(uint64(d) >> 56);
        ct += 1;
        data[ct] = byte(uint64(d) >> 48);
        ct += 1;
        data[ct] = byte(uint64(d) >> 40);
        ct += 1;
        data[ct] = byte(uint64(d) >> 32);
        ct += 1;
        data[ct] = byte(uint64(d) >> 24);
        ct += 1;
        data[ct] = byte(uint64(d) >> 16);
        ct += 1;
        data[ct] = byte(uint64(d) >> 8);
        ct += 1;
        data[ct] = byte(uint64(d));
        ct += 1;
    }
    m := 1540483477;
    r := uint(24);
    h := seed ^ len(data);
    len := len(data);
    len_4 := len >> 2;

    for i := 0; i < len_4; i++ {
        i_4 := i << 2;
        k := int(data[i_4 + 3]);
        k = k << 8;
        k = k | int(data[i_4 + 2] & 0xff);
        k = k << 8;
        k = k | int(data[i_4 + 1] & 0xff);
        k = k << 8;
        k = k | int(data[i_4 + 0] & 0xff);
        k *= m;
        k ^= int(uint64(k) >> r);
        k *= m;
        h *= m;
        h ^= k;
    }

    len_m := len_4 << 2;
    left := len - len_m;

    if left != 0 {
        if left >= 3 {
        h ^= int(data[len - 3] << 16);
        }
        if left >= 2 {
        h ^= int(data[len - 2] << 8);
        }
        if left >= 1 {
        h ^= int(data[len - 1]);
        }

        h *= m;
    }

    h = h ^ int(uint64(h) >> 13);
    h *= m;
    h = h ^ int(uint64(h) >> 15);

    return int64(h) % this.tablesize;
};
