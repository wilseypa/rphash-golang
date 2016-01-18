package utils;

func RightShiftZeroExtension(leftOperand int64, rightOperand int64) int64{
  if leftOperand > -1 {
    return leftOperand >> uint64(rightOperand);
  } else {
    return (leftOperand >> uint64(rightOperand)) + (2 << uint64(^rightOperand));
  }
};
