package generater

import "testing"

func TestQuestionDesc(t *testing.T) {
	type args struct {
		q    string
		path string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"q1",
			args{"perfect-squares", "./tmp.md"},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := QuestionDesc(tt.args.q, tt.args.path); (err == nil) != tt.wantErr {
				t.Errorf("QuestionDesc() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
