package clusterer;

import (
    "fmt"
    "math/rand"
    "github.com/wenkesj/rphash/reader"
    "github.com/wenkesj/rphash/utils"
    "github.com/wenkesj/rphash/projector"
    "github.com/wenkesj/rphash/types"
);

type KMeans struct {
    k int;
    n int;
    data [][]float64;
    projdim int;
    means [][]float64;
    clusters [][]int;
    weights []int32;
};

func NewKMeans(k int, data [][]float64, weights []int32) *KMeans{
    return &KMeans{
        k: k,
        data: data,
        projdim: 0,
        clusters: nil,
        weights: weights,
    };
};

func (this *KMeans) ComputerCentroid(vectors []int, data [][]float64) []float64 {
    d := len(data[0]);
    centroid := make([]float64, d);
    for i := 0; i < d; i++ {
        centroid[i] = 0.0;
    }
    var w_total int32 = 0;
    for _, v := range vectors {
        w_total += this.weights[v];
    }
    for _, v := range vectors {
        vec := data[v];
        weight := float64(this.weights[v])/float64(w_total);
        for i := 0; i < d; i++ {
            centroid[i] += (vec[i] * weight);
        }
    }
    return centroid;
};

func (this *KMeans) UpdateMeans(data [][]float64) {
    for i := 0; i < this.k; i++ {
        this.means[i] = this.ComputerCentroid(this.clusters[i], data);
    }
};

func (this *KMeans) AssignClusters(data [][]float64) int {
    swaps := 0;
    newClusters := [][]int{};
    for j := 0; j < this.k; j++ {
        newClusterList := []int{};
        newClusters = append(newClusters, newClusterList);
    }
    for clusterid := 0; clusterid < this.k; clusterid++ {
        for _, member := range this.clusters[clusterid] {
            nearest := utils.FindNearestDistance(data[member], this.means);
            newClusters[nearest] = append(newClusters[nearest], member);
            if nearest != clusterid {
                swaps++;
            }
        }
    }
    this.clusters = newClusters;
    return swaps;
};

func (this *KMeans) Run() {
    maxiters := 10000;
    swaps := 3;
    fulldata := this.data;
    data := [][]float64{};
    var p types.Projector = nil;
    if this.projdim != 0 {
        p = projector.NewDBFriendly(len(fulldata[0]), this.projdim, rand.Int63());
    }
    for _, v := range fulldata {
        if p != nil {
            data = append(data, p.Project(v));
        } else {
            data = append(data, v);
        }
    }
    this.n = len(data);
    this.means = make([][]float64, this.k);
    for i := 0; i < this.k; i++ {
        this.means = append(this.means, data[i * (this.n / this.k)]);
    }
    this.clusters = make([][]int, this.k);
    for i := 0; i < this.k; i++ {
        tmp := make([]int, this.n / this.k);
        start := i * (this.n / this.k);
        for j := 0; j < this.n / this.k; j++ {
            tmp = append(tmp, j + start);
        }
        this.clusters = append(this.clusters, tmp);
    }
    for swaps > 2 && maxiters > 0 {
        maxiters--;
        this.UpdateMeans(data);
        swaps = this.AssignClusters(data);
    }
    if maxiters == 0 {
        fmt.Println("Warning: Max Iterations Reached");
    }
    data = fulldata;
    this.UpdateMeans(data);
};

func (this *KMeans) GetCentroids() [][]float64 {
    if this.means == nil {
        this.Run();
    }
    return this.means;
};

func (this *KMeans) GetParam() types.RPHashObject {
    return reader.NewSimpleArray(this.data, this.k);
};
