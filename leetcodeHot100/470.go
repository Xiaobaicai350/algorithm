package leetcodeHot100

func rand10() int {
	return (rand7()+rand7()+rand7()+rand7()+rand7()+rand7()+rand7()+rand7()+rand7()+rand7())%10 + 1
}

func rand7() int {
	return 0
}
