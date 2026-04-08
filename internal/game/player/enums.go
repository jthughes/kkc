package player

import "github.com/jthughes/kkc/internal/game/items"

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

	// General

	Sabotage

	// Item

	UseNahlrout
	UseMommet
	UseTenaculum
	UsePlumBob
	UseBonetar
	UseWard
	UseThievesLamp

	// Field
	// - Linguistics

	HandDelivery
	MysteriousBulletins
	BribeTheMessenger
	LinguisticAnalysis

	// - Arithmetics

	ReducedInterest
	Pickpocket
	GreatDeals
	DecreasedTuition

	// - Rhetoric & Logic

	ArgumentumAdNauseam
	ProficientInHyperbole
	PersuasiveArguments
	LawOfContraposition

	// - Archives

	OmenRecognition
	SchoolRecords
	BannedBooks
	FaeLore

	// - Sympathy

	MommetMaking
	MalfeasanceProtection

	// - Physicking

	MedicaEmergency
	MedicaDetainment
	PsychologicalCounselling
	CheatingDeath

	// - Alchemy

	CreateTenaculum
	CreateFirestop
	CreatePlumBob
	CreateBoneTar

	// - Artificery

	CreateWard
	CreateBloodless
	CreateThievesLamp
	CreateGram

	// - Naming

)

var ActionTypeName = map[ActionType]string{
	// General
	None:     "None",
	Sabotage: "Sabotage",
	// Item
	UseNahlrout:    "Use Nahlrout",
	UseMommet:      "Use Mommet",
	UseTenaculum:   "User Tenaculum",
	UsePlumBob:     "Use Plum bob",
	UseBonetar:     "Use Bone-tar",
	UseWard:        "Use Ward",
	UseThievesLamp: "Use Thieves Lamp",
	// Field
	// - Linguistics
	HandDelivery:        "",
	MysteriousBulletins: "",
	BribeTheMessenger:   "",
	LinguisticAnalysis:  "",
	// - Arithmetics
	ReducedInterest:  "Reduced Interest",
	Pickpocket:       "Pickpocket",
	GreatDeals:       "Great Deals",
	DecreasedTuition: "Decreased Tuition",
	// - Rhetoric & Logic
	ArgumentumAdNauseam:   "Argumentum Ad Nauseum",
	ProficientInHyperbole: "Proficient in Hyperbole",
	PersuasiveArguments:   "Persuasive Arguments",
	LawOfContraposition:   "Law of Contraposition",
	// - Archives
	OmenRecognition: "",
	SchoolRecords:   "",
	BannedBooks:     "",
	FaeLore:         "",
	// - Sympathy
	MommetMaking:          "",
	MalfeasanceProtection: "",
	// - Physicking
	MedicaEmergency:          "Medica Emergency",
	MedicaDetainment:         "Medica Detainment",
	PsychologicalCounselling: "Psychological Counselling",
	CheatingDeath:            "Cheating Death",
	// - Alchemy
	CreateTenaculum: "",
	CreateFirestop:  "",
	CreatePlumBob:   "",
	CreateBoneTar:   "",
	// - Artificery
	CreateWard:        "Create Ward",
	CreateBloodless:   "Create Bloodless",
	CreateThievesLamp: "Create Thieves Lamp",
	CreateGram:        "Create Gram",
}

var ItemActions = map[items.ItemType]ActionType{
	items.Nahlrout:    UseNahlrout,
	items.Mommet:      UseMommet,
	items.Tenaculum:   UseTenaculum,
	items.PlumBob:     UsePlumBob,
	items.BoneTar:     UseBonetar,
	items.Ward:        UseWard,
	items.ThievesLamp: UseThievesLamp,
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
	Item
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
