package contract

import "github.com/grafana/grafana-plugin-sdk-go/data"

func IsBoolField(field *data.Field) bool {
	if field == nil {
		return false
	}
	return true &&
		field.Type() == data.FieldTypeBool ||
		field.Type() == data.FieldTypeNullableBool
}

func IsStringField(field *data.Field) bool {
	if field == nil {
		return false
	}
	return true &&
		field.Type() == data.FieldTypeString ||
		field.Type() == data.FieldTypeNullableString
}

func IsTimeField(field *data.Field) bool {
	if field == nil {
		return false
	}
	return true &&
		field.Type() == data.FieldTypeTime ||
		field.Type() == data.FieldTypeNullableTime
}

func IsNumericField(field *data.Field) bool {
	if field == nil {
		return false
	}
	return true &&
		// Float types
		field.Type() == data.FieldTypeFloat64 ||
		field.Type() == data.FieldTypeFloat32 ||
		field.Type() == data.FieldTypeNullableFloat64 ||
		field.Type() == data.FieldTypeNullableFloat32 ||
		// Int types
		field.Type() == data.FieldTypeInt64 ||
		field.Type() == data.FieldTypeInt32 ||
		field.Type() == data.FieldTypeInt16 ||
		field.Type() == data.FieldTypeInt8 ||
		field.Type() == data.FieldTypeNullableInt64 ||
		field.Type() == data.FieldTypeNullableInt32 ||
		field.Type() == data.FieldTypeNullableInt16 ||
		field.Type() == data.FieldTypeNullableInt8 ||
		// UInt types
		field.Type() == data.FieldTypeUint64 ||
		field.Type() == data.FieldTypeUint32 ||
		field.Type() == data.FieldTypeUint16 ||
		field.Type() == data.FieldTypeUint8 ||
		field.Type() == data.FieldTypeNullableUint64 ||
		field.Type() == data.FieldTypeNullableUint32 ||
		field.Type() == data.FieldTypeNullableUint16 ||
		field.Type() == data.FieldTypeNullableUint8
}
