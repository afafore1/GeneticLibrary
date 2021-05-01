package domain

// Chromosome contains genetic information and fitness of an entity
type Chromosome struct {
	Gene []interface{}
	Fitness interface{}
}

// Population represents a group of chromosomes
type Population struct {
	Chromosomes []Chromosome
	TopChromosomes []Chromosome
}

