package main

import "fmt"

/*
Шаблон проектирования, позволяющий скрыть сложность иерархии, сводя все вызовы к одному объекту,
делегирующему их соответствующим объектам системы.
Шаблон применяется для установки политики по отношению к другой группе объектов. Фасад подойдет, если политика должна быть яркой и заметной.

Плюсы:
1. Упрощение работы с системой.

Минусы:
1. Нужно хорошо продумать реализуемый набор интерфейсов для клиента.
2. Фасад может привязаться ко всем классам программы.
*/

func main() {
	var orchestra Orchestra
	orchestra.PlayMusic()
}

type Orchestra struct {
	t Trumpet
	d Drums
}

func (d *Orchestra) PlayMusic() {
	fmt.Println("Orchestra plays music:")
	d.t.DoTheTuTu()
	d.d.DoTheBomBom()
}

type Trumpet struct {
}

func (d *Trumpet) DoTheTuTu() {
	fmt.Println("Tu-Tu")
}

type Drums struct {
}

func (d *Drums) DoTheBomBom() {
	fmt.Println("Bom-Bom")
}
