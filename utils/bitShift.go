package utils;
// Java like zero extension shift see http://docs.oracle.com/javase/specs/jls/se7/html/jls-15.html#jls-15.19 >>> operator
func RightShiftZeroExtension(leftOperand int64, rightOperand int64) int64{
  if leftOperand > -1 {
    return leftOperand >> uint64(rightOperand);
  } else {
    return (leftOperand >> uint64(rightOperand)) + (2 << uint64(^rightOperand));
  }
};
