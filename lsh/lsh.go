package lsh;

import (
    "math/rand"
    "github.com/wenkesj/rphash/types"
);

type LSH struct {
    hash types.Hash;
    decoder types.Decoder;
    projector types.Projector;
};

func New(hash types.Hash,
            decoder types.Decoder,
            projector types.Projector) *LSH {

    return &LSH{
        hash: hash,
        decoder: decoder,
        projector: projector,
    };
};

func (this *LSH) MinHash(r []float64, radius float64, randomseed int64, n int) ([]int32, int) {

    var noise [][]float64;

    /* Generate a new source of random numbers */
    random := rand.New(rand.NewSource(randomseed));

    /* Project a vector into a smaller dimension
     * Decode the vector to determine its counterpart
     * Calculate lengths */
    projectedVector := this.projector.Project(r);
    noNoise := this.decoder.Decode(projectedVector);
    nLength, rLength, pLength := len(noNoise), len(r), len(projectedVector);

    /* Create a matrix of random vectors which will
     * symbolize a noise matrix. */
    for h := 1; h < n; h++ {
        tempVector := make([]float64, rLength);
        for i := 0; i < rLength; i++ {
            tempVector[i] = random.NormFloat64() * radius;
        }
        noise = append(noise, tempVector);
    }

    /* Formulate a result. */
    result := make([]int32, n * nLength);
    count := copy(result, noNoise);
    rTempVector := make([]float64, pLength);
    for j := 1; j < n; j++ {
        count = copy(rTempVector, projectedVector);
        tempVector := noise[j - 1];
        for k := 0; k < pLength; k++ {
            rTempVector[k] = rTempVector[k] + tempVector[k];
        }
        noNoise = this.decoder.Decode(rTempVector);
        nLength = len(noNoise);
        count = copy(result[j * nLength : j * nLength + nLength], noNoise);
    }
    return result, count;
};
