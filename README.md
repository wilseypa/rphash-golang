# Scalable Big Data Clustering by Random Projection Hashing #
+ Sam Wenke, Jacob Franklin, Sadiq Quasem
+ Advised by Dr. Phillip A. Wilsey

## Description ##
Clustering aims to solve a significant problem; finding structure in a collection of unlabeled and unconstrained streams of data. Scalable clusters of large sets of data are widely applied in today’s cumulonimbus of networking. The importance of cloud computing is tangible, and is growing at an exponential rate. The answer to clustering can only be approached by the most carefully optimized and secured algorithm. We are producing an open source streaming clustering algorithm with log-linear complexity growth and intrinsic data security geared to ease scalability and security of large scale data applications.

Parallel algorithms and modern methods will be used. The aforementioned includes developing the source in the Go<sup>1</sup> language, deploying on Hadoop<sup>2</sup> containers, and producing an unrivaled open source technology.

## Goal ##
The largest online storage and service companies (consider Google, Amazon, Microsoft and Facebook) are estimated to store at least 1,200 petabytes of data between them. Security and efficiency can be overlooked as growing pains in modern fast paced application architectures. Our team is focused on providing a Scalable Big Data Clustering open source project to accommodate large scale Medical applications that are HIPAA<sup>3</sup> compliant and transfer sensitive patient information to secure databases. As medical technologies grow and patient care becomes more and more dependent on secure, reliable, and fast data, this project will provide for all of these.

<sub><sup>1</sup>Go is a programming language developed by Google Inc., it aims to provide memory safety, garbage collection, and light weight threading for parallel computing.</sub>

<sub><sup>2</sup>Hadoop allows for the distributed processing of large data sets across clusters of computers.</sub>

<sub><sup>3</sup>HIPAA Health Insurance Portability and Accountability Act</sub>

## Project ##

```
rphash/
|
|___pkg/
|   |
|   [...installed go packages]
|
|___src/
|   |
|   [...source]
|
|_______tests/
|       |
|       [...source tests]
|_______utils/
|       |
|       [...utilities]

```
### RPHash ###
_Random Projection Hashing_
### Leach Array Decoder ###
### FJLT Projection ###
_Fast Johnson Lindenstrauss Transform Projection_
### LSH ###
_Locality Sensitive Hashing_
