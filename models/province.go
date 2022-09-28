package models

import (
	"gorm.io/gorm"
)

type Province struct {
	gorm.Model
	Name        string     `gorm:"type:varchar(50);unique" json:"nome"`  // unique, indexes, keys, etc cannot be to long
	Area        float32    `gorm:"type:decimal(12, 5)" json:"area"`      // Km^2
	Population  uint32     `json:"populacao"`                            // hab
	Density     float32    `gorm:"type:decimal(12, 5)" json:"densidade"` // hab/Km^2
	Description string     `gorm:"type:text" json:"descricao"`
	Elevation   float32    `json:"altitude"`                    // m
	Coordenate  Coordenate `gorm:"embedded" json:"coordenadas"` // (lat: º, lon: º)
}

type APIProvince struct {
	ID          uint
	Name        string     `json:"nome"`
	Area        float32    `json:"area"`
	Population  uint32     `json:"populacao"`
	Density     float32    `json:"densidade"`
	Description string     `json:"descricao"`
	Elevation   float32    `json:"altitude"`
	Coordenate  Coordenate `gorm:"embedded" json:"coordenadas"`
}

func IntialData(db *gorm.DB) {
	var provinces = []Province{
		{
			Name:       "Luanda",
			Area:       18_826,
			Population: 8_330_000,
			Density:    377,
			Description: "Luanda é uma das 18 províncias de Angola, localizada na região centro-norte do país. Tem como capital a cidade e município de Luanda, a capital nacional.\n" +
				"É a mais rica e desenvolvida província da nação, sede de grandes conglomerados industriais, comerciais e de serviços, sendo também aquela que dispõe de mais recursos de infraestrutura.",
			Elevation:  6,
			Coordenate: Coordenate{Latitude: -8.838333, Longitude: 13.234444},
		},
		{
			Name:       "Cabinda",
			Area:       7_283,
			Population: 801_374,
			Density:    20.6,
			Description: "Cabinda é uma das 18 províncias de Angola, localizada na região norte do país, sendo a mais setentrional e também único exclave da nação. A capital é a cidade e município de Cabinda.\n" +
				"Administrativamente, a província é constituída pelos municípios de Cabinda, Cacongo (anteriormente Landana), Buco-Zau e Belize.",
			Elevation:  24,
			Coordenate: Coordenate{Latitude: -5.5500000, Longitude: 12.2000000},
		},
		{
			Name:       "Benguela",
			Area:       39_827,
			Population: 2_477_595,
			// Density:    nil,
			Description: "Benguela é uma das 18 províncias de Angola, localizada na região central do país. Tem como capital a cidade e município de Benguela.\n" +
				"A província é constituída pelos seguintes municípios: Baía Farta, Balombo, Benguela, Bocoio, Caimbambo, Catumbela, Chongoroi, Cubal, Ganda e Lobito.",
			Elevation:  11,
			Coordenate: Coordenate{Latitude: -12.5762600, Longitude: 13.4054700},
		},
		{
			Name:       "Kwanza-Norte",
			Area:       24_110,
			Population: 495_810,
			// Density:    nil,
			Description: "Cuanza Norte é uma das 18 províncias de Angola, localizada na região centro-norte do país. Sua capital está na cidade de Nadalatando, no município de Cazengo.\n" +
				"É constituída de 10 municípios, sendo os de Ambaca, Banga, Bolongongo, Cambambe, Cazengo, Golungo Alto, Gonguembo, Lucala, Quiculungo e Samba Caju.",
			Elevation:  1500,
			Coordenate: Coordenate{Latitude: -9.2399, Longitude: 14.6588},
		},
		{
			Name:       "Kwanza-Sul",
			Area:       55_660,
			Population: 2_109_999,
			// Density:    nil,
			Description: "Cuanza Sul é uma das 18 províncias de Angola, localizada na região central do país. Sua capital está na cidade e município de Sumbe.\n" +
				"A província é dividida administrativamente em 12 municípios e 32 comunas, sendo constituída pelos municípios de Amboim, Cassongue, Cela, Conda, Ebo, Libolo, Mussende, Porto Amboim, Quilenda, Quibala, Seles e Sumbe.",
			// Elevation:  1500,
			Coordenate: Coordenate{Latitude: -10.82696, Longitude: 15.03197},
		},
	}
	db.Create(&provinces)
}

func GetProvince() Province {
	var province Province
	return province
}

func GetProvinces(db *gorm.DB) []APIProvince {
	var provinces []APIProvince
	_ = db.Model(&Province{}).Find(&provinces)
	return provinces
}
