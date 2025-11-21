package main

import (
	"flag"

	"github.com/NikitaAksenov/roadmap-task_tracker/internal/validator"
)

func ValidateParameterDescription(v *validator.Validator, flagSet *flag.FlagSet, parameterDescription *string) {
	category := "[-description]"

	v.Check(IsFlagPassedInSet(flagSet, "description"), category, "must be provided")
	v.Check(*parameterDescription != "", category, "must not be empty")
}

func ValidateParameterID(v *validator.Validator, flagSet *flag.FlagSet, parameterID *int) {
	category := "[-id]"

	v.Check(IsFlagPassedInSet(flagSet, "id"), category, "must be provided")
	v.Check(*parameterID > 0, category, "must be > 0")
}

func ValidateParameterStatus(v *validator.Validator, flagSet *flag.FlagSet, parameterStatus *string) {
	category := "[-status]"

	v.Check(IsFlagPassedInSet(flagSet, "status"), category, "must be provided")
}
