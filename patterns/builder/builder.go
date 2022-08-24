package main

type PenthouseBuilder struct {
	house Penthouse
}

func (pb *PenthouseBuilder) BuildWall(s string) {
	pb.house.walls = append(pb.house.walls, Wall{s})
}

func (pb *PenthouseBuilder) BuildWindow(s string) {
	pb.house.windows = append(pb.house.windows, Window{s})
}

func (pb *PenthouseBuilder) BuildDoor(s string) {
	pb.house.doors = append(pb.house.doors, Door{s})
}

func (pb *PenthouseBuilder) BuildRoof(s string) {
	pb.house.roof = Roof{s}
}

func (pb *PenthouseBuilder) Build() *Penthouse {
	return &pb.house
}

type DormBuilder struct {
	house Dorm
}

func (db *DormBuilder) BuildWall(s string) {
	db.house.walls = append(db.house.walls, Wall{s})
}

func (db *DormBuilder) BuildWindow(s string) {
	db.house.windows = append(db.house.windows, Window{s})
}

func (db *DormBuilder) BuildDoor(s string) {
	db.house.doors = append(db.house.doors, Door{s})
}

func (db *DormBuilder) BuildRoof(s string) {
	db.house.roof = Roof{s}
}

func (db *DormBuilder) Build() *Dorm {
	return &db.house
}
