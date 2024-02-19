package pandas

import (
	"reflect"
	"testing"
)

func TestNDArray_Copy(t *testing.T) {
	type fields struct {
		typ      Type
		rows     int
		nilCount int
		name     string
		data     any
	}
	tests := []struct {
		name   string
		fields fields
		want   Series
	}{
		{
			name: "copy",
			fields: fields{
				typ:      SERIES_TYPE_STRING,
				rows:     1,
				nilCount: 0,
				name:     "a",
				data:     Vector[string]{"1"},
			},
			want: &NDFrame{
				typ:      SERIES_TYPE_STRING,
				rows:     1,
				nilCount: 0,
				name:     "a",
				data:     ToSeries("1"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &NDFrame{
				typ:      tt.fields.typ,
				rows:     tt.fields.rows,
				nilCount: tt.fields.nilCount,
				name:     tt.fields.name,
				data:     tt.fields.data,
			}
			if got := this.Copy(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Copy() = %v(%T), want %v(%T)", got, got, tt.want, tt.want)
			}
		})
	}
}
