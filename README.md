# collapse
playing with wave function collapse

Having got the basic version running with cards with all sides the same, I'd like to play with the idea of controlling the chance for each card to appear -- so for example, the swamp card should be less common than the forest despite 
them having the same number of allowed neighbours.

One way to do this would be to have multiple copies of forest in the cards that link to it, but that seems crude, so I shall attempt to add a new field to the cards with a relative chance of that card appearing.

I think that the default should be 1, and then we could have the chance(swamp) = 0.8 and chance(forest) = 1.0.

Then I just need to work out the best way to change the overly simplistic randomisation to take this into account...


## to build
./make

which will build and copy the game code to c:\temp






## comments
Each time a card is added the board will change the entropy of the four cards surrounding it (assuming no walls etc).