# Boid Storm

A game of emergent behavior.

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

## Design

Boid movement is defined by a position and a velocity. At each simulation tick
the game logic is responsible for producing a new velocity, and the integrator
is responsible for producing a new position.
