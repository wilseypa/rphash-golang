/** Test package for randomprojection */
package rphash;

import (
    "testing"
    "log"
    "github.com/wenkesj/rphash/projection/rp"
);

func TestModule1(t *testing.T) {
    data := []float64{1.0,0.0,2.0,7.0,4.0,0.0,8.0,3.0,2.0,1.0};
    var d, k int = 10, 2;
    var n int64 = 15;
    RP := rp.New(d, k, n);
    log.Println(RP.Project(data));
    log.Println("rp\x1b[32;1m √\x1b[0m");
}
