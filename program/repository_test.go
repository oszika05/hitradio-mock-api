package program

import (
	"gopkg.in/guregu/null.v4"
	"reflect"
	"testing"
)

func TestRepository_GetPeople(t *testing.T) {
	type fields struct {
		programs []Program
		episodes []Episode
		people   []Person
	}
	type args struct {
		from       int
		pageSize   int
		search     null.String
		personType PersonType
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []Person
		wantErr bool
	}{
		{
			name: "returns empty array, when pageSize is 0",
			fields: fields{
				programs: mockPrograms(),
				episodes: mockEpisodes(),
				people:   mockPerson(),
			},
			args: args{
				from:       0,
				pageSize:   0,
				search:     null.String{},
				personType: "",
			},
			want:    []Person{},
			wantErr: false,
		},
		{
			name: "returns empty array, when from is out of bounds",
			fields: fields{
				programs: mockPrograms(),
				episodes: mockEpisodes(),
				people:   mockPerson(),
			},
			args: args{
				from:       99,
				pageSize:   1,
				search:     null.String{},
				personType: "",
			},
			want:    []Person{},
			wantErr: false,
		},
		{
			name: "first page works",
			fields: fields{
				programs: mockPrograms(),
				episodes: mockEpisodes(),
				people:   mockPerson(),
			},
			args: args{
				from:       0,
				pageSize:   5,
				search:     null.String{},
				personType: "",
			},
			want:    mockPerson()[0:5],
			wantErr: false,
		},
		{
			name: "seconds page works",
			fields: fields{
				programs: mockPrograms(),
				episodes: mockEpisodes(),
				people:   mockPerson(),
			},
			args: args{
				from:       5,
				pageSize:   2,
				search:     null.String{},
				personType: "",
			},
			want:    mockPerson()[5:7],
			wantErr: false,
		},
		{
			name: "last page is shorter if needed",
			fields: fields{
				programs: mockPrograms(),
				episodes: mockEpisodes(),
				people:   mockPerson(),
			},
			args: args{
				from:       5,
				pageSize:   10,
				search:     null.String{},
				personType: "",
			},
			want:    mockPerson()[5:],
			wantErr: false,
		},
		{
			name: "search works",
			fields: fields{
				programs: mockPrograms(),
				episodes: mockEpisodes(),
				people:   mockPerson(),
			},
			args: args{
				from:       0,
				pageSize:   2,
				search:     null.StringFrom("Németh"),
				personType: "",
			},
			want: []Person{
				{
					Id:          "4",
					Name:        "Németh Sándor",
					Type:        PersonTypeGuest,
					Picture:     null.StringFrom("https://mandiner.hu/attachment/0346/345825_nemeth_sandor1.jpg"),
					Description: null.StringFrom("A hit gyülekezete vezető lelkésze"),
				},
			},
			wantErr: false,
		},
		{
			name: "search works with no results",
			fields: fields{
				programs: mockPrograms(),
				episodes: mockEpisodes(),
				people:   mockPerson(),
			},
			args: args{
				from:       0,
				pageSize:   2,
				search:     null.StringFrom("jhgfksjsfhjadlkfjdsalk"),
				personType: "",
			},
			want:    []Person{},
			wantErr: false,
		},
		{
			name: "personType filter works",
			fields: fields{
				programs: mockPrograms(),
				episodes: mockEpisodes(),
				people:   mockPerson(),
			},
			args: args{
				from:       0,
				pageSize:   1,
				search:     null.String{},
				personType: PersonTypeHost,
			},
			want: []Person{
				{
					Id:          "1",
					Name:        "Király Tamás",
					Type:        PersonTypeHost,
					Picture:     null.StringFrom("https://scontent-vie1-1.xx.fbcdn.net/v/t1.6435-9/67699021_112428760098874_5586161939806945280_n.jpg?_nc_cat=107&ccb=1-5&_nc_sid=8bfeb9&_nc_ohc=jLjDglC87FcAX9m2ubg&_nc_ht=scontent-vie1-1.xx&oh=21c960625e3f537eb29cf2c19a2ce3c2&oe=6183179C"),
					Description: null.String{},
				},
			},
			wantErr: false,
		},
		{
			name: "personType and search filter works",
			fields: fields{
				programs: mockPrograms(),
				episodes: mockEpisodes(),
				people:   mockPerson(),
			},
			args: args{
				from:       0,
				pageSize:   1,
				search:     null.StringFrom("Alexandra"),
				personType: PersonTypeHost,
			},
			want: []Person{
				{
					Id:          "2",
					Name:        "Kauzál Alexandra",
					Type:        PersonTypeHost,
					Picture:     null.StringFrom("https://archivum.hitradio.hu/data/img//2018/01/Családi%20magazin%20blank-KYI-2hUoi0-1516026696.jpg"),
					Description: null.String{},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := Repository{
				programs: tt.fields.programs,
				episodes: tt.fields.episodes,
				people:   tt.fields.people,
			}
			got, err := repo.GetPeople(tt.args.from, tt.args.pageSize, tt.args.search, tt.args.personType)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPeople() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPeople() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepository_GetPrograms(t *testing.T) {
	type fields struct {
		programs []Program
		episodes []Episode
		people   []Person
	}
	type args struct {
		from     int
		pageSize int
		search   null.String
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []Program
		wantErr bool
	}{
		{
			name: "returns empty array when pageSize is 0",
			fields: fields{
				programs: mockPrograms(),
				episodes: mockEpisodes(),
				people:   mockPerson(),
			},
			args: args{
				from:     0,
				pageSize: 0,
				search:   null.String{},
			},
			want:    []Program{},
			wantErr: false,
		},
		{
			name: "returns empty array, when from is out of bounds",
			fields: fields{
				programs: mockPrograms(),
				episodes: mockEpisodes(),
				people:   mockPerson(),
			},
			args: args{
				from:     9999,
				pageSize: 10,
				search:   null.String{},
			},
			want:    []Program{},
			wantErr: false,
		},
		{
			name: "first page works",
			fields: fields{
				programs: mockPrograms(),
				episodes: mockEpisodes(),
				people:   mockPerson(),
			},
			args: args{
				from:     0,
				pageSize: 2,
				search:   null.String{},
			},
			want:    mockPrograms()[0:2],
			wantErr: false,
		},
		{
			name: "seconds page works",
			fields: fields{
				programs: mockPrograms(),
				episodes: mockEpisodes(),
				people:   mockPerson(),
			},
			args: args{
				from:     2,
				pageSize: 2,
				search:   null.String{},
			},
			want:    mockPrograms()[2:4],
			wantErr: false,
		},
		{
			name: "last page returns less then pageSize if required",
			fields: fields{
				programs: mockPrograms(),
				episodes: mockEpisodes(),
				people:   mockPerson(),
			},
			args: args{
				from:     2,
				pageSize: 3,
				search:   null.String{},
			},
			want:    mockPrograms()[2:4],
			wantErr: false,
		},
		{
			name: "search filter works",
			fields: fields{
				programs: mockPrograms(),
				episodes: mockEpisodes(),
				people:   mockPerson(),
			},
			args: args{
				from:     0,
				pageSize: 2,
				search:   null.StringFrom("nyugov"),
			},
			want:    mockPrograms()[3:4],
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := Repository{
				programs: tt.fields.programs,
				episodes: tt.fields.episodes,
				people:   tt.fields.people,
			}
			got, err := repo.GetPrograms(tt.args.from, tt.args.pageSize, tt.args.search)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPrograms() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPrograms() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepository_GetProgram(t *testing.T) {
	type fields struct {
		programs []Program
		episodes []Episode
		people   []Person
	}
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Program
		want1   bool
		wantErr bool
	}{
		{
			name: "programId filter works",
			fields: fields{
				programs: mockPrograms(),
				episodes: mockEpisodes(),
				people:   mockPerson(),
			},
			args: args{
				id: "kozeppont",
			},
			want: Program{
				Id:          "kozeppont",
				Name:        "Középpont",
				Picture:     "https://archivum.hitradio.hu/data/img//2020/03/Középpont-aaEjVuEeG7-1584445579.jpg",
				Description: null.String{},
			},
			want1:   true,
			wantErr: false,
		},
		{
			name: "not found works",
			fields: fields{
				programs: mockPrograms(),
				episodes: mockEpisodes(),
				people:   mockPerson(),
			},
			args: args{
				id: "alma",
			},
			want:    Program{},
			want1:   false,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := Repository{
				programs: tt.fields.programs,
				episodes: tt.fields.episodes,
				people:   tt.fields.people,
			}
			got, got1, err := repo.GetProgram(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetProgram() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetProgram() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetProgram() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestRepository_GetEpisodes(t *testing.T) {
	type fields struct {
		programs []Program
		episodes []Episode
		people   []Person
	}
	type args struct {
		from      int
		pageSize  int
		programId null.String
		search    null.String
		tags      []string
		people    []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []Episode
		wantErr bool
	}{
		{
			name: "returns empty array if pageSize is 0",
			fields: fields{
				programs: mockPrograms(),
				episodes: mockEpisodes(),
				people:   mockPerson(),
			},
			args: args{
				from:      0,
				pageSize:  0,
				programId: null.String{},
				search:    null.String{},
				tags:      nil,
				people:    nil,
			},
			want:    []Episode{},
			wantErr: false,
		},
		{
			name: "returns empty array if form is out of bounds",
			fields: fields{
				programs: mockPrograms(),
				episodes: mockEpisodes(),
				people:   mockPerson(),
			},
			args: args{
				from:      999,
				pageSize:  10,
				programId: null.String{},
				search:    null.String{},
				tags:      nil,
				people:    nil,
			},
			want:    []Episode{},
			wantErr: false,
		},
		{
			name: "first page works",
			fields: fields{
				programs: mockPrograms(),
				episodes: mockEpisodes(),
				people:   mockPerson(),
			},
			args: args{
				from:      0,
				pageSize:  2,
				programId: null.String{},
				search:    null.String{},
				tags:      nil,
				people:    nil,
			},
			want:    mockEpisodes()[0:2],
			wantErr: false,
		},
		{
			name: "seconds page works",
			fields: fields{
				programs: mockPrograms(),
				episodes: mockEpisodes(),
				people:   mockPerson(),
			},
			args: args{
				from:      2,
				pageSize:  2,
				programId: null.String{},
				search:    null.String{},
				tags:      nil,
				people:    nil,
			},
			want:    mockEpisodes()[2:4],
			wantErr: false,
		},
		{
			name: "seconds page returns less than pageSize if needed",
			fields: fields{
				programs: mockPrograms(),
				episodes: mockEpisodes(),
				people:   mockPerson(),
			},
			args: args{
				from:      10,
				pageSize:  10,
				programId: null.String{},
				search:    null.String{},
				tags:      nil,
				people:    nil,
			},
			want:    mockEpisodes()[10:12],
			wantErr: false,
		},
		{
			name: "programId filter works",
			fields: fields{
				programs: mockPrograms(),
				episodes: mockEpisodes(),
				people:   mockPerson(),
			},
			args: args{
				from:      0,
				pageSize:  5,
				programId: null.StringFrom("kozeppont"),
				search:    null.String{},
				tags:      nil,
				people:    nil,
			},
			want:    mockEpisodes()[0:3],
			wantErr: false,
		},
		{
			name: "search filter works",
			fields: fields{
				programs: mockPrograms(),
				episodes: mockEpisodes(),
				people:   mockPerson(),
			},
			args: args{
				from:      0,
				pageSize:  5,
				programId: null.String{},
				search:    null.StringFrom("kampánypénzből kéretlen hívások"),
				tags:      nil,
				people:    nil,
			},
			want:    mockEpisodes()[1:2],
			wantErr: false,
		},
		{
			name: "tag filter works",
			fields: fields{
				programs: mockPrograms(),
				episodes: mockEpisodes(),
				people:   mockPerson(),
			},
			args: args{
				from:      0,
				pageSize:  5,
				programId: null.String{},
				search:    null.String{},
				tags:      []string{"digitalis"},
				people:    nil,
			},
			want:    mockEpisodes()[0:1],
			wantErr: false,
		},
		{
			name: "person filter works",
			fields: fields{
				programs: mockPrograms(),
				episodes: mockEpisodes(),
				people:   mockPerson(),
			},
			args: args{
				from:      0,
				pageSize:  5,
				programId: null.String{},
				search:    null.String{},
				tags:      nil,
				people:    []string{"Németh"},
			},
			want:    mockEpisodes()[5:9],
			wantErr: false,
		},
		{
			name: "all filter works at the same time",
			fields: fields{
				programs: mockPrograms(),
				episodes: mockEpisodes(),
				people:   mockPerson(),
			},
			args: args{
				from:      0,
				pageSize:  5,
				programId: null.StringFrom("hitkoznapok"),
				search:    null.StringFrom("kapcsolatfüggőség"),
				tags:      []string{"nemethsandor"},
				people:    []string{"Németh"},
			},
			want:    mockEpisodes()[6:7],
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := Repository{
				programs: tt.fields.programs,
				episodes: tt.fields.episodes,
				people:   tt.fields.people,
			}
			got, err := repo.GetEpisodes(tt.args.from, tt.args.pageSize, tt.args.programId, tt.args.search, tt.args.tags, tt.args.people)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetEpisodes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetEpisodes() got = %v, want %v", got, tt.want)
			}
		})
	}
}
