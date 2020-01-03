# Picker
Picker is a package that randomly picks an element based on its weight

## Documentation

**Add** an element to picker


| Element | Weight |
| ------------- | -------------:|
| A | 80 |
| B | 90 |
| C | 10 |


```golang
p := picker.New()
p.Add("a", 80)
p.Add("b", 90)
p.Add("c", 10)
```

**Roll Dice** selects an element based on its probability of being taken. 
Higher weight elements are more likely to be choosen.

```golang
p := picker.New()
p.Add("a", 80)
p.Add("b", 90)
p.Add("c", 10)
element, err := p.RollDice()
```

**Reset** reset all variables of picker

```golang
p.Reset()
```