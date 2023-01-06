package nanoflakes

const (
	TimestampBits    = 14
	GeneratorIdBits  = 10
	SequenceBits     = 12
	MaxGeneratorId   = 1023
	MaxSequence      = 4095
	GeneratorIdShift = SequenceBits
	TimestampShift   = SequenceBits + GeneratorIdShift
)

func TimeStampValue(value int64) int64 {
	return value >> TimestampBits
}

func GeneratorVaule(id int64) int64 {
	return id >> GeneratorIdShift & MaxGeneratorId
}

func SequenceValue(id int64) int64 {
	return id & MaxSequence
}
