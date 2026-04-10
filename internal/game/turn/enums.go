package turn

type Punishment int

const (
	None Punishment = iota
	FormalApology
	PublicLashing
	Expulsion
)

var PunishmentName = map[Punishment]string{
	FormalApology: "Formal Apology",
	PublicLashing: "Public Lashing",
	Expulsion:     "Expulsion",
}

type Charge int

const (
	ChargesDropped Charge = iota
	UndignifiedMischief
	RecklessUseOfSympathy
	ConductUnbecomingLashing
	ConductUnbecomingExpulsion
)

var ChargeName = map[Charge]string{
	ChargesDropped:             "Charges Dropped",
	UndignifiedMischief:        "Undignified Mischief",
	RecklessUseOfSympathy:      "Reckless Use of Sympathy",
	ConductUnbecomingLashing:   "Conduct Unbecoming a Member of the Arcanum",
	ConductUnbecomingExpulsion: "Conduct Unbecoming a Member of the Arcanum",
}

var PunishmentForCharge = map[Charge]Punishment{
	ChargesDropped:             None,
	UndignifiedMischief:        FormalApology,
	RecklessUseOfSympathy:      PublicLashing,
	ConductUnbecomingLashing:   PublicLashing,
	ConductUnbecomingExpulsion: Expulsion,
}
