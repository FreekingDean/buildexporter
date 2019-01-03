package main

import (
	"fmt"
	"strings"
	"time"

	build "github.com/FreekingDean/buildengine"
	collada "github.com/FreekingDean/go-collada"
)

func newCollada() *collada.Collada {
	return &collada.Collada{
		Xmlns:   collada.Uri("http://www.collada.org/2005/11/COLLADASchema"),
		Version: collada.Version("1.4.1"),
		HasAsset: collada.HasAsset{Asset: &collada.Asset{
			Contributor: []*collada.Contributor{
				{Author: "DeanGalvin", AuthoringTool: "buildexport"},
			},
			Created:  time.Now().Format(time.RFC3339),
			Modified: time.Now().Format(time.RFC3339),
			Unit:     &collada.Unit{HasName: collada.HasName{"meter"}, Meter: 0.01},
			UpAxis:   collada.UpAxis("Z_UP"),
		}},
		LibraryEffects: []*collada.LibraryEffects{
			{
				Effect: []*collada.Effect{
					{
						HasId: collada.HasId{collada.Id("material-effect")},
						ProfileCommon: &collada.ProfileCommon{
							HasTechniqueFx: collada.HasTechniqueFx{
								TechniqueFx: []*collada.TechniqueFx{
									{
										HasSid: collada.HasSid{"common"},
										Phong: &collada.Phong{
											Emission: &collada.FxCommonColorOrTextureType{
												Color: &collada.Color{
													HasSid: collada.HasSid{"emission"},
													Float3: collada.Float3{collada.Floats{Values: collada.Values{"0 0 0 1"}}},
												},
											},
											AmbientFx: &collada.FxCommonColorOrTextureType{
												Color: &collada.Color{
													HasSid: collada.HasSid{"Ambient"},
													Float3: collada.Float3{collada.Floats{Values: collada.Values{"0 0 0 1"}}},
												},
											},
											Diffuse: &collada.FxCommonColorOrTextureType{
												Color: &collada.Color{
													HasSid: collada.HasSid{"diffuse"},
													Float3: collada.Float3{collada.Floats{Values: collada.Values{"0.64 0.64 0.64 1"}}},
												},
											},
											Specular: &collada.FxCommonColorOrTextureType{
												Color: &collada.Color{
													HasSid: collada.HasSid{"specular"},
													Float3: collada.Float3{collada.Floats{Values: collada.Values{"0.5 0.5 0.5 1"}}},
												},
											},
											Shininess: &collada.FxCommonFloatOrParamType{
												Float: &collada.Float{
													HasSid: collada.HasSid{"shininess"},
													Value:  50,
												},
											},
											IndexOfRefraction: &collada.FxCommonFloatOrParamType{
												Float: &collada.Float{
													HasSid: collada.HasSid{"index_of_refraction"},
													Value:  1,
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		LibraryMaterials: []*collada.LibraryMaterials{
			{
				Material: []*collada.Material{
					{
						HasId: collada.HasId{collada.Id("material")},
						InstanceEffect: collada.InstanceEffect{
							HasUrl: collada.HasUrl{collada.Uri("#material-effect")},
						},
					},
				},
			},
		},
		LibraryVisualScenes: []*collada.LibraryVisualScenes{
			{
				VisualScene: []*collada.VisualScene{
					{
						HasId: collada.HasId{"scene"},
						HasNodes: collada.HasNodes{
							Node: []*collada.Node{
								{
									HasId:   collada.HasId{"node"},
									HasType: collada.HasType{"NODE"},
									Translate: []*collada.Translate{
										{
											HasSid: collada.HasSid{"location"},
											Float3: collada.Float3{collada.Floats{Values: collada.Values{"0 0 0"}}},
										},
									},
									Scale: []*collada.Scale{
										{
											HasSid: collada.HasSid{"scale"},
											Float3: collada.Float3{collada.Floats{Values: collada.Values{"1 1 1"}}},
										},
									},
									Rotate: []*collada.Rotate{
										{
											HasSid: collada.HasSid{"rotationZ"},
											Float4: collada.Float4{collada.Floats{Values: collada.Values{"0 0 0 0"}}},
										},
										{
											HasSid: collada.HasSid{"rotationY"},
											Float4: collada.Float4{collada.Floats{Values: collada.Values{"0 0 0 0"}}},
										},
										{
											HasSid: collada.HasSid{"rotationX"},
											Float4: collada.Float4{collada.Floats{Values: collada.Values{"0 0 0 0"}}},
										},
									},
									InstanceGeometry: []*collada.InstanceGeometry{
										{
											HasUrl: collada.HasUrl{"#SECTOR-0"},
											BindMaterial: &collada.BindMaterial{
												HasTechniqueCommon: collada.HasTechniqueCommon{
													collada.TechniqueCommon{
														InstanceMaterial: &collada.InstanceMaterial{
															Symbol: "material",
															Target: "#material",
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		Scene: &collada.Scene{
			InstanceVisualScene: &collada.InstanceVisualScene{
				HasUrl: collada.HasUrl{"#scene"},
			},
		},
	}
}

func sectorToColladaGeom(sectorID int, sector *build.Sector) *collada.LibraryGeometries {
	walls := sector.Walls()[0:2]

	geom := &collada.Geometry{}
	geom.Id = collada.Id(fmt.Sprintf("SECTOR-%d", sectorID))
	geom.Mesh = &collada.Mesh{}
	points := &collada.Source{}
	points.Id = collada.Id(fmt.Sprintf("SECTOR-%d-WALL-POSITIONS", sectorID))
	points.FloatArray = &collada.FloatArray{}
	points.FloatArray.Id = collada.Id(fmt.Sprintf("SECTOR-%d-POSITIONS", sectorID))
	points.FloatArray.Floats = collada.Floats{Values: collada.Values{V: ""}}
	points.TechniqueCommon = collada.TechniqueCommon{Accessor: &collada.Accessor{}}
	points.TechniqueCommon.Accessor.Count = len(walls) * 2
	points.TechniqueCommon.Accessor.Stride = 3
	points.TechniqueCommon.Accessor.Params = []*collada.ParamCore{
		{HasName: collada.HasName{"X"}, HasType: collada.HasType{"float"}},
		{HasName: collada.HasName{"Y"}, HasType: collada.HasType{"float"}},
		{HasName: collada.HasName{"Z"}, HasType: collada.HasType{"float"}},
	}
	geom.Mesh.Source = []*collada.Source{points}
	geom.Mesh.Vertices = collada.Vertices{
		HasId: collada.HasId{collada.Id("verts")},
		Input: []*collada.InputUnshared{
			{
				Semantic: "POSITION",
				Source:   collada.Uri(fmt.Sprintf("#SECTOR-%d-WALL-POSITIONS", sectorID)),
			},
		},
	}
	geom.Mesh.Polylist = []*collada.Polylist{
		{
			HasSharedInput: collada.HasSharedInput{[]*collada.InputShared{
				{
					Semantic: "VERTEX",
					Source:   collada.Uri("#verts"),
				},
			}},
			VCount: &collada.Ints{Values: collada.Values{V: ""}},
			HasP:   collada.HasP{P: &collada.P{collada.Ints{Values: collada.Values{V: ""}}}},
		},
	}
	baseP := []string{}
	baseV := []string{}
	ceilingP := []string{}
	floorP := []string{}
	c := len(walls)
	wall := walls[0]
	i := 0
	for c > 0 {
		addInt(points, wall.Left().X)
		addInt(points, wall.Left().Y)
		addInt(points, sector.CeilingZ())
		addInt(points, wall.Left().X)
		addInt(points, wall.Left().Y)
		addInt(points, sector.FloorZ())
		ceilingP = append(ceilingP, fmt.Sprintf("%d %d %d", i*6, i*6+1, i*6+2))
		floorP = append(floorP, fmt.Sprintf("%d %d %d", i*6+3, i*6+4, i*6+5))
		if i+1 < len(walls) {
			baseV = append(baseV, "4")
			for j := 0; j < 12; j++ {
				baseP = append(baseP, fmt.Sprintf("%d", (i*6)+j))
			}
		}
		wall = wall.Right()
		i++
		c--
	}
	baseV = append(baseV, fmt.Sprintf("%d", len(walls)))
	baseV = append(baseV, fmt.Sprintf("%d", len(walls)))
	baseP = append(baseP, ceilingP...)
	baseP = append(baseP, floorP...)
	geom.Mesh.Polylist[0].Count = 1
	geom.Mesh.Polylist[0].VCount.Values.V = strings.Join(baseV, " ")
	geom.Mesh.Polylist[0].P.Values.V = strings.Join(baseP, " ")
	return &collada.LibraryGeometries{
		Geometry: []*collada.Geometry{geom},
	}
}

func addInt(s *collada.Source, v int) {
	f := []string{}
	if s.FloatArray.Floats.Values.V != "" {
		f = strings.Split(s.FloatArray.Floats.Values.V, " ")
	}
	f = append(f, fmt.Sprintf("%d", v))
	s.FloatArray.Count = len(f)
	s.FloatArray.Floats.Values.V = strings.Join(f, " ")
}
