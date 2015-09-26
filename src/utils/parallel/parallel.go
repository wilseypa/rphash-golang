/**
 * @author Sam Wenke
 * @reference Daniël de Kok
 */

/**
 * TODO Add methods for Map[DataType].
 * TODO Write unit tests for Map.
 */
package parallel;

import (
    "runtime"
);

type empty struct{};
type semaphore chan empty;

/**
 * Parallel For.
 */
func For(begin, end, step uint, f func(uint)) {
    cpus := uint(runtime.GOMAXPROCS(0));
    sem := make(semaphore, cpus);

    for i := uint(0); i < cpus; i++ {
        go func(sem semaphore, cpus, begin, end, step uint, f func(uint)) {
            for i := begin; i < end; i += (cpus * step) {
                f(i);
            }
            sem <- empty{};
        }(sem, cpus, begin+(i*step), end, step, f);
    }

    for i := uint(0); i < cpus; i++ {
        <- sem;
    }
};

/**
 * Parallel Map.
 */
func MapFloat64(f func(float64) float64, l []float64) []float64 {
    result := make([]float64, len(l));
    For(0, uint(len(l)), 1, func(idx uint) {
        result[idx] = f(l[idx]);
    });
    return result;
};
