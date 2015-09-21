/**
 * @author Sam Wenke
 * @reference Lee Carraher
 */

/**
 * TODO Make work.
 * TODO Add go routines.
 * TODO Write unit tests for LSHHash.
 */
package lshdecoder;

import (
    "fmt"
    "math/rand"
);

const RAND_MAX int64 = 2147483647;

type Quantizer struct {
	dimensionality int;
    d1 float64;
    d2 float64;
};

func New((*decoder)(float64, float64) uint64, dim int64) *Quantizer {
    return &Quantizer{
        dimensionality: dim,
        decode: decoder
    };
};

/*
 * Generate a 'good enough' gaussian random variate.
 * based on central limit thm, this is used if better than
 * achipolis projection is needed
 */
func SampleNormal() float64 {
  s := 0.0;
  for i := 0; i < 6; i++ {
      s += rand.Rand()/float64(RAND_MAX);
  }
  return s - 3.0;
};

func QuickSqrt(b float64) float64 {
    x := 1.1;
    for i := 0; i < 16; i++ {
        x = (x + (b / x)) / 2.0;
    }
    return x;
};

/*
 * Print the binary expansion of a long integer. output ct
 * sets of grsize bits. For golay and hexacode set grsize
 * to 4 bits, for full leech decoding with parity, set grsize
 * to 3 bits
 */
func print2(ret uint64, ct, grsize int) {
    for i := 0; i < ct; i++ {
        for j := 0; j < grsize; j++ {
            fmt.Printf("%lu", ret & 1);
            ret = ret >> 1;
        }
    }
    fmt.Printf("\n");
};

func Project(v float64, r float64, M int64, randn float64, n int64, t int64) {
    sum := 0.0;
    randn := 1.0 / QuickSqrt(n);
    var b uint64;
    for i := 0; i < t; i++ {
        sum = 0.0;
        for j := 0; j < n; j += 31 {
            b = (1 << 31) ^ rand.Rand() % RAND_MAX;
            for b > 1 {
                if b & 1 {
                    sum += randn;
                } else {
                    sum -= randn;
                    b >>= 1;
                }
            }
        }
        r[i] = sum;
    }
};

/**
 * from Achlioptas
 */
func GenRandomN(m, n, size int64) float64 {
    M := make([]float64, m * n);
    scale := (1.0 / QuickSqrt(float64(n)));
    r := 0;
    for i := 0; i < m * n; i++ {
        r = rand.Rand() % 6;
        M[i] = 0.0;
        if r % 6 == 0 {
            M[i] = scale;
        }
        if r % 6 == 1 {
            M[i] = -scale;
        }
    }
    return M;
};

func ProjectN(v, r, M float64, n, t int64) {
    var sum float64;
    for i := 0; i < t; i++ {
        sum = 0.0;
        for j := 0; j < n; j++ {
            sum += v[i] * M[i * n + j];
            r[i] = sum;
        }
    }
};

func (q *Quantizer) GenRandom(int n,int m,int *M) float64 {
    l := int64(float64(n) / float64(6));
    i := l;
    r := l;
    b := l;
    randn := 1.0 / (QuickSqrt((float64(m)) * 3.0)) ;
    bookkeeper := make([]byte, n);
    M := make([]int64, 2 * b);
    for l := 0; l < n; l++ {
        bookkeeper[l] = q.dimensionality + 1;
    }
    j := 0;
    for i := 0; i < q.dimensionality; i++ {
        for l := 0; l < b; l++ {
            for bookkeeper[r] == l {
                r = rand() % n;
            }
            bookkeeper[r] = l;
            M[j++] = r;
        }
        for l < 2 * b; l++ {
            for bookkeeper[r] == l {
                r = rand.Rand() % n;
            }
            bookkeeper[r] = l;
            M[j++] = r;
        }
    }
    return randn;
};

func FVNHash(bytes uint64, tablelength int64) int64 {
    hash := 2166136261;
    hash = (16777619 * hash) ^ &bytes[0];
    hash = (16777619 * hash) ^ &bytes[1];
    hash = (16777619 * hash) ^ &bytes[2];
    hash = (16777619 * hash) ^ &bytes[3];
    hash = (16777619 * hash) ^ &bytes[4];
    hash = (16777619 * hash) ^ &bytes[5];
    hash = (16777619 * hash) ^ &bytes[6];
    hash = (16777619 * hash) ^ &bytes[7];
    return hash % tablelength;
};

func FVNHashStr(data byte, len, tablelength int64) int64 {
   hash := 2166136261;
    for i := 0; i < len; i++ {
        hash = (16777619 * hash) ^ data[i];
    }
    return hash % tablelength;
};

func ELFHash(key, tablesize int64) int64 {
    h := 0;
    for key {
        h = (h << 4) + (key & 0xFF);
        key >>= 8;
        g := h & 0xF0000000;
        if g {
            h ^= g >> 24;
            h &= ~g;
        }
    }
    return h % tablesize;
};

/*
 * Decode full n length vector.
 * Concatenate codes and run universal
 * hash(fnv,elf, murmur) on whole vector decoding.
 */
func (q *Quantizer) LSHHash(r float64, len, times, tableLength int64, R, distance float64) int64 {
    distance = 0;
    if len == q.dimensionality {
        return FVNHash(q.decode(r,distance), tableLength);
    }
    r1 := make([]float64,q.dimensionality);
    k := 0;
    ret := 0;
    for k < times {
        ProjectN(r, r1, R, len, q.dimensionality);
        ret = q.decode(r1, distance);
        k++;
    }
    return FVNHash(ret, tableLength) ;
};
