package tests;

import (
    "testing"
    "github.com/wenkesj/rphash/hash"
    "github.com/wenkesj/rphash/decoder"
    "github.com/wenkesj/rphash/projector"
    "github.com/wenkesj/rphash/lsh"
);

func TestLSHSimple(t *testing.T) {
  var seed int64 = 0;
  var d, k, l int = 64, 6, 4;
  data := []float64{1.0,0.0,2.0,7.0,4.0,0.0,8.0,3.0,2.0,1.0};
  var inDimensions, outDimentions int = 10, 2;
  hash := hash.NewMurmur(1 << 63 - 1);
  decoder := decoder.NewSpherical(d, k, l);
  projector := projector.NewDBFriendly(inDimensions, outDimentions, seed);
  lsh := lsh.NewLSH(hash, decoder, projector);
  lsh.LSHHashSimple(data);
  t.Log("√ LSH Simple test complete");
};

func TestLSHStream(t *testing.T) {
  var seed int64 = 0;
  var d, k, l int = 64, 6, 4;
  data := []float64{1.0,0.0,2.0,7.0,4.0,0.0,8.0,3.0,2.0,1.0};
  var inDimensions, outDimentions int = 10, 2;
  hash := hash.NewMurmur(1 << 63 - 1);
  decoder := decoder.NewSpherical(d, k, l);
  projector := projector.NewDBFriendly(inDimensions, outDimentions, seed);
  lsh := lsh.NewLSH(hash, decoder, projector);
  lsh.LSHHashStream(data, 1);
  t.Log("√ LSH Stream test complete");
};

func BenchmarkSimple(b *testing.B) {
  var seed int64 = 0;
  var d, k, l int = 64, 6, 4;
  data := []float64{1.0,0.0,2.0,7.0,4.0,0.0,8.0,3.0,2.0,1.0};
  var inDimensions, outDimentions int = 10, 2;
  hash := hash.NewMurmur(1 << 63 - 1);
  decoder := decoder.NewSpherical(d, k, l);
  projector := projector.NewDBFriendly(inDimensions, outDimentions, seed);
  for i := 0; i < b.N; i++ {
      lsh := lsh.NewLSH(hash, decoder, projector);
      b.StopTimer();
      lsh.LSHHashSimple(data);
      b.StartTimer();
  }
};

func BenchmarkStream(b *testing.B) {
  var seed int64 = 0;
  var d, k, l int = 64, 6, 4;
  data := []float64{1.0,0.0,2.0,7.0,4.0,0.0,8.0,3.0,2.0,1.0};
  var inDimensions, outDimentions int = 10, 2;
  hash := hash.NewMurmur(1 << 63 - 1);
  decoder := decoder.NewSpherical(d, k, l);
  projector := projector.NewDBFriendly(inDimensions, outDimentions, seed);
  for i := 0; i < b.N; i++ {
      lsh := lsh.NewLSH(hash, decoder, projector);
      b.StopTimer();
      lsh.LSHHashStream(data, 1);
      b.StartTimer();
  }
};
