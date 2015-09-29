/** Test package for fjlt */
package rphash;

import (
    "log"
    "testing"
    "github.com/wenkesj/rphash/projection/fjlt"
);

func TestModule6(t *testing.T) {
    data := []float64{1.0,0.0,2.0,7.0,4.0,0.0,8.0,3.0,2.0,1.0};
    var d, k, n int64 = 10, 2, 15;
    f := fjlt.New(d, k, n);
    log.Println(f.Project(data));
    log.Println("fjlt\x1b[32;1m √\x1b[0m");
}
