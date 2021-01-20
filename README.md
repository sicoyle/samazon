# samazon
Samazon is a repo for me to practice Go on an Amazon/UPS delivery mock project.

The purpose of Samazon is to gain exposure to working more with websockets and practice creating a sort of event stream. The event stream here will be maintaining the driver's route and maintaining the drivers state by updating clients as the driver moves along their route. The event stream will maintain a time ticker and keep record of all stops, potentially dropping previously made stops.

As of now, distance will remain arbitrary. I'm thinking an iota of A to Z for locations. This will simulate high level that the driver is moving from house A -> house B -> ... -> house N, but its more finite here since I'm going to house Z. I'm going to use a random generator to establish the distances between houses and make it miles arbitrarily. Miles = min for the initial design.

Note: This is very much a work in progress and the readme contains some of my ideas for this, but not all as I'm expanding the scope.

Fun note: Samazon = Sam + Amazon if that wasn't obvious :)
