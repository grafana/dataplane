package numeric

import (
	"fmt"
	"sort"

	"github.com/grafana/dataplane/sdata"
	"github.com/grafana/grafana-plugin-sdk-go/data"
)

type LongFrame struct {
	*data.Frame
}

func (lf *LongFrame) Frames() data.Frames {
	return data.Frames{lf.Frame}
}

var LongFrameVersionLatest = LongFrameVersions()[len(LongFrameVersions())-1]

func LongFrameVersions() []data.FrameTypeVersion {
	return []data.FrameTypeVersion{{0, 1}}
}

func NewLongFrame(refID string, v data.FrameTypeVersion) (*LongFrame, error) {
	if v.Greater(LongFrameVersionLatest) {
		return nil, fmt.Errorf("can not create LongFrame of version %s because it is newer than library version %v", v, LongFrameVersionLatest)
	}
	return &LongFrame{emptyFrameWithTypeMD(refID, data.FrameTypeNumericLong, v)}, nil
}

func (lf *LongFrame) GetCollection(validateData bool) (Collection, error) {
	return validateAndGetRefsLong(lf, validateData)
}

// TODO: Update with current rules to match(ish) time series
func validateAndGetRefsLong(lf *LongFrame, validateData bool) (Collection, error) {
	var c Collection
	if validateData {
		panic("validateData option is not implemented")
	}

	if lf == nil || lf.Frame == nil {
		return c, fmt.Errorf("nil frame is invalid")
	}

	c.RefID = lf.Frame.RefID

	if !frameHasType(lf.Frame, data.FrameTypeNumericLong) {
		return c, fmt.Errorf("frame is missing %s type indicator", data.FrameTypeNumericLong)
	}

	if lf.Meta.TypeVersion != LongFrameVersionLatest {
		c.Warning = &sdata.VersionWarning{DataVersion: lf.Meta.TypeVersion, LibraryVersion: LongFrameVersionLatest, DataType: data.FrameTypeNumericLong}
	}

	if len(lf.Fields) == 0 { // TODO: Error differently if nil and not zero length?
		c.Refs = []MetricRef{}
		return c, nil // empty response
	}

	stringFieldIds, numericFieldIds := []int{}, []int{}
	stringFieldNames, numericFieldNames := []string{}, []string{}

	for i, field := range lf.Fields {
		fType := field.Type()
		switch {
		case fType.Numeric():
			numericFieldIds = append(numericFieldIds, i)
			numericFieldNames = append(numericFieldNames, field.Name)
		case fType == data.FieldTypeString || fType == data.FieldTypeNullableString:
			stringFieldIds = append(stringFieldIds, i)
			stringFieldNames = append(stringFieldNames, field.Name)
		}
	}

	for rowIdx := 0; rowIdx < lf.Rows(); rowIdx++ {
		l := data.Labels{}
		for i := range stringFieldIds {
			key := stringFieldNames[i]
			val, _ := lf.ConcreteAt(stringFieldIds[i], rowIdx)
			l[key] = val.(string)
		}

		for i, fieldIdx := range numericFieldIds {
			fType := lf.Fields[fieldIdx].Type()
			field := data.NewFieldFromFieldType(fType, 1)
			field.Name = numericFieldNames[i]
			field.Labels = l
			field.Set(0, lf.Fields[fieldIdx].At(rowIdx))
			c.Refs = append(c.Refs, MetricRef{
				ValueField: field,
			})
		}
	}
	return c, nil
}

// SortNumericMetricRef sorts refs by metric name, then by label string, with refs that have no
// labels first. It builds each label string once, because Labels.String allocates.
func SortNumericMetricRef(refs []MetricRef) {
	type keyedRef struct {
		ref       MetricRef
		name      string
		hasLabels bool
		labels    string
	}
	keyed := make([]keyedRef, len(refs))
	for i, r := range refs {
		keyed[i] = keyedRef{ref: r, name: r.GetMetricName()}
		if l := r.GetLabels(); l != nil {
			keyed[i].hasLabels = true
			keyed[i].labels = l.String()
		}
	}
	sort.SliceStable(keyed, func(i, j int) bool {
		a, b := &keyed[i], &keyed[j]
		if a.name != b.name {
			return a.name < b.name
		}
		if !a.hasLabels {
			return true // no labels first
		}
		if !b.hasLabels {
			return false
		}
		return a.labels < b.labels
	})
	for i := range keyed {
		refs[i] = keyed[i].ref
	}
}
