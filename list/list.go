package list

import (
    "github.com/lucasew/allegro_blasteroids_go/interfaces"
    "github.com/lucasew/allegro_blasteroids_go/collision"
    "sync"
)

// NewList cria uma lista vazia
func NewList() List {
    i := 0
    return List{
        v: map[int]interfaces.ListNodePayload{},
        i: &i,
        mutex: &sync.Mutex{},
    }
}

// List lista responsavel por guardar as entidades
type List struct {
    v map[int]interfaces.ListNodePayload
    i *int
    mutex *sync.Mutex
}

// Push adiciona mais um valor na lista
func (l List) Push(what interfaces.ListNodePayload) int {
    i := *l.i
    l.mutex.Lock()
    l.v[i] = what
    *l.i++
    l.mutex.Unlock()
    println(i)
    return i
}

// Len retorna o tamanho da lista
func (l List) Len() int {
    return len(l.v)
}

// GC busca e remove os valores que já "morreram" e retorna quantos
func (l List) GC() int {
    removed := 0
    l.mutex.Lock()
    for i := range l.v {
        if l.v[i].IsDead() {
            l.v[i].Die()
            delete(l.v, i)
            removed++
        }
    }
    l.mutex.Unlock()
    return removed
}

func (l List) ProcessCollisions(extraEntities ...interfaces.Collisionable) int {
    pts := 0
    l.mutex.Lock()
    for i := range extraEntities {
        for j := range l.v {
            pts += collision.CheckCollision(extraEntities[i], l.v[j])
        }
    }
    for i := range l.v {
        for j := range l.v {
            if (i != j) {
                pts += collision.CheckCollision(l.v[i], l.v[j])
            }
        }
    }
    l.mutex.Unlock()
    return pts
}

func (l List) DrawAll() {
    l.mutex.Lock()
    for _, e := range l.v {
        e.Draw()
    }
    l.mutex.Unlock()
}

func (l List) TickAll(tick float32, w, h int) {
    l.mutex.Lock()
    for _, e := range l.v {
        e.Tick(tick, w, h)
    }
    l.mutex.Unlock()
}
