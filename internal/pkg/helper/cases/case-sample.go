package casesHelper

import (
	"fmt"
	"github.com/aaronchen2k/deeptest/internal/pkg/consts"
	"github.com/getkin/kin-openapi/openapi3"
	"math"
)

func getRequiredSample() (ret string) {
	return ExampleEmpty
}

func getTypeSample(typ OasFieldType) (sample interface{}) {
	if typ == OasFieldTypeBoolean || typ == OasFieldTypeNumber || typ == OasFieldTypeArray {
		sample = RandStr()
	} else if typ == OasFieldTypeInteger {
		sample = RandFloat32()
	}

	return
}

func getEnumSample() (sample interface{}) {
	return RandStr()
}

func getFormatSample(format OasFieldFormat, typ OasFieldType) (sample interface{}, ok bool) {
	if typ == OasFieldTypeInteger {
		if format == OasFieldFormatInt32 {
			sample = RandInt64()
			ok = true
		} else if format == OasFieldFormatInt64 {
			sample = RandStr()
			ok = true
		}
	} else if typ == OasFieldTypeNumber {
		if format == OasFieldFormatFloat {
			sample = RandFloat64()
			ok = true
		} else if format == OasFieldFormatDouble {
			sample = RandStr()
			ok = true
		}
	} else if typ == OasFieldTypeString { // others
		sample = RandStr()
		ok = true
	}

	return
}

func getRuleSamples(schema *openapi3.Schema, name string) (ret [][]interface{}) {
	typ := OasFieldType(schema.Type)

	if typ == OasFieldTypeInteger || typ == OasFieldTypeNumber {
		if schema.Min != nil && *schema.Min != 0 {
			sample := *schema.Min - 1
			tag := fmt.Sprintf("%v", *schema.Min)

			if schema.ExclusiveMin {
				sample = *schema.Min
				tag = tag + " exclusive"
			}

			item := []interface{}{name, sample, typ, tag, consts.AlternativeCaseRulesMin}
			ret = append(ret, item)
		}

		if schema.Max != nil && *schema.Max != 0 {
			sample := *schema.Max + 1
			tag := fmt.Sprintf("%v", *schema.Max)

			if schema.ExclusiveMax {
				sample = *schema.Max
				tag = tag + " exclusive"
			}

			item := []interface{}{name, sample, typ, tag, consts.AlternativeCaseRulesMax}
			ret = append(ret, item)
		}

		if schema.MaxLength != nil && *schema.MaxLength > 0 {
			var sample interface{}

			if typ == OasFieldTypeInteger {
				sample = 1 * math.Pow(10, float64(*schema.MaxLength))
			} else {
				if *schema.MaxLength <= 3 {
					sample = 1
				} else {
					maxLen := *schema.MaxLength
					if maxLen < 1 {
						maxLen = 1
					}

					sample = 1/math.Pow(10, float64(maxLen-1)) + 1
				}
			}

			tag := *schema.MaxLength
			item := []interface{}{name, sample, typ, tag, consts.AlternativeCaseRulesMaxLength}
			ret = append(ret, item)
		}

		if schema.MinLength > 0 {
			var sample interface{}

			if typ == OasFieldTypeInteger {
				minLen := schema.MinLength
				if minLen < 2 {
					minLen = 2
				}
				sample = 1 * math.Pow(10, float64(minLen-2))
			} else {
				minLen := schema.MinLength
				if minLen < 3 {
					minLen = 3
				}

				sample = 1/math.Pow(10, float64(minLen-3)) + 1
			}

			tag := schema.MinLength
			item := []interface{}{name, sample, typ, tag, consts.AlternativeCaseRulesMinLength}
			ret = append(ret, item)
		}

		if schema.MultipleOf != nil && *schema.MultipleOf != 0 {
			var sample interface{}

			if typ == OasFieldTypeInteger {
				sample = *schema.MultipleOf + 1
			} else {
				sample = *schema.MultipleOf + *schema.MultipleOf*0.1
			}

			tag := *schema.MultipleOf
			item := []interface{}{name, sample, typ, tag, consts.AlternativeCaseRulesMultipleOf}
			ret = append(ret, item)
		}

	} else {
		if schema.Pattern != "" {
			sample := RandStrSpecial()

			tag := schema.Pattern
			item := []interface{}{name, sample, typ, tag, consts.AlternativeCaseRulesPattern}
			ret = append(ret, item)
		}

		if schema.MaxLength != nil && *(schema.MaxLength) > 0 {
			sample := RandStrWithLen(int(*(schema.MaxLength) + 1))

			tag := *schema.MaxLength
			item := []interface{}{name, sample, typ, tag, consts.AlternativeCaseRulesMaxLength}
			ret = append(ret, item)
		}

		if schema.MinLength > 0 {
			minLen := schema.MinLength
			if minLen < 1 {
				minLen = 1
			}

			sample := RandStrWithLen(int(minLen - 1))

			tag := schema.MinLength
			item := []interface{}{name, sample, typ, tag, consts.AlternativeCaseRulesMinLength}
			ret = append(ret, item)
		}
	}

	return
}
