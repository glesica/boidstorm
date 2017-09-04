[![Build Status](https://travis-ci.org/glesica/boidstorm.svg?branch=master)](https://travis-ci.org/glesica/boidstorm)

# Boid Storm

A game of emergent behavior.

## Back Story

In the future, the population has been so decimated by global war that the
remaining factions can no longer afford to risk human lives in battle. Instead,
sophisticated, semi-autonomous robots fight over territory and resources on
behalf of human masters.

The human commanders who control these legions of robotic terror do so from
afar through computerized interfaces. They make split-second decisions about
the number of robotic soldiers to deploy and how best to configure each group
they send into battle.

Poorly configured robots will fail to engage the enemy, fail to defend their
own territory, and ultimately lose battles.

The best commanders build libraries of strategies and configurations for rapid
use in future battles. They also adapt to the tactics of their enemies,
discarding obsolete ideas and dreaming up new ones.

## Ideas

Each player deploys a limited supply of boids with varying behavioral parameters
and they interact to achieve game outcomes.

Players should be able to save various configurations for quick use in-game. For
example, a "scout", or "defender", etc.

How does boid combat work? Maybe they automatically fire at enemies that are
within range? Maybe they have some kind of aggression parameter that can make
them fire at different ranges and be more or less likely to give chase?

Speed, armor, and weaponry should be trade-offs. Less weaponry and armor means
more speed, etc. Can we work sensors into this? It would be cool if some boids
could "see" further than others, but how would the trade-off work? Should other
parameters also involve trade-offs? Maybe there's a common pool of "points" and
they're split between all parameters?

Can boids communicate directly in any way? For example, can they "see" through
other boids at all? If so, what are the limitations? Maybe they get updated
infrequently, or the information is quantified or otherwise inexact. Maybe
communication occurs as a set of signals with no other contextual information
attached. For example, "I see enemies", or "I'm under attack".

Should the game be played in real time or should it pause occasionally to
allow new deployments / orders? Should existing boids be allowed to receive
new parameters? Alternatively, they could be recalled and then reprogrammed.

The goal of the boid fleets might be to capture "planets" (megaboids?). Not
sure how that would work. Maybe they "land" on the planet and are absorbed,
once a certain number are absorbed the planet becomes captured? Should the
number to capture be fixed, or should you be able to insulate your control
by pumping more boids in? Can they just fly into the planet or is there some
kind of landing process during which time they are vulnerable? Should defenders
land if they are about to lose a planet? We probably want a parameter of some
kind for whether or not, or how likely, a boid will land / capture if the
act causes the boid to be lost. What about orbiting instead of landing? That
way the orbiting boids can continue to defend? If only your boids are in orbit
then the planet is yours?

Maybe planets generate resources that let you produce more boids? Model those
mechanics after Konquest maybe? Should planets be random or should there be
"maps"? What about other obstacles? Maybe not, at least at first, because the
behavior is random, too many possible weird cases.

Possible boid parameters:

  * Aggressiveness
  * Capture / orbit affinity
  * Defensiveness
  * Adventurousness / willingness to scout (fog of war?)
  * Cohesiveness (stick together or spread out)

## Design

Boid movement is defined by a position and a velocity. At each simulation tick
the game logic is responsible for producing a new velocity, and the integrator
is responsible for producing a new position.
