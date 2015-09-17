/**
 * @author Sam Wenke
 * Modified version of the Leach Array Decoder written in C by Lee Carraher.
 * @from Lee Carraher:
 *  Leech Decoder uses a rotated Z2 lattice, so to find leading cosets
 *  just find the nearest point in 64QAM, A,B ; odd, even| to the rotated
 *  input vector
 *  rotate using the standard 2d rotation transform
 *                       [cos x -sin x ]
 *                   R = [sin x  cos x ]    cos(pi/4) = sin(pi/4)=1/sqrt(2)
 *  for faster C implementation use these binary fp constants
 *  1/sqrt(2) = cc3b7f669ea0e63f ieee fp little endian
 *            = 3fe6a09e667f3bcc ieee fp big endian
 *            = 0.7071067811865475244008
 *
 *  v' = v * R
 *  integer lattice
 *
 *   4 A000 B000 A110 B110 | A000 B000 A110 B110
 *   3 B101 A010 B010 A101 | B101 A010 B010 A101
 *   2 A111 B111 A001 B001 | A111 B111 A001 B001
 *   1 B011 A100 B100 A011 | B011 A100 B100 A011
 *     --------------------|---------------------
 *  -1 A000 B000 A110 B110 | A000 B000 A110 B110
 *  -2 B101 A010 B010 A101 | B101 A010 B010 A101
 *  -3 A111 B111v A001 B001 | A111 B111 A001 B001
 *  -4 B011 A100 B100 A011 | B011 A100 B100 A011
 *  even pts {000,110,111,001}
 *  odd  pts {010,101,100,011}
 */

package leacharraydecoder;
