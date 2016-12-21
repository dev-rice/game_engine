package unit_builder

type UnitBuilder interface {
	buildUnitWithResources(u Unit, r *Resources)
}