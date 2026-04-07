package player

type Class int

const (
	EdemaRuh Class = iota
	CealdishCommoner
	YllishCommoner
	AturanNoble
	VintishNoble
)

var ClassName = map[Class]string{
	EdemaRuh:         "Edema Ruh",
	CealdishCommoner: "Cealdish Commoner",
	YllishCommoner:   "Yllish Commoner",
	AturanNoble:      "Aturan Noble",
	VintishNoble:     "Vintish Noble",
}

type Rank int

const (
	Student Rank = iota
	Elir
	Relar
	Elthe
	Master
)

var RankName = map[Rank]string{
	Student: "Student",
	Elir:    "E'lir",
	Relar:   "Re'lar",
	Elthe:   "El'the",
	Master:  "Master",
}

type ActionType int

const (
	None ActionType = iota
	Sabotage
	LawOfContraposition
)

var ActionTypeName = map[ActionType]string{
	None:                "None",
	Sabotage:            "Sabotage",
	LawOfContraposition: "Law of Contraposition",
}

type ActionSource int

const (
	Alchemy ActionSource = iota
	Archives
	Arithmetics
	Artificery
	Linguistics
	Naming
	Physicking
	RhetoricAndLogic
	Sympathy
)

type ActionCategory int

const (
	General ActionCategory = iota
	Roleblock
)

var ActionCategoryName = map[ActionCategory]string{
	General:   "General",
	Roleblock: "Roleblock",
}
