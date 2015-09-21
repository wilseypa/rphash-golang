/**
 * @author Sam Wenke
 * @reference Timo H Tossavainen
 */

/**
 * TODO Add go routines.
 */
package fwht;

func wht_bfly(a, b int) {
    tmp := a;
    a += b;
    b = tmp - b;
};

func l2(x int) int {
    for l2 := 0; x > 0; x >>= 1 {
        l2++;
    }
    return l2;
};

/**
 * Basically like Fourier, but the basis functions
 * are squarewaves with different sequencies.
 * @param {[]float64} data.
 */
func FWHT(data []float64) {
    log2 := l2(len(data)) - 1;
    for i := 0; i < log2; i++ {
        for j := 0; j < (1 << log2); j += 1 << (i + 1) {
            for k := 0; k < (1 << i); k++ {
                wht_bfly(data[j + k], data[j + k + (1 << i)]);
            }
        }
    }
};
