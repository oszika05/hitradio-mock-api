package program

import (
	"gopkg.in/guregu/null.v4"
	"time"
)

var program_kozeppont = Program{
	Id:          "kozeppont",
	Name:        "Középpont",
	Picture:     "https://archivum.hitradio.hu/data/img//2020/03/Középpont-aaEjVuEeG7-1584445579.jpg",
	Description: null.String{},
}

var program_gyujtopont = Program{
	Id:          "gyujtopont",
	Name:        "Gyújtópont",
	Picture:     "https://archivum.hitradio.hu/data/img//2020/03/Gyujtópont-PI54fRP-RM-1584445602.jpg",
	Description: null.String{},
}

var program_hitkoznapok = Program{
	Id:          "hitkoznapok",
	Name:        "Hitköznapok",
	Picture:     "https://archivum.hitradio.hu/data/img//2019/06/Hitköznapok%20blank-97vwVVSN1d-1559568542.jpg",
	Description: null.String{},
}

var program_nyugovora = Program{
	Id:          "nyugovora",
	Name:        "Nyugovóra",
	Picture:     "https://archivum.hitradio.hu/data/img//2020/03/nyugovóra-MJLmTxqt5k-1584446182.jpg",
	Description: null.String{},
}

var person_1 = Person{
	Id:          "1",
	Name:        "Király Tamás",
	Type:        PersonTypeHost,
	Picture:     null.StringFrom("https://scontent-vie1-1.xx.fbcdn.net/v/t1.6435-9/67699021_112428760098874_5586161939806945280_n.jpg?_nc_cat=107&ccb=1-5&_nc_sid=8bfeb9&_nc_ohc=jLjDglC87FcAX9m2ubg&_nc_ht=scontent-vie1-1.xx&oh=21c960625e3f537eb29cf2c19a2ce3c2&oe=6183179C"),
	Description: null.String{},
}

var person_2 = Person{
	Id:          "2",
	Name:        "Kauzál Alexandra",
	Type:        PersonTypeHost,
	Picture:     null.StringFrom("https://archivum.hitradio.hu/data/img//2018/01/Családi%20magazin%20blank-KYI-2hUoi0-1516026696.jpg"),
	Description: null.String{},
}

var person_3 = Person{
	Id:          "3",
	Name:        "Petrőcz Katalin",
	Type:        PersonTypeHost,
	Picture:     null.StringFrom("https://i.ytimg.com/vi/BoYYN5oaly8/maxresdefault.jpg"),
	Description: null.String{},
}

var guest_1 = Person{
	Id:          "4",
	Name:        "Németh Sándor",
	Type:        PersonTypeGuest,
	Picture:     null.StringFrom("https://mandiner.hu/attachment/0346/345825_nemeth_sandor1.jpg"),
	Description: null.StringFrom("A hit gyülekezete vezető lelkésze"),
}

var guest_2 = Person{
	Id:          "5",
	Name:        "Balla Attila",
	Type:        PersonTypeGuest,
	Picture:     null.String{},
	Description: null.StringFrom("A Digitális Jólét Program (DJP) vállalkozásfejlesztési ügyvezető-helyettese"),
}

var guest_3 = Person{
	Id:          "6",
	Name:        "Dr. Péterfalvi Attila",
	Type:        PersonTypeGuest,
	Picture:     null.String{},
	Description: null.StringFrom("Nemzeti Adatvédelmi és Információszabadság Hatóság (NAIH) elnöke"),
}

var guest_4 = Person{
	Id:          "7",
	Name:        "Szarvas Szilveszter",
	Type:        PersonTypeGuest,
	Picture:     null.String{},
	Description: null.StringFrom("PestiSrácok.hu hírigazgatója"),
}

func mockPrograms() []Program {
	return []Program{
		program_kozeppont,
		program_gyujtopont,
		program_hitkoznapok,
		program_nyugovora,
	}
}

func mockEpisodes() []Episode {
	return []Episode{
		{
			Id:          "kozeppont-1",
			Title:       "Az év digitális falui",
			Date:        time.Date(2021, 9, 29, 11, 9, 0, 0, time.Local),
			Description: null.StringFrom("Balla Attila- a Digitális Jólét Program (DJP) vállalkozásfejlesztési ügyvezető-helyettese, Köcse Tibor- Nagypáli polgármestere, Visy László -Rábapordány polgármestere ,Ősz Ferenc -Bakonysárkány polgármestere"),
			Tags:        []string{"digitalis"},
			Program:     program_kozeppont,
			AudioUrl:    "https://anchor.fm/s/4891e7c0/podcast/play/39752599/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2021-8-3%2F880cefb0-99a9-0771-a1c9-84e09286f844.mp3",
			Hosts: []Person{
				person_1,
			},
			Guests: []Person{
				guest_2,
			},
		},
		{
			Id:          "kozeppont-2",
			Title:       "Kampánypénzből kéretlen hívások – mi várható, ha élesedik a választási kampány?",
			Date:        time.Date(2021, 9, 29, 10, 9, 0, 0, time.Local),
			Description: null.StringFrom("Dr. Péterfalvi Attila - Nemzeti Adatvédelmi és Információszabadság Hatóság (NAIH) elnöke, Szarvas Szilveszter - PestiSrácok.hu hírigazgatója"),
			Tags:        []string{},
			Program:     program_kozeppont,
			AudioUrl:    "https://anchor.fm/s/4891e7c0/podcast/play/39752599/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2021-8-3%2F880cefb0-99a9-0771-a1c9-84e09286f844.mp3",
			Hosts: []Person{
				person_1,
			},
			Guests: []Person{
				guest_3,
				guest_4,
			},
		},
		{
			Id:          "kozeppont-3",
			Title:       "Hogyan látták a német választásokat a helyi lakosok?",
			Date:        time.Date(2021, 9, 28, 11, 9, 0, 0, time.Local),
			Description: null.StringFrom("A Középpont műsorán belül a 2021 német választások eredményeit vitattuk meg a helyi német lakosok szemszögéből. Műsorvezető: Petrőcz Katalin Szerkesztő: Czinege Johannán Benjámin"),
			Tags:        []string{},
			Program:     program_kozeppont,
			AudioUrl:    "https://anchor.fm/s/4891e7c0/podcast/play/39752599/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2021-8-3%2F880cefb0-99a9-0771-a1c9-84e09286f844.mp3",
			Hosts: []Person{
				person_1,
			},
			Guests: nil,
		},
		{
			Id:          "gyujtopont-1",
			Title:       "Lerántani a hallgatás leplét- Jobst Bittner,interjú",
			Date:        time.Date(2021, 9, 24, 16, 10, 0, 0, time.Local),
			Description: null.StringFrom("Jobst Bittner, a Lerántani a hallgatás leplét című könyv szerzője volt a Hit Rádió litvániai különkiadásának vendége. A német lelkész saját családi drámáján keresztül döbbent rá arra, hogy a feldolgozatlan múlt milyen komoly szellemi-erkölcsi károkat okoz egy-egy nemzet és város életében. A hallgatások típusairól, a feldolgozás módszeréről, a németországi munkájukról kérdeztük Vilniusban, a vasútállomáson, egy holokauszt-megemlékezést követően. Műsorvezető: Király Tamás Fordította: Dezsényi István #Litvánia #zsidóság #holokauszt"),
			Tags:        []string{},
			Program:     program_gyujtopont,
			AudioUrl:    "https://anchor.fm/s/4891e7c0/podcast/play/39752599/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2021-8-3%2F880cefb0-99a9-0771-a1c9-84e09286f844.mp3",
			Hosts: []Person{
				person_2,
			},
			Guests: []Person{
				guest_4,
			},
		},
		{
			Id:          "gyujtopont-2",
			Title:       "Oroszországi választások",
			Date:        time.Date(2021, 9, 21, 16, 9, 0, 0, time.Local),
			Description: null.StringFrom("Bendarzsevszkij Anton- oroszország szakértő"),
			Tags:        []string{},
			Program:     program_gyujtopont,
			AudioUrl:    "https://anchor.fm/s/4891e7c0/podcast/play/39752599/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2021-8-3%2F880cefb0-99a9-0771-a1c9-84e09286f844.mp3",
			Hosts: []Person{
				person_3,
			},
			Guests: []Person{
				guest_4,
			},
		},
		{
			Id:          "gyujtopont-3",
			Title:       "Helyre lehet állítani a jogállamiságot?",
			Date:        time.Date(2021, 9, 21, 16, 9, 0, 0, time.Local),
			Description: null.StringFrom("Tóth Norbert- nemzetközi jogász, Nemzeti Közszolgálati Egyetem docense"),
			Tags:        []string{},
			Program:     program_gyujtopont,
			AudioUrl:    "https://anchor.fm/s/4891e7c0/podcast/play/39752599/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2021-8-3%2F880cefb0-99a9-0771-a1c9-84e09286f844.mp3",
			Hosts: []Person{
				person_2,
			},
			Guests: []Person{
				guest_1,
			},
		},
		{
			Id:          "hitkoznapok-1",
			Title:       "Kiút a kapcsolatfüggőségből - az igazi szeretethez vezető út",
			Date:        time.Date(2021, 9, 29, 12, 9, 0, 0, time.Local),
			Description: null.StringFrom("Kapcsolatfüggőség - függőség vagy valódi szeretet? - 3. rész Kiút a kapcsolatfüggőségből - az igazi szeretethez vezető út"),
			Tags:        []string{"nemethsandor"},
			Program:     program_hitkoznapok,
			AudioUrl:    "https://anchor.fm/s/4891e7c0/podcast/play/39752599/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2021-8-3%2F880cefb0-99a9-0771-a1c9-84e09286f844.mp3",
			Hosts: []Person{
				person_3,
			},
			Guests: []Person{
				guest_1,
			},
		},
		{
			Id:          "hitkoznapok-2",
			Title:       "Világnézetek a Biblia tükrében: Szabadkőművesség",
			Date:        time.Date(2021, 9, 27, 12, 9, 0, 0, time.Local),
			Description: null.StringFrom("Árvai Ábrahám- történész"),
			Tags:        []string{},
			Program:     program_hitkoznapok,
			AudioUrl:    "https://anchor.fm/s/4891e7c0/podcast/play/39752599/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2021-8-3%2F880cefb0-99a9-0771-a1c9-84e09286f844.mp3",
			Hosts: []Person{
				person_2,
			},
			Guests: []Person{
				guest_1,
			},
		},
		{
			Id:          "hitkoznapok-3",
			Title:       "Harc a normalitásért",
			Date:        time.Date(2021, 9, 24, 12, 9, 0, 0, time.Local),
			Description: null.StringFrom("Surjányi Csaba- a Hit Gyülekezete lelkésze, a Szent Pál Akadémia tanára"),
			Tags:        []string{},
			Program:     program_hitkoznapok,
			AudioUrl:    "https://anchor.fm/s/4891e7c0/podcast/play/39752599/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2021-8-3%2F880cefb0-99a9-0771-a1c9-84e09286f844.mp3",
			Hosts: []Person{
				person_1,
			},
			Guests: []Person{
				guest_1,
			},
		},
		{
			Id:          "nyugovora-1",
			Title:       "30. Moabita Ruth",
			Date:        time.Date(2020, 2, 10, 20, 2, 0, 0, time.Local),
			Description: null.String{},
			Tags:        []string{},
			Program:     program_nyugovora,
			AudioUrl:    "https://anchor.fm/s/4891e7c0/podcast/play/39752599/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2021-8-3%2F880cefb0-99a9-0771-a1c9-84e09286f844.mp3",
			Hosts:       nil,
			Guests:      nil,
		},
		{
			Id:          "nyugovora-2",
			Title:       "29. Sámson",
			Date:        time.Date(2020, 2, 9, 20, 2, 0, 0, time.Local),
			Description: null.String{},
			Tags:        []string{},
			Program:     program_nyugovora,
			AudioUrl:    "https://anchor.fm/s/4891e7c0/podcast/play/39752599/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2021-8-3%2F880cefb0-99a9-0771-a1c9-84e09286f844.mp3",
			Hosts:       nil,
			Guests:      nil,
		},
		{
			Id:          "nyugovora-3",
			Title:       "28. Gedeon",
			Date:        time.Date(2020, 2, 7, 20, 2, 0, 0, time.Local),
			Description: null.String{},
			Tags:        []string{},
			Program:     program_nyugovora,
			AudioUrl:    "https://anchor.fm/s/4891e7c0/podcast/play/39752599/https%3A%2F%2Fd3ctxlq1ktw2nl.cloudfront.net%2Fstaging%2F2021-8-3%2F880cefb0-99a9-0771-a1c9-84e09286f844.mp3",
			Hosts:       nil,
			Guests:      nil,
		},
	}
}

func mockPerson() []Person {
	return []Person{
		person_1,
		person_2,
		person_3,
		guest_1,
		guest_2,
		guest_3,
		guest_4,
	}
}
