#### listleaguerules

This example retrieves all league rule modifiers from the API and prints their
names and descriptions.

Output:

```
$ go run main.go
Private: League requires a password to join.
Hardcore: A character killed in Hardcore is moved to its parent league.
Drop equipped items on death: Items are dropped on death.
Instance invasion: Allows you to select other people's instances in the instance manager.
Harsh death experience penalty: Increases the death experience penalty by 30% on all difficulty levels.
Hostile by default: Non-partymembers are hostile by default when you are not partied.
Death penalty awarded to slayer: When killing a player, their death penalty is awarded to the player doing the killing.
Increased player caps: Doubles player capacity in non-town instances. Does not increase the party size.
Turbo: Monsters move, attack and cast 60% faster.
Solo: You may not party in this league.
```