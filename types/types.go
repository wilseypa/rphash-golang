/**
 * Types package.
 * Contacts interfaces for the core steps.
 * @author Sam Wenke
 * @author Jacob Franklin
 */
package types;

type Decoder interface {
    SetVariance(parameterObject float64);
    GetDimensionality() int;
    Decode(f []float64) []int32;
    GetErrorRadius() float64;
    GetDistance() float64;
};

type Projector interface {
    Project(v []float64) []float64;
};

type Hash interface {
    Hash(k []int32) int32;
};
