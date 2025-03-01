package proto

type StageDifficulty int32

const (
	StageDifficulty_None        = 0
	StageDifficulty_Normal      = 1
	StageDifficulty_Hard        = 2
	StageDifficulty_VeryHard    = 3
	StageDifficulty_VeryHard_Ex = 4
)

var (
	StageDifficulty_name = map[int32]string{
		0: "None",
		1: "Normal",
		2: "Hard",
		3: "VeryHard",
		4: "VeryHard_Ex",
	}
	StageDifficulty_value = map[string]int32{
		"None":        0,
		"Normal":      1,
		"Hard":        2,
		"VeryHard":    3,
		"VeryHard_Ex": 4,
	}
)

func (x StageDifficulty) String() string {
	return StageDifficulty_name[int32(x)]
}
