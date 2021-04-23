package fibonacci

func FiboLessThanN(num uint64) uint64{
	a := uint64(0)
	b := uint64(1)
	c := uint64(0)
	for b < num {
		c = a + b
		a = b
		b = c
	}
	return a	
}