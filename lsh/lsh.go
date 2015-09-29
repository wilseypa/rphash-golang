/**
 * Locality Sensitive Hashing
 * 2nd Step
 * @author Sam Wenke
 * @author Jacob Franklin
 */
package lsh;

import (
    "github.com/wenkesj/rphash/types"
);

func LSHHash(r []float64, hash types.Hash, decoder types.Decoder, projector types.Projector) int32{
    return hash.Hash(decoder.Decode(projector.Project(r)));
};
