package Labyrinth

func AiMove() {
	AIData[Ipos][Jpos] = 1
	Run("")

	// Tip: 3为Ai标记
	for Ipos != 9 || Jpos != 9 {
		if (Ipos-1 >= 0 && AIData[Ipos-1][Jpos] == 3) {
			AIData[Ipos-1][Jpos] = 0
			Run("w")
		}
		if (Ipos+1 <= 9 && AIData[Ipos+1][Jpos] == 3) {
			AIData[Ipos+1][Jpos] = 0
			Run("s")
		}
		if (Jpos+1 <= 9 && AIData[Ipos][Jpos+1] == 3) {
			AIData[Ipos][Jpos+1] = 0
			Run("d")
		}
		if (Jpos-1 >= 0 && AIData[Ipos][Jpos-1] == 3) {
			AIData[Ipos][Jpos-1] = 0
			Run("a")
		}
	}
}
