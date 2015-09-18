Each element type c<sub>j</sub> in the data stream can be transformed to a
distinct random variable R(c<sub>j</sub>) ∼ F to adequate approximation via
a pseudo-random number generator as follows:
1. Hash c<sub>j</sub> to an integer (or vector of integers) via a
deterministic hash function with low collision probability.
2. Use these integers to seed a pseudo-random number
generator.
3. Use the seeded generator to simulate a sequence of
independent random variables with distribution F; set R(c<sub>j</sub>) to
the value at a fixed position in the sequence.

~from [Random Projections](http://www.statslab.cam.ac.uk/CSI/Cosma.pdf)
