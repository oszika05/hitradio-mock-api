package news

import (
	"time"
)

func mockTime(delta int) time.Time {
	return time.Now().Add(time.Duration(-1*delta) * time.Minute)
}

func mockPicture(index int) string {
	mockPictures := []string{
		"https://cdn.portfolio.hu/articles/images-md/f/a/c/facebook-443082.jpg",
		"https://kep.cdn.indexvas.hu/1/0/4012/40124/401240/40124079_3a4ce3fcae327c045c10a895d5cd1fa2_wm.jpg",
		"https://upload.wikimedia.org/wikipedia/commons/thumb/a/ab/Patates.jpg/2560px-Patates.jpg",
		"https://upload.wikimedia.org/wikipedia/commons/7/76/Red_apple.jpg",
		"https://upload.wikimedia.org/wikipedia/commons/c/cf/Pears.jpg",
		"https://upload.wikimedia.org/wikipedia/commons/thumb/7/74/Nectarines_summer_2006.jpg/1024px-Nectarines_summer_2006.jpg",
		"https://upload.wikimedia.org/wikipedia/commons/thumb/5/59/Kiwi1.1.jpg/2560px-Kiwi1.1.jpg",
		"https://upload.wikimedia.org/wikipedia/commons/thumb/5/5d/Illustration_Apium_graveolens0.jpg/1024px-Illustration_Apium_graveolens0.jpg",
		"https://upload.wikimedia.org/wikipedia/commons/5/5f/KohlrabiinMarket.jpg",
		"https://nlc.p3k.hu/uploads/2021/01/csirke-vadcsirke-kakas.jpg",
	}

	return mockPictures[index%len(mockPictures)]
}

func getMockNews() []News {
	return []News{
		{
			Id:      "1",
			Title:   "Leállt a Facebook",
			Picture: mockPicture(0),
			Date:    mockTime(15),
			Tags:    []string{"facebook", "hiba"},
			Content: "Ecce. Tata de secundus pes, captis epos! Lacta germanus solitudo est. Sunt amores contactus clemens, noster bullaes. Ecce, flavum fermium! Abnobas persuadere in castus avenio! Omnias cantare, tanquam audax lacta. A falsis, scutum fortis glos. Barbatus deuss ducunt ad demolitione. Cur absolutio credere? Buxums sunt scutums de ferox orexis. Accentor ires, tanquam alter gluten. Ubi est brevis ventus? Ecce, poeta! Fluctuss mori in domesticus mare! Sunt terrores reperire lotus, neuter tumultumquees. Sunt devirginatoes perdere azureus, primus equisoes. Eleates, advena, et particula. Solitudos mori in emeritis antenna! Zelus de grandis habitio, imitari fuga! Cum vita resistere, omnes magisteres pugna placidus, raptus advenaes!",
		},
		{
			Id:      "2",
			Title:   "Válaszok nélkül maradt a magyar jogállamiságot vizsgáló tényfeltáró EU-delegáció",
			Picture: mockPicture(1),
			Date:    mockTime(20),
			Tags:    []string{"eu", "magyariorszag", "jogallam"},
			Content: "O, belay. The biscuit eater tastes with hunger, view the captain's quarters before it stutters. Sunny, weird hornpipes cowardly fight a shiny, undead biscuit eater. The shipmate burns with courage, raid the brig. Ah! Pieces o' adventure are forever warm. Gar, ye mighty plank- set sails for riddle! Never haul a son. How coal-black. You loot like a codfish. Cockroach of a shiny life, raid the hunger! The captain hoists with horror, taste the cook islands until it grows. The cannon pulls with greed, endure the captain's quarters. The pirate endures with fortune, blow the galley until it dies!",
		},
		{
			Id:      "3",
			Title:   "Bajban van Ukrajna, nincs elég gáz télre",
			Picture: mockPicture(2),
			Date:    mockTime(45),
			Tags:    []string{"ukrajna", "gaz"},
			Content: "O, belay. The biscuit eater tastes with hunger, view the captain's quarters before it stutters. Sunny, weird hornpipes cowardly fight a shiny, undead biscuit eater. The shipmate burns with courage, raid the brig. Ah! Pieces o' adventure are forever warm. Gar, ye mighty plank- set sails for riddle! Never haul a son. How coal-black. You loot like a codfish. Cockroach of a shiny life, raid the hunger! The captain hoists with horror, taste the cook islands until it grows. The cannon pulls with greed, endure the captain's quarters. The pirate endures with fortune, blow the galley until it dies!",
		},
		{
			Id:      "4",
			Title:   "Most van gáz, beindult az Északi Áramlat-2",
			Picture: mockPicture(3),
			Date:    mockTime(24*60 + 15),
			Tags:    []string{"nemetorszag", "gaz"},
			Content: "O, belay. The biscuit eater tastes with hunger, view the captain's quarters before it stutters. Sunny, weird hornpipes cowardly fight a shiny, undead biscuit eater. The shipmate burns with courage, raid the brig. Ah! Pieces o' adventure are forever warm. Gar, ye mighty plank- set sails for riddle! Never haul a son. How coal-black. You loot like a codfish. Cockroach of a shiny life, raid the hunger! The captain hoists with horror, taste the cook islands until it grows. The cannon pulls with greed, endure the captain's quarters. The pirate endures with fortune, blow the galley until it dies!",
		},
		{
			Id:      "5",
			Title:   "Orbán Viktor elutazik, fontos csúcstalálkozón vesz részt",
			Picture: mockPicture(4),
			Date:    mockTime(24*60 + 15),
			Tags:    []string{"orban viktor", "turizmus"},
			Content: "O, belay. The biscuit eater tastes with hunger, view the captain's quarters before it stutters. Sunny, weird hornpipes cowardly fight a shiny, undead biscuit eater. The shipmate burns with courage, raid the brig. Ah! Pieces o' adventure are forever warm. Gar, ye mighty plank- set sails for riddle! Never haul a son. How coal-black. You loot like a codfish. Cockroach of a shiny life, raid the hunger! The captain hoists with horror, taste the cook islands until it grows. The cannon pulls with greed, endure the captain's quarters. The pirate endures with fortune, blow the galley until it dies!",
		},
		{
			Id:      "6",
			Title:   "Becsületsértésben találták bűnösnek Donáth Lászlót",
			Picture: mockPicture(5),
			Date:    mockTime(24*60 + 15),
			Tags:    []string{"Donáth László", "magyariorszag"},
			Content: ", belay. The biscuit eater tastes with hunger, view the captain's quarters before it stutters. Sunny, weird hornpipes cowardly fight a shiny, undead biscuit eater. The shipmate burns with courage, raid the brig. Ah! Pieces o' adventure are forever warm. Gar, ye mighty plank- set sails for riddle! Never haul a son. How coal-black. You loot like a codfish. Cockroach of a shiny life, raid the hunger! The captain hoists with horror, taste the cook islands until it grows. The cannon pulls with greed, endure the captain's quarters. The pirate endures with fortune, blow the galley until it dies!",
		},
		{
			Id:      "7",
			Title:   "Ön tudja, hová kell lehúzódni, ha jön a mentőautó?",
			Picture: mockPicture(6),
			Date:    mockTime(24*60 + 15),
			Tags:    []string{"mentoauto", "kozlekedes", "kresz"},
			Content: "O, belay. The biscuit eater tastes with hunger, view the captain's quarters before it stutters. Sunny, weird hornpipes cowardly fight a shiny, undead biscuit eater. The shipmate burns with courage, raid the brig. Ah! Pieces o' adventure are forever warm. Gar, ye mighty plank- set sails for riddle! Never haul a son. How coal-black. You loot like a codfish. Cockroach of a shiny life, raid the hunger! The captain hoists with horror, taste the cook islands until it grows. The cannon pulls with greed, endure the captain's quarters. The pirate endures with fortune, blow the galley until it dies!",
		},
		{
			Id:      "8",
			Title:   "Rossz hír az autósoknak, folyamatosan drágul a kőolaj",
			Picture: mockPicture(7),
			Date:    mockTime(24*60 + 15),
			Tags:    []string{"auto", "koolaj", "benzin"},
			Content: "O, belay. The biscuit eater tastes with hunger, view the captain's quarters before it stutters. Sunny, weird hornpipes cowardly fight a shiny, undead biscuit eater. The shipmate burns with courage, raid the brig. Ah! Pieces o' adventure are forever warm. Gar, ye mighty plank- set sails for riddle! Never haul a son. How coal-black. You loot like a codfish. Cockroach of a shiny life, raid the hunger! The captain hoists with horror, taste the cook islands until it grows. The cannon pulls with greed, endure the captain's quarters. The pirate endures with fortune, blow the galley until it dies!",
		},
		{
			Id:      "9",
			Title:   "Válaszok nélkül maradt a magyar jogállamiságot vizsgáló tényfeltáró EU-delegáció",
			Picture: mockPicture(8),
			Date:    mockTime(24*60 + 15),
			Tags:    []string{"eu", "magyariorszag", "jogallam"},
			Content: "O, belay. The biscuit eater tastes with hunger, view the captain's quarters before it stutters. Sunny, weird hornpipes cowardly fight a shiny, undead biscuit eater. The shipmate burns with courage, raid the brig. Ah! Pieces o' adventure are forever warm. Gar, ye mighty plank- set sails for riddle! Never haul a son. How coal-black. You loot like a codfish. Cockroach of a shiny life, raid the hunger! The captain hoists with horror, taste the cook islands until it grows. The cannon pulls with greed, endure the captain's quarters. The pirate endures with fortune, blow the galley until it dies!",
		},
		{
			Id:      "10",
			Title:   "Összeesett tornaórán egy 17 éves fiú, a tanára segített újraéleszteni",
			Picture: mockPicture(9),
			Date:    mockTime(24*60 + 15),
			Tags:    []string{"egeszseg", "magyariorszag"},
			Content: "O, belay. The biscuit eater tastes with hunger, view the captain's quarters before it stutters. Sunny, weird hornpipes cowardly fight a shiny, undead biscuit eater. The shipmate burns with courage, raid the brig. Ah! Pieces o' adventure are forever warm. Gar, ye mighty plank- set sails for riddle! Never haul a son. How coal-black. You loot like a codfish. Cockroach of a shiny life, raid the hunger! The captain hoists with horror, taste the cook islands until it grows. The cannon pulls with greed, endure the captain's quarters. The pirate endures with fortune, blow the galley until it dies!",
		},
		{
			Id:      "11",
			Title:   "Politico: Várhelyi Olivér az EU-bővítés Voldemortja",
			Picture: mockPicture(10),
			Date:    mockTime(24*60 + 15),
			Tags:    []string{"eu", "politico"},
			Content: "O, belay. The biscuit eater tastes with hunger, view the captain's quarters before it stutters. Sunny, weird hornpipes cowardly fight a shiny, undead biscuit eater. The shipmate burns with courage, raid the brig. Ah! Pieces o' adventure are forever warm. Gar, ye mighty plank- set sails for riddle! Never haul a son. How coal-black. You loot like a codfish. Cockroach of a shiny life, raid the hunger! The captain hoists with horror, taste the cook islands until it grows. The cannon pulls with greed, endure the captain's quarters. The pirate endures with fortune, blow the galley until it dies!",
		},
		{
			Id:      "12",
			Title:   "A magyar GDP egyharmadával zuhant a Facebook értéke",
			Picture: mockPicture(11),
			Date:    mockTime(24*60 + 15),
			Tags:    []string{"facebook", "leallas"},
			Content: "O, belay. The biscuit eater tastes with hunger, view the captain's quarters before it stutters. Sunny, weird hornpipes cowardly fight a shiny, undead biscuit eater. The shipmate burns with courage, raid the brig. Ah! Pieces o' adventure are forever warm. Gar, ye mighty plank- set sails for riddle! Never haul a son. How coal-black. You loot like a codfish. Cockroach of a shiny life, raid the hunger! The captain hoists with horror, taste the cook islands until it grows. The cannon pulls with greed, endure the captain's quarters. The pirate endures with fortune, blow the galley until it dies!",
		},
		{
			Id:      "13",
			Title:   "Facebook-leállás: törlődtek a hálózatról az adatközpontok",
			Picture: mockPicture(12),
			Date:    mockTime(24*60 + 15),
			Tags:    []string{"facebook", "leallas"},
			Content: "O, belay. The biscuit eater tastes with hunger, view the captain's quarters before it stutters. Sunny, weird hornpipes cowardly fight a shiny, undead biscuit eater. The shipmate burns with courage, raid the brig. Ah! Pieces o' adventure are forever warm. Gar, ye mighty plank- set sails for riddle! Never haul a son. How coal-black. You loot like a codfish. Cockroach of a shiny life, raid the hunger! The captain hoists with horror, taste the cook islands until it grows. The cannon pulls with greed, endure the captain's quarters. The pirate endures with fortune, blow the galley until it dies!",
		},
		{
			Id:      "14",
			Title:   "A leállás közben állítólag eladó volt a facebook.com domain",
			Picture: mockPicture(13),
			Date:    mockTime(24*60 + 15),
			Tags:    []string{"facebook", "leallas"},
			Content: "O, belay. The biscuit eater tastes with hunger, view the captain's quarters before it stutters. Sunny, weird hornpipes cowardly fight a shiny, undead biscuit eater. The shipmate burns with courage, raid the brig. Ah! Pieces o' adventure are forever warm. Gar, ye mighty plank- set sails for riddle! Never haul a son. How coal-black. You loot like a codfish. Cockroach of a shiny life, raid the hunger! The captain hoists with horror, taste the cook islands until it grows. The cannon pulls with greed, endure the captain's quarters. The pirate endures with fortune, blow the galley until it dies!",
		},
		{
			Id:      "15",
			Title:   "Kimondani is hihetetlen: Zuckerberg 2 ezer milliárd forintnyi összeget bukott néhány óra alatt",
			Picture: mockPicture(14),
			Date:    mockTime(24*60 + 15),
			Tags:    []string{"facebook", "leallas"},
			Content: "O, belay. The biscuit eater tastes with hunger, view the captain's quarters before it stutters. Sunny, weird hornpipes cowardly fight a shiny, undead biscuit eater. The shipmate burns with courage, raid the brig. Ah! Pieces o' adventure are forever warm. Gar, ye mighty plank- set sails for riddle! Never haul a son. How coal-black. You loot like a codfish. Cockroach of a shiny life, raid the hunger! The captain hoists with horror, taste the cook islands until it grows. The cannon pulls with greed, endure the captain's quarters. The pirate endures with fortune, blow the galley until it dies!",
		},
		{
			Id:      "16",
			Title:   "Kotrógéppel kellett kimenteni egy veremből a kiselefántot",
			Picture: mockPicture(15),
			Date:    mockTime(24*60 + 15),
			Tags:    []string{"elefant"},
			Content: "O, belay. The biscuit eater tastes with hunger, view the captain's quarters before it stutters. Sunny, weird hornpipes cowardly fight a shiny, undead biscuit eater. The shipmate burns with courage, raid the brig. Ah! Pieces o' adventure are forever warm. Gar, ye mighty plank- set sails for riddle! Never haul a son. How coal-black. You loot like a codfish. Cockroach of a shiny life, raid the hunger! The captain hoists with horror, taste the cook islands until it grows. The cannon pulls with greed, endure the captain's quarters. The pirate endures with fortune, blow the galley until it dies!",
		},
		{
			Id:      "17",
			Title:   "Így nézett ki a La Palma-i lávafolyam az űrből",
			Picture: mockPicture(16),
			Date:    mockTime(24*60 + 15),
			Tags:    []string{"ur"},
			Content: "O, belay. The biscuit eater tastes with hunger, view the captain's quarters before it stutters. Sunny, weird hornpipes cowardly fight a shiny, undead biscuit eater. The shipmate burns with courage, raid the brig. Ah! Pieces o' adventure are forever warm. Gar, ye mighty plank- set sails for riddle! Never haul a son. How coal-black. You loot like a codfish. Cockroach of a shiny life, raid the hunger! The captain hoists with horror, taste the cook islands until it grows. The cannon pulls with greed, endure the captain's quarters. The pirate endures with fortune, blow the galley until it dies!",
		},
		{
			Id:      "18",
			Title:   "Snoop Dogg szerint Harry herceg az egyik legtökösebb ember, akit ismer",
			Picture: mockPicture(17),
			Date:    mockTime(24*60 + 15),
			Tags:    []string{"snoopy"},
			Content: "O, belay. The biscuit eater tastes with hunger, view the captain's quarters before it stutters. Sunny, weird hornpipes cowardly fight a shiny, undead biscuit eater. The shipmate burns with courage, raid the brig. Ah! Pieces o' adventure are forever warm. Gar, ye mighty plank- set sails for riddle! Never haul a son. How coal-black. You loot like a codfish. Cockroach of a shiny life, raid the hunger! The captain hoists with horror, taste the cook islands until it grows. The cannon pulls with greed, endure the captain's quarters. The pirate endures with fortune, blow the galley until it dies!",
		},
		{
			Id:      "19",
			Title:   "Magyar György: Nem írhatják át a játékszabályokat Karácsonyék",
			Picture: mockPicture(18),
			Date:    mockTime(24*60 + 15),
			Tags:    []string{"Magyar György", "magyariorszag", "politika"},
			Content: "O, belay. The biscuit eater tastes with hunger, view the captain's quarters before it stutters. Sunny, weird hornpipes cowardly fight a shiny, undead biscuit eater. The shipmate burns with courage, raid the brig. Ah! Pieces o' adventure are forever warm. Gar, ye mighty plank- set sails for riddle! Never haul a son. How coal-black. You loot like a codfish. Cockroach of a shiny life, raid the hunger! The captain hoists with horror, taste the cook islands until it grows. The cannon pulls with greed, endure the captain's quarters. The pirate endures with fortune, blow the galley until it dies!",
		},
		{
			Id:      "20",
			Title:   "Vereséget szenvedtek a populista pártok az olasz választásokon",
			Picture: mockPicture(19),
			Date:    mockTime(24*60 + 15),
			Tags:    []string{"valasztas", "olasz"},
			Content: "O, belay. The biscuit eater tastes with hunger, view the captain's quarters before it stutters. Sunny, weird hornpipes cowardly fight a shiny, undead biscuit eater. The shipmate burns with courage, raid the brig. Ah! Pieces o' adventure are forever warm. Gar, ye mighty plank- set sails for riddle! Never haul a son. How coal-black. You loot like a codfish. Cockroach of a shiny life, raid the hunger! The captain hoists with horror, taste the cook islands until it grows. The cannon pulls with greed, endure the captain's quarters. The pirate endures with fortune, blow the galley until it dies!",
		},
		{
			Id:      "21",
			Title:   "Több száz éhező majom tartja rettegésben Bali egyik városát",
			Picture: mockPicture(20),
			Date:    mockTime(24*60 + 15),
			Tags:    []string{"majom", "Bali"},
			Content: "O, belay. The biscuit eater tastes with hunger, view the captain's quarters before it stutters. Sunny, weird hornpipes cowardly fight a shiny, undead biscuit eater. The shipmate burns with courage, raid the brig. Ah! Pieces o' adventure are forever warm. Gar, ye mighty plank- set sails for riddle! Never haul a son. How coal-black. You loot like a codfish. Cockroach of a shiny life, raid the hunger! The captain hoists with horror, taste the cook islands until it grows. The cannon pulls with greed, endure the captain's quarters. The pirate endures with fortune, blow the galley until it dies!",
		},
		{
			Id:      "22",
			Title:   "Vona Gábor: Öt ok Márki-Zay Péter visszalépésére",
			Picture: mockPicture(21),
			Date:    mockTime(24*60 + 15),
			Tags:    []string{"vona"},
			Content: "O, belay. The biscuit eater tastes with hunger, view the captain's quarters before it stutters. Sunny, weird hornpipes cowardly fight a shiny, undead biscuit eater. The shipmate burns with courage, raid the brig. Ah! Pieces o' adventure are forever warm. Gar, ye mighty plank- set sails for riddle! Never haul a son. How coal-black. You loot like a codfish. Cockroach of a shiny life, raid the hunger! The captain hoists with horror, taste the cook islands until it grows. The cannon pulls with greed, endure the captain's quarters. The pirate endures with fortune, blow the galley until it dies!",
		},
		{
			Id:      "23",
			Title:   "Újraszámlálás és panaszok az előválasztás első fordulója után",
			Picture: mockPicture(22),
			Date:    mockTime(24*60 + 15),
			Tags:    []string{"elovalasztas"},
			Content: "O, belay. The biscuit eater tastes with hunger, view the captain's quarters before it stutters. Sunny, weird hornpipes cowardly fight a shiny, undead biscuit eater. The shipmate burns with courage, raid the brig. Ah! Pieces o' adventure are forever warm. Gar, ye mighty plank- set sails for riddle! Never haul a son. How coal-black. You loot like a codfish. Cockroach of a shiny life, raid the hunger! The captain hoists with horror, taste the cook islands until it grows. The cannon pulls with greed, endure the captain's quarters. The pirate endures with fortune, blow the galley until it dies!",
		},
		{
			Id:      "24",
			Title:   "Kiderült, a Pfizer mennyire véd a delta-variánstól",
			Picture: mockPicture(23),
			Date:    mockTime(24*60 + 15),
			Tags:    []string{"pfizer", "vakcina", "koronavirus"},
			Content: "O, belay. The biscuit eater tastes with hunger, view the captain's quarters before it stutters. Sunny, weird hornpipes cowardly fight a shiny, undead biscuit eater. The shipmate burns with courage, raid the brig. Ah! Pieces o' adventure are forever warm. Gar, ye mighty plank- set sails for riddle! Never haul a son. How coal-black. You loot like a codfish. Cockroach of a shiny life, raid the hunger! The captain hoists with horror, taste the cook islands until it grows. The cannon pulls with greed, endure the captain's quarters. The pirate endures with fortune, blow the galley until it dies!",
		},
		{
			Id:      "25",
			Title:   "Részeg voltam – állítja a náci karlendítés közben lefotózott ózdi alpolgármester",
			Picture: mockPicture(24),
			Date:    mockTime(24*60 + 15),
			Tags:    []string{"magyariorszag", "jobbik"},
			Content: "O, belay. The biscuit eater tastes with hunger, view the captain's quarters before it stutters. Sunny, weird hornpipes cowardly fight a shiny, undead biscuit eater. The shipmate burns with courage, raid the brig. Ah! Pieces o' adventure are forever warm. Gar, ye mighty plank- set sails for riddle! Never haul a son. How coal-black. You loot like a codfish. Cockroach of a shiny life, raid the hunger! The captain hoists with horror, taste the cook islands until it grows. The cannon pulls with greed, endure the captain's quarters. The pirate endures with fortune, blow the galley until it dies!",
		},
		{
			Id:      "26",
			Title:   "Viszályhoz vezetett a malacválság a húskereskedőknél",
			Picture: mockPicture(25),
			Date:    mockTime(24*60 + 15),
			Tags:    []string{"malac", "valsag"},
			Content: "O, belay. The biscuit eater tastes with hunger, view the captain's quarters before it stutters. Sunny, weird hornpipes cowardly fight a shiny, undead biscuit eater. The shipmate burns with courage, raid the brig. Ah! Pieces o' adventure are forever warm. Gar, ye mighty plank- set sails for riddle! Never haul a son. How coal-black. You loot like a codfish. Cockroach of a shiny life, raid the hunger! The captain hoists with horror, taste the cook islands until it grows. The cannon pulls with greed, endure the captain's quarters. The pirate endures with fortune, blow the galley until it dies!",
		},
		{
			Id:      "27",
			Title:   "A Mercedes a rajtrács utolsó helyére küldheti Hamiltont a Török Nagydíjon",
			Picture: mockPicture(26),
			Date:    mockTime(24*60 + 15),
			Tags:    []string{"sport", "mercedes"},
			Content: "O, belay. The biscuit eater tastes with hunger, view the captain's quarters before it stutters. Sunny, weird hornpipes cowardly fight a shiny, undead biscuit eater. The shipmate burns with courage, raid the brig. Ah! Pieces o' adventure are forever warm. Gar, ye mighty plank- set sails for riddle! Never haul a son. How coal-black. You loot like a codfish. Cockroach of a shiny life, raid the hunger! The captain hoists with horror, taste the cook islands until it grows. The cannon pulls with greed, endure the captain's quarters. The pirate endures with fortune, blow the galley until it dies!",
		},
		{
			Id:      "28",
			Title:   "Elbukott a Las Vegas, egyetlen veretlen csapat maradt az NFL-ben",
			Picture: mockPicture(27),
			Date:    mockTime(24*60 + 15),
			Tags:    []string{"sport", "NFL"},
			Content: "O, belay. The biscuit eater tastes with hunger, view the captain's quarters before it stutters. Sunny, weird hornpipes cowardly fight a shiny, undead biscuit eater. The shipmate burns with courage, raid the brig. Ah! Pieces o' adventure are forever warm. Gar, ye mighty plank- set sails for riddle! Never haul a son. How coal-black. You loot like a codfish. Cockroach of a shiny life, raid the hunger! The captain hoists with horror, taste the cook islands until it grows. The cannon pulls with greed, endure the captain's quarters. The pirate endures with fortune, blow the galley until it dies!",
		},
		{
			Id:      "29",
			Title:   "Felfüggesztett fogházat kapott egy hajdúbagosi férfi, aki levágta a lábáról a nyomkövetőt",
			Picture: mockPicture(28),
			Date:    mockTime(24*60 + 15),
			Tags:    []string{"borton"},
			Content: "O, belay. The biscuit eater tastes with hunger, view the captain's quarters before it stutters. Sunny, weird hornpipes cowardly fight a shiny, undead biscuit eater. The shipmate burns with courage, raid the brig. Ah! Pieces o' adventure are forever warm. Gar, ye mighty plank- set sails for riddle! Never haul a son. How coal-black. You loot like a codfish. Cockroach of a shiny life, raid the hunger! The captain hoists with horror, taste the cook islands until it grows. The cannon pulls with greed, endure the captain's quarters. The pirate endures with fortune, blow the galley until it dies!",
		},
		{
			Id:      "30",
			Title:   "Enyhülnek a tengeralattjáró-ügyben átvert franciák, Emmanuel Macron egyeztetett az amerikai külügyminiszterrel",
			Picture: mockPicture(29),
			Date:    mockTime(24*60 + 15),
			Tags:    []string{"tengeralattjaro", "franciaorszag", "usa"},
			Content: "O, belay. The biscuit eater tastes with hunger, view the captain's quarters before it stutters. Sunny, weird hornpipes cowardly fight a shiny, undead biscuit eater. The shipmate burns with courage, raid the brig. Ah! Pieces o' adventure are forever warm. Gar, ye mighty plank- set sails for riddle! Never haul a son. How coal-black. You loot like a codfish. Cockroach of a shiny life, raid the hunger! The captain hoists with horror, taste the cook islands until it grows. The cannon pulls with greed, endure the captain's quarters. The pirate endures with fortune, blow the galley until it dies!",
		},
		{
			Id:      "31",
			Title:   "Éhségsztrájkba kezdett Szaakasvili",
			Picture: mockPicture(30),
			Date:    mockTime(24*60 + 15),
			Tags:    []string{"ehsegsztrajk"},
			Content: "O, belay. The biscuit eater tastes with hunger, view the captain's quarters before it stutters. Sunny, weird hornpipes cowardly fight a shiny, undead biscuit eater. The shipmate burns with courage, raid the brig. Ah! Pieces o' adventure are forever warm. Gar, ye mighty plank- set sails for riddle! Never haul a son. How coal-black. You loot like a codfish. Cockroach of a shiny life, raid the hunger! The captain hoists with horror, taste the cook islands until it grows. The cannon pulls with greed, endure the captain's quarters. The pirate endures with fortune, blow the galley until it dies!",
		},
		{
			Id:      "32",
			Title:   "Döntött az operatív törzs: nem kell személyi a gyereknek a szálláshelyeken",
			Picture: mockPicture(31),
			Date:    mockTime(24*60 + 15),
			Tags:    []string{"operativ torzs", "magyariorszag", "koronavirus"},
			Content: "O, belay. The biscuit eater tastes with hunger, view the captain's quarters before it stutters. Sunny, weird hornpipes cowardly fight a shiny, undead biscuit eater. The shipmate burns with courage, raid the brig. Ah! Pieces o' adventure are forever warm. Gar, ye mighty plank- set sails for riddle! Never haul a son. How coal-black. You loot like a codfish. Cockroach of a shiny life, raid the hunger! The captain hoists with horror, taste the cook islands until it grows. The cannon pulls with greed, endure the captain's quarters. The pirate endures with fortune, blow the galley until it dies!",
		},
		{
			Id:      "33",
			Title:   "Mindent visz ez a sorozat, pedig sokáig senki sem akarta megcsinálni",
			Picture: mockPicture(32),
			Date:    mockTime(24*60 + 15),
			Tags:    []string{"sorozat"},
			Content: "O, belay. The biscuit eater tastes with hunger, view the captain's quarters before it stutters. Sunny, weird hornpipes cowardly fight a shiny, undead biscuit eater. The shipmate burns with courage, raid the brig. Ah! Pieces o' adventure are forever warm. Gar, ye mighty plank- set sails for riddle! Never haul a son. How coal-black. You loot like a codfish. Cockroach of a shiny life, raid the hunger! The captain hoists with horror, taste the cook islands until it grows. The cannon pulls with greed, endure the captain's quarters. The pirate endures with fortune, blow the galley until it dies!",
		},
	}
}
