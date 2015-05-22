package ProjectEular

func Questions1() int{
	sum := 0
	for i := 0; i < 1000; i++ {
		if(i%3 == 0 || i%5 == 0){
			sum += i
		}
	}
	return sum
}

func Questions2() int{
	sum := 2
	first := 1
	last := 2
	for last<4000000 {
		temp := first
		first = last
		last = temp + last
		if(last%2 == 0){
			sum += last
		}		
	}
	return sum
}

func Questions3() uint64{
	var number uint64 = 600851475143
	largestFactor := uint64(2)
	for i := uint64(2); i<=number; i++ {
		for number % i == 0{
			largestFactor = i
			number /= largestFactor 
		}
	}
	return largestFactor
}

func Questions4() int{
	bigest := 0
	for i := 100; i <= 999; i++ {
		for k := 100; k <= 999; k++ {
			if IsItPalindrome(i * k) {
				number := i * k
				if number > bigest {
					bigest = number
				}
			}
		}
	}
	return bigest
}

