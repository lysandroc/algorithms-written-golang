//Space has some methods to calculate the orbital age using earth as a paremeter
package space

type Planet string

const earthYearInSecs float64 = 31557600

const (
	orbitalOnEarth   float64 = 1.0
	orbitalOnMercury         = 0.2408467
	orbitalOnVenus           = 0.61519726
	orbitalOnMars            = 1.8808158
	orbitalOnJupiter         = 11.862615
	orbitalOnSaturn          = 29.447498
	orbitalOnUranus          = 84.016846
	orbitalOnNeptune         = 164.79132
)

const (
	Earth   Planet = "Earth"
	Mercury Planet = "Mercury"
	Venus   Planet = "Venus"
	Mars    Planet = "Mars"
	Jupiter Planet = "Jupiter"
	Saturn  Planet = "Saturn"
	Uranus  Planet = "Uranus"
	Neptune Planet = "Neptune"
)

// type OrbitalPlanetKnow struct {
// 	Planet        Planet
// 	OrbitalPeriod float64
// }
//
// Earth := OrbitalPlanetKnow {"Earth", 1}
// Mercury := OrbitalPlanetKnow {"Mercury",0.2408467}
// Venus := OrbitalPlanetKnow {"Venus",0.61519726}
// Mars := OrbitalPlanetKnow {"Mars",1.8808158}
// Jupiter := OrbitalPlanetKnow {"Jupiter",11.862615}
// Saturn := OrbitalPlanetKnow {"Saturn",29.447498}
// Uranus := OrbitalPlanetKnow {"Uranus",84.016846}
// Neptune := OrbitalPlanetKnow {"Neptune",164.79132}

type Orbital interface {
	calculateOrbitalPeriod() float64
	calculateAge() float64
}

type orbitalPlanet struct {
	planet Planet
}

func (o orbitalPlanet) calculateOrbitalPeriod() float64 {
	var orbitalPeriodInSecs float64
	switch o.planet {
	case Earth:
		orbitalPeriodInSecs = orbitalOnEarth
	case Mercury:
		orbitalPeriodInSecs = orbitalOnMercury
	case Venus:
		orbitalPeriodInSecs = orbitalOnVenus
	case Mars:
		orbitalPeriodInSecs = orbitalOnMars
	case Jupiter:
		orbitalPeriodInSecs = orbitalOnJupiter
	case Saturn:
		orbitalPeriodInSecs = orbitalOnSaturn
	case Uranus:
		orbitalPeriodInSecs = orbitalOnUranus
	case Neptune:
		orbitalPeriodInSecs = orbitalOnNeptune
	}

	return orbitalPeriodInSecs
}

func (o orbitalPlanet) calculateAge(seconds float64, orbitalPeriodInSecs float64) float64 {
	return seconds / earthYearInSecs / orbitalPeriodInSecs
}

//Age calculates the age on the orbita given the seconds and the planet
func Age(seconds float64, planet Planet) float64 {
	p := orbitalPlanet{planet: planet}
	orbitalPeriodInSecs := p.calculateOrbitalPeriod()

	return p.calculateAge(seconds, orbitalPeriodInSecs)
}
