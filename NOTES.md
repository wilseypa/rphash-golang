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

Laplacian Eigenmaps (LE) dimensionality reduction

~from [Improved time series land cover classification by missing-observation-adaptive nonlinear dimensionality reduction](http://ac.els-cdn.com/S0034425714004763/1-s2.0-S0034425714004763-main.pdf?_tid=42611a9e-609d-11e5-af6e-00000aab0f6b&acdnat=1442866604_9e32fc021246cb8c7c7f2db17d7fed3c)

~from [A discrete approach to stochastic parametrization and dimensional reduction in nonlinear dynamics](http://arxiv.org/pdf/1503.08766v1.pdf)
