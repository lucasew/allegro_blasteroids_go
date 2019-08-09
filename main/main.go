package main

import (
    "github.com/dradtke/go-allegro/allegro"
    "github.com/dradtke/go-allegro/allegro/primitives"
    "github.com/lucasew/golog"
    "github.com/lucasew/allegro_blasteroids_go/list"
    "github.com/lucasew/allegro_blasteroids_go/elements"
    "github.com/lucasew/allegro_blasteroids_go/point"
    "time"
    "os"
)

var log = golog.Default.NewLogger("main")
var l = list.NewList()
var pts = 0
var asteroidCreator = make(chan(interface{}))
const asteroidSpawnInterval = time.Second*10
var DISPLAY_WIDTH int
var DISPLAY_HEIGHT int

var spaceship = &elements.Spaceship{
    Health: 100,
    Speed: 20,
    Position: point.NewHeadedPoint(40, 40, 1),
}

func main() {
    err := primitives.Install()
    if err != nil {
        log.Panic(err.Error())
    }
    allegro.Run(App)
}

var stop = false

func App() {
    d, err := allegro.CreateDisplay(800, 600)
    if err != nil {
        log.Panic(err.Error())
    }
    DISPLAY_WIDTH = d.Width()
    DISPLAY_HEIGHT = d.Height()
    queue, err := allegro.CreateEventQueue()
    if err != nil {
        log.Panic(err.Error())
    }
    defer queue.Destroy()
    queue.Register(d)
    defer d.Destroy()
    hearthbeat := time.Tick(time.Millisecond*(1000/30))
    go handleEvents(queue)
    go asteroidSpawner()
    for !stop {
        log.Info("tick %d (%d) HP: %d", pts, l.Len(), spaceship.GetLife())
        <-hearthbeat
        pts += l.ProcessCollisions(spaceship)
        if (spaceship.IsDead()) {
            break
        }
        if (l.GC() > 0) {
            l.Push(elements.NewRandomAsteroid(DISPLAY_WIDTH, DISPLAY_HEIGHT))
        }
        l.TickAll(1.0/30, DISPLAY_WIDTH, DISPLAY_HEIGHT)
        spaceship.Tick(0, DISPLAY_WIDTH, DISPLAY_HEIGHT)
        allegro.ClearToColor(allegro.MapRGB(0, 0, 0))
        spaceship.Draw()
        l.DrawAll()
        allegro.FlipDisplay()
    }
    log.Info("GAME OVER!")
    log.Info("Você fez %d pontos!", pts)
    os.Exit(0)
}

func asteroidSpawner() {
        timer := time.Tick(asteroidSpawnInterval)
        for !stop {
            <-timer
            l.Push(elements.NewRandomAsteroid(DISPLAY_WIDTH, DISPLAY_HEIGHT))
        }
}

func handleEvents(q *allegro.EventQueue) {
    llog := log.NewLogger("handleEvents")
    err := allegro.InstallKeyboard()
    if err != nil {
        log.Panic(err.Error())
    }
    keyboard, err := allegro.KeyboardEventSource()
    if err != nil {
        log.Panic(err.Error())
    }
    defer keyboard.DestroyUserEventSource()
    keyboard.InitUserEventSource()
    q.RegisterEventSource(keyboard)
    var ev allegro.Event
    defer os.Exit(0)
    for !stop {
        v := q.WaitForEvent(&ev)
        switch e := v.(type) {
            case allegro.KeyDownEvent:
                switch (e.KeyCode()) {
                    case allegro.KEY_LEFT: 
                    llog.Info("LEFT %s", spaceship.ToString())
                    spaceship.TurnLeft()
                case allegro.KEY_RIGHT:
                    llog.Info("RIGHT %s", spaceship.ToString())
                    spaceship.TurnRight()
                case allegro.KEY_UP:
                    llog.Info("UP %s", spaceship.ToString())
                    spaceship.MoveAhead()
                case allegro.KEY_DOWN:
                    llog.Info("DOWN %s", spaceship.ToString())
                    spaceship.MoveReverse()
                case allegro.KEY_SPACE:
                    llog.Info("SPACE %s", spaceship.ToString())
                    l.Push(spaceship.NewBullet())
                case allegro.KEY_ESCAPE:
                    stop = true
                case allegro.KEY_S:
                    l.Push(elements.NewRandomAsteroid(DISPLAY_WIDTH, DISPLAY_HEIGHT))
            }
            case allegro.DisplayCloseEvent:
                stop = true
            case allegro.DisplayResizeEvent:
                e.Source().Resize(e.Width(), e.Height())
                DISPLAY_WIDTH = e.Width()
                DISPLAY_HEIGHT = e.Height()
            }
        }
    }
