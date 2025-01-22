package dynamodb

import (
	"testing"
	"time"
)

func Test_buildUpdateExpression(t *testing.T) {
	type args struct {
		kvPairs map[string]interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "happy path",
			args: args{
				kvPairs: map[string]interface{}{
					TitleKey:         "new title",
					DescriptionKey:   "new description",
					StatusKey:        "COMPLETE",
					TagsKey:          []string{"tag1", "tag2"},
					ParentsKey:       []string{},
					DueDateKey:       time.Now().Unix(),
					RecurringRuleKey: nil,
				},
			},
			wantErr: false,
		},
		{
			name: "happy path - non nil recurring rule",
			args: args{
				kvPairs: map[string]interface{}{
					TitleKey:       "new title",
					DescriptionKey: "new description",
					StatusKey:      "COMPLETE",
					TagsKey:        []string{"tag1", "tag2"},
					ParentsKey:     []string{},
					DueDateKey:     time.Now().Unix(),
					RecurringRuleKey: &RecurringRule{
						CronExpression: "* * * 1 *", // whatever this means
					},
				},
			},
			wantErr: false,
		},
		{
			name: "not allowed to update user id",
			args: args{
				kvPairs: map[string]interface{}{
					UserIDKey:        "user_id",
					TitleKey:         "new title",
					DescriptionKey:   "new description",
					StatusKey:        "COMPLETE",
					TagsKey:          []string{"tag1", "tag2"},
					ParentsKey:       []string{},
					DueDateKey:       time.Now().Unix(),
					RecurringRuleKey: nil,
				},
			},
			wantErr: true,
		},
		{
			name: "not allowed to update task id",
			args: args{
				kvPairs: map[string]interface{}{
					TaskIDKey:        "task_id",
					TitleKey:         "new title",
					DescriptionKey:   "new description",
					StatusKey:        "COMPLETE",
					TagsKey:          []string{"tag1", "tag2"},
					ParentsKey:       []string{},
					DueDateKey:       time.Now().Unix(),
					RecurringRuleKey: nil,
				},
			},
			wantErr: true,
		},
		{
			name: "unknown attribute",
			args: args{
				kvPairs: map[string]interface{}{
					TitleKey:         "new title",
					DescriptionKey:   "new description",
					StatusKey:        "COMPLETE",
					TagsKey:          []string{"tag1", "tag2"},
					ParentsKey:       []string{},
					DueDateKey:       time.Now().Unix(),
					RecurringRuleKey: nil,
					"unknown_attr":   "doesn't matter the value",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := buildUpdateExpression(tt.args.kvPairs)
			if (err != nil) != tt.wantErr {
				t.Errorf("buildUpdateExpression() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
