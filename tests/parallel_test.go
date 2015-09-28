/** Test package parallel */
package rphash;

import (
    "log"
    "time"
    "runtime"
    "testing"
    "utils/parallel"
);

func primeSieve(n uint) []bool {
    sieve := make([]bool, n+1);
    half := n / 2;

    for idx := range sieve {
        sieve[idx] = true;
    }

    for i := uint(2); i <= half; i++ {
        if sieve[i] {
            for j := i + i; j <= n; j += i {
                sieve[j] = false;
            }
        }
    }

    return sieve;
}

func primes(n uint) []uint {
    r := make([]uint, 0);

    sieve := primeSieve(n);

    for i := uint(2); i <= n; i++ {
        if sieve[i] {
            r = append(r, i);
        }
    }

    return r
}

func primeFactors(primes []uint, n uint) []uint {
    factors := make([]uint, 0);

    for {
        if n <= 1 {
            break;
        }

        d := n / primes[0];
        r := n % primes[0];

        if r == 0 {
            factors = append(factors, primes[0]);
            n = d;
        } else {
            if len(primes) == 1 {
                panic("not enough primes available to factorize number");
            }
            primes = primes[1:];
        }
    }

    return factors
}

func forLinear(begin, end, step uint, f func(uint)) error {
    for i := begin; i < end; i += step {
        f(i);
    }

    return nil;
}

func TestModule7(t *testing.T) {
    nC := runtime.NumCPU();
    runtime.GOMAXPROCS(nC);

    log.Println("The number of CPUs allocated: ",nC);

    var max uint = 1000;

    p := primes(uint(max));
    factors := make([][]uint, max);

    f := func(n uint) {
        factors[n] = primeFactors(p, uint(n));
    };

    timeStart := time.Now();
    parallel.For(0, max, 2, f);
    log.Println("Interleaved (Parallel): ",time.Now().Sub(timeStart));

    timeStart = time.Now();
    forLinear(0, max, 2, f);
    log.Println("Linear (Non-Parallel): ",time.Now().Sub(timeStart));
    log.Println("parallel\x1b[32;1m √\x1b[0m");
}
