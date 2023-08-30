package entity

import (
	"github.com/go-playground/validator/v10"
	"log"
)

type UnitType int

const (
	// todo: rename to UnitTypeNone
	UnitTypeNone UnitType = iota
	UnitTypeA
	UnitTypeB
	UnitTypeC
	UnitTypeD
	UnitTypeE
	UnitTypeF
	UnitTypeG
	UnitTypeH
	UnitTypeCustom = 99
)

const (
	ValidateUserDefinedRecordTypeName = "validate_user_defined_record_type"
)

var UnitTypeMap = map[UnitType]string{
	UnitTypeNone:   "単位なし",
	UnitTypeA:      "A",
	UnitTypeB:      "B",
	UnitTypeC:      "C",
	UnitTypeD:      "D",
	UnitTypeE:      "E",
	UnitTypeF:      "F",
	UnitTypeG:      "G",
	UnitTypeH:      "H",
	UnitTypeCustom: "カスタム",
}

func (u UnitType) String() string {
	v, ok := UnitTypeMap[u]
	if !ok {
		return "Unknown unit type"
	}
	return v
}

func ValidateUserDefinedRecordTypeFunc(field validator.FieldLevel) bool {
	unitType := field.Field().Interface().(UnitType)
	if _, ok := UnitTypeMap[unitType]; ok {
		return true
	}
	log.Printf("Error validating user defined record type: %d", unitType)
	return false
}
