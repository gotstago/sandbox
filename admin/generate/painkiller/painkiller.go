package painkiller

//go:generate stringer -type=Pill

type Pill int

const (
    Placebo Pill = iota
    Aspirin
    Ibuprofen
    Paracetamol
    Benadryl
    Acetaminophen = Paracetamol
)