package main

import "testing"

func Test_human_getAge(t *testing.T) {
	type fields struct {
		FullName string
		Sex      string
		Age      uint8
		Height   uint16
		Weight   uint16
	}
	tests := []struct {
		name   string
		fields fields
		want   uint8
	}{
		{
			"ok",
			fields{
				Age: 33,
			},
			33,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := human{
				FullName: tt.fields.FullName,
				Sex:      tt.fields.Sex,
				Age:      tt.fields.Age,
				Height:   tt.fields.Height,
				Weight:   tt.fields.Weight,
			}
			if got := h.getAge(); got != tt.want {
				t.Errorf("getAge() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_human_getFullName(t *testing.T) {
	type fields struct {
		FullName string
		Sex      string
		Age      uint8
		Height   uint16
		Weight   uint16
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			"ok",
			fields{
				FullName: "Иван",
			},
			"Иван",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := human{
				FullName: tt.fields.FullName,
				Sex:      tt.fields.Sex,
				Age:      tt.fields.Age,
				Height:   tt.fields.Height,
				Weight:   tt.fields.Weight,
			}
			if got := h.getFullName(); got != tt.want {
				t.Errorf("getFullName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_human_getHeight(t *testing.T) {
	type fields struct {
		FullName string
		Sex      string
		Age      uint8
		Height   uint16
		Weight   uint16
	}
	tests := []struct {
		name   string
		fields fields
		want   uint16
	}{
		{
			"ok",
			fields{
				Height: 176,
			},
			176,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := human{
				FullName: tt.fields.FullName,
				Sex:      tt.fields.Sex,
				Age:      tt.fields.Age,
				Height:   tt.fields.Height,
				Weight:   tt.fields.Weight,
			}
			if got := h.getHeight(); got != tt.want {
				t.Errorf("getHeight() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_human_getSex(t *testing.T) {
	type fields struct {
		FullName string
		Sex      string
		Age      uint8
		Height   uint16
		Weight   uint16
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			"ok",
			fields{
				Sex: "Мужчина",
			},
			"Мужчина",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := human{
				FullName: tt.fields.FullName,
				Sex:      tt.fields.Sex,
				Age:      tt.fields.Age,
				Height:   tt.fields.Height,
				Weight:   tt.fields.Weight,
			}
			if got := h.getSex(); got != tt.want {
				t.Errorf("getSex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_human_getWeight(t *testing.T) {
	type fields struct {
		FullName string
		Sex      string
		Age      uint8
		Height   uint16
		Weight   uint16
	}
	tests := []struct {
		name   string
		fields fields
		want   uint16
	}{
		{
			"ok",
			fields{
				Weight: 88,
			},
			88,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := human{
				FullName: tt.fields.FullName,
				Sex:      tt.fields.Sex,
				Age:      tt.fields.Age,
				Height:   tt.fields.Height,
				Weight:   tt.fields.Weight,
			}
			if got := h.getWeight(); got != tt.want {
				t.Errorf("getWeight() = %v, want %v", got, tt.want)
			}
		})
	}
}
