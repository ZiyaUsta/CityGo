package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// CityModel class...
type CityModel struct {
	XMLName        xml.Name `xml:"CityModel"`
	Text           string   `xml:",chardata"`
	Xmlns          string   `xml:"xmlns,attr"`
	Core           string   `xml:"core,attr"`
	Gml            string   `xml:"gml,attr"`
	Grp            string   `xml:"grp,attr"`
	Gen            string   `xml:"gen,attr"`
	Bldg           string   `xml:"bldg,attr"`
	App            string   `xml:"app,attr"`
	Xlink          string   `xml:"xlink,attr"`
	Xsi            string   `xml:"xsi,attr"`
	SchemaLocation string   `xml:"schemaLocation,attr"`
	BoundedBy      struct {
		Text     string `xml:",chardata"`
		Envelope struct {
			Text         string `xml:",chardata"`
			SrsName      string `xml:"srsName,attr"`
			SrsDimension string `xml:"srsDimension,attr"`
			LowerCorner  string `xml:"lowerCorner"`
			UpperCorner  string `xml:"upperCorner"`
		} `xml:"Envelope"`
	} `xml:"boundedBy"`
	CityObjectMember []struct {
		Text     string `xml:",chardata"`
		Building struct {
			Text            string `xml:",chardata"`
			ID              string `xml:"id,attr"`
			Name            string `xml:"name"`
			StringAttribute []struct {
				Text  string `xml:",chardata"`
				Name  string `xml:"name,attr"`
				Value string `xml:"value"`
			} `xml:"stringAttribute"`
			IntAttribute []struct {
				Text  string `xml:",chardata"`
				Name  string `xml:"name,attr"`
				Value string `xml:"value"`
			} `xml:"intAttribute"`
			DoubleAttribute []struct {
				Text  string `xml:",chardata"`
				Name  string `xml:"name,attr"`
				Value string `xml:"value"`
			} `xml:"doubleAttribute"`
			DateAttribute struct {
				Text  string `xml:",chardata"`
				Name  string `xml:"name,attr"`
				Value string `xml:"value"`
			} `xml:"dateAttribute"`
			Class              string `xml:"class"`
			StoreysAboveGround string `xml:"storeysAboveGround"`
			StoreysBelowGround string `xml:"storeysBelowGround"`
			Lod2MultiSurface   struct {
				Text         string `xml:",chardata"`
				MultiSurface struct {
					Text          string `xml:",chardata"`
					SurfaceMember []struct {
						Text string `xml:",chardata"`
						Href string `xml:"href,attr"`
					} `xml:"surfaceMember"`
				} `xml:"MultiSurface"`
			} `xml:"lod2MultiSurface"`
			Lod2TerrainIntersection struct {
				Text       string `xml:",chardata"`
				MultiCurve struct {
					Text        string `xml:",chardata"`
					CurveMember struct {
						Text       string `xml:",chardata"`
						LineString struct {
							Text    string `xml:",chardata"`
							PosList string `xml:"posList"`
						} `xml:"LineString"`
					} `xml:"curveMember"`
				} `xml:"MultiCurve"`
			} `xml:"lod2TerrainIntersection"`
			OuterBuildingInstallation []struct {
				Text                 string `xml:",chardata"`
				BuildingInstallation struct {
					Text            string `xml:",chardata"`
					ID              string `xml:"id,attr"`
					Name            string `xml:"name"`
					StringAttribute []struct {
						Text  string `xml:",chardata"`
						Name  string `xml:"name,attr"`
						Value string `xml:"value"`
					} `xml:"stringAttribute"`
					IntAttribute struct {
						Text  string `xml:",chardata"`
						Name  string `xml:"name,attr"`
						Value string `xml:"value"`
					} `xml:"intAttribute"`
					DoubleAttribute []struct {
						Text  string `xml:",chardata"`
						Name  string `xml:"name,attr"`
						Value string `xml:"value"`
					} `xml:"doubleAttribute"`
					Class        string `xml:"class"`
					Lod2Geometry struct {
						Text         string `xml:",chardata"`
						MultiSurface struct {
							Text          string `xml:",chardata"`
							SrsName       string `xml:"srsName,attr"`
							SurfaceMember []struct {
								Text string `xml:",chardata"`
								Href string `xml:"href,attr"`
							} `xml:"surfaceMember"`
						} `xml:"MultiSurface"`
					} `xml:"lod2Geometry"`
					BoundedBy []struct {
						Text         string `xml:",chardata"`
						FloorSurface struct {
							Text             string `xml:",chardata"`
							Lod2MultiSurface struct {
								Text         string `xml:",chardata"`
								MultiSurface struct {
									Text          string `xml:",chardata"`
									SrsName       string `xml:"srsName,attr"`
									SurfaceMember []struct {
										Text    string `xml:",chardata"`
										Polygon struct {
											Text     string `xml:",chardata"`
											ID       string `xml:"id,attr"`
											Exterior struct {
												Text       string `xml:",chardata"`
												LinearRing struct {
													Text    string `xml:",chardata"`
													PosList string `xml:"posList"`
												} `xml:"LinearRing"`
											} `xml:"exterior"`
										} `xml:"Polygon"`
									} `xml:"surfaceMember"`
								} `xml:"MultiSurface"`
							} `xml:"lod2MultiSurface"`
						} `xml:"FloorSurface"`
						WallSurface struct {
							Text             string `xml:",chardata"`
							Lod2MultiSurface struct {
								Text         string `xml:",chardata"`
								MultiSurface struct {
									Text          string `xml:",chardata"`
									SrsName       string `xml:"srsName,attr"`
									SurfaceMember []struct {
										Text    string `xml:",chardata"`
										Polygon struct {
											Text     string `xml:",chardata"`
											ID       string `xml:"id,attr"`
											Exterior struct {
												Text       string `xml:",chardata"`
												LinearRing struct {
													Text    string `xml:",chardata"`
													PosList string `xml:"posList"`
												} `xml:"LinearRing"`
											} `xml:"exterior"`
										} `xml:"Polygon"`
									} `xml:"surfaceMember"`
								} `xml:"MultiSurface"`
							} `xml:"lod2MultiSurface"`
						} `xml:"WallSurface"`
					} `xml:"boundedBy"`
				} `xml:"BuildingInstallation"`
			} `xml:"outerBuildingInstallation"`
			BoundedBy []struct {
				Text        string `xml:",chardata"`
				RoofSurface struct {
					Text             string `xml:",chardata"`
					Lod2MultiSurface struct {
						Text         string `xml:",chardata"`
						MultiSurface struct {
							Text          string `xml:",chardata"`
							SrsName       string `xml:"srsName,attr"`
							SurfaceMember struct {
								Text    string `xml:",chardata"`
								Polygon struct {
									Text     string `xml:",chardata"`
									ID       string `xml:"id,attr"`
									Exterior struct {
										Text       string `xml:",chardata"`
										LinearRing struct {
											Text    string `xml:",chardata"`
											PosList string `xml:"posList"`
										} `xml:"LinearRing"`
									} `xml:"exterior"`
								} `xml:"Polygon"`
							} `xml:"surfaceMember"`
						} `xml:"MultiSurface"`
					} `xml:"lod2MultiSurface"`
				} `xml:"RoofSurface"`
				WallSurface struct {
					Text             string `xml:",chardata"`
					Lod2MultiSurface struct {
						Text         string `xml:",chardata"`
						MultiSurface struct {
							Text          string `xml:",chardata"`
							SrsName       string `xml:"srsName,attr"`
							SurfaceMember []struct {
								Text    string `xml:",chardata"`
								Polygon struct {
									Text     string `xml:",chardata"`
									ID       string `xml:"id,attr"`
									Exterior struct {
										Text       string `xml:",chardata"`
										LinearRing struct {
											Text    string `xml:",chardata"`
											PosList string `xml:"posList"`
										} `xml:"LinearRing"`
									} `xml:"exterior"`
								} `xml:"Polygon"`
							} `xml:"surfaceMember"`
						} `xml:"MultiSurface"`
					} `xml:"lod2MultiSurface"`
				} `xml:"WallSurface"`
			} `xml:"boundedBy"`
			InteriorRoom []struct {
				Text string `xml:",chardata"`
				Room struct {
					Text            string `xml:",chardata"`
					ID              string `xml:"id,attr"`
					Name            string `xml:"name"`
					StringAttribute []struct {
						Text  string `xml:",chardata"`
						Name  string `xml:"name,attr"`
						Value string `xml:"value"`
					} `xml:"stringAttribute"`
					IntAttribute []struct {
						Text  string `xml:",chardata"`
						Name  string `xml:"name,attr"`
						Value string `xml:"value"`
					} `xml:"intAttribute"`
					DoubleAttribute []struct {
						Text  string `xml:",chardata"`
						Name  string `xml:"name,attr"`
						Value string `xml:"value"`
					} `xml:"doubleAttribute"`
					Class            string `xml:"class"`
					Lod4MultiSurface struct {
						Text         string `xml:",chardata"`
						MultiSurface struct {
							Text          string `xml:",chardata"`
							SrsName       string `xml:"srsName,attr"`
							SurfaceMember []struct {
								Text string `xml:",chardata"`
								Href string `xml:"href,attr"`
							} `xml:"surfaceMember"`
						} `xml:"MultiSurface"`
					} `xml:"lod4MultiSurface"`
					BoundedBy []struct {
						Text         string `xml:",chardata"`
						FloorSurface struct {
							Text             string `xml:",chardata"`
							Lod2MultiSurface struct {
								Text         string `xml:",chardata"`
								MultiSurface struct {
									Text          string `xml:",chardata"`
									SrsName       string `xml:"srsName,attr"`
									SurfaceMember []struct {
										Text    string `xml:",chardata"`
										Polygon struct {
											Text     string `xml:",chardata"`
											ID       string `xml:"id,attr"`
											Exterior struct {
												Text       string `xml:",chardata"`
												LinearRing struct {
													Text    string `xml:",chardata"`
													PosList string `xml:"posList"`
												} `xml:"LinearRing"`
											} `xml:"exterior"`
										} `xml:"Polygon"`
									} `xml:"surfaceMember"`
								} `xml:"MultiSurface"`
							} `xml:"lod2MultiSurface"`
						} `xml:"FloorSurface"`
						InteriorWallSurface struct {
							Text             string `xml:",chardata"`
							Lod2MultiSurface struct {
								Text         string `xml:",chardata"`
								MultiSurface struct {
									Text          string `xml:",chardata"`
									SrsName       string `xml:"srsName,attr"`
									SurfaceMember []struct {
										Text    string `xml:",chardata"`
										Polygon struct {
											Text     string `xml:",chardata"`
											ID       string `xml:"id,attr"`
											Exterior struct {
												Text       string `xml:",chardata"`
												LinearRing struct {
													Text    string `xml:",chardata"`
													PosList string `xml:"posList"`
												} `xml:"LinearRing"`
											} `xml:"exterior"`
										} `xml:"Polygon"`
									} `xml:"surfaceMember"`
								} `xml:"MultiSurface"`
							} `xml:"lod2MultiSurface"`
						} `xml:"InteriorWallSurface"`
						CeilingSurface struct {
							Text             string `xml:",chardata"`
							Lod2MultiSurface struct {
								Text         string `xml:",chardata"`
								MultiSurface struct {
									Text          string `xml:",chardata"`
									SrsName       string `xml:"srsName,attr"`
									SurfaceMember []struct {
										Text    string `xml:",chardata"`
										Polygon struct {
											Text     string `xml:",chardata"`
											ID       string `xml:"id,attr"`
											Exterior struct {
												Text       string `xml:",chardata"`
												LinearRing struct {
													Text    string `xml:",chardata"`
													PosList string `xml:"posList"`
												} `xml:"LinearRing"`
											} `xml:"exterior"`
										} `xml:"Polygon"`
									} `xml:"surfaceMember"`
								} `xml:"MultiSurface"`
							} `xml:"lod2MultiSurface"`
						} `xml:"CeilingSurface"`
					} `xml:"boundedBy"`
				} `xml:"Room"`
			} `xml:"interiorRoom"`
		} `xml:"Building"`
		CityObjectGroup struct {
			Text            string `xml:",chardata"`
			ID              string `xml:"id,attr"`
			Name            string `xml:"name"`
			StringAttribute struct {
				Text  string `xml:",chardata"`
				Name  string `xml:"name,attr"`
				Value string `xml:"value"`
			} `xml:"stringAttribute"`
			IntAttribute []struct {
				Text  string `xml:",chardata"`
				Name  string `xml:"name,attr"`
				Value string `xml:"value"`
			} `xml:"intAttribute"`
			Class       string `xml:"class"`
			GroupMember []struct {
				Text string `xml:",chardata"`
				Href string `xml:"href,attr"`
			} `xml:"groupMember"`
			Parent struct {
				Text string `xml:",chardata"`
				Href string `xml:"href,attr"`
			} `xml:"parent"`
			Geometry struct {
				Text         string `xml:",chardata"`
				MultiSurface struct {
					Text          string `xml:",chardata"`
					SrsName       string `xml:"srsName,attr"`
					SurfaceMember []struct {
						Text    string `xml:",chardata"`
						Polygon struct {
							Text     string `xml:",chardata"`
							ID       string `xml:"id,attr"`
							Exterior struct {
								Text       string `xml:",chardata"`
								LinearRing struct {
									Text    string `xml:",chardata"`
									PosList string `xml:"posList"`
								} `xml:"LinearRing"`
							} `xml:"exterior"`
						} `xml:"Polygon"`
					} `xml:"surfaceMember"`
				} `xml:"MultiSurface"`
			} `xml:"geometry"`
		} `xml:"CityObjectGroup"`
		GenericCityObject struct {
			Text            string `xml:",chardata"`
			ID              string `xml:"id,attr"`
			Name            string `xml:"name"`
			StringAttribute []struct {
				Text  string `xml:",chardata"`
				Name  string `xml:"name,attr"`
				Value string `xml:"value"`
			} `xml:"stringAttribute"`
			IntAttribute []struct {
				Text  string `xml:",chardata"`
				Name  string `xml:"name,attr"`
				Value string `xml:"value"`
			} `xml:"intAttribute"`
			DoubleAttribute []struct {
				Text  string `xml:",chardata"`
				Name  string `xml:"name,attr"`
				Value string `xml:"value"`
			} `xml:"doubleAttribute"`
			Class        string `xml:"class"`
			Lod2Geometry struct {
				Text         string `xml:",chardata"`
				MultiSurface struct {
					Text          string `xml:",chardata"`
					SrsName       string `xml:"srsName,attr"`
					SurfaceMember []struct {
						Text    string `xml:",chardata"`
						Polygon struct {
							Text     string `xml:",chardata"`
							ID       string `xml:"id,attr"`
							Exterior struct {
								Text       string `xml:",chardata"`
								LinearRing struct {
									Text    string `xml:",chardata"`
									PosList string `xml:"posList"`
								} `xml:"LinearRing"`
							} `xml:"exterior"`
							Interior struct {
								Text       string `xml:",chardata"`
								LinearRing struct {
									Text    string `xml:",chardata"`
									PosList string `xml:"posList"`
								} `xml:"LinearRing"`
							} `xml:"interior"`
						} `xml:"Polygon"`
					} `xml:"surfaceMember"`
				} `xml:"MultiSurface"`
			} `xml:"lod2Geometry"`
		} `xml:"GenericCityObject"`
	} `xml:"cityObjectMember"`
}

func main() {
	/*	v0 := Vertex{}
		v1 := Vertex{}
		v2 := Vertex{}

		t := Triangle{v0, v1, v2}
		s := Sphere{v0, 3.9}
		fmt.Println(t.V0)
		fmt.Println(s.radius)
	*/
	xmlFile, err := os.Open("src/binaTKGM.gml")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened CityGML file")

	defer xmlFile.Close()

	byteValue, _ := ioutil.ReadAll(xmlFile)

	var cityModel CityModel

	xml.Unmarshal(byteValue, &cityModel)

	//fmt.Println(cityModel.CityObjectMember[0].Building.BoundedBy[0].RoofSurface.Lod2MultiSurface.MultiSurface.SurfaceMember.Polygon.Exterior.LinearRing.PosList)

	bvertices := strings.Fields(cityModel.CityObjectMember[0].Building.BoundedBy[0].RoofSurface.Lod2MultiSurface.MultiSurface.SurfaceMember.Polygon.Exterior.LinearRing.PosList)
	fmt.Println(bvertices[0])
	a := len(bvertices) / 3
	yler := make([]string, a)

	for i := 1; i < len(bvertices); i = i + 2 {

		yler = append(yler, bvertices[i])
	}

	fmt.Println(len(bvertices) / 3)
}
