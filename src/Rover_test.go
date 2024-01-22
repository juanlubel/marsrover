package src

import (
	"reflect"
	"testing"
)

func TestNewRover(t *testing.T) {
	tests := []struct {
		name string
		want *Rover
	}{
		// TODO: Add test cases.
		{"Testing default values of a new Rover instance",
			&Rover{
				Initial:  Coordinate{X: 0, Y: 0},
				Moves:    nil,
				Position: Coordinate{X: 0, Y: 0},
				Faced:    North,
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRover(InitialPosition); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRover() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRover_GetPosition(t *testing.T) {
	type fields struct {
		Initial  Coordinate
		Moves    []string
		Position Coordinate
		Faced    CardinalDirection
	}
	initialField := fields{
		Coordinate{X: 0, Y: 0},
		nil,
		Coordinate{X: 0, Y: 0},
		North,
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
		{name: "Testing the finale position of a new Rover instance",
			fields: initialField,
			want:   "0:0:N",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Rover{
				Initial:  tt.fields.Initial,
				Moves:    tt.fields.Moves,
				Position: tt.fields.Position,
				Faced:    tt.fields.Faced,
			}
			if got := r.GetPosition(); got != tt.want {
				t.Errorf("GetPosition() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStartsRover(t *testing.T) {
	tests := []struct {
		name     string
		position Coordinate
		want     *Rover
	}{
		// TODO: Add test cases.
		{
			name:     "Testing expected rover when start one",
			position: Coordinate{X: 0, Y: 0},
			want: &Rover{
				Initial:  Coordinate{X: 0, Y: 0},
				Moves:    nil,
				Position: Coordinate{X: 0, Y: 0},
				Faced:    North,
			}},
		{
			name:     "Testing expected rover when start one in a different position",
			position: Coordinate{X: 1, Y: 1},
			want: &Rover{
				Initial:  Coordinate{X: 1, Y: 1},
				Moves:    nil,
				Position: Coordinate{X: 1, Y: 1},
				Faced:    North,
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rover := NewRover(tt.position)
			if got := rover.StartsRover(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StartsRover() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRover_moveFacing(t *testing.T) {
	type fields struct {
		Initial  Coordinate
		Moves    []string
		Position Coordinate
		Faced    CardinalDirection
	}
	type args struct {
		move string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   CardinalDirection
	}{
		{
			name: "Moving left from N end at E",
			fields: fields{
				Initial:  Coordinate{X: 1, Y: 1},
				Moves:    nil,
				Position: Coordinate{X: 1, Y: 1},
				Faced:    North,
			},
			args: args{move: "L"},
			want: East,
		},
		{
			name: "Moving left from E end at S",
			fields: fields{
				Initial:  Coordinate{X: 1, Y: 1},
				Moves:    nil,
				Position: Coordinate{X: 1, Y: 1},
				Faced:    East,
			},
			args: args{move: "L"},
			want: South,
		},
		{
			name: "Moving left from S end at W",
			fields: fields{
				Initial:  Coordinate{X: 1, Y: 1},
				Moves:    nil,
				Position: Coordinate{X: 1, Y: 1},
				Faced:    South,
			},
			args: args{move: "L"},
			want: West,
		},
		{
			name: "Moving left from W end at N",
			fields: fields{
				Initial:  Coordinate{X: 1, Y: 1},
				Moves:    nil,
				Position: Coordinate{X: 1, Y: 1},
				Faced:    West,
			},
			args: args{move: "L"},
			want: North,
		},
		{
			name: "Moving right from N end at W",
			fields: fields{
				Initial:  Coordinate{X: 1, Y: 1},
				Moves:    nil,
				Position: Coordinate{X: 1, Y: 1},
				Faced:    North,
			},
			args: args{move: "R"},
			want: West,
		},
		{
			name: "Moving right from W end at S",
			fields: fields{
				Initial:  Coordinate{X: 1, Y: 1},
				Moves:    nil,
				Position: Coordinate{X: 1, Y: 1},
				Faced:    West,
			},
			args: args{move: "R"},
			want: South,
		},
		{
			name: "Moving right from S end at E",
			fields: fields{
				Initial:  Coordinate{X: 1, Y: 1},
				Moves:    nil,
				Position: Coordinate{X: 1, Y: 1},
				Faced:    South,
			},
			args: args{move: "R"},
			want: East,
		},
		{
			name: "Moving right from E end at N",
			fields: fields{
				Initial:  Coordinate{X: 1, Y: 1},
				Moves:    nil,
				Position: Coordinate{X: 1, Y: 1},
				Faced:    East,
			},
			args: args{move: "R"},
			want: North,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Rover{
				Initial:  tt.fields.Initial,
				Moves:    tt.fields.Moves,
				Position: tt.fields.Position,
				Faced:    tt.fields.Faced,
			}
			r.moveFacing(tt.args.move)
			if got := r.Faced; got != tt.want {
				t.Errorf("The rover didn't faced correctly")
			}
		})
	}
}

func TestRover_Move(t *testing.T) {
	type fields struct {
		Initial  Coordinate
		Moves    []string
		Position Coordinate
		Faced    CardinalDirection
	}
	type args struct {
		moves string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   CardinalDirection
	}{

		{
			name: "Get moves only faced, expects North",
			fields: fields{
				Initial:  Coordinate{X: 1, Y: 1},
				Moves:    nil,
				Position: Coordinate{X: 1, Y: 1},
				Faced:    North,
			},
			args: args{moves: "RLRL"},
			want: North,
		},
		{
			name: "Get moves only faced full left, expects West",
			fields: fields{
				Initial:  Coordinate{X: 1, Y: 1},
				Moves:    nil,
				Position: Coordinate{X: 1, Y: 1},
				Faced:    North,
			},
			args: args{moves: "LLL"},
			want: West,
		},
		{
			name: "Get moves only faced full right, expects South",
			fields: fields{
				Initial:  Coordinate{X: 1, Y: 1},
				Moves:    nil,
				Position: Coordinate{X: 1, Y: 1},
				Faced:    North,
			},
			args: args{moves: "RRR"},
			want: East,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Rover{
				Initial:  tt.fields.Initial,
				Moves:    tt.fields.Moves,
				Position: tt.fields.Position,
				Faced:    tt.fields.Faced,
			}
			r.Move(tt.args.moves)
			if got := r.Faced; got != tt.want {
				t.Errorf("The rover didn't faced correctly")
			}
		})
	}
}

func TestRover_moveBackwardOrForward(t *testing.T) {
	type fields struct {
		Initial  Coordinate
		Moves    []string
		Position Coordinate
		Faced    CardinalDirection
	}
	type args struct {
		move string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Coordinate
	}{
		{
			name: "Moves forward on North faced",
			fields: fields{
				Initial:  Coordinate{X: 1, Y: 1},
				Moves:    nil,
				Position: Coordinate{X: 1, Y: 1},
				Faced:    North,
			},
			args: args{move: "F"},
			want: Coordinate{X: 1, Y: 2},
		},
		{
			name: "It moves backwards on the north face, collides at the down Y borderline.",
			fields: fields{
				Initial:  Coordinate{X: 1, Y: 1},
				Moves:    nil,
				Position: Coordinate{X: 1, Y: 1},
				Faced:    North,
			},
			args: args{move: "B"},
			want: Coordinate{X: 1, Y: 1},
		},
		{
			name: "It moves backwards on the north face, collides at the left X borderline.",
			fields: fields{
				Initial:  Coordinate{X: 1, Y: 1},
				Moves:    nil,
				Position: Coordinate{X: 1, Y: 1},
				Faced:    East,
			},
			args: args{move: "B"},
			want: Coordinate{X: 1, Y: 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Rover{
				Initial:  tt.fields.Initial,
				Moves:    tt.fields.Moves,
				Position: tt.fields.Position,
				Faced:    tt.fields.Faced,
			}
			r.moveBackwardOrForward(tt.args.move)
			if got := r.Position; got != tt.want {
				t.Errorf("Position is not correct")
			}
		})
	}
}
