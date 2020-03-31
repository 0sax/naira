package naira

import "testing"

func TestKobo_KobotoNGNString(t *testing.T) {
	tests := []struct {
		name string
		m    Kobo
		want string
	}{
		// Test Cases
		{
			name: "Positive Ten K",
			m:    9700067,
			want: "97,000.67",
		}, {
			name: "Positive Hundred K",
			m:    50200067,
			want: "502,000.67",
		},
		{
			name: "Positive Million",
			m:    500000067,
			want: "5,000,000.67",
		},
		{
			name: "Negative Million",
			m:    -500000067,
			want: "-5,000,000.67",
		},
		{
			name: "Negative Ten K",
			m:    -9700067,
			want: "-97,000.67",
		}, {
			name: "Negative Hundred K",
			m:    -50200067,
			want: "-502,000.67",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.KobotoPrettyNGNString(); got != tt.want {
				t.Errorf("KobotoPrettyNGNString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToKobo(t *testing.T) {
	type args struct {
		f         float64
		direction string
	}
	tests := []struct {
		name string
		args args
		want Kobo
	}{
		//test cases.
		{
			name: "RoundUp",
			args: args{
				f:         2045485.21547861,
				direction: "u",
			},
			want: 204548522,
		},
		{
			name: "ROundDown",
			args: args{
				f:         2045485.21547861,
				direction: "d",
			},
			want: 204548521,
		}, {
			name: "RoundUpNEgative",
			args: args{
				f:         -2045485.21547861,
				direction: "u",
			},
			want: -204548521,
		},
		{
			name: "ROundDownNegative",
			args: args{
				f:         -2045485.21547861,
				direction: "d",
			},
			want: -204548522,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToKobo(tt.args.f, tt.args.direction); got != tt.want {
				t.Errorf("ToKobo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKobo_KoboToFloat(t *testing.T) {
	tests := []struct {
		name string
		m    Kobo
		want float64
	}{
		// Add test cases.
		{
			name: "BasicTest",
			m:    63435342323,
			want: 634353423.23,
		},
		{
			name: "Double00",
			m:    63435342300,
			want: 634353423.00,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.KoboToFloat(); got != tt.want {
				t.Errorf("KoboToFloat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKobo_Multiply(t *testing.T) {
	type args struct {
		f float64
	}
	tests := []struct {
		name string
		m    Kobo
		args args
		want Kobo
	}{
		// Add test cases.
		{
			name: "Multi",
			m:    5000025,
			args: args{
				f: 0.05,
			},
			want: 250001,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.Multiply(tt.args.f); got != tt.want {
				t.Errorf("Multiply() = %v, want %v", got, tt.want)
			}
		})
	}
}
