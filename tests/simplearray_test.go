package tests;

import (
  "math"
  "math/rand"
  "reflect"
  "testing"
  "github.com/stretchr/testify/assert"
  "github.com/wenkesj/rphash/reader"
  "github.com/wenkesj/rphash/types"
  "github.com/wenkesj/rphash/utils"
);

func TestSimpleArray(t *testing.T) {
  var k = 4;
  var dimensionality = 100;
  var numBlurs = 2;
  var numProjections = 2;
  var numDataPoints = 8;
  var origVariance float64 = 1;
  var testDecoderType types.Decoder;
  var newNumProjections = 4;
  var newHashModulus int64 = rand.Int63();
  var newRandomSeed int64 = rand.Int63();

  newVarianceSample, newCentroidList := make([][]float64, numDataPoints), make([][]float64, numDataPoints);
  for i := 0; i < numDataPoints; i++ {
    newVarianceSample[i], newCentroidList[i] = make([]float64, dimensionality), make([]float64, dimensionality);
    for j := 0; j < dimensionality; j++ {
      newVarianceSample[i][j], newCentroidList[i][j] = float64(i), float64(i);
    }
  }

  newCentroid := make([]float64, dimensionality);
  for i := 0; i < dimensionality; i++ {
    newCentroid[i] = float64(i);
  }

  newTopId := make([]int64, dimensionality);
  for i := 0; i < dimensionality; i++ {
    newTopId[i] = int64(i);
  }

  RPHashObject := reader.NewSimpleArray(newCentroidList, k);

  // K.
  assert.Equal(t, k, RPHashObject.GetK(), "Expected K equal to Stream K.");

  // Dimensionality.
  assert.Equal(t, dimensionality, RPHashObject.GetDimensions(), "Expected dimensionality equal to Stream dimensionality.");

  // Iterator.
  assert.NotNil(t, RPHashObject.GetVectorIterator(), "Vector iterator should be initially not be nil.");

  // Blurs.
  assert.Equal(t, numBlurs, RPHashObject.GetNumberOfBlurs(), "Number of blurs should be initially 2.");

  // Variance.
  assert.Equal(t, origVariance, RPHashObject.GetVariance(), "Variance should be equal to the new variance value.");
  RPHashObject.SetVariance(newVarianceSample);
  newVariance := utils.VarianceSample(newVarianceSample, 0.01);
  assert.Equal(t, newVariance, RPHashObject.GetVariance(), "Variance should be equal to the new variance value.");

  // Decoders.
  origDecoderType := RPHashObject.GetDecoderType();
  assert.NotNil(t, origDecoderType);
  assert.Equal(t, reflect.ValueOf(&testDecoderType).Elem().Type(), reflect.ValueOf(&origDecoderType).Elem().Type(), "Decoder should implement the Decoder interface.");
  RPHashObject.SetDecoderType(testDecoderType);
  assert.Equal(t, testDecoderType, RPHashObject.GetDecoderType(), "Decoder should be set to a new Decoder.");

  // Projections.
  assert.Equal(t, numProjections, RPHashObject.GetNumberOfProjections(), "Number of projections should be initially 2.");
  RPHashObject.SetNumberOfProjections(newNumProjections);
  assert.Equal(t, newNumProjections, RPHashObject.GetNumberOfProjections(), "Number of projections should be equal to the new number of projections.");

  // Hash modulus.
  assert.Equal(t, int64(math.MaxInt32), RPHashObject.GetHashModulus(), "Hash modulus should be equal to the maximum 32 bit integer value.");
  RPHashObject.SetHashModulus(newHashModulus);
  assert.Equal(t, newHashModulus, RPHashObject.GetHashModulus(), "Hash modulus should be equal to the new hash modulus.");

  // Random seed.
  assert.Equal(t, int64(0), RPHashObject.GetRandomSeed(), "Random seed expected is 0.");
  RPHashObject.SetRandomSeed(newRandomSeed);
  assert.Equal(t, newRandomSeed, RPHashObject.GetRandomSeed(), "Random seed expected to be equal to new random seed.");

  // Centroids.
  assert.Empty(t, RPHashObject.GetCentroids(), "Centroids should initially be empty.");
  RPHashObject.AddCentroid(newCentroid);
  assert.Equal(t, newCentroid, RPHashObject.GetCentroids()[0], "First centroid should be the new centroid.");
  RPHashObject.SetCentroids(newCentroidList);
  assert.Equal(t, newCentroidList, RPHashObject.GetCentroids(), "Centroids should be equal to the new centroid list.");

  // Top IDs
  assert.Empty(t, RPHashObject.GetPreviousTopID(), "Previous top ID should initially be empty.");
  RPHashObject.SetPreviousTopID(newTopId);
  assert.Equal(t, newTopId, RPHashObject.GetPreviousTopID(), "Previous top ID should be equal to the new top centroid.");
}
