```mermaid
classDiagram
    direction LR

    class Match {
        +Team teamA
        +Team teamB
        +MatchTypeInterface matchType
        +List<Innings> innings
        +GameEventManager eventManager
        +startMatch()
    }

    class Team {
        +String name
        +List<Player> players
    }

    class Player {
        +String name
        +PlayerType type
    }

    class Innings {
        +Team battingTeam
        +Team bowlingTeam
        +BattingController battingController
        +BowlingController bowlingController
        +List<Over> overs
        +startInnings()
    }

    class Over {
        +int overNumber
        +List<Ball> balls
        +startOver()
    }

    class Ball {
        +Player batsman
        +Player bowler
        +RunType runType
        +BallType ballType
        +Wicket wicket
    }

    class GameEventManager {
        -List<GameEventObserver> observers
        +registerObserver(obs)
        +unregisterObserver(obs)
        +publish(ball)
    }

    class GameEventObserver {
        <<Interface>>
        +update(Ball ball)
    }

    class BattingScorecardObserver {
        -Map<Player, BattingScore> scores
        +update(Ball ball)
    }

    class BowlingScorecardObserver {
        -Map<Player, BowlingScore> scores
        +update(Ball ball)
    }

    class WicketObserver {
        -BattingController battingController
        +update(Ball ball)
    }

    class OverEndObserver {
        -BattingController battingController
        -BowlingController bowlingController
        +update(Ball ball)
    }

    class BattingController {
        -Queue<Player> battingQueue
        -Player striker
        -Player nonStriker
        +getNextBatsman()
        +swapStriker()
    }

    class BowlingController {
        -List<Player> bowlers
        -Player currentBowler
        +getNextBowler()
    }

    class MatchTypeInterface {
        <<Interface>>
        +getTotalOvers()
        +getMaxOversPerBowler()
    }
    class T20 {
      +getTotalOvers()
    }
    class ODI {
      +getTotalOvers()
    }

    %% --- Relationships ---
    Match "1" *-- "1" GameEventManager : creates
    Match "1" *-- "2" Team : has
    Match "1" *-- "1..2" Innings : has
    Match o-- MatchTypeInterface : uses (Strategy)

    Innings "1" *-- "1" BattingController : creates
    Innings "1" *-- "1" BowlingController : creates
    Innings "1" *-- "*" Over : has
    Innings ..> GameEventManager : uses (passes ref)

    Over "1" *-- "1..6" Ball : has
    %% --- MODIFIED RELATIONSHIP ---
    Over ..> GameEventManager : uses (to publish Ball)

    GameEventManager "1" o-- "*" GameEventObserver : notifies (Observer)
    GameEventObserver <|.. BattingScorecardObserver : implements
    GameEventObserver <|.. BowlingScorecardObserver : implements
    GameEventObserver <|.. WicketObserver : implements
    GameEventObserver <|.. OverEndObserver : implements

    WicketObserver ..> BattingController : commands
    OverEndObserver ..> BattingController : commands
    OverEndObserver ..> BowlingController : commands

    Innings ..> BattingScorecardObserver : creates & registers
    Innings ..> BowlingScorecardObserver : creates & registers
    Innings ..> WicketObserver : creates & registers
    Innings ..> OverEndObserver : creates & registers

    MatchTypeInterface <|.. T20 : implements
    MatchTypeInterface <|.. ODI : implements